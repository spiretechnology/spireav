// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SineBuilder Generate sine wave audio signal.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#sine
type SineBuilder interface {
	filter.Filter
	// Frequency set the sine frequency (from 0 to DBL_MAX) (default 440).
	Frequency(frequency float64) SineBuilder
	// F set the sine frequency (from 0 to DBL_MAX) (default 440).
	F(f float64) SineBuilder
	// BeepFactor set the beep frequency factor (from 0 to DBL_MAX) (default 0).
	BeepFactor(beepFactor float64) SineBuilder
	// B set the beep frequency factor (from 0 to DBL_MAX) (default 0).
	B(b float64) SineBuilder
	// SampleRate set the sample rate (from 1 to INT_MAX) (default 44100).
	SampleRate(sampleRate int) SineBuilder
	// R set the sample rate (from 1 to INT_MAX) (default 44100).
	R(r int) SineBuilder
	// Duration set the audio duration (default 0).
	Duration(duration time.Duration) SineBuilder
	// D set the audio duration (default 0).
	D(d time.Duration) SineBuilder
	// SamplesPerFrame set the number of samples per frame (default "1024").
	SamplesPerFrame(samplesPerFrame string) SineBuilder
}

// Sine creates a new SineBuilder to configure the "sine" filter.
func Sine(opts ...filter.Option) SineBuilder {
	return &implSineBuilder{
		f: filter.New("sine", 1, opts...),
	}
}

type implSineBuilder struct {
	f filter.Filter
}

func (sineBuilder *implSineBuilder) String() string {
	return sineBuilder.f.String()
}

func (sineBuilder *implSineBuilder) Outputs() int {
	return sineBuilder.f.Outputs()
}

func (sineBuilder *implSineBuilder) With(key string, value expr.Expr) filter.Filter {
	return sineBuilder.withOption(key, value)
}

func (sineBuilder *implSineBuilder) withOption(key string, value expr.Expr) SineBuilder {
	bCopy := *sineBuilder
	bCopy.f = sineBuilder.f.With(key, value)
	return &bCopy
}

func (sineBuilder *implSineBuilder) Frequency(frequency float64) SineBuilder {
	return sineBuilder.withOption("frequency", expr.Double(frequency))
}

func (sineBuilder *implSineBuilder) F(f float64) SineBuilder {
	return sineBuilder.withOption("f", expr.Double(f))
}

func (sineBuilder *implSineBuilder) BeepFactor(beepFactor float64) SineBuilder {
	return sineBuilder.withOption("beep_factor", expr.Double(beepFactor))
}

func (sineBuilder *implSineBuilder) B(b float64) SineBuilder {
	return sineBuilder.withOption("b", expr.Double(b))
}

func (sineBuilder *implSineBuilder) SampleRate(sampleRate int) SineBuilder {
	return sineBuilder.withOption("sample_rate", expr.Int(sampleRate))
}

func (sineBuilder *implSineBuilder) R(r int) SineBuilder {
	return sineBuilder.withOption("r", expr.Int(r))
}

func (sineBuilder *implSineBuilder) Duration(duration time.Duration) SineBuilder {
	return sineBuilder.withOption("duration", expr.Duration(duration))
}

func (sineBuilder *implSineBuilder) D(d time.Duration) SineBuilder {
	return sineBuilder.withOption("d", expr.Duration(d))
}

func (sineBuilder *implSineBuilder) SamplesPerFrame(samplesPerFrame string) SineBuilder {
	return sineBuilder.withOption("samples_per_frame", expr.String(samplesPerFrame))
}
