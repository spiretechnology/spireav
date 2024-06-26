// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AemphasisBuilder Audio emphasis.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#aemphasis
type AemphasisBuilder interface {
	filter.Filter
	// LevelIn set input gain (from 0 to 64) (default 1).
	LevelIn(levelIn float64) AemphasisBuilder
	// LevelInExpr set input gain (from 0 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) AemphasisBuilder
	// LevelOut set output gain (from 0 to 64) (default 1).
	LevelOut(levelOut float64) AemphasisBuilder
	// LevelOutExpr set output gain (from 0 to 64) (default 1).
	LevelOutExpr(levelOut expr.Expr) AemphasisBuilder
	// Mode set filter mode (from 0 to 1) (default reproduction).
	Mode(mode int) AemphasisBuilder
	// ModeExpr set filter mode (from 0 to 1) (default reproduction).
	ModeExpr(mode expr.Expr) AemphasisBuilder
	// Type set filter type (from 0 to 8) (default cd).
	Type(typ int) AemphasisBuilder
	// TypeExpr set filter type (from 0 to 8) (default cd).
	TypeExpr(typ expr.Expr) AemphasisBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AemphasisBuilder
}

// Aemphasis creates a new AemphasisBuilder to configure the "aemphasis" filter.
func Aemphasis(opts ...filter.Option) AemphasisBuilder {
	return &implAemphasisBuilder{
		f: filter.New("aemphasis", 1, opts...),
	}
}

type implAemphasisBuilder struct {
	f filter.Filter
}

func (aemphasisBuilder *implAemphasisBuilder) String() string {
	return aemphasisBuilder.f.String()
}

func (aemphasisBuilder *implAemphasisBuilder) Outputs() int {
	return aemphasisBuilder.f.Outputs()
}

func (aemphasisBuilder *implAemphasisBuilder) With(key string, value expr.Expr) filter.Filter {
	return aemphasisBuilder.withOption(key, value)
}

func (aemphasisBuilder *implAemphasisBuilder) withOption(key string, value expr.Expr) AemphasisBuilder {
	bCopy := *aemphasisBuilder
	bCopy.f = aemphasisBuilder.f.With(key, value)
	return &bCopy
}

func (aemphasisBuilder *implAemphasisBuilder) LevelIn(levelIn float64) AemphasisBuilder {
	return aemphasisBuilder.withOption("level_in", expr.Double(levelIn))
}

func (aemphasisBuilder *implAemphasisBuilder) LevelInExpr(levelIn expr.Expr) AemphasisBuilder {
	return aemphasisBuilder.withOption("level_in", levelIn)
}

func (aemphasisBuilder *implAemphasisBuilder) LevelOut(levelOut float64) AemphasisBuilder {
	return aemphasisBuilder.withOption("level_out", expr.Double(levelOut))
}

func (aemphasisBuilder *implAemphasisBuilder) LevelOutExpr(levelOut expr.Expr) AemphasisBuilder {
	return aemphasisBuilder.withOption("level_out", levelOut)
}

func (aemphasisBuilder *implAemphasisBuilder) Mode(mode int) AemphasisBuilder {
	return aemphasisBuilder.withOption("mode", expr.Int(mode))
}

func (aemphasisBuilder *implAemphasisBuilder) ModeExpr(mode expr.Expr) AemphasisBuilder {
	return aemphasisBuilder.withOption("mode", mode)
}

func (aemphasisBuilder *implAemphasisBuilder) Type(typ int) AemphasisBuilder {
	return aemphasisBuilder.withOption("type", expr.Int(typ))
}

func (aemphasisBuilder *implAemphasisBuilder) TypeExpr(typ expr.Expr) AemphasisBuilder {
	return aemphasisBuilder.withOption("type", typ)
}

func (aemphasisBuilder *implAemphasisBuilder) Enable(enable expr.Expr) AemphasisBuilder {
	return aemphasisBuilder.withOption("enable", enable)
}
