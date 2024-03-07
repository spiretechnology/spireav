package audiosink

import (
	"github.com/spiretechnology/spireav/filter"
)

// https://ffmpeg.org/ffmpeg-filters.html#anullsink

func NullSink(opts ...filter.Option) filter.Filter {
	return filter.New("anullsink", 1, opts...)
}
