package main

import (
	"fmt"
	"os"

	"github.com/spiretechnology/spireav/graph"
)

func main() {

	// Create a new graph
	g := graph.NewGraph()

	// Open the input file
	fileIn, err := os.Open("/Users/conner/Downloads/BigBuckBunny.mp4")
	if err != nil {
		panic(err)
	}
	defer fileIn.Close()

	input, err := graph.NewInputNode(
		fileIn,
		graph.FormatWithName("mp4"),
	)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// Create the output
	fileOut, err := os.OpenFile("/Users/conner/Desktop/graph-out.mp4", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer fileOut.Close()
	output, err := graph.NewOutputNode(
		fileOut,
		graph.FormatWithName("mp4"),
		[]*graph.StreamDescription{
			graph.NewStreamDescription("libx264", func(stream *graph.OutputStream) {
				stream.SetWidth(640)
				stream.SetHeight(360)
				stream.SetAspectRatio(16, 9)
				stream.SetFrameRate(24, 1)
				stream.SetPixFmt(stream.GetEncoder().GetDefaultPixFmt())
				stream.SetTimeBase(1, 90000)
			}),
		},
	)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// Add inputs and outputs to the graph
	inputNode := g.AddNode(input)
	outputNode := g.AddNode(output)

	// Create a text overlay node
	textOverlay := g.AddNode(graph.TextOverlayNode{
		Text: "SpireAV Test",
	})

	// Link together the nodes to create a video processing pipeline
	g.AddLink(inputNode, 0, textOverlay, 0)
	g.AddLink(textOverlay, 0, outputNode, 0)

	// Produce the output
	if err := g.NewJob().Run(); err != nil {
		fmt.Println(err.Error())
	}

}
