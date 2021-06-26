package meta

import (
	"encoding/json"
	"os/exec"
)

func GetMetadata(
	filename string,
	ffprobePath string,
) (*Meta, error) {

	// If there is no probe path, use the default
	if len(ffprobePath) == 0 {
		ffprobePath = "ffprobe"
	}

	// Create the command
	cmd := exec.Command(ffprobePath, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filename)

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
