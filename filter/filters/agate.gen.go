// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AgateBuilder Audio gate.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#agate
type AgateBuilder interface {
	filter.Filter
	// LevelIn set input level (from 0.015625 to 64) (default 1).
	LevelIn(levelIn float64) AgateBuilder
	// LevelInExpr set input level (from 0.015625 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) AgateBuilder
	// Mode set mode (from 0 to 1) (default downward).
	Mode(mode int) AgateBuilder
	// ModeExpr set mode (from 0 to 1) (default downward).
	ModeExpr(mode expr.Expr) AgateBuilder
	// Range set max gain reduction (from 0 to 1) (default 0.06125).
	Range(rng float64) AgateBuilder
	// RangeExpr set max gain reduction (from 0 to 1) (default 0.06125).
	RangeExpr(rng expr.Expr) AgateBuilder
	// Threshold set threshold (from 0 to 1) (default 0.125).
	Threshold(threshold float64) AgateBuilder
	// ThresholdExpr set threshold (from 0 to 1) (default 0.125).
	ThresholdExpr(threshold expr.Expr) AgateBuilder
	// Ratio set ratio (from 1 to 9000) (default 2).
	Ratio(ratio float64) AgateBuilder
	// RatioExpr set ratio (from 1 to 9000) (default 2).
	RatioExpr(ratio expr.Expr) AgateBuilder
	// Attack set attack (from 0.01 to 9000) (default 20).
	Attack(attack float64) AgateBuilder
	// AttackExpr set attack (from 0.01 to 9000) (default 20).
	AttackExpr(attack expr.Expr) AgateBuilder
	// Release set release (from 0.01 to 9000) (default 250).
	Release(release float64) AgateBuilder
	// ReleaseExpr set release (from 0.01 to 9000) (default 250).
	ReleaseExpr(release expr.Expr) AgateBuilder
	// Makeup set makeup gain (from 1 to 64) (default 1).
	Makeup(makeup float64) AgateBuilder
	// MakeupExpr set makeup gain (from 1 to 64) (default 1).
	MakeupExpr(makeup expr.Expr) AgateBuilder
	// Knee set knee (from 1 to 8) (default 2.82843).
	Knee(knee float64) AgateBuilder
	// KneeExpr set knee (from 1 to 8) (default 2.82843).
	KneeExpr(knee expr.Expr) AgateBuilder
	// Detection set detection (from 0 to 1) (default rms).
	Detection(detection int) AgateBuilder
	// DetectionExpr set detection (from 0 to 1) (default rms).
	DetectionExpr(detection expr.Expr) AgateBuilder
	// Link set link (from 0 to 1) (default average).
	Link(link int) AgateBuilder
	// LinkExpr set link (from 0 to 1) (default average).
	LinkExpr(link expr.Expr) AgateBuilder
	// LevelSc set sidechain gain (from 0.015625 to 64) (default 1).
	LevelSc(levelSc float64) AgateBuilder
	// LevelScExpr set sidechain gain (from 0.015625 to 64) (default 1).
	LevelScExpr(levelSc expr.Expr) AgateBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AgateBuilder
}

// Agate creates a new AgateBuilder to configure the "agate" filter.
func Agate(opts ...filter.Option) AgateBuilder {
	return &implAgateBuilder{
		f: filter.New("agate", 1, opts...),
	}
}

type implAgateBuilder struct {
	f filter.Filter
}

func (agateBuilder *implAgateBuilder) String() string {
	return agateBuilder.f.String()
}

func (agateBuilder *implAgateBuilder) Outputs() int {
	return agateBuilder.f.Outputs()
}

func (agateBuilder *implAgateBuilder) With(key string, value expr.Expr) filter.Filter {
	return agateBuilder.withOption(key, value)
}

func (agateBuilder *implAgateBuilder) withOption(key string, value expr.Expr) AgateBuilder {
	bCopy := *agateBuilder
	bCopy.f = agateBuilder.f.With(key, value)
	return &bCopy
}

