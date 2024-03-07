package main

import (
	"context"
	"fmt"

	"github.com/spiretechnology/spireav/metadata"
)

func main() {

	// Get the metadata for a file
	metadata, err := metadata.GetMetadata(
		context.Background(),
		"reference-media/BigBuckBunny.mp4",
	)
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Printf("%+v\n", metadata)

}
