package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/proc"
)

func main() {

	// Create a new graph
	g := graph.Graph{}
	inputNode := g.AddInput(&graph.FileInputNode{
		Filename: "reference-media/BigBuckBunny.mp4",
	})
	outputNode := g.AddOutput(&graph.OutputNodeMP4{
		Filename: "reference-outputs/BigBuckBunny-GRAPH.mp4",
		// FrameRate: "0.5",
	})

	// Create a text overlay node
	textOverlay := g.AddTransform(&transform.TextOverlay{
		Text: "SpireAV Test",
	})

	// Create a text overlay node
	scale := g.AddTransform(&transform.Scale{
		Width:  200,
		Height: 200,
	})

	// Link together the nodes to create a video processing pipeline
	g.AddLink(inputNode, 0, scale, 0)
	g.AddLink(scale, 0, textOverlay, 0)
	g.AddLink(textOverlay, 0, outputNode, 0)
	g.AddLink(inputNode, 1, outputNode, 1)

	// Create the process
	p := proc.Proc{
		Graph:                 &g,
		EstimatedLengthFrames: 14315,
	}

	// Create a progress handler function
	progressFunc := func(progress proc.Progress) {
		fmt.Printf("P: %+v\n", progress)
		fmt.Printf("E: %+v\n", progress.Estimate)
	}

	// Create a context for the transcoding job
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Run the transcoding job
	if err := p.RunWithProgressContext(ctx, progressFunc); err != nil {
		fmt.Println(err.Error())
	}

}
