package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/graph/output"
	"github.com/spiretechnology/spireav/graph/transform/drawtext"
	"github.com/spiretechnology/spireav/graph/transform/expr"
	"github.com/spiretechnology/spireav/graph/transform/scale"
)

func main() {

	// Create a new graph
	g := graph.New()
	inputNode := g.NewInput("reference-media/BigBuckBunny.mp4")
	outputNode := g.NewOutput(
		"reference-outputs/BigBuckBunny-GRAPH.mp4",
		output.WithFormatMP4(),
	)

	// Create a text overlay node
	textOverlay := g.AddTransform(drawtext.DrawText(
		drawtext.WithText("SpireAV Test"),
	))

	// Create a text overlay node
	scaleNode := g.AddTransform(scale.Scale(
		scale.WithWidth(expr.Int(200)),
		scale.WithHeight(expr.Int(200)),
	))

	// Link together the nodes to create a video processing pipeline
	g.AddLink(inputNode, 0, scaleNode, 0)
	g.AddLink(scaleNode, 0, textOverlay, 0)
	g.AddLink(textOverlay, 0, outputNode, 0)
	g.AddLink(inputNode, 1, outputNode, 1)

	// Create the process
	p := spireav.Process{
		FfmpegArger:           g,
		EstimatedLengthFrames: 14315,
	}

	// Create a progress handler function
	progressFunc := func(progress spireav.Progress) {
		fmt.Printf("P: %+v\n", progress)
		fmt.Printf("E: %+v\n", progress.Estimate)
	}

	// Create a context for the transcoding job
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Run the transcoding job
	if err := p.RunWithProgress(ctx, progressFunc); err != nil {
		fmt.Println(err.Error())
	}

}
