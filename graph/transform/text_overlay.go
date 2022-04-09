package transform

import "fmt"

// TextOverlay is a transform that takes in a single video stream and draws static text on top of it
type TextOverlay struct {
	Text           string
	X              string
	Y              string
	FontColor      string
	FontFile       string
	FontSize       int
	Box            bool
	BoxColor       string
	BoxBorderWidth int
}

func (node *TextOverlay) OutputsCount() int {
	return 1
}

func (node *TextOverlay) getFilterOptions() map[string]string {
	opts := map[string]string{
		"text": node.Text,
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
	if node.FontSize > 0 {
		opts["fontsize"] = fmt.Sprintf("%d", node.FontSize)
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

func (node *TextOverlay) FilterString() string {
	return FormatTransform("drawtext", node.getFilterOptions())
}
