// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AdynamicequalizerBuilder Apply Dynamic Equalization of input audio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#adynamicequalizer
type AdynamicequalizerBuilder interface {
	filter.Filter
	// Threshold set detection threshold (from 0 to 100) (default 0).
	Threshold(threshold float64) AdynamicequalizerBuilder
	// ThresholdExpr set detection threshold (from 0 to 100) (default 0).
	ThresholdExpr(threshold expr.Expr) AdynamicequalizerBuilder
	// Dfrequency set detection frequency (from 2 to 1e+06) (default 1000).
	Dfrequency(dfrequency float64) AdynamicequalizerBuilder
	// DfrequencyExpr set detection frequency (from 2 to 1e+06) (default 1000).
	DfrequencyExpr(dfrequency expr.Expr) AdynamicequalizerBuilder
	// Dqfactor set detection Q factor (from 0.001 to 1000) (default 1).
	Dqfactor(dqfactor float64) AdynamicequalizerBuilder
	// DqfactorExpr set detection Q factor (from 0.001 to 1000) (default 1).
	DqfactorExpr(dqfactor expr.Expr) AdynamicequalizerBuilder
	// Tfrequency set target frequency (from 2 to 1e+06) (default 1000).
	Tfrequency(tfrequency float64) AdynamicequalizerBuilder
	// TfrequencyExpr set target frequency (from 2 to 1e+06) (default 1000).
	TfrequencyExpr(tfrequency expr.Expr) AdynamicequalizerBuilder
	// Tqfactor set target Q factor (from 0.001 to 1000) (default 1).
	Tqfactor(tqfactor float64) AdynamicequalizerBuilder
	// TqfactorExpr set target Q factor (from 0.001 to 1000) (default 1).
	TqfactorExpr(tqfactor expr.Expr) AdynamicequalizerBuilder
	// Attack set detection attack duration (from 0.01 to 2000) (default 20).
	Attack(attack float64) AdynamicequalizerBuilder
	// AttackExpr set detection attack duration (from 0.01 to 2000) (default 20).
	AttackExpr(attack expr.Expr) AdynamicequalizerBuilder
	// Release set detection release duration (from 0.01 to 2000) (default 200).
	Release(release float64) AdynamicequalizerBuilder
	// ReleaseExpr set detection release duration (from 0.01 to 2000) (default 200).
	ReleaseExpr(release expr.Expr) AdynamicequalizerBuilder
	// Ratio set ratio factor (from 0 to 30) (default 1).
	Ratio(ratio float64) AdynamicequalizerBuilder
	// RatioExpr set ratio factor (from 0 to 30) (default 1).
	RatioExpr(ratio expr.Expr) AdynamicequalizerBuilder
	// Makeup set makeup gain (from 0 to 1000) (default 0).
	Makeup(makeup float64) AdynamicequalizerBuilder
	// MakeupExpr set makeup gain (from 0 to 1000) (default 0).
	MakeupExpr(makeup expr.Expr) AdynamicequalizerBuilder
	// Range set max gain (from 1 to 2000) (default 50).
	Range(rng float64) AdynamicequalizerBuilder
	// RangeExpr set max gain (from 1 to 2000) (default 50).
	RangeExpr(rng expr.Expr) AdynamicequalizerBuilder
	// Mode set mode (from -1 to 3) (default cutbelow).
	Mode(mode int) AdynamicequalizerBuilder
	// ModeExpr set mode (from -1 to 3) (default cutbelow).
	ModeExpr(mode expr.Expr) AdynamicequalizerBuilder
	// Dftype set detection filter type (from 0 to 3) (default bandpass).
	Dftype(dftype int) AdynamicequalizerBuilder
	// DftypeExpr set detection filter type (from 0 to 3) (default bandpass).
	DftypeExpr(dftype expr.Expr) AdynamicequalizerBuilder
	// Tftype set target filter type (from 0 to 2) (default bell).
	Tftype(tftype int) AdynamicequalizerBuilder
	// TftypeExpr set target filter type (from 0 to 2) (default bell).
	TftypeExpr(tftype expr.Expr) AdynamicequalizerBuilder
	// Auto set auto threshold (from 1 to 4) (default off).
	Auto(auto int) AdynamicequalizerBuilder
	// AutoExpr set auto threshold (from 1 to 4) (default off).
	AutoExpr(auto expr.Expr) AdynamicequalizerBuilder
	// Precision set processing precision (from 0 to 2) (default auto).
	Precision(precision int) AdynamicequalizerBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AdynamicequalizerBuilder
}

// Adynamicequalizer creates a new AdynamicequalizerBuilder to configure the "adynamicequalizer" filter.
func Adynamicequalizer(opts ...filter.Option) AdynamicequalizerBuilder {
	return &implAdynamicequalizerBuilder{
		f: filter.New("adynamicequalizer", 1, opts...),
	}
}

type implAdynamicequalizerBuilder struct {
	f filter.Filter
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) String() string {
	return adynamicequalizerBuilder.f.String()
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Outputs() int {
	return adynamicequalizerBuilder.f.Outputs()
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) With(key string, value expr.Expr) filter.Filter {
	return adynamicequalizerBuilder.withOption(key, value)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) withOption(key string, value expr.Expr) AdynamicequalizerBuilder {
	bCopy := *adynamicequalizerBuilder
	bCopy.f = adynamicequalizerBuilder.f.With(key, value)
	return &bCopy
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Threshold(threshold float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("threshold", expr.Double(threshold))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) ThresholdExpr(threshold expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("threshold", threshold)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Dfrequency(dfrequency float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("dfrequency", expr.Double(dfrequency))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) DfrequencyExpr(dfrequency expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("dfrequency", dfrequency)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Dqfactor(dqfactor float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("dqfactor", expr.Double(dqfactor))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) DqfactorExpr(dqfactor expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("dqfactor", dqfactor)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Tfrequency(tfrequency float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("tfrequency", expr.Double(tfrequency))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) TfrequencyExpr(tfrequency expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("tfrequency", tfrequency)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Tqfactor(tqfactor float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("tqfactor", expr.Double(tqfactor))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) TqfactorExpr(tqfactor expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("tqfactor", tqfactor)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Attack(attack float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("attack", expr.Double(attack))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) AttackExpr(attack expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("attack", attack)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Release(release float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("release", expr.Double(release))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) ReleaseExpr(release expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("release", release)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Ratio(ratio float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("ratio", expr.Double(ratio))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) RatioExpr(ratio expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("ratio", ratio)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Makeup(makeup float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("makeup", expr.Double(makeup))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) MakeupExpr(makeup expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("makeup", makeup)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Range(rng float64) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("range", expr.Double(rng))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) RangeExpr(rng expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("range", rng)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Mode(mode int) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("mode", expr.Int(mode))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) ModeExpr(mode expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("mode", mode)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Dftype(dftype int) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("dftype", expr.Int(dftype))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) DftypeExpr(dftype expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("dftype", dftype)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Tftype(tftype int) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("tftype", expr.Int(tftype))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) TftypeExpr(tftype expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("tftype", tftype)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Auto(auto int) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("auto", expr.Int(auto))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) AutoExpr(auto expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("auto", auto)
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Precision(precision int) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("precision", expr.Int(precision))
}

func (adynamicequalizerBuilder *implAdynamicequalizerBuilder) Enable(enable expr.Expr) AdynamicequalizerBuilder {
	return adynamicequalizerBuilder.withOption("enable", enable)
}
