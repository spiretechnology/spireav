// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AcompressorBuilder Audio compressor.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#acompressor
type AcompressorBuilder interface {
	filter.Filter
	// LevelIn set input gain (from 0.015625 to 64) (default 1).
	LevelIn(levelIn float64) AcompressorBuilder
	// LevelInExpr set input gain (from 0.015625 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) AcompressorBuilder
	// Mode set mode (from 0 to 1) (default downward).
	Mode(mode int) AcompressorBuilder
	// ModeExpr set mode (from 0 to 1) (default downward).
	ModeExpr(mode expr.Expr) AcompressorBuilder
	// Threshold set threshold (from 0.000976563 to 1) (default 0.125).
	Threshold(threshold float64) AcompressorBuilder
	// ThresholdExpr set threshold (from 0.000976563 to 1) (default 0.125).
	ThresholdExpr(threshold expr.Expr) AcompressorBuilder
	// Ratio set ratio (from 1 to 20) (default 2).
	Ratio(ratio float64) AcompressorBuilder
	// RatioExpr set ratio (from 1 to 20) (default 2).
	RatioExpr(ratio expr.Expr) AcompressorBuilder
	// Attack set attack (from 0.01 to 2000) (default 20).
	Attack(attack float64) AcompressorBuilder
	// AttackExpr set attack (from 0.01 to 2000) (default 20).
	AttackExpr(attack expr.Expr) AcompressorBuilder
	// Release set release (from 0.01 to 9000) (default 250).
	Release(release float64) AcompressorBuilder
	// ReleaseExpr set release (from 0.01 to 9000) (default 250).
	ReleaseExpr(release expr.Expr) AcompressorBuilder
	// Makeup set make up gain (from 1 to 64) (default 1).
	Makeup(makeup float64) AcompressorBuilder
	// MakeupExpr set make up gain (from 1 to 64) (default 1).
	MakeupExpr(makeup expr.Expr) AcompressorBuilder
	// Knee set knee (from 1 to 8) (default 2.82843).
	Knee(knee float64) AcompressorBuilder
	// KneeExpr set knee (from 1 to 8) (default 2.82843).
	KneeExpr(knee expr.Expr) AcompressorBuilder
	// Link set link type (from 0 to 1) (default average).
	Link(link int) AcompressorBuilder
	// LinkExpr set link type (from 0 to 1) (default average).
	LinkExpr(link expr.Expr) AcompressorBuilder
	// Detection set detection (from 0 to 1) (default rms).
	Detection(detection int) AcompressorBuilder
	// DetectionExpr set detection (from 0 to 1) (default rms).
	DetectionExpr(detection expr.Expr) AcompressorBuilder
	// LevelSc set sidechain gain (from 0.015625 to 64) (default 1).
	LevelSc(levelSc float64) AcompressorBuilder
	// LevelScExpr set sidechain gain (from 0.015625 to 64) (default 1).
	LevelScExpr(levelSc expr.Expr) AcompressorBuilder
	// Mix set mix (from 0 to 1) (default 1).
	Mix(mix float64) AcompressorBuilder
	// MixExpr set mix (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) AcompressorBuilder
}

// Acompressor creates a new AcompressorBuilder to configure the "acompressor" filter.
func Acompressor(opts ...filter.Option) AcompressorBuilder {
	return &implAcompressorBuilder{
		f: filter.New("acompressor", 1, opts...),
	}
}

type implAcompressorBuilder struct {
	f filter.Filter
}

func (acompressorBuilder *implAcompressorBuilder) String() string {
	return acompressorBuilder.f.String()
}

func (acompressorBuilder *implAcompressorBuilder) Outputs() int {
	return acompressorBuilder.f.Outputs()
}

func (acompressorBuilder *implAcompressorBuilder) With(key string, value expr.Expr) filter.Filter {
	return acompressorBuilder.withOption(key, value)
}

