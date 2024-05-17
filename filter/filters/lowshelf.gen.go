// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// LowshelfBuilder Apply a low shelf filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#lowshelf
type LowshelfBuilder interface {
	filter.Filter
	// Frequency set central frequency (from 0 to 999999) (default 100).
	Frequency(frequency float64) LowshelfBuilder
	// FrequencyExpr set central frequency (from 0 to 999999) (default 100).
	FrequencyExpr(frequency expr.Expr) LowshelfBuilder
	// F set central frequency (from 0 to 999999) (default 100).
	F(f float64) LowshelfBuilder
	// FExpr set central frequency (from 0 to 999999) (default 100).
	FExpr(f expr.Expr) LowshelfBuilder
	// WidthType set filter-width type (from 1 to 5) (default q).
	WidthType(widthType int) LowshelfBuilder
	// WidthTypeExpr set filter-width type (from 1 to 5) (default q).
	WidthTypeExpr(widthType expr.Expr) LowshelfBuilder
	// T set filter-width type (from 1 to 5) (default q).
	T(t int) LowshelfBuilder
	// TExpr set filter-width type (from 1 to 5) (default q).
	TExpr(t expr.Expr) LowshelfBuilder
	// Width set width (from 0 to 99999) (default 0.5).
	Width(width float64) LowshelfBuilder
	// WidthExpr set width (from 0 to 99999) (default 0.5).
	WidthExpr(width expr.Expr) LowshelfBuilder
	// W set width (from 0 to 99999) (default 0.5).
	W(w float64) LowshelfBuilder
	// WExpr set width (from 0 to 99999) (default 0.5).
	WExpr(w expr.Expr) LowshelfBuilder
	// Gain set gain (from -900 to 900) (default 0).
	Gain(gain float64) LowshelfBuilder
	// GainExpr set gain (from -900 to 900) (default 0).
	GainExpr(gain expr.Expr) LowshelfBuilder
	// G set gain (from -900 to 900) (default 0).
	G(g float64) LowshelfBuilder
	// GExpr set gain (from -900 to 900) (default 0).
	GExpr(g expr.Expr) LowshelfBuilder
	// Poles set number of poles (from 1 to 2) (default 2).
	Poles(poles int) LowshelfBuilder
	// P set number of poles (from 1 to 2) (default 2).
	P(p int) LowshelfBuilder
	// Mix set mix (from 0 to 1) (default 1).
	Mix(mix float64) LowshelfBuilder
	// MixExpr set mix (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) LowshelfBuilder
	// M set mix (from 0 to 1) (default 1).
	M(m float64) LowshelfBuilder
	// MExpr set mix (from 0 to 1) (default 1).
	MExpr(m expr.Expr) LowshelfBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) LowshelfBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) LowshelfBuilder
	// C set channels to filter (default "all").
	C(c string) LowshelfBuilder
	// CExpr set channels to filter (default "all").
	CExpr(c expr.Expr) LowshelfBuilder
	// Normalize normalize coefficients (default false).
	Normalize(normalize bool) LowshelfBuilder
	// NormalizeExpr normalize coefficients (default false).
	NormalizeExpr(normalize expr.Expr) LowshelfBuilder
	// N normalize coefficients (default false).
	N(n bool) LowshelfBuilder
	// NExpr normalize coefficients (default false).
	NExpr(n expr.Expr) LowshelfBuilder
	// Transform set transform type (from 0 to 6) (default di).
	Transform(transform int) LowshelfBuilder
	// A set transform type (from 0 to 6) (default di).
	A(a int) LowshelfBuilder
	// Precision set filtering precision (from -1 to 3) (default auto).
	Precision(precision int) LowshelfBuilder
	// R set filtering precision (from -1 to 3) (default auto).
	R(r int) LowshelfBuilder
	// Blocksize set the block size (from 0 to 32768) (default 0).
	Blocksize(blocksize int) LowshelfBuilder
	// B set the block size (from 0 to 32768) (default 0).
	B(b int) LowshelfBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) LowshelfBuilder
}

// Lowshelf creates a new LowshelfBuilder to configure the "lowshelf" filter.
func Lowshelf(opts ...filter.Option) LowshelfBuilder {
	return &implLowshelfBuilder{
		f: filter.New("lowshelf", 1, opts...),
	}
}

type implLowshelfBuilder struct {
	f filter.Filter
}

func (lowshelfBuilder *implLowshelfBuilder) String() string {
	return lowshelfBuilder.f.String()
}

func (lowshelfBuilder *implLowshelfBuilder) Outputs() int {
	return lowshelfBuilder.f.Outputs()
}

func (lowshelfBuilder *implLowshelfBuilder) With(key string, value expr.Expr) filter.Filter {
	return lowshelfBuilder.withOption(key, value)
}

func (lowshelfBuilder *implLowshelfBuilder) withOption(key string, value expr.Expr) LowshelfBuilder {
	bCopy := *lowshelfBuilder
	bCopy.f = lowshelfBuilder.f.With(key, value)
	return &bCopy
}

func (lowshelfBuilder *implLowshelfBuilder) Frequency(frequency float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("frequency", expr.Double(frequency))
}

