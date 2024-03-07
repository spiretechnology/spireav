package spireav

import (
	"fmt"
	"strings"
)

type FfmpegError struct {
	ExitCode int
	Logs     []string
	child    error
}

func (e *FfmpegError) Error() string {
	if len(e.Logs) > 0 {
		return fmt.Sprintf("ffmpeg exit code %d: %s", e.ExitCode, strings.Join(e.Logs, "\n"))
	} else {
		return fmt.Sprintf("ffmpeg exit code %d", e.ExitCode)
	}
}

func (e *FfmpegError) Unwrap() error {
	return e.child
}