func (agateBuilder *implAgateBuilder) LevelIn(levelIn float64) AgateBuilder {
	return agateBuilder.withOption("level_in", expr.Double(levelIn))
}

func (agateBuilder *implAgateBuilder) LevelInExpr(levelIn expr.Expr) AgateBuilder {
	return agateBuilder.withOption("level_in", levelIn)
}

func (agateBuilder *implAgateBuilder) Mode(mode int) AgateBuilder {
	return agateBuilder.withOption("mode", expr.Int(mode))
}

func (agateBuilder *implAgateBuilder) ModeExpr(mode expr.Expr) AgateBuilder {
	return agateBuilder.withOption("mode", mode)
}

func (agateBuilder *implAgateBuilder) Range(rng float64) AgateBuilder {
	return agateBuilder.withOption("range", expr.Double(rng))
}

func (agateBuilder *implAgateBuilder) RangeExpr(rng expr.Expr) AgateBuilder {
	return agateBuilder.withOption("range", rng)
}

func (agateBuilder *implAgateBuilder) Threshold(threshold float64) AgateBuilder {
	return agateBuilder.withOption("threshold", expr.Double(threshold))
}

func (agateBuilder *implAgateBuilder) ThresholdExpr(threshold expr.Expr) AgateBuilder {
	return agateBuilder.withOption("threshold", threshold)
}

func (agateBuilder *implAgateBuilder) Ratio(ratio float64) AgateBuilder {
	return agateBuilder.withOption("ratio", expr.Double(ratio))
}

func (agateBuilder *implAgateBuilder) RatioExpr(ratio expr.Expr) AgateBuilder {
	return agateBuilder.withOption("ratio", ratio)
}

func (agateBuilder *implAgateBuilder) Attack(attack float64) AgateBuilder {
	return agateBuilder.withOption("attack", expr.Double(attack))
}

func (agateBuilder *implAgateBuilder) AttackExpr(attack expr.Expr) AgateBuilder {
	return agateBuilder.withOption("attack", attack)
}

func (agateBuilder *implAgateBuilder) Release(release float64) AgateBuilder {
	return agateBuilder.withOption("release", expr.Double(release))
}

func (agateBuilder *implAgateBuilder) ReleaseExpr(release expr.Expr) AgateBuilder {
	return agateBuilder.withOption("release", release)
}

func (agateBuilder *implAgateBuilder) Makeup(makeup float64) AgateBuilder {
	return agateBuilder.withOption("makeup", expr.Double(makeup))
}

func (agateBuilder *implAgateBuilder) MakeupExpr(makeup expr.Expr) AgateBuilder {
	return agateBuilder.withOption("makeup", makeup)
}

func (agateBuilder *implAgateBuilder) Knee(knee float64) AgateBuilder {
	return agateBuilder.withOption("knee", expr.Double(knee))
}

func (agateBuilder *implAgateBuilder) KneeExpr(knee expr.Expr) AgateBuilder {
	return agateBuilder.withOption("knee", knee)
}

func (agateBuilder *implAgateBuilder) Detection(detection int) AgateBuilder {
	return agateBuilder.withOption("detection", expr.Int(detection))
}

func (agateBuilder *implAgateBuilder) DetectionExpr(detection expr.Expr) AgateBuilder {
	return agateBuilder.withOption("detection", detection)
}

func (agateBuilder *implAgateBuilder) Link(link int) AgateBuilder {
	return agateBuilder.withOption("link", expr.Int(link))
}

func (agateBuilder *implAgateBuilder) LinkExpr(link expr.Expr) AgateBuilder {
	return agateBuilder.withOption("link", link)
}

func (agateBuilder *implAgateBuilder) LevelSc(levelSc float64) AgateBuilder {
	return agateBuilder.withOption("level_sc", expr.Double(levelSc))
}

func (agateBuilder *implAgateBuilder) LevelScExpr(levelSc expr.Expr) AgateBuilder {
	return agateBuilder.withOption("level_sc", levelSc)
}

func (agateBuilder *implAgateBuilder) Enable(enable expr.Expr) AgateBuilder {
	return agateBuilder.withOption("enable", enable)
}
