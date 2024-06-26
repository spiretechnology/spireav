// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AmplifyBuilder Amplify changes between successive video frames.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#amplify
type AmplifyBuilder interface {
	filter.Filter
	// Radius set radius (from 1 to 63) (default 2).
	Radius(radius int) AmplifyBuilder
	// Factor set factor (from 0 to 65535) (default 2).
	Factor(factor float32) AmplifyBuilder
	// FactorExpr set factor (from 0 to 65535) (default 2).
	FactorExpr(factor expr.Expr) AmplifyBuilder
	// Threshold set threshold (from 0 to 65535) (default 10).
	Threshold(threshold float32) AmplifyBuilder
	// ThresholdExpr set threshold (from 0 to 65535) (default 10).
	ThresholdExpr(threshold expr.Expr) AmplifyBuilder
	// Tolerance set tolerance (from 0 to 65535) (default 0).
	Tolerance(tolerance float32) AmplifyBuilder
	// ToleranceExpr set tolerance (from 0 to 65535) (default 0).
	ToleranceExpr(tolerance expr.Expr) AmplifyBuilder
	// Low set low limit for amplification (from 0 to 65535) (default 65535).
	Low(low float32) AmplifyBuilder
	// LowExpr set low limit for amplification (from 0 to 65535) (default 65535).
	LowExpr(low expr.Expr) AmplifyBuilder
	// High set high limit for amplification (from 0 to 65535) (default 65535).
	High(high float32) AmplifyBuilder
	// HighExpr set high limit for amplification (from 0 to 65535) (default 65535).
	HighExpr(high expr.Expr) AmplifyBuilder
	// Planes set what planes to filter (default 7).
	Planes(planes ...string) AmplifyBuilder
	// PlanesExpr set what planes to filter (default 7).
	PlanesExpr(planes expr.Expr) AmplifyBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AmplifyBuilder
}

// Amplify creates a new AmplifyBuilder to configure the "amplify" filter.
func Amplify(opts ...filter.Option) AmplifyBuilder {
	return &implAmplifyBuilder{
		f: filter.New("amplify", 1, opts...),
	}
}

type implAmplifyBuilder struct {
	f filter.Filter
}

func (amplifyBuilder *implAmplifyBuilder) String() string {
	return amplifyBuilder.f.String()
}

func (amplifyBuilder *implAmplifyBuilder) Outputs() int {
	return amplifyBuilder.f.Outputs()
}

func (amplifyBuilder *implAmplifyBuilder) With(key string, value expr.Expr) filter.Filter {
	return amplifyBuilder.withOption(key, value)
}

func (amplifyBuilder *implAmplifyBuilder) withOption(key string, value expr.Expr) AmplifyBuilder {
	bCopy := *amplifyBuilder
	bCopy.f = amplifyBuilder.f.With(key, value)
	return &bCopy
}

func (amplifyBuilder *implAmplifyBuilder) Radius(radius int) AmplifyBuilder {
	return amplifyBuilder.withOption("radius", expr.Int(radius))
}

func (amplifyBuilder *implAmplifyBuilder) Factor(factor float32) AmplifyBuilder {
	return amplifyBuilder.withOption("factor", expr.Float(factor))
}

func (amplifyBuilder *implAmplifyBuilder) FactorExpr(factor expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("factor", factor)
}

func (amplifyBuilder *implAmplifyBuilder) Threshold(threshold float32) AmplifyBuilder {
	return amplifyBuilder.withOption("threshold", expr.Float(threshold))
}

func (amplifyBuilder *implAmplifyBuilder) ThresholdExpr(threshold expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("threshold", threshold)
}

func (amplifyBuilder *implAmplifyBuilder) Tolerance(tolerance float32) AmplifyBuilder {
	return amplifyBuilder.withOption("tolerance", expr.Float(tolerance))
}

func (amplifyBuilder *implAmplifyBuilder) ToleranceExpr(tolerance expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("tolerance", tolerance)
}

func (amplifyBuilder *implAmplifyBuilder) Low(low float32) AmplifyBuilder {
	return amplifyBuilder.withOption("low", expr.Float(low))
}

func (amplifyBuilder *implAmplifyBuilder) LowExpr(low expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("low", low)
}

func (amplifyBuilder *implAmplifyBuilder) High(high float32) AmplifyBuilder {
	return amplifyBuilder.withOption("high", expr.Float(high))
}

func (amplifyBuilder *implAmplifyBuilder) HighExpr(high expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("high", high)
}

func (amplifyBuilder *implAmplifyBuilder) Planes(planes ...string) AmplifyBuilder {
	return amplifyBuilder.withOption("planes", expr.Flags(planes))
}

func (amplifyBuilder *implAmplifyBuilder) PlanesExpr(planes expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("planes", planes)
}

func (amplifyBuilder *implAmplifyBuilder) Enable(enable expr.Expr) AmplifyBuilder {
	return amplifyBuilder.withOption("enable", enable)
}
