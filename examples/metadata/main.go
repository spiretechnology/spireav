package main

import (
	"fmt"

	"github.com/spiretechnology/spireav"
)

func main() {

	// Get the metadata for a file
	metadata, err := spireav.GetMetadata("reference-media/BigBuckBunny.mp4", "")
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Printf("%+v\n", metadata)

}
