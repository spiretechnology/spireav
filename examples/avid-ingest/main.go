package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav/routines/avid"
)

func main() {
	// Create a context for the metadata operation
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Get the metadata for the MXF file
	avidMeta, _, err := avid.GetMetadata(ctx, "reference-media/WS1108AA/WS1108AAA01.DDAF2E31D4D43EA.mxf")
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Printf("%+v\n", avidMeta)
}
