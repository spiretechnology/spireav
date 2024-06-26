// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TonemapBuilder Conversion to/from different dynamic ranges.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#tonemap
type TonemapBuilder interface {
	filter.Filter
	// Tonemap tonemap algorithm selection (from 0 to 6) (default none).
	Tonemap(tonemap int) TonemapBuilder
	// Param tonemap parameter (from DBL_MIN to DBL_MAX) (default nan).
	Param(param float64) TonemapBuilder
	// Desat desaturation strength (from 0 to DBL_MAX) (default 2).
	Desat(desat float64) TonemapBuilder
	// Peak signal peak override (from 0 to DBL_MAX) (default 0).
	Peak(peak float64) TonemapBuilder
}

// Tonemap creates a new TonemapBuilder to configure the "tonemap" filter.
func Tonemap(opts ...filter.Option) TonemapBuilder {
	return &implTonemapBuilder{
		f: filter.New("tonemap", 1, opts...),
	}
}

type implTonemapBuilder struct {
	f filter.Filter
}

func (tonemapBuilder *implTonemapBuilder) String() string {
	return tonemapBuilder.f.String()
}

func (tonemapBuilder *implTonemapBuilder) Outputs() int {
	return tonemapBuilder.f.Outputs()
}

func (tonemapBuilder *implTonemapBuilder) With(key string, value expr.Expr) filter.Filter {
	return tonemapBuilder.withOption(key, value)
}

func (tonemapBuilder *implTonemapBuilder) withOption(key string, value expr.Expr) TonemapBuilder {
	bCopy := *tonemapBuilder
	bCopy.f = tonemapBuilder.f.With(key, value)
	return &bCopy
}

func (tonemapBuilder *implTonemapBuilder) Tonemap(tonemap int) TonemapBuilder {
	return tonemapBuilder.withOption("tonemap", expr.Int(tonemap))
}

func (tonemapBuilder *implTonemapBuilder) Param(param float64) TonemapBuilder {
	return tonemapBuilder.withOption("param", expr.Double(param))
}

func (tonemapBuilder *implTonemapBuilder) Desat(desat float64) TonemapBuilder {
	return tonemapBuilder.withOption("desat", expr.Double(desat))
}

func (tonemapBuilder *implTonemapBuilder) Peak(peak float64) TonemapBuilder {
	return tonemapBuilder.withOption("peak", expr.Double(peak))
}
