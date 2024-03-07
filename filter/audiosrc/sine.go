package audiosrc

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

func SineSource(opts ...filter.Option) filter.Filter {
	return filter.New("sine", 1, opts...)
}

func WithFrequency(frequency expr.Expr) filter.Option {
	return filter.WithOption("f", frequency)
}
