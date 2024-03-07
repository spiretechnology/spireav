package amerge

import (
	"github.com/spiretechnology/spireav/graph/filter"
	"github.com/spiretechnology/spireav/graph/filter/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#amerge

func AudioMerge(opts ...filter.Option) filter.Filter {
	return filter.New("amerge", 1, opts...)
}

func WithInputs(inputs int) filter.Option {
	return filter.WithOption("inputs", expr.Int(inputs))
}
