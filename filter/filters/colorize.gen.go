// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColorizeBuilder Overlay a solid color on the video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colorize
type ColorizeBuilder interface {
	filter.Filter
	// Hue set the hue (from 0 to 360) (default 0).
	Hue(hue float32) ColorizeBuilder
	// HueExpr set the hue (from 0 to 360) (default 0).
	HueExpr(hue expr.Expr) ColorizeBuilder
	// Saturation set the saturation (from 0 to 1) (default 0.5).
	Saturation(saturation float32) ColorizeBuilder
	// SaturationExpr set the saturation (from 0 to 1) (default 0.5).
	SaturationExpr(saturation expr.Expr) ColorizeBuilder
	// Lightness set the lightness (from 0 to 1) (default 0.5).
	Lightness(lightness float32) ColorizeBuilder
	// LightnessExpr set the lightness (from 0 to 1) (default 0.5).
	LightnessExpr(lightness expr.Expr) ColorizeBuilder
	// Mix set the mix of source lightness (from 0 to 1) (default 1).
	Mix(mix float32) ColorizeBuilder
	// MixExpr set the mix of source lightness (from 0 to 1) (default 1).
	MixExpr(mix expr.Expr) ColorizeBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ColorizeBuilder
}

// Colorize creates a new ColorizeBuilder to configure the "colorize" filter.
func Colorize(opts ...filter.Option) ColorizeBuilder {
	return &implColorizeBuilder{
		f: filter.New("colorize", 1, opts...),
	}
}

type implColorizeBuilder struct {
	f filter.Filter
}

func (colorizeBuilder *implColorizeBuilder) String() string {
	return colorizeBuilder.f.String()
}

func (colorizeBuilder *implColorizeBuilder) Outputs() int {
	return colorizeBuilder.f.Outputs()
}

func (colorizeBuilder *implColorizeBuilder) With(key string, value expr.Expr) filter.Filter {
	return colorizeBuilder.withOption(key, value)
}

func (colorizeBuilder *implColorizeBuilder) withOption(key string, value expr.Expr) ColorizeBuilder {
	bCopy := *colorizeBuilder
	bCopy.f = colorizeBuilder.f.With(key, value)
	return &bCopy
}

func (colorizeBuilder *implColorizeBuilder) Hue(hue float32) ColorizeBuilder {
	return colorizeBuilder.withOption("hue", expr.Float(hue))
}

func (colorizeBuilder *implColorizeBuilder) HueExpr(hue expr.Expr) ColorizeBuilder {
	return colorizeBuilder.withOption("hue", hue)
}

func (colorizeBuilder *implColorizeBuilder) Saturation(saturation float32) ColorizeBuilder {
	return colorizeBuilder.withOption("saturation", expr.Float(saturation))
}

func (colorizeBuilder *implColorizeBuilder) SaturationExpr(saturation expr.Expr) ColorizeBuilder {
	return colorizeBuilder.withOption("saturation", saturation)
}

func (colorizeBuilder *implColorizeBuilder) Lightness(lightness float32) ColorizeBuilder {
	return colorizeBuilder.withOption("lightness", expr.Float(lightness))
}

func (colorizeBuilder *implColorizeBuilder) LightnessExpr(lightness expr.Expr) ColorizeBuilder {
	return colorizeBuilder.withOption("lightness", lightness)
}

func (colorizeBuilder *implColorizeBuilder) Mix(mix float32) ColorizeBuilder {
	return colorizeBuilder.withOption("mix", expr.Float(mix))
}

func (colorizeBuilder *implColorizeBuilder) MixExpr(mix expr.Expr) ColorizeBuilder {
	return colorizeBuilder.withOption("mix", mix)
}

func (colorizeBuilder *implColorizeBuilder) Enable(enable expr.Expr) ColorizeBuilder {
	return colorizeBuilder.withOption("enable", enable)
}
