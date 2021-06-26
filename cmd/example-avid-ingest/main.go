package main

import (
	"fmt"

	"github.com/spiretechnology/spireav/meta"
	"github.com/spiretechnology/spireav/routines/avid"
)

func main() {

	// Get the metadata for a file
	metadata, err := meta.GetMetadata("/Users/conner/Desktop/Sample Transcodes/SC0808GB.01/SC0808GB_AV01.D9722BAE086DV.mxf", "")
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
