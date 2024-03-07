package scale

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

const (
	EvalInit  = "init"
	EvalFrame = "frame"
)

func Scale(opts ...transform.Option) transform.Transform {
	return transform.New("scale", 1, opts...)
}

func WithWidth(width expr.Expr) transform.Option {
	return transform.WithOption("w", width)
}

func WithWidthPreserveAspectRatio(width expr.Expr) transform.Option {
	return transform.WithMulti(
		WithWidth(width),
		WithHeight(expr.Int(-1)),
	)
}

func WithHeight(height expr.Expr) transform.Option {
	return transform.WithOption("h", height)
}

func WithHeightPreserveAspectRatio(height expr.Expr) transform.Option {
	return transform.WithMulti(
		WithHeight(height),
		WithWidth(expr.Int(-1)),
	)
}

func WithEval(eval string) transform.Option {
	return transform.WithOption("eval", expr.String(eval))
}

func WithInterlacingMode(mode string) transform.Option {
	return transform.WithOption("interl", expr.String(mode))
}

func WithForceOriginalAspectRatio(mode string) transform.Option {
	return transform.WithOption("force_original_aspect_ratio", expr.String(mode))
}

func WithPreserveAspectRatioDecrease() transform.Option {
	return WithForceOriginalAspectRatio("decrease")
}

func WithPreserveAspectRatioIncrease() transform.Option {
	return WithForceOriginalAspectRatio("increase")
}

func WithIgnoreAspectRatio() transform.Option {
	return WithForceOriginalAspectRatio("disable")
}

func WithForceDivisibleBy(divisor int) transform.Option {
	return transform.WithOption("force_divisible_by", expr.Int(divisor))
}
