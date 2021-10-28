package main

import (
	"fmt"

	"github.com/spiretechnology/spireav/meta"
)

func main() {

	// Get the metadata for a file
	metadata, err := meta.GetMetadata("reference-media/BigBuckBunny.mp4", "")
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Printf("%+v\n", metadata)

}
