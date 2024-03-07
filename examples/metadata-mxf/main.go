package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/metadata"
)

func main() {
	// Setup the mxf2raw path
	spireav.Mxf2RawPath = "/Users/conner/Desktop/bmx/out/build/apps/mxf2raw/mxf2raw"

	// Create the context for the process
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	// Get the metadata for a file
	metadata, err := metadata.GetMxfMetadata(ctx, "/Users/conner/Downloads/Stage AAF/AA01B7EAC60C.mxf")
	if err != nil {
		panic(err)
	}

	// Log the metadata
	fmt.Println("Clip name:       ", metadata.Clip.Name)
	fmt.Println("Start timecode:  ", metadata.Clip.StartTimecodes.PhysicalSource.Timecode)
}
