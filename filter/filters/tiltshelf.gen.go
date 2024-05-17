// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TiltshelfBuilder Apply a tilt shelf filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#tiltshelf
type TiltshelfBuilder interface {
	filter.Filter
	// Frequency set central frequency (from 0 to 999999) (default 3000).
	Frequency(frequency float64) TiltshelfBuilder
	// FrequencyExpr set central frequency (from 0 to 999999) (default 3000).
	FrequencyExpr(frequency expr.Expr) TiltshelfBuilder
	// F set central frequency (from 0 to 999999) (default 3000).
	F(f float64) TiltshelfBuilder
	// FExpr set central frequency (from 0 to 999999) (default 3000).
	FExpr(f expr.Expr) TiltshelfBuilder
	// WidthType set filter-width type (from 1 to 5) (default q).
	WidthType(widthType int) TiltshelfBuilder
	// WidthTypeExpr set filter-width type (from 1 to 5) (default q).
	WidthTypeExpr(widthType expr.Expr) TiltshelfBuilder
	// T set filter-width type (from 1 to 5) (default q).
	T(t int) TiltshelfBuilder
	// TExpr set filter-width type (from 1 to 5) (default q).
	TExpr(t expr.Expr) TiltshelfBuilder
	// Width set width (from 0 to 99999) (default 0.5).
	Width(width float64) TiltshelfBuilder
	// WidthExpr set width (from 0 to 99999) (default 0.5).
	WidthExpr(width expr.Expr) TiltshelfBuilder
	// W set width (from 0 to 99999) (default 0.5).
	W(w float64) TiltshelfBuilder
	// WExpr set width (from 0 to 99999) (default 0.5).
	WExpr(w expr.Expr) TiltshelfBuilder
	// Gain set gain (from -900 to 900) (default 0).
	Gain(gain float64) TiltshelfBuilder
	// GainExpr set gain (from -900 to 900) (default 0).
	GainExpr(gain expr.Expr) TiltshelfBuilder
	// G set gain (from -900 to 900) (default 0).
	G(g float64) TiltshelfBuilder
	// GExpr set gain (from -900 to 900) (default 0).
	GExpr(g expr.Expr) TiltshelfBuilder
	// Poles set number of poles (from 1 to 2) (default 2).
	Poles(poles int) TiltshelfBuilder
	// P set number of poles (from 1 to 2) (default 2).
	P(p int) TiltshelfBuilder
	// Mix set mix (from 0 to 1) (default 1).
	Mix(mix float64) TiltshelfBuilder
	// MixExpr set mix (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) TiltshelfBuilder
	// M set mix (from 0 to 1) (default 1).
	M(m float64) TiltshelfBuilder
	// MExpr set mix (from 0 to 1) (default 1).
	MExpr(m expr.Expr) TiltshelfBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) TiltshelfBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) TiltshelfBuilder
	// C set channels to filter (default "all").
	C(c string) TiltshelfBuilder
	// CExpr set channels to filter (default "all").
	CExpr(c expr.Expr) TiltshelfBuilder
	// Normalize normalize coefficients (default false).
	Normalize(normalize bool) TiltshelfBuilder
	// NormalizeExpr normalize coefficients (default false).
	NormalizeExpr(normalize expr.Expr) TiltshelfBuilder
	// N normalize coefficients (default false).
	N(n bool) TiltshelfBuilder
	// NExpr normalize coefficients (default false).
	NExpr(n expr.Expr) TiltshelfBuilder
	// Transform set transform type (from 0 to 6) (default di).
	Transform(transform int) TiltshelfBuilder
	// A set transform type (from 0 to 6) (default di).
	A(a int) TiltshelfBuilder
	// Precision set filtering precision (from -1 to 3) (default auto).
	Precision(precision int) TiltshelfBuilder
	// R set filtering precision (from -1 to 3) (default auto).
	R(r int) TiltshelfBuilder
	// Blocksize set the block size (from 0 to 32768) (default 0).
	Blocksize(blocksize int) TiltshelfBuilder
	// B set the block size (from 0 to 32768) (default 0).
	B(b int) TiltshelfBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) TiltshelfBuilder
}

// Tiltshelf creates a new TiltshelfBuilder to configure the "tiltshelf" filter.
func Tiltshelf(opts ...filter.Option) TiltshelfBuilder {
	return &implTiltshelfBuilder{
		f: filter.New("tiltshelf", 1, opts...),
	}
}

type implTiltshelfBuilder struct {
	f filter.Filter
}

func (tiltshelfBuilder *implTiltshelfBuilder) String() string {
	return tiltshelfBuilder.f.String()
}

func (tiltshelfBuilder *implTiltshelfBuilder) Outputs() int {
	return tiltshelfBuilder.f.Outputs()
}

