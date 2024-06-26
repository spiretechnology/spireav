// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// HighshelfBuilder Apply a high shelf filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#highshelf
type HighshelfBuilder interface {
	filter.Filter
	// Frequency set central frequency (from 0 to 999999) (default 3000).
	Frequency(frequency float64) HighshelfBuilder
	// FrequencyExpr set central frequency (from 0 to 999999) (default 3000).
	FrequencyExpr(frequency expr.Expr) HighshelfBuilder
	// F set central frequency (from 0 to 999999) (default 3000).
	F(f float64) HighshelfBuilder
	// FExpr set central frequency (from 0 to 999999) (default 3000).
	FExpr(f expr.Expr) HighshelfBuilder
	// WidthType set filter-width type (from 1 to 5) (default q).
	WidthType(widthType int) HighshelfBuilder
	// WidthTypeExpr set filter-width type (from 1 to 5) (default q).
	WidthTypeExpr(widthType expr.Expr) HighshelfBuilder
	// T set filter-width type (from 1 to 5) (default q).
	T(t int) HighshelfBuilder
	// TExpr set filter-width type (from 1 to 5) (default q).
	TExpr(t expr.Expr) HighshelfBuilder
	// Width set width (from 0 to 99999) (default 0.5).
	Width(width float64) HighshelfBuilder
	// WidthExpr set width (from 0 to 99999) (default 0.5).
	WidthExpr(width expr.Expr) HighshelfBuilder
	// W set width (from 0 to 99999) (default 0.5).
	W(w float64) HighshelfBuilder
	// WExpr set width (from 0 to 99999) (default 0.5).
	WExpr(w expr.Expr) HighshelfBuilder
	// Gain set gain (from -900 to 900) (default 0).
	Gain(gain float64) HighshelfBuilder
	// GainExpr set gain (from -900 to 900) (default 0).
	GainExpr(gain expr.Expr) HighshelfBuilder
	// G set gain (from -900 to 900) (default 0).
	G(g float64) HighshelfBuilder
	// GExpr set gain (from -900 to 900) (default 0).
	GExpr(g expr.Expr) HighshelfBuilder
	// Poles set number of poles (from 1 to 2) (default 2).
	Poles(poles int) HighshelfBuilder
	// P set number of poles (from 1 to 2) (default 2).
	P(p int) HighshelfBuilder
	// Mix set mix (from 0 to 1) (default 1).
	Mix(mix float64) HighshelfBuilder
	// MixExpr set mix (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) HighshelfBuilder
	// M set mix (from 0 to 1) (default 1).
	M(m float64) HighshelfBuilder
	// MExpr set mix (from 0 to 1) (default 1).
	MExpr(m expr.Expr) HighshelfBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) HighshelfBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) HighshelfBuilder
	// C set channels to filter (default "all").
	C(c string) HighshelfBuilder
	// CExpr set channels to filter (default "all").
	CExpr(c expr.Expr) HighshelfBuilder
	// Normalize normalize coefficients (default false).
	Normalize(normalize bool) HighshelfBuilder
	// NormalizeExpr normalize coefficients (default false).
	NormalizeExpr(normalize expr.Expr) HighshelfBuilder
	// N normalize coefficients (default false).
	N(n bool) HighshelfBuilder
	// NExpr normalize coefficients (default false).
	NExpr(n expr.Expr) HighshelfBuilder
	// Transform set transform type (from 0 to 6) (default di).
	Transform(transform int) HighshelfBuilder
	// A set transform type (from 0 to 6) (default di).
	A(a int) HighshelfBuilder
	// Precision set filtering precision (from -1 to 3) (default auto).
	Precision(precision int) HighshelfBuilder
	// R set filtering precision (from -1 to 3) (default auto).
	R(r int) HighshelfBuilder
	// Blocksize set the block size (from 0 to 32768) (default 0).
	Blocksize(blocksize int) HighshelfBuilder
	// B set the block size (from 0 to 32768) (default 0).
	B(b int) HighshelfBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) HighshelfBuilder
}

// Highshelf creates a new HighshelfBuilder to configure the "highshelf" filter.
func Highshelf(opts ...filter.Option) HighshelfBuilder {
	return &implHighshelfBuilder{
		f: filter.New("highshelf", 1, opts...),
	}
}

type implHighshelfBuilder struct {
	f filter.Filter
}

func (highshelfBuilder *implHighshelfBuilder) String() string {
	return highshelfBuilder.f.String()
}

func (highshelfBuilder *implHighshelfBuilder) Outputs() int {
	return highshelfBuilder.f.Outputs()
}

func (highshelfBuilder *implHighshelfBuilder) With(key string, value expr.Expr) filter.Filter {
	return highshelfBuilder.withOption(key, value)
}

func (highshelfBuilder *implHighshelfBuilder) withOption(key string, value expr.Expr) HighshelfBuilder {
	bCopy := *highshelfBuilder
	bCopy.f = highshelfBuilder.f.With(key, value)
	return &bCopy
}

func (highshelfBuilder *implHighshelfBuilder) Frequency(frequency float64) HighshelfBuilder {
	return highshelfBuilder.withOption("frequency", expr.Double(frequency))
}

func (highshelfBuilder *implHighshelfBuilder) FrequencyExpr(frequency expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("frequency", frequency)
}

