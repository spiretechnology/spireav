package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/filter/filters"
	"github.com/spiretechnology/spireav/output"
)

func main() {
	// Create a new graph
	g := spireav.New()

	background := g.Filter(filters.NullSource().Duration("10").Size(1280, 720))

	foreground := g.Input("reference-media/BigBuckBunny.mp4")
	scaleForeground := g.Filter(filters.Scale().WidthInt(200).HeightInt(200))
	foreground.Pipe(scaleForeground, 0)

	rotateForeground := g.Filter(filters.Rotate().Angle("t*PI/4"))
	scaleForeground.Pipe(rotateForeground, 0)

	overlayFilter := g.Filter(filters.Overlay().PosX("0").PosY("0"))
	background.Pipe(overlayFilter, 0)
	rotateForeground.Pipe(overlayFilter, 1)

	outputFile := g.Output("reference-outputs/out-overlay.mp4", output.WithFormatMP4())

	overlayFilter.Pipe(outputFile, 0)
	foreground.Stream(1).Pipe(outputFile, 1)

	// Create a progress handler function
	progressFunc := func(progress spireav.Progress) {
		fmt.Printf("%+v\n", progress)
	}

	// Create a context for the transcoding job
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Create the process
	runner := spireav.NewRunner(g, spireav.WithProgressCallback(progressFunc))

	// Run the transcoding job
	if err := runner.Run(ctx); err != nil {
		fmt.Println(err.Error())
	}
}
