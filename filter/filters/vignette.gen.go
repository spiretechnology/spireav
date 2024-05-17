// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// VignetteBuilder Make or reverse a vignette effect.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#vignette
type VignetteBuilder interface {
	filter.Filter
	// Angle set lens angle (default "PI/5").
	Angle(angle string) VignetteBuilder
	// A set lens angle (default "PI/5").
	A(a string) VignetteBuilder
	// X0 set circle center position on x-axis (default "w/2").
	X0(x0 string) VignetteBuilder
	// Y0 set circle center position on y-axis (default "h/2").
	Y0(y0 string) VignetteBuilder
	// Mode set forward/backward mode (from 0 to 1) (default forward).
	Mode(mode int) VignetteBuilder
	// Eval specify when to evaluate expressions (from 0 to 1) (default init).
	Eval(eval int) VignetteBuilder
	// Dither set dithering (default true).
	Dither(dither bool) VignetteBuilder
	// Aspect set aspect ratio (from 0 to DBL_MAX) (default 1/1).
	Aspect(aspect expr.Rational) VignetteBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) VignetteBuilder
}

// Vignette creates a new VignetteBuilder to configure the "vignette" filter.
func Vignette(opts ...filter.Option) VignetteBuilder {
	return &implVignetteBuilder{
		f: filter.New("vignette", 1, opts...),
	}
}

type implVignetteBuilder struct {
	f filter.Filter
}

func (vignetteBuilder *implVignetteBuilder) String() string {
	return vignetteBuilder.f.String()
}

func (vignetteBuilder *implVignetteBuilder) Outputs() int {
	return vignetteBuilder.f.Outputs()
}

func (vignetteBuilder *implVignetteBuilder) With(key string, value expr.Expr) filter.Filter {
	return vignetteBuilder.withOption(key, value)
}

func (vignetteBuilder *implVignetteBuilder) withOption(key string, value expr.Expr) VignetteBuilder {
	bCopy := *vignetteBuilder
	bCopy.f = vignetteBuilder.f.With(key, value)
	return &bCopy
}

func (vignetteBuilder *implVignetteBuilder) Angle(angle string) VignetteBuilder {
	return vignetteBuilder.withOption("angle", expr.String(angle))
}

func (vignetteBuilder *implVignetteBuilder) A(a string) VignetteBuilder {
	return vignetteBuilder.withOption("a", expr.String(a))
}

func (vignetteBuilder *implVignetteBuilder) X0(x0 string) VignetteBuilder {
	return vignetteBuilder.withOption("x0", expr.String(x0))
}

func (vignetteBuilder *implVignetteBuilder) Y0(y0 string) VignetteBuilder {
	return vignetteBuilder.withOption("y0", expr.String(y0))
}

func (vignetteBuilder *implVignetteBuilder) Mode(mode int) VignetteBuilder {
	return vignetteBuilder.withOption("mode", expr.Int(mode))
}

func (vignetteBuilder *implVignetteBuilder) Eval(eval int) VignetteBuilder {
	return vignetteBuilder.withOption("eval", expr.Int(eval))
}

func (vignetteBuilder *implVignetteBuilder) Dither(dither bool) VignetteBuilder {
	return vignetteBuilder.withOption("dither", expr.Bool(dither))
}

func (vignetteBuilder *implVignetteBuilder) Aspect(aspect expr.Rational) VignetteBuilder {
	return vignetteBuilder.withOption("aspect", aspect)
}

func (vignetteBuilder *implVignetteBuilder) Enable(enable expr.Expr) VignetteBuilder {
	return vignetteBuilder.withOption("enable", enable)
}