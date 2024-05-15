// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// WaveformBuilder Video waveform monitor.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#waveform
type WaveformBuilder interface {
	filter.Filter
	// Mode set mode (from 0 to 1) (default column).
	Mode(mode int) WaveformBuilder
	// M set mode (from 0 to 1) (default column).
	M(m int) WaveformBuilder
	// Intensity set intensity (from 0 to 1) (default 0.04).
	Intensity(intensity float32) WaveformBuilder
	// IntensityExpr set intensity (from 0 to 1) (default 0.04).
	IntensityExpr(intensity expr.Expr) WaveformBuilder
	// I set intensity (from 0 to 1) (default 0.04).
	I(i float32) WaveformBuilder
	// IExpr set intensity (from 0 to 1) (default 0.04).
	IExpr(i expr.Expr) WaveformBuilder
	// Mirror set mirroring (default true).
	Mirror(mirror bool) WaveformBuilder
	// R set mirroring (default true).
	R(r bool) WaveformBuilder
	// Display set display mode (from 0 to 2) (default stack).
	Display(display int) WaveformBuilder
	// D set display mode (from 0 to 2) (default stack).
	D(d int) WaveformBuilder
	// Components set components to display (from 1 to 15) (default 1).
	Components(components int) WaveformBuilder
	// C set components to display (from 1 to 15) (default 1).
	C(c int) WaveformBuilder
	// Envelope set envelope to display (from 0 to 3) (default none).
	Envelope(envelope int) WaveformBuilder
	// EnvelopeExpr set envelope to display (from 0 to 3) (default none).
	EnvelopeExpr(envelope expr.Expr) WaveformBuilder
	// E set envelope to display (from 0 to 3) (default none).
	E(e int) WaveformBuilder
	// EExpr set envelope to display (from 0 to 3) (default none).
	EExpr(e expr.Expr) WaveformBuilder
	// Filter set filter (from 0 to 7) (default lowpass).
	Filter(filter int) WaveformBuilder
	// F set filter (from 0 to 7) (default lowpass).
	F(f int) WaveformBuilder
	// Graticule set graticule (from 0 to 3) (default none).
	Graticule(graticule int) WaveformBuilder
	// G set graticule (from 0 to 3) (default none).
	G(g int) WaveformBuilder
	// Opacity set graticule opacity (from 0 to 1) (default 0.75).
	Opacity(opacity float32) WaveformBuilder
	// OpacityExpr set graticule opacity (from 0 to 1) (default 0.75).
	OpacityExpr(opacity expr.Expr) WaveformBuilder
	// O set graticule opacity (from 0 to 1) (default 0.75).
	O(o float32) WaveformBuilder
	// OExpr set graticule opacity (from 0 to 1) (default 0.75).
	OExpr(o expr.Expr) WaveformBuilder
	// Flags set graticule flags (default numbers).
	Flags(flags string) WaveformBuilder
	// FlagsExpr set graticule flags (default numbers).
	FlagsExpr(flags expr.Expr) WaveformBuilder
	// Fl set graticule flags (default numbers).
	Fl(fl string) WaveformBuilder
	// FlExpr set graticule flags (default numbers).
	FlExpr(fl expr.Expr) WaveformBuilder
	// Scale set scale (from 0 to 2) (default digital).
	Scale(scale int) WaveformBuilder
	// S set scale (from 0 to 2) (default digital).
	S(s int) WaveformBuilder
	// Bgopacity set background opacity (from 0 to 1) (default 0.75).
	Bgopacity(bgopacity float32) WaveformBuilder
	// BgopacityExpr set background opacity (from 0 to 1) (default 0.75).
	BgopacityExpr(bgopacity expr.Expr) WaveformBuilder
	// B set background opacity (from 0 to 1) (default 0.75).
	B(b float32) WaveformBuilder
	// BExpr set background opacity (from 0 to 1) (default 0.75).
	BExpr(b expr.Expr) WaveformBuilder
	// Tint0 set 1st tint (from -1 to 1) (default 0).
	Tint0(tint0 float32) WaveformBuilder
	// Tint0Expr set 1st tint (from -1 to 1) (default 0).
	Tint0Expr(tint0 expr.Expr) WaveformBuilder
	// T0 set 1st tint (from -1 to 1) (default 0).
	T0(t0 float32) WaveformBuilder
	// T0Expr set 1st tint (from -1 to 1) (default 0).
	T0Expr(t0 expr.Expr) WaveformBuilder
	// Tint1 set 2nd tint (from -1 to 1) (default 0).
	Tint1(tint1 float32) WaveformBuilder
	// Tint1Expr set 2nd tint (from -1 to 1) (default 0).
	Tint1Expr(tint1 expr.Expr) WaveformBuilder
	// T1 set 2nd tint (from -1 to 1) (default 0).
	T1(t1 float32) WaveformBuilder
	// T1Expr set 2nd tint (from -1 to 1) (default 0).
	T1Expr(t1 expr.Expr) WaveformBuilder
	// Fitmode set fit mode (from 0 to 1) (default none).
	Fitmode(fitmode int) WaveformBuilder
	// Fm set fit mode (from 0 to 1) (default none).
	Fm(fm int) WaveformBuilder
	// Input set input formats selection (from 0 to 1) (default first).
	Input(input int) WaveformBuilder
}

