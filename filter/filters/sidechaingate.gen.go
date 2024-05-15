// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SidechaingateBuilder Audio sidechain gate.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#sidechaingate
type SidechaingateBuilder interface {
	filter.Filter
	// LevelIn set input level (from 0.015625 to 64) (default 1).
	LevelIn(levelIn float64) SidechaingateBuilder
	// LevelInExpr set input level (from 0.015625 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) SidechaingateBuilder
	// Mode set mode (from 0 to 1) (default downward).
	Mode(mode int) SidechaingateBuilder
	// ModeExpr set mode (from 0 to 1) (default downward).
	ModeExpr(mode expr.Expr) SidechaingateBuilder
	// Range set max gain reduction (from 0 to 1) (default 0.06125).
	Range(rng float64) SidechaingateBuilder
	// RangeExpr set max gain reduction (from 0 to 1) (default 0.06125).
	RangeExpr(rng expr.Expr) SidechaingateBuilder
	// Threshold set threshold (from 0 to 1) (default 0.125).
	Threshold(threshold float64) SidechaingateBuilder
	// ThresholdExpr set threshold (from 0 to 1) (default 0.125).
	ThresholdExpr(threshold expr.Expr) SidechaingateBuilder
	// Ratio set ratio (from 1 to 9000) (default 2).
	Ratio(ratio float64) SidechaingateBuilder
	// RatioExpr set ratio (from 1 to 9000) (default 2).
	RatioExpr(ratio expr.Expr) SidechaingateBuilder
	// Attack set attack (from 0.01 to 9000) (default 20).
	Attack(attack float64) SidechaingateBuilder
	// AttackExpr set attack (from 0.01 to 9000) (default 20).
	AttackExpr(attack expr.Expr) SidechaingateBuilder
	// Release set release (from 0.01 to 9000) (default 250).
	Release(release float64) SidechaingateBuilder
	// ReleaseExpr set release (from 0.01 to 9000) (default 250).
	ReleaseExpr(release expr.Expr) SidechaingateBuilder
	// Makeup set makeup gain (from 1 to 64) (default 1).
	Makeup(makeup float64) SidechaingateBuilder
	// MakeupExpr set makeup gain (from 1 to 64) (default 1).
	MakeupExpr(makeup expr.Expr) SidechaingateBuilder
	// Knee set knee (from 1 to 8) (default 2.82843).
	Knee(knee float64) SidechaingateBuilder
	// KneeExpr set knee (from 1 to 8) (default 2.82843).
	KneeExpr(knee expr.Expr) SidechaingateBuilder
	// Detection set detection (from 0 to 1) (default rms).
	Detection(detection int) SidechaingateBuilder
	// DetectionExpr set detection (from 0 to 1) (default rms).
	DetectionExpr(detection expr.Expr) SidechaingateBuilder
	// Link set link (from 0 to 1) (default average).
	Link(link int) SidechaingateBuilder
	// LinkExpr set link (from 0 to 1) (default average).
	LinkExpr(link expr.Expr) SidechaingateBuilder
	// LevelSc set sidechain gain (from 0.015625 to 64) (default 1).
	LevelSc(levelSc float64) SidechaingateBuilder
	// LevelScExpr set sidechain gain (from 0.015625 to 64) (default 1).
	LevelScExpr(levelSc expr.Expr) SidechaingateBuilder
}

// Sidechaingate creates a new SidechaingateBuilder to configure the "sidechaingate" filter.
func Sidechaingate(opts ...filter.Option) SidechaingateBuilder {
	return &implSidechaingateBuilder{
		f: filter.New("sidechaingate", 1, opts...),
	}
}

type implSidechaingateBuilder struct {
	f filter.Filter
}

func (sidechaingateBuilder *implSidechaingateBuilder) String() string {
	return sidechaingateBuilder.f.String()
}

func (sidechaingateBuilder *implSidechaingateBuilder) Outputs() int {
	return sidechaingateBuilder.f.Outputs()
}

func (sidechaingateBuilder *implSidechaingateBuilder) With(key string, value expr.Expr) filter.Filter {
	return sidechaingateBuilder.withOption(key, value)
}

func (sidechaingateBuilder *implSidechaingateBuilder) withOption(key string, value expr.Expr) SidechaingateBuilder {
	bCopy := *sidechaingateBuilder
	bCopy.f = sidechaingateBuilder.f.With(key, value)
	return &bCopy
}

func (sidechaingateBuilder *implSidechaingateBuilder) LevelIn(levelIn float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("level_in", expr.Double(levelIn))
}

func (sidechaingateBuilder *implSidechaingateBuilder) LevelInExpr(levelIn expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("level_in", levelIn)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Mode(mode int) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("mode", expr.Int(mode))
}

func (sidechaingateBuilder *implSidechaingateBuilder) ModeExpr(mode expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("mode", mode)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Range(rng float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("range", expr.Double(rng))
}

func (sidechaingateBuilder *implSidechaingateBuilder) RangeExpr(rng expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("range", rng)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Threshold(threshold float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("threshold", expr.Double(threshold))
}

func (sidechaingateBuilder *implSidechaingateBuilder) ThresholdExpr(threshold expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("threshold", threshold)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Ratio(ratio float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("ratio", expr.Double(ratio))
}

func (sidechaingateBuilder *implSidechaingateBuilder) RatioExpr(ratio expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("ratio", ratio)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Attack(attack float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("attack", expr.Double(attack))
}

func (sidechaingateBuilder *implSidechaingateBuilder) AttackExpr(attack expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("attack", attack)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Release(release float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("release", expr.Double(release))
}

func (sidechaingateBuilder *implSidechaingateBuilder) ReleaseExpr(release expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("release", release)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Makeup(makeup float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("makeup", expr.Double(makeup))
}

func (sidechaingateBuilder *implSidechaingateBuilder) MakeupExpr(makeup expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("makeup", makeup)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Knee(knee float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("knee", expr.Double(knee))
}

func (sidechaingateBuilder *implSidechaingateBuilder) KneeExpr(knee expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("knee", knee)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Detection(detection int) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("detection", expr.Int(detection))
}

func (sidechaingateBuilder *implSidechaingateBuilder) DetectionExpr(detection expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("detection", detection)
}

func (sidechaingateBuilder *implSidechaingateBuilder) Link(link int) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("link", expr.Int(link))
}

func (sidechaingateBuilder *implSidechaingateBuilder) LinkExpr(link expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("link", link)
}

func (sidechaingateBuilder *implSidechaingateBuilder) LevelSc(levelSc float64) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("level_sc", expr.Double(levelSc))
}

func (sidechaingateBuilder *implSidechaingateBuilder) LevelScExpr(levelSc expr.Expr) SidechaingateBuilder {
	return sidechaingateBuilder.withOption("level_sc", levelSc)
}
