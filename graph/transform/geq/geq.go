package geq

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#geq

// GenericEquation applies a generic equation to each pixel.
func GenericEquation(opts ...transform.Option) transform.Transform {
	return transform.New("geq", 1, opts...)
}

func WithLuma(luma expr.Expr) transform.Option {
	return transform.WithOption("luma", luma)
}

func WithChrominanceBlue(cb expr.Expr) transform.Option {
	return transform.WithOption("cb", cb)
}

func WithChrominanceRed(cr expr.Expr) transform.Option {
	return transform.WithOption("cr", cr)
}

func WithAlpha(alpha expr.Expr) transform.Option {
	return transform.WithOption("a", alpha)
}

func WithRed(red expr.Expr) transform.Option {
	return transform.WithOption("r", red)
}

func WithGreen(green expr.Expr) transform.Option {
	return transform.WithOption("g", green)
}

func WithBlue(blue expr.Expr) transform.Option {
	return transform.WithOption("b", blue)
}
