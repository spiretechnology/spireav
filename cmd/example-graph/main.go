package main

import (
	"fmt"
	"log"

	"github.com/spiretechnology/spireav/graph"
)

func main() {

	// Create a new graph
	g := graph.NewGraph()

	// Create the input
	input := graph.NewFileInputNode("/Users/conner/Downloads/BigBuckBunny.mp4")
	inStream, err := input.GetStream(0)
	if err != nil {
		log.Panicln(err)
	}
	inStream.OpenDecoderContext()
	if err != nil {
		log.Panicln(err)
	}

	// Create the output
	output := graph.NewFileOutputNode("/Users/conner/Desktop/graph-out.mp4")
	encoder, err := inStream.GetCodec().GetEncoder()
	if err != nil {
		log.Panicln(err)
	}
	outStream, err := output.AddStream(encoder)
	if err != nil {
		log.Panicln(err)
	}
	outStream.OpenEncoderContext()
	if err != nil {
		log.Panicln(err)
	}

	graph.CopyStreamSettings(inStream, outStream)

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
	err = g.NewJob().Run()
	if err != nil {
		fmt.Println(err.Error())
	}

}
