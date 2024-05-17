// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TrebleBuilder Boost or cut upper frequencies.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#treble
type TrebleBuilder interface {
	filter.Filter
	// Frequency set central frequency (from 0 to 999999) (default 3000).
	Frequency(frequency float64) TrebleBuilder
	// FrequencyExpr set central frequency (from 0 to 999999) (default 3000).
	FrequencyExpr(frequency expr.Expr) TrebleBuilder
	// F set central frequency (from 0 to 999999) (default 3000).
	F(f float64) TrebleBuilder
	// FExpr set central frequency (from 0 to 999999) (default 3000).
	FExpr(f expr.Expr) TrebleBuilder
	// WidthType set filter-width type (from 1 to 5) (default q).
	WidthType(widthType int) TrebleBuilder
	// WidthTypeExpr set filter-width type (from 1 to 5) (default q).
	WidthTypeExpr(widthType expr.Expr) TrebleBuilder
	// T set filter-width type (from 1 to 5) (default q).
	T(t int) TrebleBuilder
	// TExpr set filter-width type (from 1 to 5) (default q).
	TExpr(t expr.Expr) TrebleBuilder
	// Width set width (from 0 to 99999) (default 0.5).
	Width(width float64) TrebleBuilder
	// WidthExpr set width (from 0 to 99999) (default 0.5).
	WidthExpr(width expr.Expr) TrebleBuilder
	// W set width (from 0 to 99999) (default 0.5).
	W(w float64) TrebleBuilder
	// WExpr set width (from 0 to 99999) (default 0.5).
	WExpr(w expr.Expr) TrebleBuilder
	// Gain set gain (from -900 to 900) (default 0).
	Gain(gain float64) TrebleBuilder
	// GainExpr set gain (from -900 to 900) (default 0).
	GainExpr(gain expr.Expr) TrebleBuilder
	// G set gain (from -900 to 900) (default 0).
	G(g float64) TrebleBuilder
	// GExpr set gain (from -900 to 900) (default 0).
	GExpr(g expr.Expr) TrebleBuilder
	// Poles set number of poles (from 1 to 2) (default 2).
	Poles(poles int) TrebleBuilder
	// P set number of poles (from 1 to 2) (default 2).
	P(p int) TrebleBuilder
	// Mix set mix (from 0 to 1) (default 1).
	Mix(mix float64) TrebleBuilder
	// MixExpr set mix (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) TrebleBuilder
	// M set mix (from 0 to 1) (default 1).
	M(m float64) TrebleBuilder
	// MExpr set mix (from 0 to 1) (default 1).
	MExpr(m expr.Expr) TrebleBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) TrebleBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) TrebleBuilder
	// C set channels to filter (default "all").
	C(c string) TrebleBuilder
	// CExpr set channels to filter (default "all").
	CExpr(c expr.Expr) TrebleBuilder
	// Normalize normalize coefficients (default false).
	Normalize(normalize bool) TrebleBuilder
	// NormalizeExpr normalize coefficients (default false).
	NormalizeExpr(normalize expr.Expr) TrebleBuilder
	// N normalize coefficients (default false).
	N(n bool) TrebleBuilder
	// NExpr normalize coefficients (default false).
	NExpr(n expr.Expr) TrebleBuilder
	// Transform set transform type (from 0 to 6) (default di).
	Transform(transform int) TrebleBuilder
	// A set transform type (from 0 to 6) (default di).
	A(a int) TrebleBuilder
	// Precision set filtering precision (from -1 to 3) (default auto).
	Precision(precision int) TrebleBuilder
	// R set filtering precision (from -1 to 3) (default auto).
	R(r int) TrebleBuilder
	// Blocksize set the block size (from 0 to 32768) (default 0).
	Blocksize(blocksize int) TrebleBuilder
	// B set the block size (from 0 to 32768) (default 0).
	B(b int) TrebleBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) TrebleBuilder
}

// Treble creates a new TrebleBuilder to configure the "treble" filter.
func Treble(opts ...filter.Option) TrebleBuilder {
	return &implTrebleBuilder{
		f: filter.New("treble", 1, opts...),
	}
}

type implTrebleBuilder struct {
	f filter.Filter
}

func (trebleBuilder *implTrebleBuilder) String() string {
	return trebleBuilder.f.String()
}

func (trebleBuilder *implTrebleBuilder) Outputs() int {
	return trebleBuilder.f.Outputs()
}

func (trebleBuilder *implTrebleBuilder) With(key string, value expr.Expr) filter.Filter {
	return trebleBuilder.withOption(key, value)
}

func (trebleBuilder *implTrebleBuilder) withOption(key string, value expr.Expr) TrebleBuilder {
	bCopy := *trebleBuilder
	bCopy.f = trebleBuilder.f.With(key, value)
	return &bCopy
}

func (trebleBuilder *implTrebleBuilder) Frequency(frequency float64) TrebleBuilder {
	return trebleBuilder.withOption("frequency", expr.Double(frequency))
}

