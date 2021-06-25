package main

import (
	"fmt"

	"github.com/spiretechnology/spireav/graph"
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/proc"
)

func main() {

	// Create a new graph
	g := graph.Graph{}
	inputNode := g.AddInput(&graph.FileInputNode{
		Filename: "/Users/conner/Downloads/BigBuckBunny.mp4",
	})
	outputNode := g.AddOutput(&graph.OutputNodeMP4{
		Filename: "/Users/conner/Desktop/graph-out-new.mp4",
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

	// Run the transcoding job
	if err := p.RunWithProgress(progressFunc); err != nil {
		fmt.Println(err.Error())
	}

}
