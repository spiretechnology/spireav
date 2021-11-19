package spireav

import (
	"context"
	"encoding/json"
	"os/exec"
	"syscall"
)

type MetadataOptions struct {
	SysProcAttr *syscall.SysProcAttr
}

func GetMetadata(ctx context.Context, filename string) (*Meta, error) {
	return GetMetadataWithOptions(ctx, filename, nil)
}

func GetMetadataWithOptions(
	ctx context.Context,
	filename string,
	opts *MetadataOptions,
) (*Meta, error) {

	// Create the command
	cmd := exec.CommandContext(ctx, FfprobePath, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filename)
	if opts != nil && opts.SysProcAttr != nil {
		cmd.SysProcAttr = opts.SysProcAttr
	}

	// Run the command
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON output
	var meta Meta
	if err := json.Unmarshal(output, &meta); err != nil {
		return nil, err
	}

	// Return the metadata
	return &meta, nil

}
