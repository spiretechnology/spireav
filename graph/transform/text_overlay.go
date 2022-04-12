package transform

import (
	"fmt"

	"github.com/spiretechnology/spireav/graph/transform/expr"
)

func NewTextOverlay(opts ...Option) Transform {
	return New("drawtext", 1, opts...)
}

func WithText(text string) Option {
	return withOption("text", expr.String(text))
}

func WithX(x expr.Expr) Option {
	return withOption("x", x)
}

func WithY(y expr.Expr) Option {
	return withOption("y", y)
}

func WithFontColor(color string) Option {
	return withOption("fontcolor", expr.String(color))
}

func WithFontFile(file string) Option {
	return withOption("fontfile", expr.String(file))
}

func WithFontSize(size expr.Expr) Option {
	return withOption("fontsize", size)
}

func WithTimecode(tc, framerate string) Option {
	if len(framerate) == 0 {
		framerate = "23.976"
	}
	return withMulti(
		withOption("timecode", expr.String(tc)),
		withOption("r", expr.String(framerate)),
	)
}

func WithBox(color string) Option {
	return withMulti(
		withOption("box", expr.String("1")),
		withOption("boxcolor", expr.String(color)),
	)
}

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
