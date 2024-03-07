package overlay

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

const (
	EvalInit  = "init"
	EvalFrame = "frame"

	AlphaStraight      = "straight"
	AlphaPremultiplied = "premultipied"
)

// https://ffmpeg.org/ffmpeg-filters.html#overlay

func Overlay(opts ...filter.Option) filter.Filter {
	return filter.New("overlay", 1, opts...)
}

func WithX(x expr.Expr) filter.Option {
	return filter.WithOption("x", x)
}

func WithY(y expr.Expr) filter.Option {
	return filter.WithOption("y", y)
}

func WithEval(eval string) filter.Option {
	return filter.WithOption("eval", expr.String(eval))
}

func WithAlpha(alpha string) filter.Option {
	return filter.WithOption("alpha", expr.String(alpha))
}

func WithEnable(enable expr.Expr) filter.Option {
	return filter.WithOption("enable", enable)
}
