package scale

import (
	"github.com/spiretechnology/spireav/graph/filter"
	"github.com/spiretechnology/spireav/graph/filter/expr"
)

const (
	EvalInit  = "init"
	EvalFrame = "frame"
)

func Scale(opts ...filter.Option) filter.Filter {
	return filter.New("scale", 1, opts...)
}

func WithWidth(width expr.Expr) filter.Option {
	return filter.WithOption("w", width)
}

func WithWidthPreserveAspectRatio(width expr.Expr) filter.Option {
	return filter.WithMulti(
		WithWidth(width),
		WithHeight(expr.Int(-1)),
	)
}

func WithHeight(height expr.Expr) filter.Option {
	return filter.WithOption("h", height)
}

func WithHeightPreserveAspectRatio(height expr.Expr) filter.Option {
	return filter.WithMulti(
		WithHeight(height),
		WithWidth(expr.Int(-1)),
	)
}

func WithEval(eval string) filter.Option {
	return filter.WithOption("eval", expr.String(eval))
}

func WithInterlacingMode(mode string) filter.Option {
	return filter.WithOption("interl", expr.String(mode))
}

func WithForceOriginalAspectRatio(mode string) filter.Option {
	return filter.WithOption("force_original_aspect_ratio", expr.String(mode))
}

func WithPreserveAspectRatioDecrease() filter.Option {
	return WithForceOriginalAspectRatio("decrease")
}

func WithPreserveAspectRatioIncrease() filter.Option {
	return WithForceOriginalAspectRatio("increase")
}

func WithIgnoreAspectRatio() filter.Option {
	return WithForceOriginalAspectRatio("disable")
}

func WithForceDivisibleBy(divisor int) filter.Option {
	return filter.WithOption("force_divisible_by", expr.Int(divisor))
}