func (tiltshelfBuilder *implTiltshelfBuilder) With(key string, value expr.Expr) filter.Filter {
	return tiltshelfBuilder.withOption(key, value)
}

func (tiltshelfBuilder *implTiltshelfBuilder) withOption(key string, value expr.Expr) TiltshelfBuilder {
	bCopy := *tiltshelfBuilder
	bCopy.f = tiltshelfBuilder.f.With(key, value)
	return &bCopy
}

func (tiltshelfBuilder *implTiltshelfBuilder) Frequency(frequency float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("frequency", expr.Double(frequency))
}

func (tiltshelfBuilder *implTiltshelfBuilder) FrequencyExpr(frequency expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("frequency", frequency)
}

func (tiltshelfBuilder *implTiltshelfBuilder) F(f float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("f", expr.Double(f))
}

func (tiltshelfBuilder *implTiltshelfBuilder) FExpr(f expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("f", f)
}

func (tiltshelfBuilder *implTiltshelfBuilder) WidthType(widthType int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("width_type", expr.Int(widthType))
}

func (tiltshelfBuilder *implTiltshelfBuilder) WidthTypeExpr(widthType expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("width_type", widthType)
}

func (tiltshelfBuilder *implTiltshelfBuilder) T(t int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("t", expr.Int(t))
}

func (tiltshelfBuilder *implTiltshelfBuilder) TExpr(t expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("t", t)
}

func (tiltshelfBuilder *implTiltshelfBuilder) Width(width float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("width", expr.Double(width))
}

func (tiltshelfBuilder *implTiltshelfBuilder) WidthExpr(width expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("width", width)
}

func (tiltshelfBuilder *implTiltshelfBuilder) W(w float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("w", expr.Double(w))
}

func (tiltshelfBuilder *implTiltshelfBuilder) WExpr(w expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("w", w)
}

func (tiltshelfBuilder *implTiltshelfBuilder) Gain(gain float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("gain", expr.Double(gain))
}

func (tiltshelfBuilder *implTiltshelfBuilder) GainExpr(gain expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("gain", gain)
}

func (tiltshelfBuilder *implTiltshelfBuilder) G(g float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("g", expr.Double(g))
}

func (tiltshelfBuilder *implTiltshelfBuilder) GExpr(g expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("g", g)
}

func (tiltshelfBuilder *implTiltshelfBuilder) Poles(poles int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("poles", expr.Int(poles))
}

func (tiltshelfBuilder *implTiltshelfBuilder) P(p int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("p", expr.Int(p))
}

func (tiltshelfBuilder *implTiltshelfBuilder) Mix(mix float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("mix", expr.Double(mix))
}

func (tiltshelfBuilder *implTiltshelfBuilder) MixExpr(mix expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("mix", mix)
}

func (tiltshelfBuilder *implTiltshelfBuilder) M(m float64) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("m", expr.Double(m))
}

func (tiltshelfBuilder *implTiltshelfBuilder) MExpr(m expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("m", m)
}

func (tiltshelfBuilder *implTiltshelfBuilder) Channels(channels string) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("channels", expr.String(channels))
}

func (tiltshelfBuilder *implTiltshelfBuilder) ChannelsExpr(channels expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("channels", channels)
}

func (tiltshelfBuilder *implTiltshelfBuilder) C(c string) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("c", expr.String(c))
}

func (tiltshelfBuilder *implTiltshelfBuilder) CExpr(c expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("c", c)
}

func (tiltshelfBuilder *implTiltshelfBuilder) Normalize(normalize bool) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("normalize", expr.Bool(normalize))
}

func (tiltshelfBuilder *implTiltshelfBuilder) NormalizeExpr(normalize expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("normalize", normalize)
}

func (tiltshelfBuilder *implTiltshelfBuilder) N(n bool) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("n", expr.Bool(n))
}

func (tiltshelfBuilder *implTiltshelfBuilder) NExpr(n expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("n", n)
}

func (tiltshelfBuilder *implTiltshelfBuilder) Transform(transform int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("transform", expr.Int(transform))
}

func (tiltshelfBuilder *implTiltshelfBuilder) A(a int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("a", expr.Int(a))
}

func (tiltshelfBuilder *implTiltshelfBuilder) Precision(precision int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("precision", expr.Int(precision))
}

func (tiltshelfBuilder *implTiltshelfBuilder) R(r int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("r", expr.Int(r))
}

func (tiltshelfBuilder *implTiltshelfBuilder) Blocksize(blocksize int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("blocksize", expr.Int(blocksize))
}

func (tiltshelfBuilder *implTiltshelfBuilder) B(b int) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("b", expr.Int(b))
}

func (tiltshelfBuilder *implTiltshelfBuilder) Enable(enable expr.Expr) TiltshelfBuilder {
	return tiltshelfBuilder.withOption("enable", enable)
}