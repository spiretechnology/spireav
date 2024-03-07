package overlay

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

const (
	EvalInit  = "init"
	EvalFrame = "frame"

	AlphaStraight      = "straight"
	AlphaPremultiplied = "premultipied"
)

// https://ffmpeg.org/ffmpeg-filters.html#overlay

func Overlay(opts ...transform.Option) transform.Transform {
	return transform.New("overlay", 1, opts...)
}

func WithX(x expr.Expr) transform.Option {
	return transform.WithOption("x", x)
}

func WithY(y expr.Expr) transform.Option {
	return transform.WithOption("y", y)
}

func WithEval(eval string) transform.Option {
	return transform.WithOption("eval", expr.String(eval))
}

func WithAlpha(alpha string) transform.Option {
	return transform.WithOption("alpha", expr.String(alpha))
}

func WithEnable(enable expr.Expr) transform.Option {
	return transform.WithOption("enable", enable)
}
