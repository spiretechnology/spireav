// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AcrusherBuilder Reduce audio bit resolution.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#acrusher
type AcrusherBuilder interface {
	filter.Filter
	// LevelIn set level in (from 0.015625 to 64) (default 1).
	LevelIn(levelIn float64) AcrusherBuilder
	// LevelInExpr set level in (from 0.015625 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) AcrusherBuilder
	// LevelOut set level out (from 0.015625 to 64) (default 1).
	LevelOut(levelOut float64) AcrusherBuilder
	// LevelOutExpr set level out (from 0.015625 to 64) (default 1).
	LevelOutExpr(levelOut expr.Expr) AcrusherBuilder
	// Bits set bit reduction (from 1 to 64) (default 8).
	Bits(bits float64) AcrusherBuilder
	// BitsExpr set bit reduction (from 1 to 64) (default 8).
	BitsExpr(bits expr.Expr) AcrusherBuilder
	// Mix set mix (from 0 to 1) (default 0.5).
	Mix(mix float64) AcrusherBuilder
	// MixExpr set mix (from 0 to 1) (default 0.5).
	MixExpr(mix expr.Expr) AcrusherBuilder
	// Mode set mode (from 0 to 1) (default lin).
	Mode(mode int) AcrusherBuilder
	// ModeExpr set mode (from 0 to 1) (default lin).
	ModeExpr(mode expr.Expr) AcrusherBuilder
	// Dc set DC (from 0.25 to 4) (default 1).
	Dc(dc float64) AcrusherBuilder
	// DcExpr set DC (from 0.25 to 4) (default 1).
	DcExpr(dc expr.Expr) AcrusherBuilder
	// Aa set anti-aliasing (from 0 to 1) (default 0.5).
	Aa(aa float64) AcrusherBuilder
	// AaExpr set anti-aliasing (from 0 to 1) (default 0.5).
	AaExpr(aa expr.Expr) AcrusherBuilder
	// Samples set sample reduction (from 1 to 250) (default 1).
	Samples(samples float64) AcrusherBuilder
	// SamplesExpr set sample reduction (from 1 to 250) (default 1).
	SamplesExpr(samples expr.Expr) AcrusherBuilder
	// Lfo enable LFO (default false).
	Lfo(lfo bool) AcrusherBuilder
	// LfoExpr enable LFO (default false).
	LfoExpr(lfo expr.Expr) AcrusherBuilder
	// Lforange set LFO depth (from 1 to 250) (default 20).
	Lforange(lforange float64) AcrusherBuilder
	// LforangeExpr set LFO depth (from 1 to 250) (default 20).
	LforangeExpr(lforange expr.Expr) AcrusherBuilder
	// Lforate set LFO rate (from 0.01 to 200) (default 0.3).
	Lforate(lforate float64) AcrusherBuilder
	// LforateExpr set LFO rate (from 0.01 to 200) (default 0.3).
	LforateExpr(lforate expr.Expr) AcrusherBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AcrusherBuilder
}

// Acrusher creates a new AcrusherBuilder to configure the "acrusher" filter.
func Acrusher(opts ...filter.Option) AcrusherBuilder {
	return &implAcrusherBuilder{
		f: filter.New("acrusher", 1, opts...),
	}
}

type implAcrusherBuilder struct {
	f filter.Filter
}

func (acrusherBuilder *implAcrusherBuilder) String() string {
	return acrusherBuilder.f.String()
}

func (acrusherBuilder *implAcrusherBuilder) Outputs() int {
	return acrusherBuilder.f.Outputs()
}

func (acrusherBuilder *implAcrusherBuilder) With(key string, value expr.Expr) filter.Filter {
	return acrusherBuilder.withOption(key, value)
}

func (acrusherBuilder *implAcrusherBuilder) withOption(key string, value expr.Expr) AcrusherBuilder {
	bCopy := *acrusherBuilder
	bCopy.f = acrusherBuilder.f.With(key, value)
	return &bCopy
}

func (acrusherBuilder *implAcrusherBuilder) LevelIn(levelIn float64) AcrusherBuilder {
	return acrusherBuilder.withOption("level_in", expr.Double(levelIn))
}

func (acrusherBuilder *implAcrusherBuilder) LevelInExpr(levelIn expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("level_in", levelIn)
}

func (acrusherBuilder *implAcrusherBuilder) LevelOut(levelOut float64) AcrusherBuilder {
	return acrusherBuilder.withOption("level_out", expr.Double(levelOut))
}

func (acrusherBuilder *implAcrusherBuilder) LevelOutExpr(levelOut expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("level_out", levelOut)
}

func (acrusherBuilder *implAcrusherBuilder) Bits(bits float64) AcrusherBuilder {
	return acrusherBuilder.withOption("bits", expr.Double(bits))
}

func (acrusherBuilder *implAcrusherBuilder) BitsExpr(bits expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("bits", bits)
}

func (acrusherBuilder *implAcrusherBuilder) Mix(mix float64) AcrusherBuilder {
	return acrusherBuilder.withOption("mix", expr.Double(mix))
}

func (acrusherBuilder *implAcrusherBuilder) MixExpr(mix expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("mix", mix)
}

func (acrusherBuilder *implAcrusherBuilder) Mode(mode int) AcrusherBuilder {
	return acrusherBuilder.withOption("mode", expr.Int(mode))
}

func (acrusherBuilder *implAcrusherBuilder) ModeExpr(mode expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("mode", mode)
}

func (acrusherBuilder *implAcrusherBuilder) Dc(dc float64) AcrusherBuilder {
	return acrusherBuilder.withOption("dc", expr.Double(dc))
}

func (acrusherBuilder *implAcrusherBuilder) DcExpr(dc expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("dc", dc)
}

func (acrusherBuilder *implAcrusherBuilder) Aa(aa float64) AcrusherBuilder {
	return acrusherBuilder.withOption("aa", expr.Double(aa))
}

func (acrusherBuilder *implAcrusherBuilder) AaExpr(aa expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("aa", aa)
}

func (acrusherBuilder *implAcrusherBuilder) Samples(samples float64) AcrusherBuilder {
	return acrusherBuilder.withOption("samples", expr.Double(samples))
}

func (acrusherBuilder *implAcrusherBuilder) SamplesExpr(samples expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("samples", samples)
}

func (acrusherBuilder *implAcrusherBuilder) Lfo(lfo bool) AcrusherBuilder {
	return acrusherBuilder.withOption("lfo", expr.Bool(lfo))
}

func (acrusherBuilder *implAcrusherBuilder) LfoExpr(lfo expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("lfo", lfo)
}

func (acrusherBuilder *implAcrusherBuilder) Lforange(lforange float64) AcrusherBuilder {
	return acrusherBuilder.withOption("lforange", expr.Double(lforange))
}

func (acrusherBuilder *implAcrusherBuilder) LforangeExpr(lforange expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("lforange", lforange)
}

func (acrusherBuilder *implAcrusherBuilder) Lforate(lforate float64) AcrusherBuilder {
	return acrusherBuilder.withOption("lforate", expr.Double(lforate))
}

func (acrusherBuilder *implAcrusherBuilder) LforateExpr(lforate expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("lforate", lforate)
}

func (acrusherBuilder *implAcrusherBuilder) Enable(enable expr.Expr) AcrusherBuilder {
	return acrusherBuilder.withOption("enable", enable)
}