// Waveform creates a new WaveformBuilder to configure the "waveform" filter.
func Waveform(opts ...filter.Option) WaveformBuilder {
	return &implWaveformBuilder{
		f: filter.New("waveform", 1, opts...),
	}
}

type implWaveformBuilder struct {
	f filter.Filter
}

func (waveformBuilder *implWaveformBuilder) String() string {
	return waveformBuilder.f.String()
}

func (waveformBuilder *implWaveformBuilder) Outputs() int {
	return waveformBuilder.f.Outputs()
}

func (waveformBuilder *implWaveformBuilder) With(key string, value expr.Expr) filter.Filter {
	return waveformBuilder.withOption(key, value)
}

func (waveformBuilder *implWaveformBuilder) withOption(key string, value expr.Expr) WaveformBuilder {
	bCopy := *waveformBuilder
	bCopy.f = waveformBuilder.f.With(key, value)
	return &bCopy
}

func (waveformBuilder *implWaveformBuilder) Mode(mode int) WaveformBuilder {
	return waveformBuilder.withOption("mode", expr.Int(mode))
}

func (waveformBuilder *implWaveformBuilder) M(m int) WaveformBuilder {
	return waveformBuilder.withOption("m", expr.Int(m))
}

func (waveformBuilder *implWaveformBuilder) Intensity(intensity float32) WaveformBuilder {
	return waveformBuilder.withOption("intensity", expr.Float(intensity))
}

func (waveformBuilder *implWaveformBuilder) IntensityExpr(intensity expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("intensity", intensity)
}

func (waveformBuilder *implWaveformBuilder) I(i float32) WaveformBuilder {
	return waveformBuilder.withOption("i", expr.Float(i))
}

func (waveformBuilder *implWaveformBuilder) IExpr(i expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("i", i)
}

func (waveformBuilder *implWaveformBuilder) Mirror(mirror bool) WaveformBuilder {
	return waveformBuilder.withOption("mirror", expr.Bool(mirror))
}

func (waveformBuilder *implWaveformBuilder) R(r bool) WaveformBuilder {
	return waveformBuilder.withOption("r", expr.Bool(r))
}

func (waveformBuilder *implWaveformBuilder) Display(display int) WaveformBuilder {
	return waveformBuilder.withOption("display", expr.Int(display))
}

func (waveformBuilder *implWaveformBuilder) D(d int) WaveformBuilder {
	return waveformBuilder.withOption("d", expr.Int(d))
}

func (waveformBuilder *implWaveformBuilder) Components(components int) WaveformBuilder {
	return waveformBuilder.withOption("components", expr.Int(components))
}

func (waveformBuilder *implWaveformBuilder) C(c int) WaveformBuilder {
	return waveformBuilder.withOption("c", expr.Int(c))
}

func (waveformBuilder *implWaveformBuilder) Envelope(envelope int) WaveformBuilder {
	return waveformBuilder.withOption("envelope", expr.Int(envelope))
}

func (waveformBuilder *implWaveformBuilder) EnvelopeExpr(envelope expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("envelope", envelope)
}

func (waveformBuilder *implWaveformBuilder) E(e int) WaveformBuilder {
	return waveformBuilder.withOption("e", expr.Int(e))
}

func (waveformBuilder *implWaveformBuilder) EExpr(e expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("e", e)
}