func (trebleBuilder *implTrebleBuilder) FrequencyExpr(frequency expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("frequency", frequency)
}

func (trebleBuilder *implTrebleBuilder) F(f float64) TrebleBuilder {
	return trebleBuilder.withOption("f", expr.Double(f))
}

func (trebleBuilder *implTrebleBuilder) FExpr(f expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("f", f)
}

func (trebleBuilder *implTrebleBuilder) WidthType(widthType int) TrebleBuilder {
	return trebleBuilder.withOption("width_type", expr.Int(widthType))
}

func (trebleBuilder *implTrebleBuilder) WidthTypeExpr(widthType expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("width_type", widthType)
}

func (trebleBuilder *implTrebleBuilder) T(t int) TrebleBuilder {
	return trebleBuilder.withOption("t", expr.Int(t))
}

func (trebleBuilder *implTrebleBuilder) TExpr(t expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("t", t)
}

func (trebleBuilder *implTrebleBuilder) Width(width float64) TrebleBuilder {
	return trebleBuilder.withOption("width", expr.Double(width))
}

func (trebleBuilder *implTrebleBuilder) WidthExpr(width expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("width", width)
}

func (trebleBuilder *implTrebleBuilder) W(w float64) TrebleBuilder {
	return trebleBuilder.withOption("w", expr.Double(w))
}

func (trebleBuilder *implTrebleBuilder) WExpr(w expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("w", w)
}

func (trebleBuilder *implTrebleBuilder) Gain(gain float64) TrebleBuilder {
	return trebleBuilder.withOption("gain", expr.Double(gain))
}

func (trebleBuilder *implTrebleBuilder) GainExpr(gain expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("gain", gain)
}

func (trebleBuilder *implTrebleBuilder) G(g float64) TrebleBuilder {
	return trebleBuilder.withOption("g", expr.Double(g))
}

func (trebleBuilder *implTrebleBuilder) GExpr(g expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("g", g)
}

func (trebleBuilder *implTrebleBuilder) Poles(poles int) TrebleBuilder {
	return trebleBuilder.withOption("poles", expr.Int(poles))
}

func (trebleBuilder *implTrebleBuilder) P(p int) TrebleBuilder {
	return trebleBuilder.withOption("p", expr.Int(p))
}

func (trebleBuilder *implTrebleBuilder) Mix(mix float64) TrebleBuilder {
	return trebleBuilder.withOption("mix", expr.Double(mix))
}

func (trebleBuilder *implTrebleBuilder) MixExpr(mix expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("mix", mix)
}

func (trebleBuilder *implTrebleBuilder) M(m float64) TrebleBuilder {
	return trebleBuilder.withOption("m", expr.Double(m))
}

func (trebleBuilder *implTrebleBuilder) MExpr(m expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("m", m)
}

func (trebleBuilder *implTrebleBuilder) Channels(channels string) TrebleBuilder {
	return trebleBuilder.withOption("channels", expr.String(channels))
}

func (trebleBuilder *implTrebleBuilder) ChannelsExpr(channels expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("channels", channels)
}

func (trebleBuilder *implTrebleBuilder) C(c string) TrebleBuilder {
	return trebleBuilder.withOption("c", expr.String(c))
}

func (trebleBuilder *implTrebleBuilder) CExpr(c expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("c", c)
}

func (trebleBuilder *implTrebleBuilder) Normalize(normalize bool) TrebleBuilder {
	return trebleBuilder.withOption("normalize", expr.Bool(normalize))
}

func (trebleBuilder *implTrebleBuilder) NormalizeExpr(normalize expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("normalize", normalize)
}

func (trebleBuilder *implTrebleBuilder) N(n bool) TrebleBuilder {
	return trebleBuilder.withOption("n", expr.Bool(n))
}

func (trebleBuilder *implTrebleBuilder) NExpr(n expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("n", n)
}

func (trebleBuilder *implTrebleBuilder) Transform(transform int) TrebleBuilder {
	return trebleBuilder.withOption("transform", expr.Int(transform))
}

func (trebleBuilder *implTrebleBuilder) A(a int) TrebleBuilder {
	return trebleBuilder.withOption("a", expr.Int(a))
}

func (trebleBuilder *implTrebleBuilder) Precision(precision int) TrebleBuilder {
	return trebleBuilder.withOption("precision", expr.Int(precision))
}

func (trebleBuilder *implTrebleBuilder) R(r int) TrebleBuilder {
	return trebleBuilder.withOption("r", expr.Int(r))
}

func (trebleBuilder *implTrebleBuilder) Blocksize(blocksize int) TrebleBuilder {
	return trebleBuilder.withOption("blocksize", expr.Int(blocksize))
}

func (trebleBuilder *implTrebleBuilder) B(b int) TrebleBuilder {
	return trebleBuilder.withOption("b", expr.Int(b))
}

func (trebleBuilder *implTrebleBuilder) Enable(enable expr.Expr) TrebleBuilder {
	return trebleBuilder.withOption("enable", enable)
}
