package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spiretechnology/spireav/mxf2raw"
)

func main() {
	// Setup the mxf2raw path
	mxf2raw.Mxf2RawPath = "/Users/conner/Desktop/bmx/out/build/apps/mxf2raw/mxf2raw"

	// Create the context for the process
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	// Get the metadata for a file
	metadata, err := mxf2raw.GetMetadataWithOptions(ctx, "/Users/conner/Downloads/Stage AAF/AA01B7EAC60C.mxf", nil)
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Println("Clip name:       ", metadata.Clip.Name)
	fmt.Println("Start timecode:  ", metadata.Clip.StartTimecodes.PhysicalSource.Timecode)
}
