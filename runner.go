package spireav

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Progress struct {
	Frame      int64
	FPS        float64
	Bitrate    int64
	TotalSize  int64
	OutTime    string
	DupFrames  int64
	DropFrames int64
	Speed      float64
	Done       bool
	Started    time.Time
	Elapsed    time.Duration
}

type Runner interface {
	Run(ctx context.Context) error
}

type Option func(*implRunner)

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

func NewRunner(runnable Runnable, opts ...Option) Runner {
	runner := &implRunner{
		runnable: runnable,
	}
	for _, opt := range opts {
		opt(runner)
	}
	return runner
}

type implRunner struct {
	runnable         Runnable
	progressCallback func(Progress)
	sysProcAttr      *syscall.SysProcAttr
}

// GetCommandString is a utility function that gets the FFmpeg command string that is run by this process
func (p *implRunner) GetCommandString() (string, error) {
	args, err := p.runnable.RunnableArgs()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s %s\n",
		FfmpegPath,
		strings.Join(args, " "),
	), nil
}

// Run executes the transcoding process and blocks the calling thread until the process is completed or failed.
// If a chanProgress is provided, it will receive progress updates during the process's execution. When the process
// completes, the progress channel will be closed automatically.
func (p *implRunner) Run(ctx context.Context) error {
	// Generate the FFmpeg arguments
	args, err := p.runnable.RunnableArgs()
	if err != nil {
		return fmt.Errorf("generating ffmpeg args: %w", err)
	}

	// Create the command
	cmd := exec.CommandContext(ctx, FfmpegPath, args...)
	cmd.SysProcAttr = p.sysProcAttr

	// Stderr will contain error messages ONLY
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("acquiring stderr pipe: %w", err)
	}
	defer stderr.Close()

	// Stdout will contain a streaming progress report
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("acquiring stdout pipe: %w", err)
	}
	defer stdout.Close()

	// Capture and buffer stderr
	var stderrBuf []byte

	// Parse the output into progress reports on the channel
	// Also make sure we collect all of stderr before returning any error
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if p.progressCallback != nil {
			p.reportFFmpegProgress(stdout)
		} else {
			io.Copy(io.Discard, stdout)
		}
	}()
	go func() {
		defer wg.Done()
		stderrBuf, _ = io.ReadAll(stderr)
	}()

	// Start running the command in the background
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("starting ffmpeg process: %w", err)
	}

	// Wait for the process to end
	if err := cmd.Wait(); err != nil {
		// Wait for the error log to finish
		wg.Wait()

		// If the context triggered the process to be killed, we want to see the context's error
		// instead of the process's error
		if ctx != nil && ctx.Err() != nil {
			return ctx.Err()
		}
		return &FfmpegError{
			ExitCode: cmd.ProcessState.ExitCode(),
			Logs:     strings.Split(strings.TrimSpace(string(stderrBuf)), "\n"),
			child:    err,
		}
	}
	return nil
}

func (p *implRunner) reportFFmpegProgress(processOutput io.Reader) {
	// Initialize the progress to empty
	var progress Progress
	progress.Started = time.Now()

	// Create a scanner for the standard output
	scanner := bufio.NewScanner(processOutput)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// Read a line from the input
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		// Get the key and value
		key, value := parts[0], parts[1]
		switch key {
		case "frame":
			progress.Frame, _ = strconv.ParseInt(value, 10, 64)
		case "fps":
			progress.FPS, _ = strconv.ParseFloat(value, 64)
		case "bitrate":
			kbps, _ := strconv.ParseFloat(strings.TrimSuffix(value, "kbits/s"), 64)
			progress.Bitrate = int64(kbps * 1000)
		case "total_size":
			progress.TotalSize, _ = strconv.ParseInt(value, 10, 64)
		case "out_time":
			progress.OutTime = value
		case "dup_frames":
			progress.DupFrames, _ = strconv.ParseInt(value, 10, 64)
		case "drop_frames":
			progress.DropFrames, _ = strconv.ParseInt(value, 10, 64)
		case "speed":
			progress.Speed, _ = strconv.ParseFloat(strings.TrimSuffix(value, "x"), 64)
		case "progress":
			progress.Done = (value == "end")
		}

		// If the key is "progress", send the current progress to the handler
		if key == "progress" {
			progress.Elapsed = time.Since(progress.Started)
			p.progressCallback(progress)
		}
	}
}
