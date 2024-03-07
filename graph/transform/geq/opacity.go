package geq

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

func Opacity(opacity float64) transform.Transform {
	return GenericEquation(
		WithRed(expr.Expr("r(X,Y)")),
		WithAlpha(expr.Mul(expr.Expr("alpha(X,Y)"), expr.Float(opacity))),
	)
}
