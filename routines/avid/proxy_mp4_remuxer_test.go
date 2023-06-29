package avid_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/routines/avid"
	"github.com/stretchr/testify/assert"
)

func TestAvidRemux(t *testing.T) {

	// Create the proxy remuxer
	remux, err := avid.NewProxyMP4Remuxer([]string{
		testMedia("avid/A01.DE767F0A_179DE179DE147A.mxf"),
		testMedia("avid/A02.DE767F0B_179DE179DE157A.mxf"),
		testMedia("avid/V01.DE767F09_179DE179DE134V.mxf"),
	})
	remux.OverlayTimecode = true
	if err != nil {
		panic(err)
	}

	// Create the output directory
	var tempDir string
	if envTempDir, ok := os.LookupEnv("TEMP_DIR"); ok {
		tempDir = filepath.Join(envTempDir, "spireav-remux-test")
	} else {
		tempDir = filepath.Join(os.TempDir(), "spireav-remux-test")
	}
	err = os.MkdirAll(tempDir, 0700)
	assert.NoError(t, err, "create temp output dir")
	defer os.RemoveAll(tempDir)

	// Create the process
	p, err := remux.GenerateProc(tempDir)
	assert.NoError(t, err, "generate remux process")

	// Create a progress handler function
	var receivedProgress0 bool
	var receivedProgress100 bool
	var receivedProgressDone bool
	progressFunc := func(progress spireav.Progress) {
		fmt.Printf("Progress: %+v\n", progress)
		if progress.Estimate != nil {
			fmt.Printf("EST: %+v\n", progress.Estimate)
			if progress.Estimate.Percent < 0.0001 {
				receivedProgress0 = true
			}
			if progress.Estimate.Percent >= 0.9999 {
				receivedProgress100 = true
			}
		}
		if progress.Done {
			receivedProgressDone = true
		}
	}

	// Run the transcoding job
	err = p.RunWithProgress(context.Background(), progressFunc)
	assert.NoError(t, err, "running remux process")

	// Assert that all the progress steps were hit
	assert.True(t, receivedProgress0, "received progress value 0 pct")
	assert.True(t, receivedProgress100, "received progress value 100 pct")
	assert.True(t, receivedProgressDone, "received progress done")

	// Assert that all the outputs exist and have correct metadata
	// ...

}
