package main

import (
	"fmt"

	"github.com/spiretechnology/spireav/meta"
	"github.com/spiretechnology/spireav/routines/avid"
)

func main() {

	// Get the metadata for a file
	metadata, err := meta.GetMetadata("/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBV01.D97B84568C1FF8V.mxf", "")
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
