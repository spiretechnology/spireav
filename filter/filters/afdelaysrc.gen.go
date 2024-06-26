// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AfdelaysrcBuilder Generate a Fractional delay FIR coefficients.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#afdelaysrc
type AfdelaysrcBuilder interface {
	filter.Filter
	// Delay set fractional delay (from 0 to 32767) (default 0).
	Delay(delay float64) AfdelaysrcBuilder
	// D set fractional delay (from 0 to 32767) (default 0).
	D(d float64) AfdelaysrcBuilder
	// SampleRate set sample rate (from 1 to INT_MAX) (default 44100).
	SampleRate(sampleRate int) AfdelaysrcBuilder
	// R set sample rate (from 1 to INT_MAX) (default 44100).
	R(r int) AfdelaysrcBuilder
	// NbSamples set the number of samples per requested frame (from 1 to INT_MAX) (default 1024).
	NbSamples(nbSamples int) AfdelaysrcBuilder
	// N set the number of samples per requested frame (from 1 to INT_MAX) (default 1024).
	N(n int) AfdelaysrcBuilder
	// Taps set number of taps for delay filter (from 0 to 32768) (default 0).
	Taps(taps int) AfdelaysrcBuilder
	// T set number of taps for delay filter (from 0 to 32768) (default 0).
	T(t int) AfdelaysrcBuilder
	// ChannelLayout set channel layout (default "stereo").
	ChannelLayout(channelLayout expr.ChannelLayout) AfdelaysrcBuilder
	// C set channel layout (default "stereo").
	C(c expr.ChannelLayout) AfdelaysrcBuilder
}

// Afdelaysrc creates a new AfdelaysrcBuilder to configure the "afdelaysrc" filter.
func Afdelaysrc(opts ...filter.Option) AfdelaysrcBuilder {
	return &implAfdelaysrcBuilder{
		f: filter.New("afdelaysrc", 1, opts...),
	}
}

type implAfdelaysrcBuilder struct {
	f filter.Filter
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) String() string {
	return afdelaysrcBuilder.f.String()
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) Outputs() int {
	return afdelaysrcBuilder.f.Outputs()
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) With(key string, value expr.Expr) filter.Filter {
	return afdelaysrcBuilder.withOption(key, value)
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) withOption(key string, value expr.Expr) AfdelaysrcBuilder {
	bCopy := *afdelaysrcBuilder
	bCopy.f = afdelaysrcBuilder.f.With(key, value)
	return &bCopy
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) Delay(delay float64) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("delay", expr.Double(delay))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) D(d float64) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("d", expr.Double(d))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) SampleRate(sampleRate int) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("sample_rate", expr.Int(sampleRate))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) R(r int) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("r", expr.Int(r))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) NbSamples(nbSamples int) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("nb_samples", expr.Int(nbSamples))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) N(n int) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("n", expr.Int(n))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) Taps(taps int) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("taps", expr.Int(taps))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) T(t int) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("t", expr.Int(t))
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) ChannelLayout(channelLayout expr.ChannelLayout) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("channel_layout", channelLayout)
}

func (afdelaysrcBuilder *implAfdelaysrcBuilder) C(c expr.ChannelLayout) AfdelaysrcBuilder {
	return afdelaysrcBuilder.withOption("c", c)
}
