// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AllpassBuilder Apply a two-pole all-pass filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#allpass
type AllpassBuilder interface {
	filter.Filter
	// Frequency set central frequency (from 0 to 999999) (default 3000).
	Frequency(frequency float64) AllpassBuilder
	// FrequencyExpr set central frequency (from 0 to 999999) (default 3000).
	FrequencyExpr(frequency expr.Expr) AllpassBuilder
	// F set central frequency (from 0 to 999999) (default 3000).
	F(f float64) AllpassBuilder
	// FExpr set central frequency (from 0 to 999999) (default 3000).
	FExpr(f expr.Expr) AllpassBuilder
	// WidthType set filter-width type (from 1 to 5) (default q).
	WidthType(widthType int) AllpassBuilder
	// WidthTypeExpr set filter-width type (from 1 to 5) (default q).
	WidthTypeExpr(widthType expr.Expr) AllpassBuilder
	// T set filter-width type (from 1 to 5) (default q).
	T(t int) AllpassBuilder
	// TExpr set filter-width type (from 1 to 5) (default q).
	TExpr(t expr.Expr) AllpassBuilder
	// Width set width (from 0 to 99999) (default 0.707).
	Width(width float64) AllpassBuilder
	// WidthExpr set width (from 0 to 99999) (default 0.707).
	WidthExpr(width expr.Expr) AllpassBuilder
	// W set width (from 0 to 99999) (default 0.707).
	W(w float64) AllpassBuilder
	// WExpr set width (from 0 to 99999) (default 0.707).
	WExpr(w expr.Expr) AllpassBuilder
	// Mix set mix (from 0 to 1) (default 1).
	Mix(mix float64) AllpassBuilder
	// MixExpr set mix (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) AllpassBuilder
	// M set mix (from 0 to 1) (default 1).
	M(m float64) AllpassBuilder
	// MExpr set mix (from 0 to 1) (default 1).
	MExpr(m expr.Expr) AllpassBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) AllpassBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) AllpassBuilder
	// C set channels to filter (default "all").
	C(c string) AllpassBuilder
	// CExpr set channels to filter (default "all").
	CExpr(c expr.Expr) AllpassBuilder
	// Normalize normalize coefficients (default false).
	Normalize(normalize bool) AllpassBuilder
	// NormalizeExpr normalize coefficients (default false).
	NormalizeExpr(normalize expr.Expr) AllpassBuilder
	// N normalize coefficients (default false).
	N(n bool) AllpassBuilder
	// NExpr normalize coefficients (default false).
	NExpr(n expr.Expr) AllpassBuilder
	// Order set filter order (from 1 to 2) (default 2).
	Order(order int) AllpassBuilder
	// OrderExpr set filter order (from 1 to 2) (default 2).
	OrderExpr(order expr.Expr) AllpassBuilder
	// O set filter order (from 1 to 2) (default 2).
	O(o int) AllpassBuilder
	// OExpr set filter order (from 1 to 2) (default 2).
	OExpr(o expr.Expr) AllpassBuilder
	// Transform set transform type (from 0 to 6) (default di).
	Transform(transform int) AllpassBuilder
	// A set transform type (from 0 to 6) (default di).
	A(a int) AllpassBuilder
	// Precision set filtering precision (from -1 to 3) (default auto).
	Precision(precision int) AllpassBuilder
	// R set filtering precision (from -1 to 3) (default auto).
	R(r int) AllpassBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AllpassBuilder
}

// Allpass creates a new AllpassBuilder to configure the "allpass" filter.
func Allpass(opts ...filter.Option) AllpassBuilder {
	return &implAllpassBuilder{
		f: filter.New("allpass", 1, opts...),
	}
}

type implAllpassBuilder struct {
	f filter.Filter
}

func (allpassBuilder *implAllpassBuilder) String() string {
	return allpassBuilder.f.String()
}

func (allpassBuilder *implAllpassBuilder) Outputs() int {
	return allpassBuilder.f.Outputs()
}

func (allpassBuilder *implAllpassBuilder) With(key string, value expr.Expr) filter.Filter {
	return allpassBuilder.withOption(key, value)
}

func (allpassBuilder *implAllpassBuilder) withOption(key string, value expr.Expr) AllpassBuilder {
	bCopy := *allpassBuilder
	bCopy.f = allpassBuilder.f.With(key, value)
	return &bCopy
}

