package vidsink

import (
	"github.com/spiretechnology/spireav/filter"
)

// https://ffmpeg.org/ffmpeg-filters.html#nullsink

func NullSink(opts ...filter.Option) filter.Filter {
	return filter.New("nullsink", 1, opts...)
}
