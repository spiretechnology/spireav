// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// HaldclutsrcBuilder Provide an identity Hald CLUT.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#haldclutsrc
type HaldclutsrcBuilder interface {
	filter.Filter
	// Level set level (from 2 to 16) (default 6).
	Level(level int) HaldclutsrcBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) HaldclutsrcBuilder
	// R set video rate (default "25").
	R(r expr.Rational) HaldclutsrcBuilder
	// Duration set video duration (default -0.000001).
	Duration(duration time.Duration) HaldclutsrcBuilder
	// D set video duration (default -0.000001).
	D(d time.Duration) HaldclutsrcBuilder
	// Sar set video sample aspect ratio (from 0 to INT_MAX) (default 1/1).
	Sar(sar expr.Rational) HaldclutsrcBuilder
}

// Haldclutsrc creates a new HaldclutsrcBuilder to configure the "haldclutsrc" filter.
func Haldclutsrc(opts ...filter.Option) HaldclutsrcBuilder {
	return &implHaldclutsrcBuilder{
		f: filter.New("haldclutsrc", 1, opts...),
	}
}

type implHaldclutsrcBuilder struct {
	f filter.Filter
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) String() string {
	return haldclutsrcBuilder.f.String()
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) Outputs() int {
	return haldclutsrcBuilder.f.Outputs()
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) With(key string, value expr.Expr) filter.Filter {
	return haldclutsrcBuilder.withOption(key, value)
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) withOption(key string, value expr.Expr) HaldclutsrcBuilder {
	bCopy := *haldclutsrcBuilder
	bCopy.f = haldclutsrcBuilder.f.With(key, value)
	return &bCopy
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) Level(level int) HaldclutsrcBuilder {
	return haldclutsrcBuilder.withOption("level", expr.Int(level))
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) Rate(rate expr.Rational) HaldclutsrcBuilder {
	return haldclutsrcBuilder.withOption("rate", rate)
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) R(r expr.Rational) HaldclutsrcBuilder {
	return haldclutsrcBuilder.withOption("r", r)
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) Duration(duration time.Duration) HaldclutsrcBuilder {
	return haldclutsrcBuilder.withOption("duration", expr.Duration(duration))
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) D(d time.Duration) HaldclutsrcBuilder {
	return haldclutsrcBuilder.withOption("d", expr.Duration(d))
}

func (haldclutsrcBuilder *implHaldclutsrcBuilder) Sar(sar expr.Rational) HaldclutsrcBuilder {
	return haldclutsrcBuilder.withOption("sar", sar)
}
