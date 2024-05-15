package expr_test

import (
	"testing"

	"github.com/spiretechnology/spireav/filter/expr"
	"github.com/stretchr/testify/require"
)

func TestExpressions(t *testing.T) {
	t.Run("min and max", func(t *testing.T) {
		min1 := expr.Min(expr.Int(100))
		max1 := expr.Max(expr.Int(100))
		require.Equal(t, "100", min1.String())
		require.Equal(t, "100", max1.String())

		min2 := expr.Min(expr.Int(5), expr.Int(10))
		max2 := expr.Max(expr.Int(5), expr.Int(10))
		require.Equal(t, "min(5,10)", min2.String())
		require.Equal(t, "max(5,10)", max2.String())

		min3 := expr.Min(expr.Int(1), expr.Int(2), expr.Int(3))
		max3 := expr.Max(expr.Int(1), expr.Int(2), expr.Int(3))
		require.Equal(t, "min(1,min(2,3))", min3.String())
		require.Equal(t, "max(1,max(2,3))", max3.String())

		min4 := expr.Min(expr.Int(1), expr.Int(2), expr.Int(3), expr.Int(10))
		max4 := expr.Max(expr.Int(1), expr.Int(2), expr.Int(3), expr.Int(10))
		require.Equal(t, "min(1,min(2,min(3,10)))", min4.String())
		require.Equal(t, "max(1,max(2,max(3,10)))", max4.String())
	})
}
