// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ShowcqtBuilder Convert input audio to a CQT (Constant/Clamped Q Transform) spectrum video output.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#showcqt
type ShowcqtBuilder interface {
	filter.Filter
	// Size set video size (default "1920x1080").
	Size(size expr.Size) ShowcqtBuilder
	// S set video size (default "1920x1080").
	S(s expr.Size) ShowcqtBuilder
	// Fps set video rate (default "25").
	Fps(fps expr.Rational) ShowcqtBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) ShowcqtBuilder
	// R set video rate (default "25").
	R(r expr.Rational) ShowcqtBuilder
	// BarH set bargraph height (from -1 to INT_MAX) (default -1).
	BarH(barH int) ShowcqtBuilder
	// AxisH set axis height (from -1 to INT_MAX) (default -1).
	AxisH(axisH int) ShowcqtBuilder
	// SonoH set sonogram height (from -1 to INT_MAX) (default -1).
	SonoH(sonoH int) ShowcqtBuilder
	// Fullhd set fullhd size (default true).
	Fullhd(fullhd bool) ShowcqtBuilder
	// SonoV set sonogram volume (default "16").
	SonoV(sonoV string) ShowcqtBuilder
	// Volume set sonogram volume (default "16").
	Volume(volume string) ShowcqtBuilder
	// BarV set bargraph volume (default "sono_v").
	BarV(barV string) ShowcqtBuilder
	// Volume2 set bargraph volume (default "sono_v").
	Volume2(volume2 string) ShowcqtBuilder
	// SonoG set sonogram gamma (from 1 to 7) (default 3).
	SonoG(sonoG float32) ShowcqtBuilder
	// Gamma set sonogram gamma (from 1 to 7) (default 3).
	Gamma(gamma float32) ShowcqtBuilder
	// BarG set bargraph gamma (from 1 to 7) (default 1).
	BarG(barG float32) ShowcqtBuilder
	// Gamma2 set bargraph gamma (from 1 to 7) (default 1).
	Gamma2(gamma2 float32) ShowcqtBuilder
	// BarT set bar transparency (from 0 to 1) (default 1).
	BarT(barT float32) ShowcqtBuilder
	// Timeclamp set timeclamp (from 0.002 to 1) (default 0.17).
	Timeclamp(timeclamp float64) ShowcqtBuilder
	// Tc set timeclamp (from 0.002 to 1) (default 0.17).
	Tc(tc float64) ShowcqtBuilder
	// Attack set attack time (from 0 to 1) (default 0).
	Attack(attack float64) ShowcqtBuilder
	// Basefreq set base frequency (from 10 to 100000) (default 20.0152).
	Basefreq(basefreq float64) ShowcqtBuilder
	// Endfreq set end frequency (from 10 to 100000) (default 20495.6).
	Endfreq(endfreq float64) ShowcqtBuilder
	// Coeffclamp set coeffclamp (from 0.1 to 10) (default 1).
	Coeffclamp(coeffclamp float32) ShowcqtBuilder
	// Tlength set tlength (default "384*tc/(384+tc*f)").
	Tlength(tlength string) ShowcqtBuilder
	// Count set transform count (from 1 to 30) (default 6).
	Count(count int) ShowcqtBuilder
	// Fcount set frequency count (from 0 to 10) (default 0).
	Fcount(fcount int) ShowcqtBuilder
	// Fontfile set axis font file.
	Fontfile(fontfile string) ShowcqtBuilder
	// Font set axis font.
	Font(font string) ShowcqtBuilder
	// Fontcolor set font color (default "st(0, (midi(f)-59.5)/12);st(1, if(between(ld(0),0,1), 0.5-0.5*cos(2*PI*ld(0)), 0));r(1-ld(1)) + b(ld(1))").
	Fontcolor(fontcolor string) ShowcqtBuilder
	// Axisfile set axis image.
	Axisfile(axisfile string) ShowcqtBuilder
	// Axis draw axis (default true).
	Axis(axis bool) ShowcqtBuilder
	// Text draw axis (default true).
	Text(text bool) ShowcqtBuilder
	// Csp set color space (from 0 to INT_MAX) (default unspecified).
	Csp(csp int) ShowcqtBuilder
	// Cscheme set color scheme (default "1|0.5|0|0|0.5|1").
	Cscheme(cscheme string) ShowcqtBuilder
}

