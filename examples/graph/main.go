package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/filter/expr"
	"github.com/spiretechnology/spireav/filter/filters"
	"github.com/spiretechnology/spireav/output"
)

func main() {
	// Create a new graph
	g := spireav.New()
	inputNode := g.Input("reference-media/BigBuckBunny.mp4")
	outputNode := g.Output(
		"reference-outputs/BigBuckBunny-GRAPH.mp4",
		output.WithFormatMP4(),
	)

	// Create a pipeline of filters
	scaleAndOverlay := spireav.Pipeline(
		// Scale the video to 200x200
		g.Filter(filters.Scale().WExpr(expr.Int(200)).HExpr(expr.Int(200))),
		// Overlay text on the video
		g.Filter(filters.Drawtext().Text("SpireAV Test")),
	)

	// Pass the video into the pipeline
	inputNode.Video(0).Pipe(scaleAndOverlay, 0)

	// Pass the pipeline output to the output file
	scaleAndOverlay.Pipe(outputNode, 0)

	// Pass the audio directly to the output file
	inputNode.Audio(0).Pipe(outputNode, 1)

	// Create a progress handler function
	progressFunc := func(progress spireav.Progress) {
		fmt.Printf("%+v\n", progress)
	}

	// Create a context for the transcoding job
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Create the process
	runner := spireav.NewRunner(g, spireav.WithProgressCallback(progressFunc))

	// Run the transcoding job
	if err := runner.Run(ctx); err != nil {
		fmt.Println(err.Error())
	}
}
