package transform

import (
	"fmt"
	"strconv"
	"strings"
)

// TimecodeOverlay is a transform that takes in a single video stream and draws live timecode text on top of it
type TimecodeOverlay struct {
	StartTimecode  string
	FrameRate      float64
	X              string
	Y              string
	FontColor      string
	FontFile       string
	Box            bool
	BoxColor       string
	BoxBorderWidth int
}

func (node *TimecodeOverlay) GetOutputsCount() int {
	return 1
}

func (node *TimecodeOverlay) getFilterOptions() map[string]string {
	opts := map[string]string{}
	if len(node.StartTimecode) == 0 {
		opts["timecode"] = "00\\:00\\:00\\:00"
	} else {
		opts["timecode"] = strings.ReplaceAll(node.StartTimecode, ":", "\\:")
	}
	if node.FrameRate < 0.0001 {
		opts["r"] = "23.976"
	} else {
		opts["r"] = strconv.FormatFloat(node.FrameRate, 'f', 3, 64)
	}
	if len(node.X) > 0 {
		opts["x"] = node.X
	}
	if len(node.Y) > 0 {
		opts["y"] = node.Y
	}
	if len(node.FontFile) > 0 {
		opts["fontfile"] = node.FontFile
	}
	if len(node.FontColor) > 0 {
		opts["fontcolor"] = node.FontColor
	}
	if node.Box {
		opts["box"] = "1"
		if len(node.BoxColor) > 0 {
			opts["boxcolor"] = node.BoxColor
		}
		if node.BoxBorderWidth > 0 {
			opts["boxborderw"] = fmt.Sprintf("%d", node.BoxBorderWidth)
		}
	}
	return opts
}

func (node *TimecodeOverlay) FilterString() string {
	return FormatTransform("drawtext", node.getFilterOptions())
}
