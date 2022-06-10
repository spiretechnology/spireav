package remux

import (
	"fmt"
	"path"

	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/graph/input"
	"github.com/spiretechnology/spireav/graph/output"
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

type StreamRef struct {
	Input  int `json:"input"`
	Stream int `json:"stream"`
}

type Output struct {
	Size            int    `json:"size"`
	Format          string `json:"format"`
	OverlayTimecode bool   `json:"overlay_timecode"`
}

// Config defines the configuration for a remux operation. This includes the files to use as inputs,
// the mappings of the streams, output sizes, etc.
type Config struct {
	Inputs        []string
	VideoStream   StreamRef
	AudioStreams  []StreamRef
	StartTimecode string
	FrameRate     string
	Outputs       []Output
	OutDir        string
	TimecodeFont  string
	ThumbnailSize int
}

func roundUpNearestMult(x int) int {
	if x%2 == 1 {
		return x + 1
	}
	return x
}

// Remux creates a transcoding graph to perform a generic remux operation.
func Remux(config *Config) (graph.Graph, error) {

	// Create the transcoding graph
	g := graph.New()

	// Add all the inputs to the graph
	inputs := make([]input.Input, len(config.Inputs))
	for i, filename := range config.Inputs {
		inputs[i] = g.AddInput(input.New(filename))
	}

	type SizeContext struct {
		Options Output
		Video   transform.Transform
		Output  output.Output
	}

	// Create the MP4 output for all the output sizes
	sizeContexts := make([]SizeContext, len(config.Outputs))
	for i, out := range config.Outputs {

		// Create the scale transform
		scale := g.AddTransform(&transform.Scale{
			Width:  -2,
			Height: roundUpNearestMult(out.Size),
		})

		// Create the output for this size
		outputFilename := path.Join(config.OutDir, fmt.Sprintf("%d.%s", out.Size, out.Format))

		// Create the output for the
		var sizedOutput output.Output
		if out.Format == "mp4" {
			sizedOutput = g.AddOutput(output.New(
				outputFilename,
				output.WithFormatMP4(),
				output.WithConstantRateFactor(30),
			))
		} else if out.Format == "mov" {
			sizedOutput = g.AddOutput(output.New(
				outputFilename,
				output.WithFormatMOV(),
				output.WithTimecode(config.StartTimecode),
				output.WithConstantRateFactor(30),
			))
		} else {
			return nil, fmt.Errorf("unrecognized output format: %q", out.Format)
		}

		// Create the context for this size
		sizeContexts[i] = SizeContext{
			Options: out,
			Video:   scale,
			Output:  sizedOutput,
		}

		// Go ahead and feed the scale transform with the input video stream
		g.AddLink(inputs[config.VideoStream.Input], config.VideoStream.Stream, scale, 0)

	}

	// Create the thumbnail output
	outputThumb := g.AddOutput(output.New(
		path.Join(config.OutDir, "thumb.mp4"),
		output.WithFormatMP4(),
		output.WithFrameRate("0.5"),
	))

	// Create a scale node for the thumbnail
	scaleThumb := g.AddTransform(&transform.Scale{
		Width:  -2,
		Height: roundUpNearestMult(config.ThumbnailSize),
	})

	// Loop through all the sizes
	for _, sizeContext := range sizeContexts {

		// If we're burning timecode, add the timecode burn node
		if sizeContext.Options.OverlayTimecode {

			// Create the timecode overlay transform node
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
				transform.WithFontFile(config.TimecodeFont),
				transform.WithFontColor("white"),
				transform.WithFontSize(
					expr.Div(
						expr.Var("h"),
						expr.Int(10),
					),
				),
			))

			// Pipe the video into the timecode overlay
			g.AddLink(sizeContext.Video, 0, timecodeOverlay, 0)
			sizeContext.Video = timecodeOverlay

		}

		// Pipe the video of this size to the output
		g.AddLink(sizeContext.Video, 0, sizeContext.Output, 0)

	}

	// Link everything together for the thumbnail
	g.AddLink(inputs[config.VideoStream.Input], config.VideoStream.Stream, scaleThumb, 0)
	g.AddLink(scaleThumb, 0, outputThumb, 0)

	// Only add the audio merge if there are 1 or more audio streams
	if len(config.AudioStreams) > 0 {

		// Create a merge node for the audio
		amerge := g.AddTransform(&transform.AudioMerge{
			Inputs: len(config.AudioStreams),
		})

		// Create an audio split transform
		asplit := g.AddTransform(transform.New(
			fmt.Sprintf("asplit=%d", len(sizeContexts)),
			len(sizeContexts),
		))
		g.AddLink(amerge, 0, asplit, 0)

		// Map all the audio streams to the audio merge
		for i, audio := range config.AudioStreams {
			g.AddLink(inputs[audio.Input], audio.Stream, amerge, i)
		}

		// Loop through all the sizes and pipe the audio to the second stream
		for i, sizeContext := range sizeContexts {
			g.AddLink(asplit, i, sizeContext.Output, 1)
		}

	}

	// Return the graph
	return g, nil

}