func (allpassBuilder *implAllpassBuilder) Frequency(frequency float64) AllpassBuilder {
	return allpassBuilder.withOption("frequency", expr.Double(frequency))
}

func (allpassBuilder *implAllpassBuilder) FrequencyExpr(frequency expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("frequency", frequency)
}

func (allpassBuilder *implAllpassBuilder) F(f float64) AllpassBuilder {
	return allpassBuilder.withOption("f", expr.Double(f))
}

func (allpassBuilder *implAllpassBuilder) FExpr(f expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("f", f)
}

func (allpassBuilder *implAllpassBuilder) WidthType(widthType int) AllpassBuilder {
	return allpassBuilder.withOption("width_type", expr.Int(widthType))
}

func (allpassBuilder *implAllpassBuilder) WidthTypeExpr(widthType expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("width_type", widthType)
}

func (allpassBuilder *implAllpassBuilder) T(t int) AllpassBuilder {
	return allpassBuilder.withOption("t", expr.Int(t))
}

func (allpassBuilder *implAllpassBuilder) TExpr(t expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("t", t)
}

func (allpassBuilder *implAllpassBuilder) Width(width float64) AllpassBuilder {
	return allpassBuilder.withOption("width", expr.Double(width))
}

func (allpassBuilder *implAllpassBuilder) WidthExpr(width expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("width", width)
}

func (allpassBuilder *implAllpassBuilder) W(w float64) AllpassBuilder {
	return allpassBuilder.withOption("w", expr.Double(w))
}

func (allpassBuilder *implAllpassBuilder) WExpr(w expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("w", w)
}

func (allpassBuilder *implAllpassBuilder) Mix(mix float64) AllpassBuilder {
	return allpassBuilder.withOption("mix", expr.Double(mix))
}

func (allpassBuilder *implAllpassBuilder) MixExpr(mix expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("mix", mix)
}

func (allpassBuilder *implAllpassBuilder) M(m float64) AllpassBuilder {
	return allpassBuilder.withOption("m", expr.Double(m))
}

func (allpassBuilder *implAllpassBuilder) MExpr(m expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("m", m)
}

func (allpassBuilder *implAllpassBuilder) Channels(channels string) AllpassBuilder {
	return allpassBuilder.withOption("channels", expr.String(channels))
}

func (allpassBuilder *implAllpassBuilder) ChannelsExpr(channels expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("channels", channels)
}

func (allpassBuilder *implAllpassBuilder) C(c string) AllpassBuilder {
	return allpassBuilder.withOption("c", expr.String(c))
}

func (allpassBuilder *implAllpassBuilder) CExpr(c expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("c", c)
}

func (allpassBuilder *implAllpassBuilder) Normalize(normalize bool) AllpassBuilder {
	return allpassBuilder.withOption("normalize", expr.Bool(normalize))
}

func (allpassBuilder *implAllpassBuilder) NormalizeExpr(normalize expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("normalize", normalize)
}

func (allpassBuilder *implAllpassBuilder) N(n bool) AllpassBuilder {
	return allpassBuilder.withOption("n", expr.Bool(n))
}

func (allpassBuilder *implAllpassBuilder) NExpr(n expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("n", n)
}

func (allpassBuilder *implAllpassBuilder) Order(order int) AllpassBuilder {
	return allpassBuilder.withOption("order", expr.Int(order))
}

func (allpassBuilder *implAllpassBuilder) OrderExpr(order expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("order", order)
}

func (allpassBuilder *implAllpassBuilder) O(o int) AllpassBuilder {
	return allpassBuilder.withOption("o", expr.Int(o))
}

func (allpassBuilder *implAllpassBuilder) OExpr(o expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("o", o)
}

func (allpassBuilder *implAllpassBuilder) Transform(transform int) AllpassBuilder {
	return allpassBuilder.withOption("transform", expr.Int(transform))
}

func (allpassBuilder *implAllpassBuilder) A(a int) AllpassBuilder {
	return allpassBuilder.withOption("a", expr.Int(a))
}

func (allpassBuilder *implAllpassBuilder) Precision(precision int) AllpassBuilder {
	return allpassBuilder.withOption("precision", expr.Int(precision))
}

func (allpassBuilder *implAllpassBuilder) R(r int) AllpassBuilder {
	return allpassBuilder.withOption("r", expr.Int(r))
}

func (allpassBuilder *implAllpassBuilder) Enable(enable expr.Expr) AllpassBuilder {
	return allpassBuilder.withOption("enable", enable)
}
