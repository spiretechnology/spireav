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

	foreground := g.Input("reference-media/BigBuckBunny.mp4")

	// Scale and rotate the foreground video
	scaleAndRotate := spireav.Pipeline(
		g.Filter(filters.Scale().WidthExpr(expr.Int(200)).HeightExpr(expr.Int(200))),
		g.Filter(filters.Pad().WExpr(expr.Raw("max(iw,ih)*sqrt(2)")).HExpr(expr.Raw("max(ih,iw)*sqrt(2)")).X(-1).Y(-1).Color(expr.ColorTransparent)),
		g.Filter(
			filters.Rotate().
				AngleExpr(expr.Div(expr.Mul(expr.Var("t"), expr.PI), expr.Int(4))).
				Fillcolor(expr.ColorTransparent.String())),
	)
	foreground.Video(0).Pipe(scaleAndRotate, 0)

	// Overlay the scaled and rotated foreground video on top of a solid background
	overlayFilter := g.Filter(filters.Overlay().X(0).Y(0))
	background := g.Filter(filters.Color().Color(expr.ColorRed).Duration(10 * time.Second).Size(expr.Size{Width: 1280, Height: 720}))
	background.Pipe(overlayFilter, 0)
	scaleAndRotate.Pipe(overlayFilter, 1)

	// Create the output file
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