func (highshelfBuilder *implHighshelfBuilder) F(f float64) HighshelfBuilder {
	return highshelfBuilder.withOption("f", expr.Double(f))
}

func (highshelfBuilder *implHighshelfBuilder) FExpr(f expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("f", f)
}

func (highshelfBuilder *implHighshelfBuilder) WidthType(widthType int) HighshelfBuilder {
	return highshelfBuilder.withOption("width_type", expr.Int(widthType))
}

func (highshelfBuilder *implHighshelfBuilder) WidthTypeExpr(widthType expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("width_type", widthType)
}

func (highshelfBuilder *implHighshelfBuilder) T(t int) HighshelfBuilder {
	return highshelfBuilder.withOption("t", expr.Int(t))
}

func (highshelfBuilder *implHighshelfBuilder) TExpr(t expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("t", t)
}

func (highshelfBuilder *implHighshelfBuilder) Width(width float64) HighshelfBuilder {
	return highshelfBuilder.withOption("width", expr.Double(width))
}

func (highshelfBuilder *implHighshelfBuilder) WidthExpr(width expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("width", width)
}

func (highshelfBuilder *implHighshelfBuilder) W(w float64) HighshelfBuilder {
	return highshelfBuilder.withOption("w", expr.Double(w))
}

func (highshelfBuilder *implHighshelfBuilder) WExpr(w expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("w", w)
}

func (highshelfBuilder *implHighshelfBuilder) Gain(gain float64) HighshelfBuilder {
	return highshelfBuilder.withOption("gain", expr.Double(gain))
}

func (highshelfBuilder *implHighshelfBuilder) GainExpr(gain expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("gain", gain)
}

func (highshelfBuilder *implHighshelfBuilder) G(g float64) HighshelfBuilder {
	return highshelfBuilder.withOption("g", expr.Double(g))
}

func (highshelfBuilder *implHighshelfBuilder) GExpr(g expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("g", g)
}

func (highshelfBuilder *implHighshelfBuilder) Poles(poles int) HighshelfBuilder {
	return highshelfBuilder.withOption("poles", expr.Int(poles))
}

func (highshelfBuilder *implHighshelfBuilder) P(p int) HighshelfBuilder {
	return highshelfBuilder.withOption("p", expr.Int(p))
}

func (highshelfBuilder *implHighshelfBuilder) Mix(mix float64) HighshelfBuilder {
	return highshelfBuilder.withOption("mix", expr.Double(mix))
}

func (highshelfBuilder *implHighshelfBuilder) MixExpr(mix expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("mix", mix)
}

func (highshelfBuilder *implHighshelfBuilder) M(m float64) HighshelfBuilder {
	return highshelfBuilder.withOption("m", expr.Double(m))
}

func (highshelfBuilder *implHighshelfBuilder) MExpr(m expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("m", m)
}

func (highshelfBuilder *implHighshelfBuilder) Channels(channels string) HighshelfBuilder {
	return highshelfBuilder.withOption("channels", expr.String(channels))
}

func (highshelfBuilder *implHighshelfBuilder) ChannelsExpr(channels expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("channels", channels)
}

func (highshelfBuilder *implHighshelfBuilder) C(c string) HighshelfBuilder {
	return highshelfBuilder.withOption("c", expr.String(c))
}

func (highshelfBuilder *implHighshelfBuilder) CExpr(c expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("c", c)
}

func (highshelfBuilder *implHighshelfBuilder) Normalize(normalize bool) HighshelfBuilder {
	return highshelfBuilder.withOption("normalize", expr.Bool(normalize))
}

func (highshelfBuilder *implHighshelfBuilder) NormalizeExpr(normalize expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("normalize", normalize)
}

func (highshelfBuilder *implHighshelfBuilder) N(n bool) HighshelfBuilder {
	return highshelfBuilder.withOption("n", expr.Bool(n))
}

func (highshelfBuilder *implHighshelfBuilder) NExpr(n expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("n", n)
}

func (highshelfBuilder *implHighshelfBuilder) Transform(transform int) HighshelfBuilder {
	return highshelfBuilder.withOption("transform", expr.Int(transform))
}

func (highshelfBuilder *implHighshelfBuilder) A(a int) HighshelfBuilder {
	return highshelfBuilder.withOption("a", expr.Int(a))
}

func (highshelfBuilder *implHighshelfBuilder) Precision(precision int) HighshelfBuilder {
	return highshelfBuilder.withOption("precision", expr.Int(precision))
}

func (highshelfBuilder *implHighshelfBuilder) R(r int) HighshelfBuilder {
	return highshelfBuilder.withOption("r", expr.Int(r))
}

func (highshelfBuilder *implHighshelfBuilder) Blocksize(blocksize int) HighshelfBuilder {
	return highshelfBuilder.withOption("blocksize", expr.Int(blocksize))
}

func (highshelfBuilder *implHighshelfBuilder) B(b int) HighshelfBuilder {
	return highshelfBuilder.withOption("b", expr.Int(b))
}

func (highshelfBuilder *implHighshelfBuilder) Enable(enable expr.Expr) HighshelfBuilder {
	return highshelfBuilder.withOption("enable", enable)
}
