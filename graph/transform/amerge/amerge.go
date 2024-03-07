package amerge

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#amerge

func AudioMerge(opts ...transform.Option) transform.Transform {
	return transform.New("amerge", 1, opts...)
}

func WithInputs(inputs int) transform.Option {
	return transform.WithOption("inputs", expr.Int(inputs))
}
