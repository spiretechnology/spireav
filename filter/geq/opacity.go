package geq

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

func Opacity(opacity float64) filter.Filter {
	return GenericEquation(
		WithRed(expr.Expr("r(X,Y)")),
		WithAlpha(expr.Mul(expr.Expr("alpha(X,Y)"), expr.Float(opacity))),
	)
}
