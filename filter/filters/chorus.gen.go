// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ChorusBuilder Add a chorus effect to the audio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#chorus
type ChorusBuilder interface {
	filter.Filter
	// InGain set input gain (from 0 to 1) (default 0.4).
	InGain(inGain float32) ChorusBuilder
	// OutGain set output gain (from 0 to 1) (default 0.4).
	OutGain(outGain float32) ChorusBuilder
	// Delays set delays.
	Delays(delays string) ChorusBuilder
	// Decays set decays.
	Decays(decays string) ChorusBuilder
	// Speeds set speeds.
	Speeds(speeds string) ChorusBuilder
	// Depths set depths.
	Depths(depths string) ChorusBuilder
}

// Chorus creates a new ChorusBuilder to configure the "chorus" filter.
func Chorus(opts ...filter.Option) ChorusBuilder {
	return &implChorusBuilder{
		f: filter.New("chorus", 1, opts...),
	}
}

type implChorusBuilder struct {
	f filter.Filter
}

func (chorusBuilder *implChorusBuilder) String() string {
	return chorusBuilder.f.String()
}

func (chorusBuilder *implChorusBuilder) Outputs() int {
	return chorusBuilder.f.Outputs()
}

func (chorusBuilder *implChorusBuilder) With(key string, value expr.Expr) filter.Filter {
	return chorusBuilder.withOption(key, value)
}

func (chorusBuilder *implChorusBuilder) withOption(key string, value expr.Expr) ChorusBuilder {
	bCopy := *chorusBuilder
	bCopy.f = chorusBuilder.f.With(key, value)
	return &bCopy
}

func (chorusBuilder *implChorusBuilder) InGain(inGain float32) ChorusBuilder {
	return chorusBuilder.withOption("in_gain", expr.Float(inGain))
}

func (chorusBuilder *implChorusBuilder) OutGain(outGain float32) ChorusBuilder {
	return chorusBuilder.withOption("out_gain", expr.Float(outGain))
}

func (chorusBuilder *implChorusBuilder) Delays(delays string) ChorusBuilder {
	return chorusBuilder.withOption("delays", expr.String(delays))
}

func (chorusBuilder *implChorusBuilder) Decays(decays string) ChorusBuilder {
	return chorusBuilder.withOption("decays", expr.String(decays))
}

func (chorusBuilder *implChorusBuilder) Speeds(speeds string) ChorusBuilder {
	return chorusBuilder.withOption("speeds", expr.String(speeds))
}

func (chorusBuilder *implChorusBuilder) Depths(depths string) ChorusBuilder {
	return chorusBuilder.withOption("depths", expr.String(depths))
}