func (waveformBuilder *implWaveformBuilder) Filter(filter int) WaveformBuilder {
	return waveformBuilder.withOption("filter", expr.Int(filter))
}

func (waveformBuilder *implWaveformBuilder) F(f int) WaveformBuilder {
	return waveformBuilder.withOption("f", expr.Int(f))
}

func (waveformBuilder *implWaveformBuilder) Graticule(graticule int) WaveformBuilder {
	return waveformBuilder.withOption("graticule", expr.Int(graticule))
}

func (waveformBuilder *implWaveformBuilder) G(g int) WaveformBuilder {
	return waveformBuilder.withOption("g", expr.Int(g))
}

func (waveformBuilder *implWaveformBuilder) Opacity(opacity float32) WaveformBuilder {
	return waveformBuilder.withOption("opacity", expr.Float(opacity))
}

func (waveformBuilder *implWaveformBuilder) OpacityExpr(opacity expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("opacity", opacity)
}

func (waveformBuilder *implWaveformBuilder) O(o float32) WaveformBuilder {
	return waveformBuilder.withOption("o", expr.Float(o))
}

func (waveformBuilder *implWaveformBuilder) OExpr(o expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("o", o)
}

func (waveformBuilder *implWaveformBuilder) Flags(flags string) WaveformBuilder {
	return waveformBuilder.withOption("flags", expr.String(flags))
}

func (waveformBuilder *implWaveformBuilder) FlagsExpr(flags expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("flags", flags)
}

func (waveformBuilder *implWaveformBuilder) Fl(fl string) WaveformBuilder {
	return waveformBuilder.withOption("fl", expr.String(fl))
}

func (waveformBuilder *implWaveformBuilder) FlExpr(fl expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("fl", fl)
}

func (waveformBuilder *implWaveformBuilder) Scale(scale int) WaveformBuilder {
	return waveformBuilder.withOption("scale", expr.Int(scale))
}

func (waveformBuilder *implWaveformBuilder) S(s int) WaveformBuilder {
	return waveformBuilder.withOption("s", expr.Int(s))
}

func (waveformBuilder *implWaveformBuilder) Bgopacity(bgopacity float32) WaveformBuilder {
	return waveformBuilder.withOption("bgopacity", expr.Float(bgopacity))
}

func (waveformBuilder *implWaveformBuilder) BgopacityExpr(bgopacity expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("bgopacity", bgopacity)
}

func (waveformBuilder *implWaveformBuilder) B(b float32) WaveformBuilder {
	return waveformBuilder.withOption("b", expr.Float(b))
}

func (waveformBuilder *implWaveformBuilder) BExpr(b expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("b", b)
}

func (waveformBuilder *implWaveformBuilder) Tint0(tint0 float32) WaveformBuilder {
	return waveformBuilder.withOption("tint0", expr.Float(tint0))
}

func (waveformBuilder *implWaveformBuilder) Tint0Expr(tint0 expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("tint0", tint0)
}

func (waveformBuilder *implWaveformBuilder) T0(t0 float32) WaveformBuilder {
	return waveformBuilder.withOption("t0", expr.Float(t0))
}

func (waveformBuilder *implWaveformBuilder) T0Expr(t0 expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("t0", t0)
}

func (waveformBuilder *implWaveformBuilder) Tint1(tint1 float32) WaveformBuilder {
	return waveformBuilder.withOption("tint1", expr.Float(tint1))
}

func (waveformBuilder *implWaveformBuilder) Tint1Expr(tint1 expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("tint1", tint1)
}

func (waveformBuilder *implWaveformBuilder) T1(t1 float32) WaveformBuilder {
	return waveformBuilder.withOption("t1", expr.Float(t1))
}

func (waveformBuilder *implWaveformBuilder) T1Expr(t1 expr.Expr) WaveformBuilder {
	return waveformBuilder.withOption("t1", t1)
}

func (waveformBuilder *implWaveformBuilder) Fitmode(fitmode int) WaveformBuilder {
	return waveformBuilder.withOption("fitmode", expr.Int(fitmode))
}

func (waveformBuilder *implWaveformBuilder) Fm(fm int) WaveformBuilder {
	return waveformBuilder.withOption("fm", expr.Int(fm))
}

func (waveformBuilder *implWaveformBuilder) Input(input int) WaveformBuilder {
	return waveformBuilder.withOption("input", expr.Int(input))
}
