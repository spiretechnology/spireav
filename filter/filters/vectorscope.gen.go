// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// VectorscopeBuilder Video vectorscope.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#vectorscope
type VectorscopeBuilder interface {
	filter.Filter
	// Mode set vectorscope mode (from 0 to 5) (default gray).
	Mode(mode int) VectorscopeBuilder
	// M set vectorscope mode (from 0 to 5) (default gray).
	M(m int) VectorscopeBuilder
	// X set color component on X axis (from 0 to 2) (default 1).
	X(x int) VectorscopeBuilder
	// Y set color component on Y axis (from 0 to 2) (default 2).
	Y(y int) VectorscopeBuilder
	// Intensity set intensity (from 0 to 1) (default 0.004).
	Intensity(intensity float32) VectorscopeBuilder
	// IntensityExpr set intensity (from 0 to 1) (default 0.004).
	IntensityExpr(intensity expr.Expr) VectorscopeBuilder
	// I set intensity (from 0 to 1) (default 0.004).
	I(i float32) VectorscopeBuilder
	// IExpr set intensity (from 0 to 1) (default 0.004).
	IExpr(i expr.Expr) VectorscopeBuilder
	// Envelope set envelope (from 0 to 3) (default none).
	Envelope(envelope int) VectorscopeBuilder
	// EnvelopeExpr set envelope (from 0 to 3) (default none).
	EnvelopeExpr(envelope expr.Expr) VectorscopeBuilder
	// E set envelope (from 0 to 3) (default none).
	E(e int) VectorscopeBuilder
	// EExpr set envelope (from 0 to 3) (default none).
	EExpr(e expr.Expr) VectorscopeBuilder
	// Graticule set graticule (from 0 to 3) (default none).
	Graticule(graticule int) VectorscopeBuilder
	// G set graticule (from 0 to 3) (default none).
	G(g int) VectorscopeBuilder
	// Opacity set graticule opacity (from 0 to 1) (default 0.75).
	Opacity(opacity float32) VectorscopeBuilder
	// OpacityExpr set graticule opacity (from 0 to 1) (default 0.75).
	OpacityExpr(opacity expr.Expr) VectorscopeBuilder
	// O set graticule opacity (from 0 to 1) (default 0.75).
	O(o float32) VectorscopeBuilder
	// OExpr set graticule opacity (from 0 to 1) (default 0.75).
	OExpr(o expr.Expr) VectorscopeBuilder
	// Flags set graticule flags (default name).
	Flags(flags ...string) VectorscopeBuilder
	// FlagsExpr set graticule flags (default name).
	FlagsExpr(flags expr.Expr) VectorscopeBuilder
	// F set graticule flags (default name).
	F(f ...string) VectorscopeBuilder
	// FExpr set graticule flags (default name).
	FExpr(f expr.Expr) VectorscopeBuilder
	// Bgopacity set background opacity (from 0 to 1) (default 0.3).
	Bgopacity(bgopacity float32) VectorscopeBuilder
	// BgopacityExpr set background opacity (from 0 to 1) (default 0.3).
	BgopacityExpr(bgopacity expr.Expr) VectorscopeBuilder
	// B set background opacity (from 0 to 1) (default 0.3).
	B(b float32) VectorscopeBuilder
	// BExpr set background opacity (from 0 to 1) (default 0.3).
	BExpr(b expr.Expr) VectorscopeBuilder
	// Lthreshold set low threshold (from 0 to 1) (default 0).
	Lthreshold(lthreshold float32) VectorscopeBuilder
	// L set low threshold (from 0 to 1) (default 0).
	L(l float32) VectorscopeBuilder
	// Hthreshold set high threshold (from 0 to 1) (default 1).
	Hthreshold(hthreshold float32) VectorscopeBuilder
	// H set high threshold (from 0 to 1) (default 1).
	H(h float32) VectorscopeBuilder
	// Colorspace set colorspace (from 0 to 2) (default auto).
	Colorspace(colorspace int) VectorscopeBuilder
	// C set colorspace (from 0 to 2) (default auto).
	C(c int) VectorscopeBuilder
	// Tint0 set 1st tint (from -1 to 1) (default 0).
	Tint0(tint0 float32) VectorscopeBuilder
	// Tint0Expr set 1st tint (from -1 to 1) (default 0).
	Tint0Expr(tint0 expr.Expr) VectorscopeBuilder
	// T0 set 1st tint (from -1 to 1) (default 0).
	T0(t0 float32) VectorscopeBuilder
	// T0Expr set 1st tint (from -1 to 1) (default 0).
	T0Expr(t0 expr.Expr) VectorscopeBuilder
	// Tint1 set 2nd tint (from -1 to 1) (default 0).
	Tint1(tint1 float32) VectorscopeBuilder
	// Tint1Expr set 2nd tint (from -1 to 1) (default 0).
	Tint1Expr(tint1 expr.Expr) VectorscopeBuilder
	// T1 set 2nd tint (from -1 to 1) (default 0).
	T1(t1 float32) VectorscopeBuilder
	// T1Expr set 2nd tint (from -1 to 1) (default 0).
	T1Expr(t1 expr.Expr) VectorscopeBuilder
}

// Vectorscope creates a new VectorscopeBuilder to configure the "vectorscope" filter.
func Vectorscope(opts ...filter.Option) VectorscopeBuilder {
	return &implVectorscopeBuilder{
		f: filter.New("vectorscope", 1, opts...),
	}
}

