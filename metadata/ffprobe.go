package metadata

import (
	"context"
	"encoding/json"
	"os/exec"
	"syscall"

	"github.com/spiretechnology/spireav"
)

type MetadataOptions struct {
	SysProcAttr *syscall.SysProcAttr
}

func GetMetadata(ctx context.Context, filename string) (*Meta, error) {
	return GetMetadataWithOptions(ctx, filename, MetadataOptions{})
}

func GetMetadataWithOptions(
	ctx context.Context,
	filename string,
	opts MetadataOptions,
) (*Meta, error) {
	// Create the command
	cmd := exec.CommandContext(ctx, spireav.FfprobePath, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filename)
	if opts.SysProcAttr != nil {
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
	return &meta, nil
}
