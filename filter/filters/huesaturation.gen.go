// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// HuesaturationBuilder Apply hue-saturation-intensity adjustments.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#huesaturation
type HuesaturationBuilder interface {
	filter.Filter
	// Hue set the hue shift (from -180 to 180) (default 0).
	Hue(hue float32) HuesaturationBuilder
	// HueExpr set the hue shift (from -180 to 180) (default 0).
	HueExpr(hue expr.Expr) HuesaturationBuilder
	// Saturation set the saturation shift (from -1 to 1) (default 0).
	Saturation(saturation float32) HuesaturationBuilder
	// SaturationExpr set the saturation shift (from -1 to 1) (default 0).
	SaturationExpr(saturation expr.Expr) HuesaturationBuilder
	// Intensity set the intensity shift (from -1 to 1) (default 0).
	Intensity(intensity float32) HuesaturationBuilder
	// IntensityExpr set the intensity shift (from -1 to 1) (default 0).
	IntensityExpr(intensity expr.Expr) HuesaturationBuilder
	// Colors set colors range (default r+y+g+c+b+m+a).
	Colors(colors ...string) HuesaturationBuilder
	// ColorsExpr set colors range (default r+y+g+c+b+m+a).
	ColorsExpr(colors expr.Expr) HuesaturationBuilder
	// Strength set the filtering strength (from 0 to 100) (default 1).
	Strength(strength float32) HuesaturationBuilder
	// StrengthExpr set the filtering strength (from 0 to 100) (default 1).
	StrengthExpr(strength expr.Expr) HuesaturationBuilder
	// Rw set the red weight (from 0 to 1) (default 0.333).
	Rw(rw float32) HuesaturationBuilder
	// RwExpr set the red weight (from 0 to 1) (default 0.333).
	RwExpr(rw expr.Expr) HuesaturationBuilder
	// Gw set the green weight (from 0 to 1) (default 0.334).
	Gw(gw float32) HuesaturationBuilder
	// GwExpr set the green weight (from 0 to 1) (default 0.334).
	GwExpr(gw expr.Expr) HuesaturationBuilder
	// Bw set the blue weight (from 0 to 1) (default 0.333).
	Bw(bw float32) HuesaturationBuilder
	// BwExpr set the blue weight (from 0 to 1) (default 0.333).
	BwExpr(bw expr.Expr) HuesaturationBuilder
	// Lightness set the preserve lightness (default false).
	Lightness(lightness bool) HuesaturationBuilder
	// LightnessExpr set the preserve lightness (default false).
	LightnessExpr(lightness expr.Expr) HuesaturationBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) HuesaturationBuilder
}

// Huesaturation creates a new HuesaturationBuilder to configure the "huesaturation" filter.
func Huesaturation(opts ...filter.Option) HuesaturationBuilder {
	return &implHuesaturationBuilder{
		f: filter.New("huesaturation", 1, opts...),
	}
}

type implHuesaturationBuilder struct {
	f filter.Filter
}

func (huesaturationBuilder *implHuesaturationBuilder) String() string {
	return huesaturationBuilder.f.String()
}

func (huesaturationBuilder *implHuesaturationBuilder) Outputs() int {
	return huesaturationBuilder.f.Outputs()
}

func (huesaturationBuilder *implHuesaturationBuilder) With(key string, value expr.Expr) filter.Filter {
	return huesaturationBuilder.withOption(key, value)
}

func (huesaturationBuilder *implHuesaturationBuilder) withOption(key string, value expr.Expr) HuesaturationBuilder {
	bCopy := *huesaturationBuilder
	bCopy.f = huesaturationBuilder.f.With(key, value)
	return &bCopy
}

func (huesaturationBuilder *implHuesaturationBuilder) Hue(hue float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("hue", expr.Float(hue))
}

func (huesaturationBuilder *implHuesaturationBuilder) HueExpr(hue expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("hue", hue)
}

func (huesaturationBuilder *implHuesaturationBuilder) Saturation(saturation float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("saturation", expr.Float(saturation))
}

func (huesaturationBuilder *implHuesaturationBuilder) SaturationExpr(saturation expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("saturation", saturation)
}

func (huesaturationBuilder *implHuesaturationBuilder) Intensity(intensity float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("intensity", expr.Float(intensity))
}

func (huesaturationBuilder *implHuesaturationBuilder) IntensityExpr(intensity expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("intensity", intensity)
}

func (huesaturationBuilder *implHuesaturationBuilder) Colors(colors ...string) HuesaturationBuilder {
	return huesaturationBuilder.withOption("colors", expr.Flags(colors))
}

func (huesaturationBuilder *implHuesaturationBuilder) ColorsExpr(colors expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("colors", colors)
}

func (huesaturationBuilder *implHuesaturationBuilder) Strength(strength float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("strength", expr.Float(strength))
}

func (huesaturationBuilder *implHuesaturationBuilder) StrengthExpr(strength expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("strength", strength)
}

func (huesaturationBuilder *implHuesaturationBuilder) Rw(rw float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("rw", expr.Float(rw))
}

func (huesaturationBuilder *implHuesaturationBuilder) RwExpr(rw expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("rw", rw)
}

func (huesaturationBuilder *implHuesaturationBuilder) Gw(gw float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("gw", expr.Float(gw))
}

func (huesaturationBuilder *implHuesaturationBuilder) GwExpr(gw expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("gw", gw)
}

func (huesaturationBuilder *implHuesaturationBuilder) Bw(bw float32) HuesaturationBuilder {
	return huesaturationBuilder.withOption("bw", expr.Float(bw))
}

func (huesaturationBuilder *implHuesaturationBuilder) BwExpr(bw expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("bw", bw)
}

func (huesaturationBuilder *implHuesaturationBuilder) Lightness(lightness bool) HuesaturationBuilder {
	return huesaturationBuilder.withOption("lightness", expr.Bool(lightness))
}

func (huesaturationBuilder *implHuesaturationBuilder) LightnessExpr(lightness expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("lightness", lightness)
}

func (huesaturationBuilder *implHuesaturationBuilder) Enable(enable expr.Expr) HuesaturationBuilder {
	return huesaturationBuilder.withOption("enable", enable)
}
