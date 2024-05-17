// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// GraphmonitorBuilder Show various filtergraph stats.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#graphmonitor
type GraphmonitorBuilder interface {
	filter.Filter
	// Size set monitor size (default "hd720").
	Size(size expr.Size) GraphmonitorBuilder
	// S set monitor size (default "hd720").
	S(s expr.Size) GraphmonitorBuilder
	// Opacity set video opacity (from 0 to 1) (default 0.9).
	Opacity(opacity float32) GraphmonitorBuilder
	// OpacityExpr set video opacity (from 0 to 1) (default 0.9).
	OpacityExpr(opacity expr.Expr) GraphmonitorBuilder
	// O set video opacity (from 0 to 1) (default 0.9).
	O(o float32) GraphmonitorBuilder
	// OExpr set video opacity (from 0 to 1) (default 0.9).
	OExpr(o expr.Expr) GraphmonitorBuilder
	// Mode set mode (default 0).
	Mode(mode ...string) GraphmonitorBuilder
	// ModeExpr set mode (default 0).
	ModeExpr(mode expr.Expr) GraphmonitorBuilder
	// M set mode (default 0).
	M(m ...string) GraphmonitorBuilder
	// MExpr set mode (default 0).
	MExpr(m expr.Expr) GraphmonitorBuilder
	// Flags set flags (default all+queue).
	Flags(flags ...string) GraphmonitorBuilder
	// FlagsExpr set flags (default all+queue).
	FlagsExpr(flags expr.Expr) GraphmonitorBuilder
	// F set flags (default all+queue).
	F(f ...string) GraphmonitorBuilder
	// FExpr set flags (default all+queue).
	FExpr(f expr.Expr) GraphmonitorBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) GraphmonitorBuilder
	// R set video rate (default "25").
	R(r expr.Rational) GraphmonitorBuilder
}

// Graphmonitor creates a new GraphmonitorBuilder to configure the "graphmonitor" filter.
func Graphmonitor(opts ...filter.Option) GraphmonitorBuilder {
	return &implGraphmonitorBuilder{
		f: filter.New("graphmonitor", 1, opts...),
	}
}

type implGraphmonitorBuilder struct {
	f filter.Filter
}

func (graphmonitorBuilder *implGraphmonitorBuilder) String() string {
	return graphmonitorBuilder.f.String()
}

func (graphmonitorBuilder *implGraphmonitorBuilder) Outputs() int {
	return graphmonitorBuilder.f.Outputs()
}

func (graphmonitorBuilder *implGraphmonitorBuilder) With(key string, value expr.Expr) filter.Filter {
	return graphmonitorBuilder.withOption(key, value)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) withOption(key string, value expr.Expr) GraphmonitorBuilder {
	bCopy := *graphmonitorBuilder
	bCopy.f = graphmonitorBuilder.f.With(key, value)
	return &bCopy
}

func (graphmonitorBuilder *implGraphmonitorBuilder) Size(size expr.Size) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("size", size)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) S(s expr.Size) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("s", s)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) Opacity(opacity float32) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("opacity", expr.Float(opacity))
}

func (graphmonitorBuilder *implGraphmonitorBuilder) OpacityExpr(opacity expr.Expr) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("opacity", opacity)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) O(o float32) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("o", expr.Float(o))
}

func (graphmonitorBuilder *implGraphmonitorBuilder) OExpr(o expr.Expr) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("o", o)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) Mode(mode ...string) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("mode", expr.Flags(mode))
}

func (graphmonitorBuilder *implGraphmonitorBuilder) ModeExpr(mode expr.Expr) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("mode", mode)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) M(m ...string) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("m", expr.Flags(m))
}

func (graphmonitorBuilder *implGraphmonitorBuilder) MExpr(m expr.Expr) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("m", m)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) Flags(flags ...string) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("flags", expr.Flags(flags))
}

func (graphmonitorBuilder *implGraphmonitorBuilder) FlagsExpr(flags expr.Expr) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("flags", flags)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) F(f ...string) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("f", expr.Flags(f))
}

func (graphmonitorBuilder *implGraphmonitorBuilder) FExpr(f expr.Expr) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("f", f)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) Rate(rate expr.Rational) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("rate", rate)
}

func (graphmonitorBuilder *implGraphmonitorBuilder) R(r expr.Rational) GraphmonitorBuilder {
	return graphmonitorBuilder.withOption("r", r)
}
