package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/filter/expr"
	"github.com/spiretechnology/spireav/filter/filters"
	"github.com/spiretechnology/spireav/output"
)

func main() {
	// Create a new graph
	g := spireav.New()

	background := g.Filter(filters.Nullsrc().Duration(10 * time.Second).Size(expr.Size{Width: 1280, Height: 720}))

	foreground := g.Input("reference-media/BigBuckBunny.mp4")
	scaleForeground := g.Filter(filters.Scale().WidthExpr(expr.Int(200)).HeightExpr(expr.Int(200)))
	foreground.Video(0).Pipe(scaleForeground, 0)

	rotateForeground := g.Filter(filters.Rotate().Angle("t*PI/4"))
	scaleForeground.Pipe(rotateForeground, 0)

	overlayFilter := g.Filter(filters.Overlay().X("0").Y("0"))
	background.Pipe(overlayFilter, 0)
	rotateForeground.Pipe(overlayFilter, 1)

	outputFile := g.Output("reference-outputs/out-overlay.mp4", output.WithFormatMP4())

	overlayFilter.Pipe(outputFile, 0)
	foreground.Audio(0).Pipe(outputFile, 1)

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
