package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/audiosrc"
	"github.com/spiretechnology/spireav/filter/expr"
	"github.com/spiretechnology/spireav/filter/vidsrc"
	"github.com/spiretechnology/spireav/output"
)

var sourceNames = []string{
	// "allrgb",
	// "allyuv",
	// "color",
	// "colorchart",
	// "colorspectrum",
	// "haldclutsrc",
	// "nullsrc",
	// "pal75bars",
	// "pal100bars",
	// "rgbtestsrc",
	// "smptebars",
	"smptehdbars",
	// "testsrc",
	// "testsrc2",
	// "testsrc2",
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(sourceNames))
	for _, sourceName := range sourceNames {
		// Create a new graph
		g := spireav.New()
		// inputNode := g.Input("reference-media/BigBuckBunny.mp4")
		outputNode := g.Output(
			fmt.Sprintf("reference-outputs/out-%s.mp4", sourceName),
			output.WithFormatMP4(),
		)

		// // Create a text overlay node
		// textOverlay := g.Filter(drawtext.DrawText(
		// 	drawtext.WithText("SpireAV Test"),
		// ))

		videoSrc := g.Filter(filter.New(
			sourceName,
			1,
			vidsrc.WithFrameRate(expr.Int(30)),
			vidsrc.WithDuration(expr.Int(10)),
		))
		videoSrc.Pipe(outputNode, 0)

		// Sin wave from 300 to 1000 Hz, period of 2 seconds
		// freq := expr.Add(expr.Int(300), expr.Mul(expr.Int(700), expr.Sin0to1("t")))
		freq := expr.Int(440)
		fmt.Println(freq.String())
		fmt.Println()

		audioSrc := g.Filter(audiosrc.SineSource(
			audiosrc.WithFrequency(freq),
			audiosrc.WithDuration(expr.Int(10)),
		))
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

		go func() {
			defer wg.Done()
			// Run the transcoding job
			if err := runner.Run(ctx); err != nil {
				fmt.Println(err.Error())
			}
		}()
	}
	wg.Wait()
}
