// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AformatBuilder Convert the input audio to one of the specified formats.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#aformat
type AformatBuilder interface {
	filter.Filter
	// SampleFmts A '|'-separated list of sample formats..
	SampleFmts(sampleFmts string) AformatBuilder
	// F A '|'-separated list of sample formats..
	F(f string) AformatBuilder
	// SampleRates A '|'-separated list of sample rates..
	SampleRates(sampleRates string) AformatBuilder
	// R A '|'-separated list of sample rates..
	R(r string) AformatBuilder
	// ChannelLayouts A '|'-separated list of channel layouts..
	ChannelLayouts(channelLayouts string) AformatBuilder
	// Cl A '|'-separated list of channel layouts..
	Cl(cl string) AformatBuilder
}

// Aformat creates a new AformatBuilder to configure the "aformat" filter.
func Aformat(opts ...filter.Option) AformatBuilder {
	return &implAformatBuilder{
		f: filter.New("aformat", 1, opts...),
	}
}

type implAformatBuilder struct {
	f filter.Filter
}

func (aformatBuilder *implAformatBuilder) String() string {
	return aformatBuilder.f.String()
}

func (aformatBuilder *implAformatBuilder) Outputs() int {
	return aformatBuilder.f.Outputs()
}

func (aformatBuilder *implAformatBuilder) With(key string, value expr.Expr) filter.Filter {
	return aformatBuilder.withOption(key, value)
}

func (aformatBuilder *implAformatBuilder) withOption(key string, value expr.Expr) AformatBuilder {
	bCopy := *aformatBuilder
	bCopy.f = aformatBuilder.f.With(key, value)
	return &bCopy
}

func (aformatBuilder *implAformatBuilder) SampleFmts(sampleFmts string) AformatBuilder {
	return aformatBuilder.withOption("sample_fmts", expr.String(sampleFmts))
}

func (aformatBuilder *implAformatBuilder) F(f string) AformatBuilder {
	return aformatBuilder.withOption("f", expr.String(f))
}

func (aformatBuilder *implAformatBuilder) SampleRates(sampleRates string) AformatBuilder {
	return aformatBuilder.withOption("sample_rates", expr.String(sampleRates))
}

func (aformatBuilder *implAformatBuilder) R(r string) AformatBuilder {
	return aformatBuilder.withOption("r", expr.String(r))
}

func (aformatBuilder *implAformatBuilder) ChannelLayouts(channelLayouts string) AformatBuilder {
	return aformatBuilder.withOption("channel_layouts", expr.String(channelLayouts))
}

func (aformatBuilder *implAformatBuilder) Cl(cl string) AformatBuilder {
	return aformatBuilder.withOption("cl", expr.String(cl))
}
