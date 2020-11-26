package main

import (
	"fmt"

	"github.com/spiretechnology/spireav"
)

func main() {

	graph := spireav.NewGraph()

	input := graph.AddNode(spireav.NewFileInputNode("/Users/conner/Downloads/BigBuckBunny.mp4"))
	output := graph.AddNode(spireav.NewFileOutputNode("/Users/conner/Desktop/graph-out.mp4"))

	textOverlay := graph.AddNode(spireav.TextOverlayNode{
		Text: "SpireAV Test",
	})

	graph.AddLink(input, 0, textOverlay, 0)
	graph.AddLink(textOverlay, 0, output, 0)

	err := graph.ProduceOutput(output)
	if err != nil {
		fmt.Println(err.Error())
	}

}
