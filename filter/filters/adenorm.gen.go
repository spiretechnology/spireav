// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AdenormBuilder Remedy denormals by adding extremely low-level noise.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#adenorm
type AdenormBuilder interface {
	filter.Filter
	// Level set level (from -451 to -90) (default -351).
	Level(level float64) AdenormBuilder
	// LevelExpr set level (from -451 to -90) (default -351).
	LevelExpr(level expr.Expr) AdenormBuilder
	// Type set type (from 0 to 3) (default dc).
	Type(typ int) AdenormBuilder
	// TypeExpr set type (from 0 to 3) (default dc).
	TypeExpr(typ expr.Expr) AdenormBuilder
}

// Adenorm creates a new AdenormBuilder to configure the "adenorm" filter.
func Adenorm(opts ...filter.Option) AdenormBuilder {
	return &implAdenormBuilder{
		f: filter.New("adenorm", 1, opts...),
	}
}

type implAdenormBuilder struct {
	f filter.Filter
}

func (adenormBuilder *implAdenormBuilder) String() string {
	return adenormBuilder.f.String()
}

func (adenormBuilder *implAdenormBuilder) Outputs() int {
	return adenormBuilder.f.Outputs()
}

func (adenormBuilder *implAdenormBuilder) With(key string, value expr.Expr) filter.Filter {
	return adenormBuilder.withOption(key, value)
}

func (adenormBuilder *implAdenormBuilder) withOption(key string, value expr.Expr) AdenormBuilder {
	bCopy := *adenormBuilder
	bCopy.f = adenormBuilder.f.With(key, value)
	return &bCopy
}

func (adenormBuilder *implAdenormBuilder) Level(level float64) AdenormBuilder {
	return adenormBuilder.withOption("level", expr.Double(level))
}

func (adenormBuilder *implAdenormBuilder) LevelExpr(level expr.Expr) AdenormBuilder {
	return adenormBuilder.withOption("level", level)
}

func (adenormBuilder *implAdenormBuilder) Type(typ int) AdenormBuilder {
	return adenormBuilder.withOption("type", expr.Int(typ))
}

func (adenormBuilder *implAdenormBuilder) TypeExpr(typ expr.Expr) AdenormBuilder {
	return adenormBuilder.withOption("type", typ)
}
