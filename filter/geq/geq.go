package geq

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#geq

// GenericEquation applies a generic equation to each pixel.
func GenericEquation(opts ...filter.Option) filter.Filter {
	return filter.New("geq", 1, opts...)
}

func WithLuma(luma expr.Expr) filter.Option {
	return filter.WithOption("luma", luma)
}

func WithChrominanceBlue(cb expr.Expr) filter.Option {
	return filter.WithOption("cb", cb)
}

func WithChrominanceRed(cr expr.Expr) filter.Option {
	return filter.WithOption("cr", cr)
}

func WithAlpha(alpha expr.Expr) filter.Option {
	return filter.WithOption("a", alpha)
}

func WithRed(red expr.Expr) filter.Option {
	return filter.WithOption("r", red)
}

func WithGreen(green expr.Expr) filter.Option {
	return filter.WithOption("g", green)
}

func WithBlue(blue expr.Expr) filter.Option {
	return filter.WithOption("b", blue)
}
