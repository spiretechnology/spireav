package audiosrc

import (
	"strings"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

func EvalSource(opts ...filter.Option) filter.Filter {
	return filter.New("aevalsrc", 1, opts...)
}

func WithExpressions(exps ...expr.Expr) filter.Option {
	var expsStr []string
	for _, exp := range exps {
		expsStr = append(expsStr, exp.String())
	}
	return filter.WithOption("exprs", expr.Expr(strings.Join(expsStr, "|")))
}

func WithChannelLayouts(layouts ...string) filter.Option {
	return filter.WithOption("c", expr.String(strings.Join(layouts, "|")))
}

func WithDuration(duration expr.Expr) filter.Option {
	return filter.WithOption("d", duration)
}

func WithSampleRate(sampleRate expr.Expr) filter.Option {
	return filter.WithOption("s", sampleRate)
}
