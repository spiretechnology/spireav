package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

func Opacity(opacity float64) filter.Filter {
	return GenericEquation().
		Red(expr.Expr("r(X,Y)")).
		Alpha(expr.Mul(expr.Expr("alpha(X,Y)"), expr.Float(opacity)))
}
