package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav/meta"
	"github.com/spiretechnology/spireav/routines/avid"
)

func main() {

	// Create a context for the metadata operation
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Get the metadata for a file
	metadata, err := meta.GetMetadataContext(ctx, "reference-media/SC0808GB.01/SC0808GB_AA01.D9722BAE3008A.mxf", "")
	if err != nil {
		panic(err)
	}

	// Parse out the Avid specific metadata
	avidMeta, err := avid.ParseAvidMxfOpAtomMeta(metadata)
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Printf("%+v\n", avidMeta)

}
