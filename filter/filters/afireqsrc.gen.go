// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AfireqsrcBuilder Generate a FIR equalizer coefficients audio stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#afireqsrc
type AfireqsrcBuilder interface {
	filter.Filter
	// Preset set equalizer preset (from -1 to 17) (default flat).
	Preset(preset int) AfireqsrcBuilder
	// P set equalizer preset (from -1 to 17) (default flat).
	P(p int) AfireqsrcBuilder
	// Gains set gain values per band (default "0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0").
	Gains(gains string) AfireqsrcBuilder
	// G set gain values per band (default "0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0").
	G(g string) AfireqsrcBuilder
	// Bands set central frequency values per band (default "25 40 63 100 160 250 400 630 1000 1600 2500 4000 6300 10000 16000 24000").
	Bands(bands string) AfireqsrcBuilder
	// B set central frequency values per band (default "25 40 63 100 160 250 400 630 1000 1600 2500 4000 6300 10000 16000 24000").
	B(b string) AfireqsrcBuilder
	// Taps set number of taps (from 16 to 65535) (default 4096).
	Taps(taps int) AfireqsrcBuilder
	// T set number of taps (from 16 to 65535) (default 4096).
	T(t int) AfireqsrcBuilder
	// SampleRate set sample rate (from 1 to INT_MAX) (default 44100).
	SampleRate(sampleRate int) AfireqsrcBuilder
	// R set sample rate (from 1 to INT_MAX) (default 44100).
	R(r int) AfireqsrcBuilder
	// NbSamples set the number of samples per requested frame (from 1 to INT_MAX) (default 1024).
	NbSamples(nbSamples int) AfireqsrcBuilder
	// N set the number of samples per requested frame (from 1 to INT_MAX) (default 1024).
	N(n int) AfireqsrcBuilder
	// Interp set the interpolation (from 0 to 1) (default linear).
	Interp(interp int) AfireqsrcBuilder
	// I set the interpolation (from 0 to 1) (default linear).
	I(i int) AfireqsrcBuilder
	// Phase set the phase (from 0 to 1) (default min).
	Phase(phase int) AfireqsrcBuilder
	// H set the phase (from 0 to 1) (default min).
	H(h int) AfireqsrcBuilder
}

// Afireqsrc creates a new AfireqsrcBuilder to configure the "afireqsrc" filter.
func Afireqsrc(opts ...filter.Option) AfireqsrcBuilder {
	return &implAfireqsrcBuilder{
		f: filter.New("afireqsrc", 1, opts...),
	}
}

type implAfireqsrcBuilder struct {
	f filter.Filter
}

func (afireqsrcBuilder *implAfireqsrcBuilder) String() string {
	return afireqsrcBuilder.f.String()
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Outputs() int {
	return afireqsrcBuilder.f.Outputs()
}

func (afireqsrcBuilder *implAfireqsrcBuilder) With(key string, value expr.Expr) filter.Filter {
	return afireqsrcBuilder.withOption(key, value)
}

func (afireqsrcBuilder *implAfireqsrcBuilder) withOption(key string, value expr.Expr) AfireqsrcBuilder {
	bCopy := *afireqsrcBuilder
	bCopy.f = afireqsrcBuilder.f.With(key, value)
	return &bCopy
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Preset(preset int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("preset", expr.Int(preset))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) P(p int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("p", expr.Int(p))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Gains(gains string) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("gains", expr.String(gains))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) G(g string) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("g", expr.String(g))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Bands(bands string) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("bands", expr.String(bands))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) B(b string) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("b", expr.String(b))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Taps(taps int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("taps", expr.Int(taps))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) T(t int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("t", expr.Int(t))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) SampleRate(sampleRate int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("sample_rate", expr.Int(sampleRate))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) R(r int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("r", expr.Int(r))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) NbSamples(nbSamples int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("nb_samples", expr.Int(nbSamples))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) N(n int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("n", expr.Int(n))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Interp(interp int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("interp", expr.Int(interp))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) I(i int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("i", expr.Int(i))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) Phase(phase int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("phase", expr.Int(phase))
}

func (afireqsrcBuilder *implAfireqsrcBuilder) H(h int) AfireqsrcBuilder {
	return afireqsrcBuilder.withOption("h", expr.Int(h))
}
