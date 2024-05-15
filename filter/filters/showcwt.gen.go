// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ShowcwtBuilder Convert input audio to a CWT (Continuous Wavelet Transform) spectrum video output.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#showcwt
type ShowcwtBuilder interface {
	filter.Filter
	// Size set video size (default "640x512").
	Size(size expr.Size) ShowcwtBuilder
	// S set video size (default "640x512").
	S(s expr.Size) ShowcwtBuilder
	// Rate set video rate (default "25").
	Rate(rate string) ShowcwtBuilder
	// R set video rate (default "25").
	R(r string) ShowcwtBuilder
	// Scale set frequency scale (from 0 to 8) (default linear).
	Scale(scale int) ShowcwtBuilder
	// Iscale set intensity scale (from 0 to 4) (default log).
	Iscale(iscale int) ShowcwtBuilder
	// Min set minimum frequency (from 1 to 192000) (default 20).
	Min(min float32) ShowcwtBuilder
	// Max set maximum frequency (from 1 to 192000) (default 20000).
	Max(max float32) ShowcwtBuilder
	// Imin set minimum intensity (from 0 to 1) (default 0).
	Imin(imin float32) ShowcwtBuilder
	// Imax set maximum intensity (from 0 to 1) (default 1).
	Imax(imax float32) ShowcwtBuilder
	// Logb set logarithmic basis (from 0 to 1) (default 0.0001).
	Logb(logb float32) ShowcwtBuilder
	// Deviation set frequency deviation (from 0 to 100) (default 1).
	Deviation(deviation float32) ShowcwtBuilder
	// Pps set pixels per second (from 1 to 1024) (default 64).
	Pps(pps int) ShowcwtBuilder
	// Mode set output mode (from 0 to 4) (default magnitude).
	Mode(mode int) ShowcwtBuilder
	// Slide set slide mode (from 0 to 2) (default replace).
	Slide(slide int) ShowcwtBuilder
	// Direction set direction mode (from 0 to 3) (default lr).
	Direction(direction int) ShowcwtBuilder
	// Bar set bargraph ratio (from 0 to 1) (default 0).
	Bar(bar float32) ShowcwtBuilder
	// Rotation set color rotation (from -1 to 1) (default 0).
	Rotation(rotation float32) ShowcwtBuilder
}

// Showcwt creates a new ShowcwtBuilder to configure the "showcwt" filter.
func Showcwt(opts ...filter.Option) ShowcwtBuilder {
	return &implShowcwtBuilder{
		f: filter.New("showcwt", 1, opts...),
	}
}

type implShowcwtBuilder struct {
	f filter.Filter
}

func (showcwtBuilder *implShowcwtBuilder) String() string {
	return showcwtBuilder.f.String()
}

func (showcwtBuilder *implShowcwtBuilder) Outputs() int {
	return showcwtBuilder.f.Outputs()
}

func (showcwtBuilder *implShowcwtBuilder) With(key string, value expr.Expr) filter.Filter {
	return showcwtBuilder.withOption(key, value)
}

func (showcwtBuilder *implShowcwtBuilder) withOption(key string, value expr.Expr) ShowcwtBuilder {
	bCopy := *showcwtBuilder
	bCopy.f = showcwtBuilder.f.With(key, value)
	return &bCopy
}

func (showcwtBuilder *implShowcwtBuilder) Size(size expr.Size) ShowcwtBuilder {
	return showcwtBuilder.withOption("size", size)
}

func (showcwtBuilder *implShowcwtBuilder) S(s expr.Size) ShowcwtBuilder {
	return showcwtBuilder.withOption("s", s)
}

func (showcwtBuilder *implShowcwtBuilder) Rate(rate string) ShowcwtBuilder {
	return showcwtBuilder.withOption("rate", expr.String(rate))
}

func (showcwtBuilder *implShowcwtBuilder) R(r string) ShowcwtBuilder {
	return showcwtBuilder.withOption("r", expr.String(r))
}

func (showcwtBuilder *implShowcwtBuilder) Scale(scale int) ShowcwtBuilder {
	return showcwtBuilder.withOption("scale", expr.Int(scale))
}

func (showcwtBuilder *implShowcwtBuilder) Iscale(iscale int) ShowcwtBuilder {
	return showcwtBuilder.withOption("iscale", expr.Int(iscale))
}

func (showcwtBuilder *implShowcwtBuilder) Min(min float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("min", expr.Float(min))
}

func (showcwtBuilder *implShowcwtBuilder) Max(max float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("max", expr.Float(max))
}

func (showcwtBuilder *implShowcwtBuilder) Imin(imin float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("imin", expr.Float(imin))
}

func (showcwtBuilder *implShowcwtBuilder) Imax(imax float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("imax", expr.Float(imax))
}

func (showcwtBuilder *implShowcwtBuilder) Logb(logb float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("logb", expr.Float(logb))
}

func (showcwtBuilder *implShowcwtBuilder) Deviation(deviation float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("deviation", expr.Float(deviation))
}

func (showcwtBuilder *implShowcwtBuilder) Pps(pps int) ShowcwtBuilder {
	return showcwtBuilder.withOption("pps", expr.Int(pps))
}

func (showcwtBuilder *implShowcwtBuilder) Mode(mode int) ShowcwtBuilder {
	return showcwtBuilder.withOption("mode", expr.Int(mode))
}

func (showcwtBuilder *implShowcwtBuilder) Slide(slide int) ShowcwtBuilder {
	return showcwtBuilder.withOption("slide", expr.Int(slide))
}

func (showcwtBuilder *implShowcwtBuilder) Direction(direction int) ShowcwtBuilder {
	return showcwtBuilder.withOption("direction", expr.Int(direction))
}

func (showcwtBuilder *implShowcwtBuilder) Bar(bar float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("bar", expr.Float(bar))
}

func (showcwtBuilder *implShowcwtBuilder) Rotation(rotation float32) ShowcwtBuilder {
	return showcwtBuilder.withOption("rotation", expr.Float(rotation))
}
