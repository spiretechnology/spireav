// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// MixBuilder Mix video inputs.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#mix
type MixBuilder interface {
	filter.Filter
	// Inputs set number of inputs (from 2 to 32767) (default 2).
	Inputs(inputs int) MixBuilder
	// Weights set weight for each input (default "1 1").
	Weights(weights string) MixBuilder
	// WeightsExpr set weight for each input (default "1 1").
	WeightsExpr(weights expr.Expr) MixBuilder
	// Scale set scale (from 0 to 32767) (default 0).
	Scale(scale float32) MixBuilder
	// ScaleExpr set scale (from 0 to 32767) (default 0).
	ScaleExpr(scale expr.Expr) MixBuilder
	// Planes set what planes to filter (default F).
	Planes(planes ...string) MixBuilder
	// PlanesExpr set what planes to filter (default F).
	PlanesExpr(planes expr.Expr) MixBuilder
	// Duration how to determine end of stream (from 0 to 2) (default longest).
	Duration(duration int) MixBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) MixBuilder
}

// Mix creates a new MixBuilder to configure the "mix" filter.
func Mix(opts ...filter.Option) MixBuilder {
	return &implMixBuilder{
		f: filter.New("mix", 1, opts...),
	}
}

type implMixBuilder struct {
	f filter.Filter
}

func (mixBuilder *implMixBuilder) String() string {
	return mixBuilder.f.String()
}

func (mixBuilder *implMixBuilder) Outputs() int {
	return mixBuilder.f.Outputs()
}

func (mixBuilder *implMixBuilder) With(key string, value expr.Expr) filter.Filter {
	return mixBuilder.withOption(key, value)
}

func (mixBuilder *implMixBuilder) withOption(key string, value expr.Expr) MixBuilder {
	bCopy := *mixBuilder
	bCopy.f = mixBuilder.f.With(key, value)
	return &bCopy
}

func (mixBuilder *implMixBuilder) Inputs(inputs int) MixBuilder {
	return mixBuilder.withOption("inputs", expr.Int(inputs))
}

func (mixBuilder *implMixBuilder) Weights(weights string) MixBuilder {
	return mixBuilder.withOption("weights", expr.String(weights))
}

func (mixBuilder *implMixBuilder) WeightsExpr(weights expr.Expr) MixBuilder {
	return mixBuilder.withOption("weights", weights)
}

func (mixBuilder *implMixBuilder) Scale(scale float32) MixBuilder {
	return mixBuilder.withOption("scale", expr.Float(scale))
}

func (mixBuilder *implMixBuilder) ScaleExpr(scale expr.Expr) MixBuilder {
	return mixBuilder.withOption("scale", scale)
}

func (mixBuilder *implMixBuilder) Planes(planes ...string) MixBuilder {
	return mixBuilder.withOption("planes", expr.Flags(planes))
}

func (mixBuilder *implMixBuilder) PlanesExpr(planes expr.Expr) MixBuilder {
	return mixBuilder.withOption("planes", planes)
}

func (mixBuilder *implMixBuilder) Duration(duration int) MixBuilder {
	return mixBuilder.withOption("duration", expr.Int(duration))
}

func (mixBuilder *implMixBuilder) Enable(enable expr.Expr) MixBuilder {
	return mixBuilder.withOption("enable", enable)
}
