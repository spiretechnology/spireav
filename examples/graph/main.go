package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/filter/drawtext"
	"github.com/spiretechnology/spireav/filter/expr"
	"github.com/spiretechnology/spireav/filter/scale"
	"github.com/spiretechnology/spireav/output"
)

func main() {
	// Create a new graph
	g := spireav.New()
	inputNode := g.NewInput("reference-media/BigBuckBunny.mp4")
	outputNode := g.NewOutput(
		"reference-outputs/BigBuckBunny-GRAPH.mp4",
		output.WithFormatMP4(),
	)

	// Create a text overlay node
	textOverlay := g.NewFilter(drawtext.DrawText(
		drawtext.WithText("SpireAV Test"),
	))

	// Create a text overlay node
	scaleNode := g.NewFilter(scale.Scale(
		scale.WithWidth(expr.Int(200)),
		scale.WithHeight(expr.Int(200)),
	))

	// Link together the nodes to create a video processing pipeline
	inputNode.Stream(0).Pipe(scaleNode, 0)
	scaleNode.Stream(0).Pipe(textOverlay, 0)
	textOverlay.Stream(0).Pipe(outputNode, 0)
	inputNode.Stream(1).Pipe(outputNode, 1)

	// Create a progress handler function
	progressFunc := func(progress spireav.Progress) {
		fmt.Printf("P: %+v\n", progress)
		fmt.Printf("E: %+v\n", progress.Estimate)
	}

	// Create a context for the transcoding job
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Create the process
	runner := spireav.NewRunner(g,
		spireav.WithEstimatedLengthFrames(14315),
		spireav.WithProgressCallback(progressFunc),
	)

	// Run the transcoding job
	if err := runner.Run(ctx); err != nil {
		fmt.Println(err.Error())
	}
}
