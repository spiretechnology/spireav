package remux

import (
	"path"

	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/graph/input"
	"github.com/spiretechnology/spireav/graph/output"
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

type Input struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
}

type StreamRef struct {
	Input  int `json:"input"`
	Stream int `json:"stream"`
}

type Config struct {
	Inputs          []Input     `json:"inputs"`
	VideoStream     StreamRef   `json:"video_stream"`
	AudioStreams    []StreamRef `json:"audio_streams"`
	OutDir          string      `json:"out_dir"`
	OverlayTimecode bool        `json:"overlay_timecode"`
	StartTimecode   string      `json:"start_timecode"`
	FrameRate       string      `json:"frame_rate"`
	OutSizes        []int       `json:"out_sizes"`
}

const (
	proxyHeight = 240
	thumbHeight = 120
)

func roundUpNearestMult(x int) int {
	if x%2 == 1 {
		return x + 1
	}
	return x
}

func Remux(config *Config) (graph.Graph, error) {

	// Create the transcoding graph
	g := graph.New()

	// Add all the inputs to the graph
	inputs := make([]input.Input, len(config.Inputs))
	for i, in := range config.Inputs {
		inputs[i] = g.AddInput(input.New(in.Filename))
	}

	// Create the MP4 output
	output240 := g.AddOutput(output.New(
		path.Join(config.OutDir, "240.mp4"),
		output.WithFormatMP4(),
	))

	// Create the thumbnail output
	outputThumb := g.AddOutput(output.New(
		path.Join(config.OutDir, "thumb.mp4"),
		output.WithFormatMP4(),
		output.WithFrameRate("0.5"),
	))

	// Create a scale node for the video
	scale := g.AddTransform(&transform.Scale{
		Width:  roundUpNearestMult(proxyHeight * 16 / 9),
		Height: roundUpNearestMult(proxyHeight),
	})

	// Create a scale node for the thumbnail
	scaleThumb := g.AddTransform(&transform.Scale{
		Width:  roundUpNearestMult(thumbHeight * 16 / 9),
		Height: roundUpNearestMult(thumbHeight),
	})

	// Link everything together for the primary output
	g.AddLink(inputs[config.VideoStream.Input], config.VideoStream.Stream, scale, 0)

	// The video node is the full-size video result. It might just be the scaled video, or it might be
	// the timecode overlay result, depending on the flag
	var videoNode graph.Node = scale
	if config.OverlayTimecode {

		// Create a timecode overlay node for the video
		// timecodeOverlay := g.AddTransform(&transform.TimecodeOverlay{
		// 	StartTimecode: videoInput.avidMeta.Timecode,
		// 	FrameRate:     videoInput.fileMeta.GetFrameRate(),
		// 	X:             "(w-tw)/2",
		// 	Y:             "h-th*2",
		// 	Box:           true,
		// 	FontColor:     "white",
		// 	FontSize:      "24",
		// 	FontFile:      r.TimecodeFont,
		// 	BoxColor:      "black@0.5",
		// })

		timecodeOverlay := g.AddTransform(transform.NewTextOverlay(
			transform.WithTimecode(
				config.StartTimecode,
				config.FrameRate,
			),
			// x = (w-tw)/2
			transform.WithX(
				expr.Div(
					expr.Sub(
						expr.Var("w"),
						expr.Var("tw"),
					),
					expr.Int(2),
				),
			),
			// y = h-th*2
			transform.WithY(
				expr.Sub(
					expr.Var("h"),
					expr.Mul(
						expr.Var("th"),
						expr.Int(2),
					),
				),
			),
			transform.WithBox("black@0.5"),
			transform.WithFontColor("white"),
			transform.WithFontSize(
				expr.Div(
					expr.Int(240), // expr.Var("h") this should be
					expr.Int(10),
				),
			),
		))

		g.AddLink(scale, 0, timecodeOverlay, 0)
		videoNode = timecodeOverlay
	}

	// Continue linking everything together
	g.AddLink(inputs[config.VideoStream.Input], config.VideoStream.Stream, scaleThumb, 0)
	g.AddLink(videoNode, 0, output240, 0)
	g.AddLink(scaleThumb, 0, outputThumb, 0)

	// Only add the audio merge if there are 1 or more audio streams
	if len(config.AudioStreams) > 0 {

		// Create a merge node for the audio
		amerge := g.AddTransform(&transform.AudioMerge{
			Inputs: len(config.AudioStreams),
		})

		// Map all the audio streams to the audio merge
		for i, audio := range config.AudioStreams {
			g.AddLink(inputs[audio.Input], audio.Stream, amerge, i)
		}

		// The result of the audio merge becomes the audio stream in the
		// non-muted outputs.
		g.AddLink(amerge, 0, output240, 1)

	}

	// Return the graph
	return g, nil

}
