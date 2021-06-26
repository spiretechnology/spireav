package main

import (
	"fmt"
	"os"

	"github.com/spiretechnology/spireav/proc"
	"github.com/spiretechnology/spireav/routines/avid"
)

func main() {

	// Create the proxy remuxer
	remux, err := avid.NewProxyMP4Remuxer([]string{
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA01.D97B84568C2266A.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA02.D97B84568C2273A.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA03.D97B84568C227FA.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA04.D97B84568C228AA.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA05.D97B84568C2295A.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA06.D97B84568C22A1A.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA07.D97B84568C22ACA.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBA08.D97B84568C22B7A.mxf",
		"/Users/conner/Desktop/Sample Transcodes/SS0815KB.01/SS0815KBV01.D97B84568C1FF8V.mxf",
	})
	if err != nil {
		panic(err)
	}

	// Create the output directory
	outDir := "/Users/conner/Desktop/avid-proxy"
	if err := os.MkdirAll(outDir, 0700); err != nil {
		panic(err)
	}

	// Create the process
	p, err := remux.GenerateProc(outDir)
	if err != nil {
		panic(err)
	}

	// Print the FFmpeg command string
	fmt.Println(p.GetCommandString())

	// Create a progress handler function
	progressFunc := func(progress proc.Progress) {
		// fmt.Printf("P: %+v\n", progress)
		fmt.Printf("E: %+v\n", progress.Estimate)
	}

	// Run the transcoding job
	if err := p.RunWithProgress(progressFunc); err != nil {
		fmt.Println(err.Error())
	}

}
