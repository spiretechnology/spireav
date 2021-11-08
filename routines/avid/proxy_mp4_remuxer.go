package avid

import (
	"context"
	"errors"
	"path"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/graph/transform"
)

type proxyMP4RemuxerInput struct {
	filename string
	fileMeta *spireav.Meta
	avidMeta *AvidMxfMeta
	node     graph.InputNode
}

// ProxyMP4Remuxer manages the job of remuxing multiple op-atom MXF files from Avid back into a proxy in the
// MP4 format so it can be viewed over the web
type ProxyMP4Remuxer struct {
	mxfInputs       []*proxyMP4RemuxerInput
	OverlayTimecode bool
}

// ffmpeg -i Z:\Avid MediaFiles\MXF\Transcode3.1\WC0624CC_AV01.DCFA3E60ECBFV.mxf -i Z:\Avid MediaFiles\MXF\Transcode3.1\WC0624CC_AA02.DCFA3E60F1FAA.mxf -i Z:\Avid MediaFiles\MXF\Transcode3.1\WC0624CC_AA01.DCFA3E60F1E1A.mxf -y -filter_complex [0:v:0]scale=426:240[genczFA];[genczFA]drawtext=fontfile='C\:\\Windows\\Fonts\\Verdana.ttf':x=(w-tw)/2:y=h-th*2:fontcolor=white:fontsize=24:box=1:boxcolor=black@0.5:boxborderw=5:timecode='15\:51\:33\:04':r=23.980[genxXgD];[1:a:0][2:a:0]amerge=inputs=2[genytl5];[0:v:0]scale=214:120[genY2f9];[genY2f9]drawtext=fontfile='C\:\\Windows\\Fonts\\Verdana.ttf':x=(w-tw)/2:y=h-th*2:fontcolor=white:fontsize=24:box=1:boxcolor=black@0.5:boxborderw=5:timecode='15\:51\:33\:04':r=23.980[genojHw];[1:a:0][2:a:0]amerge=inputs=2[genngKX] -map [genxXgD] -map [genytl5] -movflags use_metadata_tags -pix_fmt yuv420p -movflags +faststart -f mp4 D:\MEDIA_OUT\TV21\WC0624CC_AMA\WC0624CC_AMA.01.new.01\240.mp4 -r 0.5 -map [genojHw] -map [genngKX] -an -movflags use_metadata_tags -pix_fmt yuv420p -movflags +faststart -f mp4 D:\MEDIA_OUT\TV21\WC0624CC_AMA\WC0624CC_AMA.01.new.01\thumb.mp4

func NewProxyMP4Remuxer(mxfFiles []string) (*ProxyMP4Remuxer, error) {

	// Create the slice of all MXF inputs
	mxfInputs := []*proxyMP4RemuxerInput{}
	for _, filename := range mxfFiles {
		input, err := proxyMP4LoadInput(filename)
		if err != nil {
			return nil, err
		}
		mxfInputs = append(mxfInputs, input)
	}

	// Return the remuxer instance
	return &ProxyMP4Remuxer{
		mxfInputs: mxfInputs,
	}, nil

}

func proxyMP4LoadInput(filename string) (*proxyMP4RemuxerInput, error) {

	// Get the metadata for the MXF file
	fileMeta, err := spireav.GetMetadata(
		context.Background(),
		filename,
	)
	if err != nil {
		return nil, err
	}
	avidMeta, err := ParseAvidMxfOpAtomMeta(fileMeta)
	if err != nil {
		return nil, err
	}

	// Create the input
	return &proxyMP4RemuxerInput{
		filename: filename,
		fileMeta: fileMeta,
		avidMeta: avidMeta,
	}, nil

}

func (r *ProxyMP4Remuxer) GenerateProc(outDir string) (*spireav.Process, error) {

	// Generate the graph
	g, err := r.GenerateGraph(outDir)
	if err != nil {
		return nil, err
	}

	// Create the process
	return &spireav.Process{
		FfmpegArger:           g,
		EstimatedLengthFrames: r.EstimateLengthFrames(),
	}, nil

}

func (r *ProxyMP4Remuxer) GenerateGraph(outDir string) (*graph.Graph, error) {

	// Create a graph
	g := graph.Graph{}

	// As we loop through below, put the video stream here
	var videoInput *proxyMP4RemuxerInput

	// And put all of the audio inputs here
	audioInputs := []*proxyMP4RemuxerInput{}

	// Register all of the inputs to the graph
	for _, input := range r.mxfInputs {

		// Create the file input node
		input.node = g.AddInput(&graph.FileInputNode{
			Filename: input.filename,
		})

		// If this is the video stream, use it
		if input.avidMeta.EssenceStream.CodecType == "video" {
			videoInput = input
		} else if input.avidMeta.EssenceStream.CodecType == "audio" {
			audioInputs = append(audioInputs, input)
		}

	}

	// If no video stream was found
	if videoInput == nil {
		return nil, errors.New("no video stream found for proxy transcode")
	}

	// Create the MP4 output
	output := g.AddOutput(&graph.OutputNodeMP4{
		Filename: path.Join(outDir, "240.mp4"),
	})

	// Create the thumbnail output
	outputThumb := g.AddOutput(&graph.OutputNodeMP4{
		Filename:  path.Join(outDir, "thumb.mp4"),
		FrameRate: "0.5",
	})

	// Create a scale node for the video
	scale := g.AddTransform(&transform.Scale{
		Width:  426,
		Height: 240,
	})

	// Create a scale node for the thumbnail
	scaleThumb := g.AddTransform(&transform.Scale{
		Width:  214,
		Height: 120,
	})

	// Create a merge node for the audio
	amerge := g.AddTransform(&transform.AudioMerge{
		Inputs: len(audioInputs),
	})

	// Link everything together for the primary output
	g.AddLink(videoInput.node, videoInput.avidMeta.EssenceStream.Index, scale, 0)

	// The video node is the full-size video result. It might just be the scaled video, or it might be
	// the timecode overlay result, depending on the flag
	var videoNode graph.Node = scale
	if r.OverlayTimecode {
		// Create a timecode overlay node for the video
		timecodeOverlay := g.AddTransform(&transform.TimecodeOverlay{
			StartTimecode: videoInput.avidMeta.Timecode,
			FrameRate:     videoInput.fileMeta.GetFrameRate(),
			X:             "(w-tw)/2",
			Y:             "h-th*2",
			Box:           true,
			FontColor:     "white",
			BoxColor:      "black@0.5",
		})
		// FontSize = 24
		g.AddLink(scale, 0, timecodeOverlay, 0)
		videoNode = timecodeOverlay
	}

	// Continue linking everything together
	g.AddLink(videoInput.node, videoInput.avidMeta.EssenceStream.Index, scaleThumb, 0)
	g.AddLink(videoNode, 0, output, 0)
	g.AddLink(scaleThumb, 0, outputThumb, 0)
	for i, audio := range audioInputs {
		g.AddLink(audio.node, audio.avidMeta.EssenceStream.Index, amerge, i)
	}
	g.AddLink(amerge, 0, output, 1)

	// Return the graph
	return &g, nil

}

func (r *ProxyMP4Remuxer) EstimateLengthFrames() int {
	for _, input := range r.mxfInputs {
		if input.avidMeta.EssenceStream.CodecType == "video" {
			return input.fileMeta.GetLengthFrames()
		}
	}
	return 0
}
