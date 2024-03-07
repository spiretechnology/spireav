package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav"
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

	// Create a text overlay node
	textOverlay := g.Filter(filters.DrawText().Text("SpireAV Test"))

	// Create a text overlay node
	scaleNode := g.Filter(filters.Scale().WidthInt(200).HeightInt(200).IgnoreAspectRatio())

	// Link together the nodes to create a video processing pipeline
	spireav.Pipeline(inputNode, scaleNode, textOverlay, outputNode)
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
