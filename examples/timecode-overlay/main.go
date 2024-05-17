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
	inputNode := g.Input("reference-media/BigBuckBunny.mp4")
	outputNode := g.Output(
		"reference-outputs/BigBuckBunny-timecode.mp4",
		output.WithFormatMP4(),
	)

	// Overlay timecode on the video
	timecode := g.Filter(
		filters.Drawtext().
			Timecode("00:00:00:00").
			TimecodeRate(expr.RateFilm).
			Fontcolor(expr.ColorWhite).
			Box(true).
			Boxcolor(expr.ColorBlack.WithOpacity(0.5)).
			FontsizeExpr(expr.Div(expr.Var("h"), expr.Int(10))).
			XExpr(expr.Div(expr.Sub(expr.Var("w"), expr.Var("tw")), expr.Int(2))).
			YExpr(expr.Sub(expr.Var("h"), expr.Var("th"), expr.Int(20))),
	)

	// Pass the video into the pipeline
	inputNode.Video(0).Pipe(timecode, 0)

	// Pass the pipeline output to the output file
	timecode.Pipe(outputNode, 0)

	// Pass the audio directly to the output file
	inputNode.Audio(0).Pipe(outputNode, 1)

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