func (acompressorBuilder *implAcompressorBuilder) withOption(key string, value expr.Expr) AcompressorBuilder {
	bCopy := *acompressorBuilder
	bCopy.f = acompressorBuilder.f.With(key, value)
	return &bCopy
}

func (acompressorBuilder *implAcompressorBuilder) LevelIn(levelIn float64) AcompressorBuilder {
	return acompressorBuilder.withOption("level_in", expr.Double(levelIn))
}

func (acompressorBuilder *implAcompressorBuilder) LevelInExpr(levelIn expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("level_in", levelIn)
}

func (acompressorBuilder *implAcompressorBuilder) Mode(mode int) AcompressorBuilder {
	return acompressorBuilder.withOption("mode", expr.Int(mode))
}

func (acompressorBuilder *implAcompressorBuilder) ModeExpr(mode expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("mode", mode)
}

func (acompressorBuilder *implAcompressorBuilder) Threshold(threshold float64) AcompressorBuilder {
	return acompressorBuilder.withOption("threshold", expr.Double(threshold))
}

func (acompressorBuilder *implAcompressorBuilder) ThresholdExpr(threshold expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("threshold", threshold)
}

func (acompressorBuilder *implAcompressorBuilder) Ratio(ratio float64) AcompressorBuilder {
	return acompressorBuilder.withOption("ratio", expr.Double(ratio))
}

func (acompressorBuilder *implAcompressorBuilder) RatioExpr(ratio expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("ratio", ratio)
}

func (acompressorBuilder *implAcompressorBuilder) Attack(attack float64) AcompressorBuilder {
	return acompressorBuilder.withOption("attack", expr.Double(attack))
}

func (acompressorBuilder *implAcompressorBuilder) AttackExpr(attack expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("attack", attack)
}

func (acompressorBuilder *implAcompressorBuilder) Release(release float64) AcompressorBuilder {
	return acompressorBuilder.withOption("release", expr.Double(release))
}

func (acompressorBuilder *implAcompressorBuilder) ReleaseExpr(release expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("release", release)
}

func (acompressorBuilder *implAcompressorBuilder) Makeup(makeup float64) AcompressorBuilder {
	return acompressorBuilder.withOption("makeup", expr.Double(makeup))
}

func (acompressorBuilder *implAcompressorBuilder) MakeupExpr(makeup expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("makeup", makeup)
}

func (acompressorBuilder *implAcompressorBuilder) Knee(knee float64) AcompressorBuilder {
	return acompressorBuilder.withOption("knee", expr.Double(knee))
}

func (acompressorBuilder *implAcompressorBuilder) KneeExpr(knee expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("knee", knee)
}

func (acompressorBuilder *implAcompressorBuilder) Link(link int) AcompressorBuilder {
	return acompressorBuilder.withOption("link", expr.Int(link))
}

func (acompressorBuilder *implAcompressorBuilder) LinkExpr(link expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("link", link)
}

func (acompressorBuilder *implAcompressorBuilder) Detection(detection int) AcompressorBuilder {
	return acompressorBuilder.withOption("detection", expr.Int(detection))
}

func (acompressorBuilder *implAcompressorBuilder) DetectionExpr(detection expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("detection", detection)
}

func (acompressorBuilder *implAcompressorBuilder) LevelSc(levelSc float64) AcompressorBuilder {
	return acompressorBuilder.withOption("level_sc", expr.Double(levelSc))
}

func (acompressorBuilder *implAcompressorBuilder) LevelScExpr(levelSc expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("level_sc", levelSc)
}

func (acompressorBuilder *implAcompressorBuilder) Mix(mix float64) AcompressorBuilder {
	return acompressorBuilder.withOption("mix", expr.Double(mix))
}

func (acompressorBuilder *implAcompressorBuilder) MixExpr(mix expr.Expr) AcompressorBuilder {
	return acompressorBuilder.withOption("mix", mix)
}