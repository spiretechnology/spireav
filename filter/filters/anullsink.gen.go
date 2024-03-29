// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AudioNullSinkBuilder corresponds to the "anullsink" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#anullsink
type AudioNullSinkBuilder interface {
	filter.Filter
}

// AudioNullSink creates a new AudioNullSinkBuilder to configure the "anullsink" filter.
func AudioNullSink(opts ...filter.Option) AudioNullSinkBuilder {
	return &implAudioNullSinkBuilder{
		f: filter.New("anullsink", 0, opts...),
	}
}

type implAudioNullSinkBuilder struct {
	f filter.Filter
}

func (b *implAudioNullSinkBuilder) String() string {
	return b.f.String()
}

func (b *implAudioNullSinkBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implAudioNullSinkBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implAudioNullSinkBuilder) withOption(key string, value expr.Expr) AudioNullSinkBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}
