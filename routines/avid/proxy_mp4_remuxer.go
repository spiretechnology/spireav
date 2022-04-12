package avid

import (
	"context"
	"errors"
	"sort"
	"strconv"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/routines/remux"
)

type proxyMP4RemuxerInput struct {
	filename string
	fileMeta *spireav.Meta
	avidMeta *AvidMxfMeta
}

// ProxyMP4Remuxer manages the job of remuxing multiple op-atom MXF files from Avid back into a proxy in the
// MP4 format so it can be viewed over the web
type ProxyMP4Remuxer struct {
	mxfInputs       []*proxyMP4RemuxerInput
	OverlayTimecode bool
	TimecodeFont    string
}

// NewProxyMP4Remuxer creates a new remuxer for Avid Media Composer proxies into MP4 files.
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

func (r *ProxyMP4Remuxer) GenerateGraph(outDir string) (graph.Graph, error) {

	// Create the base config for the remux
	config := remux.Config{
		OverlayTimecode: r.OverlayTimecode,
		OutSizes:        []int{480, 240},
		ThumbnailSize:   120,
		OutDir:          outDir,
	}

	var hasVideoStream bool

	// As we loop through below, put the video stream here
	for i, input := range r.mxfInputs {

		// If this stream isn't video or audio, skip it
		if input.avidMeta.EssenceStream.CodecType != "video" && input.avidMeta.EssenceStream.CodecType != "audio" {
			continue
		}

		// Add this as an input to the graph
		config.Inputs = append(config.Inputs, remux.Input{
			Filename: input.filename,
			Type:     input.avidMeta.EssenceStream.CodecType,
		})

		// If this is the video stream, use it
		if input.avidMeta.EssenceStream.CodecType == "video" {

			// We found a video stream
			hasVideoStream = true

			// Update the video fields for the remux
			config.StartTimecode = input.avidMeta.Timecode
			config.FrameRate = strconv.FormatFloat(input.fileMeta.GetFrameRate(), 'f', 3, 64)

			// Add the video stream reference to the config
			config.VideoStream = remux.StreamRef{
				Input:  i,
				Stream: input.avidMeta.EssenceStream.Index,
			}

		} else if input.avidMeta.EssenceStream.CodecType == "audio" {

			// Add the audio stream reference to the config
			config.AudioStreams = append(config.AudioStreams, remux.StreamRef{
				Input:  i,
				Stream: input.avidMeta.EssenceStream.Index,
			})

		}

	}

	// Sort the audio streams in ascending order by their stream index. This is a trick for
	// Avid transcoding only.
	sort.Slice(config.AudioStreams, func(i, j int) bool {
		return config.AudioStreams[i].Stream < config.AudioStreams[j].Stream
	})

	// If we didn't find a video stream
	if !hasVideoStream {
		return nil, errors.New("no video stream found for remux")
	}

	// Return the graph
	return remux.Remux(&config)

}

func (r *ProxyMP4Remuxer) EstimateLengthFrames() int {
	for _, input := range r.mxfInputs {
		if input.avidMeta.EssenceStream.CodecType == "video" {
			return input.fileMeta.GetLengthFrames()
		}
	}
	return 0
}
