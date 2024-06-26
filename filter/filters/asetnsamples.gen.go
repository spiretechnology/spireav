// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AsetnsamplesBuilder Set the number of samples for each output audio frames.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#asetnsamples
type AsetnsamplesBuilder interface {
	filter.Filter
	// NbOutSamples set the number of per-frame output samples (from 1 to INT_MAX) (default 1024).
	NbOutSamples(nbOutSamples int) AsetnsamplesBuilder
	// NbOutSamplesExpr set the number of per-frame output samples (from 1 to INT_MAX) (default 1024).
	NbOutSamplesExpr(nbOutSamples expr.Expr) AsetnsamplesBuilder
	// N set the number of per-frame output samples (from 1 to INT_MAX) (default 1024).
	N(n int) AsetnsamplesBuilder
	// NExpr set the number of per-frame output samples (from 1 to INT_MAX) (default 1024).
	NExpr(n expr.Expr) AsetnsamplesBuilder
	// Pad pad last frame with zeros (default true).
	Pad(pad bool) AsetnsamplesBuilder
	// PadExpr pad last frame with zeros (default true).
	PadExpr(pad expr.Expr) AsetnsamplesBuilder
	// P pad last frame with zeros (default true).
	P(p bool) AsetnsamplesBuilder
	// PExpr pad last frame with zeros (default true).
	PExpr(p expr.Expr) AsetnsamplesBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AsetnsamplesBuilder
}

// Asetnsamples creates a new AsetnsamplesBuilder to configure the "asetnsamples" filter.
func Asetnsamples(opts ...filter.Option) AsetnsamplesBuilder {
	return &implAsetnsamplesBuilder{
		f: filter.New("asetnsamples", 1, opts...),
	}
}

type implAsetnsamplesBuilder struct {
	f filter.Filter
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) String() string {
	return asetnsamplesBuilder.f.String()
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) Outputs() int {
	return asetnsamplesBuilder.f.Outputs()
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) With(key string, value expr.Expr) filter.Filter {
	return asetnsamplesBuilder.withOption(key, value)
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) withOption(key string, value expr.Expr) AsetnsamplesBuilder {
	bCopy := *asetnsamplesBuilder
	bCopy.f = asetnsamplesBuilder.f.With(key, value)
	return &bCopy
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) NbOutSamples(nbOutSamples int) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("nb_out_samples", expr.Int(nbOutSamples))
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) NbOutSamplesExpr(nbOutSamples expr.Expr) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("nb_out_samples", nbOutSamples)
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) N(n int) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("n", expr.Int(n))
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) NExpr(n expr.Expr) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("n", n)
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) Pad(pad bool) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("pad", expr.Bool(pad))
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) PadExpr(pad expr.Expr) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("pad", pad)
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) P(p bool) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("p", expr.Bool(p))
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) PExpr(p expr.Expr) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("p", p)
}

func (asetnsamplesBuilder *implAsetnsamplesBuilder) Enable(enable expr.Expr) AsetnsamplesBuilder {
	return asetnsamplesBuilder.withOption("enable", enable)
}
