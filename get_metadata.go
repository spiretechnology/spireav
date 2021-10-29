package spireav

import (
	"context"
	"encoding/json"
	"os/exec"
)

func GetMetadata(
	ctx context.Context,
	filename string,
) (*Meta, error) {

	// Create the command
	cmd := exec.CommandContext(ctx, FfprobePath, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filename)

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
