package drawtext

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

func DrawText(opts ...transform.Option) transform.Transform {
	return transform.New("drawtext", 1, opts...)
}

func WithText(text string) transform.Option {
	return transform.WithOption("text", expr.String(text))
}

func WithX(x expr.Expr) transform.Option {
	return transform.WithOption("x", x)
}

func WithY(y expr.Expr) transform.Option {
	return transform.WithOption("y", y)
}

func WithFontColor(color string) transform.Option {
	return transform.WithOption("fontcolor", expr.String(color))
}

func WithFontFile(file string) transform.Option {
	return transform.WithOption("fontfile", expr.String(file))
}

func WithFontSize(size expr.Expr) transform.Option {
	return transform.WithOption("fontsize", size)
}

func WithTimecode(tc, framerate string) transform.Option {
	if len(framerate) == 0 {
		framerate = "23.976"
	}
	return transform.WithMulti(
		transform.WithOption("timecode", expr.String(tc)),
		transform.WithOption("r", expr.String(framerate)),
	)
}

func WithBox(color string) transform.Option {
	return transform.WithMulti(
		transform.WithOption("box", expr.String("1")),
		transform.WithOption("boxcolor", expr.String(color)),
	)
}
