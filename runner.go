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
	"syscall"
	"time"
)

type Runner interface {
	Run(ctx context.Context) error
}

type Option func(*implRunner)

func WithEstimatedLengthFrames(frames int) Option {
	return func(r *implRunner) {
		r.estimatedLengthFrames = frames
	}
}

func WithProgressCallback(callback func(Progress)) Option {
	return func(r *implRunner) {
		r.progressCallback = callback
	}
}

func WithSysProcAttr(attr *syscall.SysProcAttr) Option {
	return func(r *implRunner) {
		r.sysProcAttr = attr
	}
}

func NewRunner(args FfmpegArgs, opts ...Option) Runner {
	runner := &implRunner{
		args: args,
	}
	for _, opt := range opts {
		opt(runner)
	}
	return runner
}

type implRunner struct {
	args                  FfmpegArgs
	estimatedLengthFrames int
	progressCallback      func(Progress)
	sysProcAttr           *syscall.SysProcAttr
	stderrLines           [20]string
	stderrLineCursor      int
}

// GetCommandString is a utility function that gets the FFmpeg command string that is run by this process
func (p *implRunner) GetCommandString() (string, error) {
	args, err := p.args.FfmpegArgs()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s %s\n",
		FfmpegPath,
		strings.Join(args, " "),
	), nil
}

func (p *implRunner) stdErrSlice() []string {
	if p.stderrLineCursor == 0 {
		return []string{}
	} else if p.stderrLineCursor < len(p.stderrLines) {
		return p.stderrLines[:p.stderrLineCursor]
	} else if p.stderrLineCursor == len(p.stderrLines) {
		return p.stderrLines[:]
	} else {
		slice := make([]string, len(p.stderrLines))
		for i := 0; i < len(p.stderrLines); i++ {
			slice[i] = p.stderrLines[(p.stderrLineCursor+i)%len(p.stderrLines)]
		}
		return slice
	}
}

// Run executes the transcoding process and blocks the calling thread until the process is completed or failed.
// If a chanProgress is provided, it will receive progress updates during the process's execution. When the process
// completes, the progress channel will be closed automatically.
func (p *implRunner) Run(ctx context.Context) error {
	// Generate the FFmpeg arguments
	args, err := p.args.FfmpegArgs()
	if err != nil {
		return fmt.Errorf("generating ffmpeg args: %w", err)
	}

	// Create the command
	cmd := exec.CommandContext(ctx, FfmpegPath, args...)
	cmd.SysProcAttr = p.sysProcAttr

	// Get a readable pipe of the command stdout
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("acquiring stderr pipe: %w", err)
	}
	defer stderr.Close()

	// Parse the output into progress reports on the channel
	// Also make sure we collect all of stderr before returning any error
	var progressWaitGroup sync.WaitGroup
	progressWaitGroup.Add(1)
	go func() {
		defer progressWaitGroup.Done()
		p.reportFFmpegProgress(stderr)
	}()

	// Start running the command in the background
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("starting ffmpeg process: %w", err)
	}

	// Wait for the process to end
	if err := cmd.Wait(); err != nil {
		// Wait for the error log to finish
		progressWaitGroup.Wait()

		// If the context triggered the process to be killed, we want to see the context's error
		// instead of the process's error
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}
		return &FfmpegError{
			ExitCode: cmd.ProcessState.ExitCode(),
			Logs:     p.stdErrSlice(),
			child:    err,
		}
	}
	return nil
}

func (p *implRunner) reportFFmpegProgress(processOutput io.Reader) {
	// Calculate the start time
	startTime := time.Now()

	// Initialize the progress to empty
	progress := emptyProgress(startTime)
	progress.Estimate = &ProgressEstimate{
		Percent: 0,
	}
	if p.progressCallback != nil {
		p.progressCallback(*progress)
	}

	// Create a scanner for the standard output
	scanner := bufio.NewScanner(processOutput)
	scanner.Split(scanLinesWithCR)
	for scanner.Scan() {
		// Read a line from the input
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\r", "\n")
		lines := strings.Split(line, "\n")
		for _, line := range lines {
			if line != "" {
				p.stderrLines[p.stderrLineCursor%len(p.stderrLines)] = line
				p.stderrLineCursor++
			}

			// Read a new progress value
			if p.progressCallback != nil {
				newProgress := parseProgressLine(
					line,
					p.estimatedLengthFrames,
					startTime,
				)
				if newProgress != nil {
					progress = newProgress
					p.progressCallback(*progress)
				}
			}
		}
	}

	// Send the 100% progress
	if p.progressCallback != nil {
		progress.Elapsed = time.Since(startTime)
		progress.FPS = 0
		progress.Speed = 0
		progress.Done = true
		progress.Estimate = &ProgressEstimate{
			Percent:   1,
			Remaining: time.Duration(0),
		}
		p.progressCallback(*progress)
	}
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
