package framemux

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#select_002c-aselect

func Select(opts ...filter.Option) filter.Filter {
	return filter.New("select", 1, opts...)
}

func AudioSelect(outputs int, opts ...filter.Option) filter.Filter {
	return filter.New("aselect", 1, opts...)
}

func WithOutputs(outputs int) filter.Option {
	return filter.WithOption("outputs", expr.Int(outputs))
}

func WithExpr(e expr.Expr) filter.Option {
	return filter.WithOption("e", e)
}
