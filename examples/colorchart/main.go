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
	// inputNode := g.Input("reference-media/BigBuckBunny.mp4")
	outputNode := g.Output("reference-outputs/out-colorchart.mp4", output.WithFormatMP4())

	textOverlay := g.Filter(filters.DrawText().Text("SpireAV Test"))
	videoSrc := g.Filter(filters.SMPTEHDBars().Duration("10").FrameRate(expr.Int(30)))
	spireav.Pipeline(videoSrc, textOverlay, outputNode)

	// Sin wave from 300 to 1000 Hz, period of 2 seconds
	// freq := expr.Add(expr.Int(300), expr.Mul(expr.Int(700), expr.Sin0to1("t")))
	freq := expr.Int(440)
	fmt.Println(freq.String())
	fmt.Println()

	audioSrc := g.Filter(filters.SineSource().Frequency(freq).Duration("10"))
	audioSrc.Pipe(outputNode, 1)

	// Create a progress handler function
	progressFunc := func(progress spireav.Progress) {
		fmt.Printf("P: %+v\n", progress)
		fmt.Printf("E: %+v\n", progress.Estimate)
	}

	// Create a context for the transcoding job
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Create the process
	runner := spireav.NewRunner(g,
		spireav.WithEstimatedLengthFrames(14315),
		spireav.WithProgressCallback(progressFunc),
	)

	// Run the transcoding job
	if err := runner.Run(ctx); err != nil {
		fmt.Println(err.Error())
	}
}
