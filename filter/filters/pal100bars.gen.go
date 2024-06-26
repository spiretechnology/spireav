// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// Pal100barsBuilder Generate PAL 100% color bars.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#pal100bars
type Pal100barsBuilder interface {
	filter.Filter
	// Size set video size (default "320x240").
	Size(size expr.Size) Pal100barsBuilder
	// S set video size (default "320x240").
	S(s expr.Size) Pal100barsBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) Pal100barsBuilder
	// R set video rate (default "25").
	R(r expr.Rational) Pal100barsBuilder
	// Duration set video duration (default -0.000001).
	Duration(duration time.Duration) Pal100barsBuilder
	// D set video duration (default -0.000001).
	D(d time.Duration) Pal100barsBuilder
	// Sar set video sample aspect ratio (from 0 to INT_MAX) (default 1/1).
	Sar(sar expr.Rational) Pal100barsBuilder
}

// Pal100bars creates a new Pal100barsBuilder to configure the "pal100bars" filter.
func Pal100bars(opts ...filter.Option) Pal100barsBuilder {
	return &implPal100barsBuilder{
		f: filter.New("pal100bars", 1, opts...),
	}
}

type implPal100barsBuilder struct {
	f filter.Filter
}

func (pal100barsBuilder *implPal100barsBuilder) String() string {
	return pal100barsBuilder.f.String()
}

func (pal100barsBuilder *implPal100barsBuilder) Outputs() int {
	return pal100barsBuilder.f.Outputs()
}

func (pal100barsBuilder *implPal100barsBuilder) With(key string, value expr.Expr) filter.Filter {
	return pal100barsBuilder.withOption(key, value)
}

func (pal100barsBuilder *implPal100barsBuilder) withOption(key string, value expr.Expr) Pal100barsBuilder {
	bCopy := *pal100barsBuilder
	bCopy.f = pal100barsBuilder.f.With(key, value)
	return &bCopy
}

func (pal100barsBuilder *implPal100barsBuilder) Size(size expr.Size) Pal100barsBuilder {
	return pal100barsBuilder.withOption("size", size)
}

func (pal100barsBuilder *implPal100barsBuilder) S(s expr.Size) Pal100barsBuilder {
	return pal100barsBuilder.withOption("s", s)
}

func (pal100barsBuilder *implPal100barsBuilder) Rate(rate expr.Rational) Pal100barsBuilder {
	return pal100barsBuilder.withOption("rate", rate)
}

func (pal100barsBuilder *implPal100barsBuilder) R(r expr.Rational) Pal100barsBuilder {
	return pal100barsBuilder.withOption("r", r)
}

func (pal100barsBuilder *implPal100barsBuilder) Duration(duration time.Duration) Pal100barsBuilder {
	return pal100barsBuilder.withOption("duration", expr.Duration(duration))
}

func (pal100barsBuilder *implPal100barsBuilder) D(d time.Duration) Pal100barsBuilder {
	return pal100barsBuilder.withOption("d", expr.Duration(d))
}

func (pal100barsBuilder *implPal100barsBuilder) Sar(sar expr.Rational) Pal100barsBuilder {
	return pal100barsBuilder.withOption("sar", sar)
}