type implVectorscopeBuilder struct {
	f filter.Filter
}

func (vectorscopeBuilder *implVectorscopeBuilder) String() string {
	return vectorscopeBuilder.f.String()
}

func (vectorscopeBuilder *implVectorscopeBuilder) Outputs() int {
	return vectorscopeBuilder.f.Outputs()
}

func (vectorscopeBuilder *implVectorscopeBuilder) With(key string, value expr.Expr) filter.Filter {
	return vectorscopeBuilder.withOption(key, value)
}

func (vectorscopeBuilder *implVectorscopeBuilder) withOption(key string, value expr.Expr) VectorscopeBuilder {
	bCopy := *vectorscopeBuilder
	bCopy.f = vectorscopeBuilder.f.With(key, value)
	return &bCopy
}

func (vectorscopeBuilder *implVectorscopeBuilder) Mode(mode int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("mode", expr.Int(mode))
}

func (vectorscopeBuilder *implVectorscopeBuilder) M(m int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("m", expr.Int(m))
}

func (vectorscopeBuilder *implVectorscopeBuilder) X(x int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("x", expr.Int(x))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Y(y int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("y", expr.Int(y))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Intensity(intensity float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("intensity", expr.Float(intensity))
}

func (vectorscopeBuilder *implVectorscopeBuilder) IntensityExpr(intensity expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("intensity", intensity)
}

func (vectorscopeBuilder *implVectorscopeBuilder) I(i float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("i", expr.Float(i))
}

func (vectorscopeBuilder *implVectorscopeBuilder) IExpr(i expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("i", i)
}

func (vectorscopeBuilder *implVectorscopeBuilder) Envelope(envelope int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("envelope", expr.Int(envelope))
}

func (vectorscopeBuilder *implVectorscopeBuilder) EnvelopeExpr(envelope expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("envelope", envelope)
}

func (vectorscopeBuilder *implVectorscopeBuilder) E(e int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("e", expr.Int(e))
}

func (vectorscopeBuilder *implVectorscopeBuilder) EExpr(e expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("e", e)
}

func (vectorscopeBuilder *implVectorscopeBuilder) Graticule(graticule int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("graticule", expr.Int(graticule))
}

func (vectorscopeBuilder *implVectorscopeBuilder) G(g int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("g", expr.Int(g))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Opacity(opacity float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("opacity", expr.Float(opacity))
}

func (vectorscopeBuilder *implVectorscopeBuilder) OpacityExpr(opacity expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("opacity", opacity)
}

func (vectorscopeBuilder *implVectorscopeBuilder) O(o float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("o", expr.Float(o))
}

func (vectorscopeBuilder *implVectorscopeBuilder) OExpr(o expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("o", o)
}

func (vectorscopeBuilder *implVectorscopeBuilder) Flags(flags ...string) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("flags", expr.Flags(flags))
}

func (vectorscopeBuilder *implVectorscopeBuilder) FlagsExpr(flags expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("flags", flags)
}

func (vectorscopeBuilder *implVectorscopeBuilder) F(f ...string) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("f", expr.Flags(f))
}

func (vectorscopeBuilder *implVectorscopeBuilder) FExpr(f expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("f", f)
}

func (vectorscopeBuilder *implVectorscopeBuilder) Bgopacity(bgopacity float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("bgopacity", expr.Float(bgopacity))
}

func (vectorscopeBuilder *implVectorscopeBuilder) BgopacityExpr(bgopacity expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("bgopacity", bgopacity)
}

func (vectorscopeBuilder *implVectorscopeBuilder) B(b float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("b", expr.Float(b))
}

func (vectorscopeBuilder *implVectorscopeBuilder) BExpr(b expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("b", b)
}

func (vectorscopeBuilder *implVectorscopeBuilder) Lthreshold(lthreshold float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("lthreshold", expr.Float(lthreshold))
}

func (vectorscopeBuilder *implVectorscopeBuilder) L(l float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("l", expr.Float(l))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Hthreshold(hthreshold float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("hthreshold", expr.Float(hthreshold))
}

func (vectorscopeBuilder *implVectorscopeBuilder) H(h float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("h", expr.Float(h))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Colorspace(colorspace int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("colorspace", expr.Int(colorspace))
}

func (vectorscopeBuilder *implVectorscopeBuilder) C(c int) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("c", expr.Int(c))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Tint0(tint0 float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("tint0", expr.Float(tint0))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Tint0Expr(tint0 expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("tint0", tint0)
}

func (vectorscopeBuilder *implVectorscopeBuilder) T0(t0 float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("t0", expr.Float(t0))
}

func (vectorscopeBuilder *implVectorscopeBuilder) T0Expr(t0 expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("t0", t0)
}

func (vectorscopeBuilder *implVectorscopeBuilder) Tint1(tint1 float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("tint1", expr.Float(tint1))
}

func (vectorscopeBuilder *implVectorscopeBuilder) Tint1Expr(tint1 expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("tint1", tint1)
}

func (vectorscopeBuilder *implVectorscopeBuilder) T1(t1 float32) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("t1", expr.Float(t1))
}

func (vectorscopeBuilder *implVectorscopeBuilder) T1Expr(t1 expr.Expr) VectorscopeBuilder {
	return vectorscopeBuilder.withOption("t1", t1)
}
