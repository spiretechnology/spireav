// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AudioEvalSourceBuilder corresponds to the "aevalsrc" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#aevalsrc
type AudioEvalSourceBuilder interface {
	filter.Filter
	// Duration sets the "d" option on the filter.
	Duration(duration string) AudioEvalSourceBuilder
	// SampleRate sets the "s" option on the filter.
	SampleRate(sampleRate expr.Expr) AudioEvalSourceBuilder
	// SampleRateInt sets the "s" option on the filter.
	SampleRateInt(sampleRate int) AudioEvalSourceBuilder
	ChannelExpressions(exps ...expr.Expr) AudioEvalSourceBuilder
	ChannelLayouts(layouts ...string) AudioEvalSourceBuilder
}

// AudioEvalSource creates a new AudioEvalSourceBuilder to configure the "aevalsrc" filter.
func AudioEvalSource(opts ...filter.Option) AudioEvalSourceBuilder {
	return &implAudioEvalSourceBuilder{
		f: filter.New("aevalsrc", 1, opts...),
	}
}

type implAudioEvalSourceBuilder struct {
	f filter.Filter
}

func (b *implAudioEvalSourceBuilder) String() string {
	return b.f.String()
}

func (b *implAudioEvalSourceBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implAudioEvalSourceBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implAudioEvalSourceBuilder) withOption(key string, value expr.Expr) AudioEvalSourceBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implAudioEvalSourceBuilder) Duration(duration string) AudioEvalSourceBuilder {
	return b.withOption("d", expr.String(duration))
}

func (b *implAudioEvalSourceBuilder) SampleRate(sampleRate expr.Expr) AudioEvalSourceBuilder {
	return b.withOption("s", sampleRate)
}

func (b *implAudioEvalSourceBuilder) SampleRateInt(sampleRate int) AudioEvalSourceBuilder {
	return b.withOption("s", expr.Int(sampleRate))
}