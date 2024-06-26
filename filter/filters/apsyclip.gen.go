// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ApsyclipBuilder Audio Psychoacoustic Clipper.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#apsyclip
type ApsyclipBuilder interface {
	filter.Filter
	// LevelIn set input level (from 0.015625 to 64) (default 1).
	LevelIn(levelIn float64) ApsyclipBuilder
	// LevelInExpr set input level (from 0.015625 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) ApsyclipBuilder
	// LevelOut set output level (from 0.015625 to 64) (default 1).
	LevelOut(levelOut float64) ApsyclipBuilder
	// LevelOutExpr set output level (from 0.015625 to 64) (default 1).
	LevelOutExpr(levelOut expr.Expr) ApsyclipBuilder
	// Clip set clip level (from 0.015625 to 1) (default 1).
	Clip(clip float64) ApsyclipBuilder
	// ClipExpr set clip level (from 0.015625 to 1) (default 1).
	ClipExpr(clip expr.Expr) ApsyclipBuilder
	// Diff enable difference (default false).
	Diff(diff bool) ApsyclipBuilder
	// DiffExpr enable difference (default false).
	DiffExpr(diff expr.Expr) ApsyclipBuilder
	// Adaptive set adaptive distortion (from 0 to 1) (default 0.5).
	Adaptive(adaptive float64) ApsyclipBuilder
	// AdaptiveExpr set adaptive distortion (from 0 to 1) (default 0.5).
	AdaptiveExpr(adaptive expr.Expr) ApsyclipBuilder
	// Iterations set iterations (from 1 to 20) (default 10).
	Iterations(iterations int) ApsyclipBuilder
	// IterationsExpr set iterations (from 1 to 20) (default 10).
	IterationsExpr(iterations expr.Expr) ApsyclipBuilder
	// Level set auto level (default false).
	Level(level bool) ApsyclipBuilder
	// LevelExpr set auto level (default false).
	LevelExpr(level expr.Expr) ApsyclipBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ApsyclipBuilder
}

// Apsyclip creates a new ApsyclipBuilder to configure the "apsyclip" filter.
func Apsyclip(opts ...filter.Option) ApsyclipBuilder {
	return &implApsyclipBuilder{
		f: filter.New("apsyclip", 1, opts...),
	}
}

type implApsyclipBuilder struct {
	f filter.Filter
}

func (apsyclipBuilder *implApsyclipBuilder) String() string {
	return apsyclipBuilder.f.String()
}

func (apsyclipBuilder *implApsyclipBuilder) Outputs() int {
	return apsyclipBuilder.f.Outputs()
}

func (apsyclipBuilder *implApsyclipBuilder) With(key string, value expr.Expr) filter.Filter {
	return apsyclipBuilder.withOption(key, value)
}

func (apsyclipBuilder *implApsyclipBuilder) withOption(key string, value expr.Expr) ApsyclipBuilder {
	bCopy := *apsyclipBuilder
	bCopy.f = apsyclipBuilder.f.With(key, value)
	return &bCopy
}

func (apsyclipBuilder *implApsyclipBuilder) LevelIn(levelIn float64) ApsyclipBuilder {
	return apsyclipBuilder.withOption("level_in", expr.Double(levelIn))
}

func (apsyclipBuilder *implApsyclipBuilder) LevelInExpr(levelIn expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("level_in", levelIn)
}

func (apsyclipBuilder *implApsyclipBuilder) LevelOut(levelOut float64) ApsyclipBuilder {
	return apsyclipBuilder.withOption("level_out", expr.Double(levelOut))
}

func (apsyclipBuilder *implApsyclipBuilder) LevelOutExpr(levelOut expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("level_out", levelOut)
}

func (apsyclipBuilder *implApsyclipBuilder) Clip(clip float64) ApsyclipBuilder {
	return apsyclipBuilder.withOption("clip", expr.Double(clip))
}

func (apsyclipBuilder *implApsyclipBuilder) ClipExpr(clip expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("clip", clip)
}

func (apsyclipBuilder *implApsyclipBuilder) Diff(diff bool) ApsyclipBuilder {
	return apsyclipBuilder.withOption("diff", expr.Bool(diff))
}

func (apsyclipBuilder *implApsyclipBuilder) DiffExpr(diff expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("diff", diff)
}

func (apsyclipBuilder *implApsyclipBuilder) Adaptive(adaptive float64) ApsyclipBuilder {
	return apsyclipBuilder.withOption("adaptive", expr.Double(adaptive))
}

func (apsyclipBuilder *implApsyclipBuilder) AdaptiveExpr(adaptive expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("adaptive", adaptive)
}

func (apsyclipBuilder *implApsyclipBuilder) Iterations(iterations int) ApsyclipBuilder {
	return apsyclipBuilder.withOption("iterations", expr.Int(iterations))
}

func (apsyclipBuilder *implApsyclipBuilder) IterationsExpr(iterations expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("iterations", iterations)
}

func (apsyclipBuilder *implApsyclipBuilder) Level(level bool) ApsyclipBuilder {
	return apsyclipBuilder.withOption("level", expr.Bool(level))
}

func (apsyclipBuilder *implApsyclipBuilder) LevelExpr(level expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("level", level)
}

func (apsyclipBuilder *implApsyclipBuilder) Enable(enable expr.Expr) ApsyclipBuilder {
	return apsyclipBuilder.withOption("enable", enable)
}