func (lowshelfBuilder *implLowshelfBuilder) FrequencyExpr(frequency expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("frequency", frequency)
}

func (lowshelfBuilder *implLowshelfBuilder) F(f float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("f", expr.Double(f))
}

func (lowshelfBuilder *implLowshelfBuilder) FExpr(f expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("f", f)
}

func (lowshelfBuilder *implLowshelfBuilder) WidthType(widthType int) LowshelfBuilder {
	return lowshelfBuilder.withOption("width_type", expr.Int(widthType))
}

func (lowshelfBuilder *implLowshelfBuilder) WidthTypeExpr(widthType expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("width_type", widthType)
}

func (lowshelfBuilder *implLowshelfBuilder) T(t int) LowshelfBuilder {
	return lowshelfBuilder.withOption("t", expr.Int(t))
}

func (lowshelfBuilder *implLowshelfBuilder) TExpr(t expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("t", t)
}

func (lowshelfBuilder *implLowshelfBuilder) Width(width float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("width", expr.Double(width))
}

func (lowshelfBuilder *implLowshelfBuilder) WidthExpr(width expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("width", width)
}

func (lowshelfBuilder *implLowshelfBuilder) W(w float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("w", expr.Double(w))
}

func (lowshelfBuilder *implLowshelfBuilder) WExpr(w expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("w", w)
}

func (lowshelfBuilder *implLowshelfBuilder) Gain(gain float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("gain", expr.Double(gain))
}

func (lowshelfBuilder *implLowshelfBuilder) GainExpr(gain expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("gain", gain)
}

func (lowshelfBuilder *implLowshelfBuilder) G(g float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("g", expr.Double(g))
}

func (lowshelfBuilder *implLowshelfBuilder) GExpr(g expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("g", g)
}

func (lowshelfBuilder *implLowshelfBuilder) Poles(poles int) LowshelfBuilder {
	return lowshelfBuilder.withOption("poles", expr.Int(poles))
}

func (lowshelfBuilder *implLowshelfBuilder) P(p int) LowshelfBuilder {
	return lowshelfBuilder.withOption("p", expr.Int(p))
}

func (lowshelfBuilder *implLowshelfBuilder) Mix(mix float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("mix", expr.Double(mix))
}

func (lowshelfBuilder *implLowshelfBuilder) MixExpr(mix expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("mix", mix)
}

func (lowshelfBuilder *implLowshelfBuilder) M(m float64) LowshelfBuilder {
	return lowshelfBuilder.withOption("m", expr.Double(m))
}

func (lowshelfBuilder *implLowshelfBuilder) MExpr(m expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("m", m)
}

func (lowshelfBuilder *implLowshelfBuilder) Channels(channels string) LowshelfBuilder {
	return lowshelfBuilder.withOption("channels", expr.String(channels))
}

func (lowshelfBuilder *implLowshelfBuilder) ChannelsExpr(channels expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("channels", channels)
}

func (lowshelfBuilder *implLowshelfBuilder) C(c string) LowshelfBuilder {
	return lowshelfBuilder.withOption("c", expr.String(c))
}

func (lowshelfBuilder *implLowshelfBuilder) CExpr(c expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("c", c)
}

func (lowshelfBuilder *implLowshelfBuilder) Normalize(normalize bool) LowshelfBuilder {
	return lowshelfBuilder.withOption("normalize", expr.Bool(normalize))
}

func (lowshelfBuilder *implLowshelfBuilder) NormalizeExpr(normalize expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("normalize", normalize)
}

func (lowshelfBuilder *implLowshelfBuilder) N(n bool) LowshelfBuilder {
	return lowshelfBuilder.withOption("n", expr.Bool(n))
}

func (lowshelfBuilder *implLowshelfBuilder) NExpr(n expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("n", n)
}

func (lowshelfBuilder *implLowshelfBuilder) Transform(transform int) LowshelfBuilder {
	return lowshelfBuilder.withOption("transform", expr.Int(transform))
}

func (lowshelfBuilder *implLowshelfBuilder) A(a int) LowshelfBuilder {
	return lowshelfBuilder.withOption("a", expr.Int(a))
}

func (lowshelfBuilder *implLowshelfBuilder) Precision(precision int) LowshelfBuilder {
	return lowshelfBuilder.withOption("precision", expr.Int(precision))
}

func (lowshelfBuilder *implLowshelfBuilder) R(r int) LowshelfBuilder {
	return lowshelfBuilder.withOption("r", expr.Int(r))
}

func (lowshelfBuilder *implLowshelfBuilder) Blocksize(blocksize int) LowshelfBuilder {
	return lowshelfBuilder.withOption("blocksize", expr.Int(blocksize))
}

func (lowshelfBuilder *implLowshelfBuilder) B(b int) LowshelfBuilder {
	return lowshelfBuilder.withOption("b", expr.Int(b))
}

func (lowshelfBuilder *implLowshelfBuilder) Enable(enable expr.Expr) LowshelfBuilder {
	return lowshelfBuilder.withOption("enable", enable)
}
