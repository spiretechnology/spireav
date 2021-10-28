package proc

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

	"github.com/spiretechnology/spireav/graph"
)

type Proc struct {
	BinPath               string
	Graph                 *graph.Graph
	EstimatedLengthFrames int
}

// generateCmdArgs generates a slice of command arguments to pass into FFmpeg
func (p *Proc) generateCmdArgs() []string {

	// Get all the inputs and outputs
	inputs := p.Graph.GetInputs()
	outputs := p.Graph.GetOutputs()

	// A slice to hold all the arguments
	args := []string{}

	// Loop through all the inputs
	for _, in := range inputs {
		args = append(args, "-i", in.GetFilename())
	}

	// Add the -y argument to skip interactive prompts
	args = append(args, "-y")

	// Add the filters string
	args = append(args, "-filter_complex", p.Graph.FilterString())

	// Loop through all the outputs
	for _, out := range outputs {

		// Get the mappings for the output
		for _, mapping := range p.Graph.GetOutputMappings(out) {
			args = append(args, "-map", mapping)
		}

		// Add all the other options, and end with the filename
		args = append(args, out.GetOptions()...)
		args = append(args, out.GetFilename())

	}

	// Return the slice of arguments
	return args

}

func (p *Proc) getBinPath() string {
	if len(p.BinPath) > 0 {
		return p.BinPath
	}
	return "ffmpeg"
}

func (p *Proc) GetCommandString() string {
	return fmt.Sprintf(
		"%s %s\n",
		p.getBinPath(),
		strings.Join(p.generateCmdArgs(), " "),
	)
}

func (p *Proc) Run(chanProgress chan<- Progress) error {
	return p.RunContext(
		context.Background(),
		chanProgress,
	)
}

func (p *Proc) RunContext(
	ctx context.Context,
	chanProgress chan<- Progress,
) error {

	// Generate the FFmpeg arguments
	args := p.generateCmdArgs()

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
		return err
	}
	return nil

}

func (p *Proc) RunWithProgress(progressFunc func(Progress)) error {
	return p.RunWithProgressContext(
		context.Background(),
		progressFunc,
	)
}

func (p *Proc) RunWithProgressContext(
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
	err := p.Run(chanProgress)
	wg.Done()
	wg.Wait()

	// Return the error, if any
	return err

}

func (p *Proc) reportFFmpegProgress(chanProgress chan<- Progress, processOutput io.Reader) {

	// Calculate the start time
	startTime := time.Now()

	// Initialize the progress to empty
	progress := EmptyProgress(startTime)
	chanProgress <- *progress

	// Create a scanner for the standard output
	scanner := bufio.NewScanner(processOutput)
	scanner.Split(ScanLinesWithCR)
	for scanner.Scan() {

		// Read a line from the input
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\r", "\n")

		// Read a new progress value
		newProgress := ParseProgressLine(
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

func ScanLinesWithCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
