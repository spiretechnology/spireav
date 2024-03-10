package metadata

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spiretechnology/spireav"
)

func GetMxfMetadata(ctx context.Context, filename string) (*Mxf2RawResult, error) {
	return GetMxfMetadataWithOptions(ctx, filename, MetadataOptions{})
}

// GetMxfMetadata gets the metadata for the given MXF file.
func GetMxfMetadataWithOptions(ctx context.Context, filename string, opts MetadataOptions) (*Mxf2RawResult, error) {
	// Make sure the path to mxf2raw is configured
	if spireav.Mxf2RawPath == "" {
		return nil, errors.New("mxf2raw path not configured")
	}

	// Create the command
	cmd := exec.CommandContext(ctx, spireav.Mxf2RawPath, "--info-format", "xml", filename)
	if opts.SysProcAttr != nil {
		cmd.SysProcAttr = opts.SysProcAttr
	}

	// Create the output buffer to capture stdout
	var stdoutBuffer, stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer

	// Run the command
	if err := cmd.Run(); err != nil {
		errStr := strings.TrimSpace(stderrBuffer.String())
		return nil, fmt.Errorf("running mxf2raw: %s: %s", err, errStr)
	}

	// Get the bytes from the output
	stdoutBytes := stdoutBuffer.Bytes()

	// Parse the XML result
	var result Mxf2RawResult
	if err := xml.Unmarshal(stdoutBytes, &result); err != nil {
		return nil, fmt.Errorf("parsing mxf2raw XML result: %s", err)
	}
	return &result, nil
}