// Showcqt creates a new ShowcqtBuilder to configure the "showcqt" filter.
func Showcqt(opts ...filter.Option) ShowcqtBuilder {
	return &implShowcqtBuilder{
		f: filter.New("showcqt", 1, opts...),
	}
}

type implShowcqtBuilder struct {
	f filter.Filter
}

func (showcqtBuilder *implShowcqtBuilder) String() string {
	return showcqtBuilder.f.String()
}

func (showcqtBuilder *implShowcqtBuilder) Outputs() int {
	return showcqtBuilder.f.Outputs()
}

func (showcqtBuilder *implShowcqtBuilder) With(key string, value expr.Expr) filter.Filter {
	return showcqtBuilder.withOption(key, value)
}

func (showcqtBuilder *implShowcqtBuilder) withOption(key string, value expr.Expr) ShowcqtBuilder {
	bCopy := *showcqtBuilder
	bCopy.f = showcqtBuilder.f.With(key, value)
	return &bCopy
}

func (showcqtBuilder *implShowcqtBuilder) Size(size expr.Size) ShowcqtBuilder {
	return showcqtBuilder.withOption("size", size)
}

func (showcqtBuilder *implShowcqtBuilder) S(s expr.Size) ShowcqtBuilder {
	return showcqtBuilder.withOption("s", s)
}

func (showcqtBuilder *implShowcqtBuilder) Fps(fps expr.Rational) ShowcqtBuilder {
	return showcqtBuilder.withOption("fps", fps)
}

func (showcqtBuilder *implShowcqtBuilder) Rate(rate expr.Rational) ShowcqtBuilder {
	return showcqtBuilder.withOption("rate", rate)
}

func (showcqtBuilder *implShowcqtBuilder) R(r expr.Rational) ShowcqtBuilder {
	return showcqtBuilder.withOption("r", r)
}

func (showcqtBuilder *implShowcqtBuilder) BarH(barH int) ShowcqtBuilder {
	return showcqtBuilder.withOption("bar_h", expr.Int(barH))
}

func (showcqtBuilder *implShowcqtBuilder) AxisH(axisH int) ShowcqtBuilder {
	return showcqtBuilder.withOption("axis_h", expr.Int(axisH))
}

func (showcqtBuilder *implShowcqtBuilder) SonoH(sonoH int) ShowcqtBuilder {
	return showcqtBuilder.withOption("sono_h", expr.Int(sonoH))
}

func (showcqtBuilder *implShowcqtBuilder) Fullhd(fullhd bool) ShowcqtBuilder {
	return showcqtBuilder.withOption("fullhd", expr.Bool(fullhd))
}

func (showcqtBuilder *implShowcqtBuilder) SonoV(sonoV string) ShowcqtBuilder {
	return showcqtBuilder.withOption("sono_v", expr.String(sonoV))
}

func (showcqtBuilder *implShowcqtBuilder) Volume(volume string) ShowcqtBuilder {
	return showcqtBuilder.withOption("volume", expr.String(volume))
}

func (showcqtBuilder *implShowcqtBuilder) BarV(barV string) ShowcqtBuilder {
	return showcqtBuilder.withOption("bar_v", expr.String(barV))
}

func (showcqtBuilder *implShowcqtBuilder) Volume2(volume2 string) ShowcqtBuilder {
	return showcqtBuilder.withOption("volume2", expr.String(volume2))
}

func (showcqtBuilder *implShowcqtBuilder) SonoG(sonoG float32) ShowcqtBuilder {
	return showcqtBuilder.withOption("sono_g", expr.Float(sonoG))
}

