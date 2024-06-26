// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CrossfeedBuilder Apply headphone crossfeed filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#crossfeed
type CrossfeedBuilder interface {
	filter.Filter
	// Strength set crossfeed strength (from 0 to 1) (default 0.2).
	Strength(strength float64) CrossfeedBuilder
	// StrengthExpr set crossfeed strength (from 0 to 1) (default 0.2).
	StrengthExpr(strength expr.Expr) CrossfeedBuilder
	// Range set soundstage wideness (from 0 to 1) (default 0.5).
	Range(rng float64) CrossfeedBuilder
	// RangeExpr set soundstage wideness (from 0 to 1) (default 0.5).
	RangeExpr(rng expr.Expr) CrossfeedBuilder
	// Slope set curve slope (from 0.01 to 1) (default 0.5).
	Slope(slope float64) CrossfeedBuilder
	// SlopeExpr set curve slope (from 0.01 to 1) (default 0.5).
	SlopeExpr(slope expr.Expr) CrossfeedBuilder
	// LevelIn set level in (from 0 to 1) (default 0.9).
	LevelIn(levelIn float64) CrossfeedBuilder
	// LevelInExpr set level in (from 0 to 1) (default 0.9).
	LevelInExpr(levelIn expr.Expr) CrossfeedBuilder
	// LevelOut set level out (from 0 to 1) (default 1).
	LevelOut(levelOut float64) CrossfeedBuilder
	// LevelOutExpr set level out (from 0 to 1) (default 1).
	LevelOutExpr(levelOut expr.Expr) CrossfeedBuilder
	// BlockSize set the block size (from 0 to 32768) (default 0).
	BlockSize(blockSize int) CrossfeedBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) CrossfeedBuilder
}

// Crossfeed creates a new CrossfeedBuilder to configure the "crossfeed" filter.
func Crossfeed(opts ...filter.Option) CrossfeedBuilder {
	return &implCrossfeedBuilder{
		f: filter.New("crossfeed", 1, opts...),
	}
}

type implCrossfeedBuilder struct {
	f filter.Filter
}

func (crossfeedBuilder *implCrossfeedBuilder) String() string {
	return crossfeedBuilder.f.String()
}

func (crossfeedBuilder *implCrossfeedBuilder) Outputs() int {
	return crossfeedBuilder.f.Outputs()
}

func (crossfeedBuilder *implCrossfeedBuilder) With(key string, value expr.Expr) filter.Filter {
	return crossfeedBuilder.withOption(key, value)
}

func (crossfeedBuilder *implCrossfeedBuilder) withOption(key string, value expr.Expr) CrossfeedBuilder {
	bCopy := *crossfeedBuilder
	bCopy.f = crossfeedBuilder.f.With(key, value)
	return &bCopy
}

func (crossfeedBuilder *implCrossfeedBuilder) Strength(strength float64) CrossfeedBuilder {
	return crossfeedBuilder.withOption("strength", expr.Double(strength))
}

func (crossfeedBuilder *implCrossfeedBuilder) StrengthExpr(strength expr.Expr) CrossfeedBuilder {
	return crossfeedBuilder.withOption("strength", strength)
}

func (crossfeedBuilder *implCrossfeedBuilder) Range(rng float64) CrossfeedBuilder {
	return crossfeedBuilder.withOption("range", expr.Double(rng))
}

func (crossfeedBuilder *implCrossfeedBuilder) RangeExpr(rng expr.Expr) CrossfeedBuilder {
	return crossfeedBuilder.withOption("range", rng)
}

func (crossfeedBuilder *implCrossfeedBuilder) Slope(slope float64) CrossfeedBuilder {
	return crossfeedBuilder.withOption("slope", expr.Double(slope))
}

func (crossfeedBuilder *implCrossfeedBuilder) SlopeExpr(slope expr.Expr) CrossfeedBuilder {
	return crossfeedBuilder.withOption("slope", slope)
}

func (crossfeedBuilder *implCrossfeedBuilder) LevelIn(levelIn float64) CrossfeedBuilder {
	return crossfeedBuilder.withOption("level_in", expr.Double(levelIn))
}

func (crossfeedBuilder *implCrossfeedBuilder) LevelInExpr(levelIn expr.Expr) CrossfeedBuilder {
	return crossfeedBuilder.withOption("level_in", levelIn)
}

func (crossfeedBuilder *implCrossfeedBuilder) LevelOut(levelOut float64) CrossfeedBuilder {
	return crossfeedBuilder.withOption("level_out", expr.Double(levelOut))
}

func (crossfeedBuilder *implCrossfeedBuilder) LevelOutExpr(levelOut expr.Expr) CrossfeedBuilder {
	return crossfeedBuilder.withOption("level_out", levelOut)
}

func (crossfeedBuilder *implCrossfeedBuilder) BlockSize(blockSize int) CrossfeedBuilder {
	return crossfeedBuilder.withOption("block_size", expr.Int(blockSize))
}

func (crossfeedBuilder *implCrossfeedBuilder) Enable(enable expr.Expr) CrossfeedBuilder {
	return crossfeedBuilder.withOption("enable", enable)
}
