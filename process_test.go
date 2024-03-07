package spireav_test

import (
	"context"
	"testing"

	"github.com/spiretechnology/spireav"
	"github.com/stretchr/testify/require"
)

func overrideFfmpegPath(path string) func() {
	original := spireav.FfmpegPath
	spireav.FfmpegPath = path
	return func() {
		spireav.FfmpegPath = original
	}
}

func TestProcessErrors(t *testing.T) {
	t.Run("invalid args", func(t *testing.T) {
		runner := spireav.NewRunner(spireav.FfmpegArgsSlice{"several", "bad", "args"})
		err := runner.Run(context.Background())
		require.Error(t, err, "expected error when running process with invalid args")
		require.ErrorContains(t, err, "Invalid argument")
	})
	t.Run("ffmpeg not in path", func(t *testing.T) {
		defer overrideFfmpegPath("something-that-does-not-exist")()
		runner := spireav.NewRunner(spireav.FfmpegArgsSlice{"-i", "something"})
		err := runner.Run(context.Background())
		require.Error(t, err, "expected error when path is missing ffmpeg binary")
		require.ErrorContains(t, err, "starting ffmpeg process: exec: \"something-that-does-not-exist\": executable file not found")
	})
}