func (showcqtBuilder *implShowcqtBuilder) Gamma(gamma float32) ShowcqtBuilder {
	return showcqtBuilder.withOption("gamma", expr.Float(gamma))
}

func (showcqtBuilder *implShowcqtBuilder) BarG(barG float32) ShowcqtBuilder {
	return showcqtBuilder.withOption("bar_g", expr.Float(barG))
}

func (showcqtBuilder *implShowcqtBuilder) Gamma2(gamma2 float32) ShowcqtBuilder {
	return showcqtBuilder.withOption("gamma2", expr.Float(gamma2))
}

func (showcqtBuilder *implShowcqtBuilder) BarT(barT float32) ShowcqtBuilder {
	return showcqtBuilder.withOption("bar_t", expr.Float(barT))
}

func (showcqtBuilder *implShowcqtBuilder) Timeclamp(timeclamp float64) ShowcqtBuilder {
	return showcqtBuilder.withOption("timeclamp", expr.Double(timeclamp))
}

func (showcqtBuilder *implShowcqtBuilder) Tc(tc float64) ShowcqtBuilder {
	return showcqtBuilder.withOption("tc", expr.Double(tc))
}

func (showcqtBuilder *implShowcqtBuilder) Attack(attack float64) ShowcqtBuilder {
	return showcqtBuilder.withOption("attack", expr.Double(attack))
}

func (showcqtBuilder *implShowcqtBuilder) Basefreq(basefreq float64) ShowcqtBuilder {
	return showcqtBuilder.withOption("basefreq", expr.Double(basefreq))
}

func (showcqtBuilder *implShowcqtBuilder) Endfreq(endfreq float64) ShowcqtBuilder {
	return showcqtBuilder.withOption("endfreq", expr.Double(endfreq))
}

func (showcqtBuilder *implShowcqtBuilder) Coeffclamp(coeffclamp float32) ShowcqtBuilder {
	return showcqtBuilder.withOption("coeffclamp", expr.Float(coeffclamp))
}

func (showcqtBuilder *implShowcqtBuilder) Tlength(tlength string) ShowcqtBuilder {
	return showcqtBuilder.withOption("tlength", expr.String(tlength))
}

func (showcqtBuilder *implShowcqtBuilder) Count(count int) ShowcqtBuilder {
	return showcqtBuilder.withOption("count", expr.Int(count))
}

func (showcqtBuilder *implShowcqtBuilder) Fcount(fcount int) ShowcqtBuilder {
	return showcqtBuilder.withOption("fcount", expr.Int(fcount))
}

func (showcqtBuilder *implShowcqtBuilder) Fontfile(fontfile string) ShowcqtBuilder {
	return showcqtBuilder.withOption("fontfile", expr.String(fontfile))
}

func (showcqtBuilder *implShowcqtBuilder) Font(font string) ShowcqtBuilder {
	return showcqtBuilder.withOption("font", expr.String(font))
}

func (showcqtBuilder *implShowcqtBuilder) Fontcolor(fontcolor string) ShowcqtBuilder {
	return showcqtBuilder.withOption("fontcolor", expr.String(fontcolor))
}

func (showcqtBuilder *implShowcqtBuilder) Axisfile(axisfile string) ShowcqtBuilder {
	return showcqtBuilder.withOption("axisfile", expr.String(axisfile))
}

func (showcqtBuilder *implShowcqtBuilder) Axis(axis bool) ShowcqtBuilder {
	return showcqtBuilder.withOption("axis", expr.Bool(axis))
}

func (showcqtBuilder *implShowcqtBuilder) Text(text bool) ShowcqtBuilder {
	return showcqtBuilder.withOption("text", expr.Bool(text))
}

func (showcqtBuilder *implShowcqtBuilder) Csp(csp int) ShowcqtBuilder {
	return showcqtBuilder.withOption("csp", expr.Int(csp))
}

func (showcqtBuilder *implShowcqtBuilder) Cscheme(cscheme string) ShowcqtBuilder {
	return showcqtBuilder.withOption("cscheme", expr.String(cscheme))
}
