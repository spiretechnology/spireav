package drawtext

import (
	"github.com/spiretechnology/spireav/graph/filter"
	"github.com/spiretechnology/spireav/graph/filter/expr"
)

func DrawText(opts ...filter.Option) filter.Filter {
	return filter.New("drawtext", 1, opts...)
}

func WithText(text string) filter.Option {
	return filter.WithOption("text", expr.String(text))
}

func WithX(x expr.Expr) filter.Option {
	return filter.WithOption("x", x)
}

func WithY(y expr.Expr) filter.Option {
	return filter.WithOption("y", y)
}

func WithFontColor(color string) filter.Option {
	return filter.WithOption("fontcolor", expr.String(color))
}

func WithFontFile(file string) filter.Option {
	return filter.WithOption("fontfile", expr.String(file))
}

func WithFontSize(size expr.Expr) filter.Option {
	return filter.WithOption("fontsize", size)
}

func WithTimecode(tc, framerate string) filter.Option {
	if len(framerate) == 0 {
		framerate = "23.976"
	}
	return filter.WithMulti(
		filter.WithOption("timecode", expr.String(tc)),
		filter.WithOption("r", expr.String(framerate)),
	)
}

func WithBox(color string) filter.Option {
	return filter.WithMulti(
		filter.WithOption("box", expr.String("1")),
		filter.WithOption("boxcolor", expr.String(color)),
	)
}
