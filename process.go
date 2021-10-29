package spireav

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// Process represents a transcoding process that runs through FFmpeg
type Process struct {
	BinPath               string
	FfmpegArger           FfmpegArger
	EstimatedLengthFrames int
}

// getBinPath gets the path to the executable to use for FFmpeg.
func (p *Process) getBinPath() string {
	if len(p.BinPath) > 0 {
		return p.BinPath
	}
	return "ffmpeg"
}

// GetCommandString is a utility function that gets the FFmpeg command string that is run by this process
func (p *Process) GetCommandString() (string, error) {
	args, err := p.FfmpegArger.FfmpegArgs()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s %s\n",
		p.getBinPath(),
		strings.Join(args, " "),
	), nil
}

// Run executes the transcoding process and blocks the calling thread until the process is completed or failed.
// If a chanProgress is provided, it will receive progress updates during the process's execution. When the process
// completes, the progress channel will be closed automatically.
func (p *Process) Run(
	ctx context.Context,
	chanProgress chan<- Progress,
) error {

	// Generate the FFmpeg arguments
	args, err := p.FfmpegArger.FfmpegArgs()
	if err != nil {
		return err
	}

	// Create the command
	cmd := exec.CommandContext(ctx, p.getBinPath(), args...)

	// If progress needs to be reported
	if chanProgress != nil {

		// Get a readable pipe of the command stdout
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}
		defer stderr.Close()

		// Parse the output into progress reports on the channel
		go p.reportFFmpegProgress(
			chanProgress,
			stderr,
		)

	}

	// Start running the command in the background
	if err := cmd.Start(); err != nil {
		return err
	}

	// Wait for the process to end
	if err := cmd.Wait(); err != nil {

		// If the context triggered the process to be killed, we want to see the context's error
		// instead of the process's error
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}
		return err

	}
	return nil

}

// RunWithProgress executes the process and reports progress back to a progress handler function
func (p *Process) RunWithProgress(
	ctx context.Context,
	progressFunc func(Progress),
) error {

	// Create a channel for progress reporting
	chanProgress := make(chan Progress)

	// Wait for both the progress handler to finish, and the transcoding
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for progress := range chanProgress {
			progressFunc(progress)
		}
	}()

	// Produce the output
	err := p.Run(ctx, chanProgress)
	wg.Done()
	wg.Wait()

	// Return the error, if any
	return err

}

func (p *Process) reportFFmpegProgress(
	chanProgress chan<- Progress,
	processOutput io.Reader,
) {

	// Calculate the start time
	startTime := time.Now()

	// Initialize the progress to empty
	progress := emptyProgress(startTime)
	chanProgress <- *progress

	// Create a scanner for the standard output
	scanner := bufio.NewScanner(processOutput)
	scanner.Split(scanLinesWithCR)
	for scanner.Scan() {

		// Read a line from the input
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\r", "\n")

		// Read a new progress value
		newProgress := parseProgressLine(
			line,
			p.EstimatedLengthFrames,
			startTime,
		)
		if newProgress != nil {
			progress = newProgress
			chanProgress <- *progress
		}

	}

	// Close the channel
	close(chanProgress)

}

func scanLinesWithCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\r'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
