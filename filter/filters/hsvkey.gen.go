// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// HsvkeyBuilder Turns a certain HSV range into transparency. Operates on YUV colors.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#hsvkey
type HsvkeyBuilder interface {
	filter.Filter
	// Hue set the hue value (from -360 to 360) (default 0).
	Hue(hue float32) HsvkeyBuilder
	// HueExpr set the hue value (from -360 to 360) (default 0).
	HueExpr(hue expr.Expr) HsvkeyBuilder
	// Sat set the saturation value (from -1 to 1) (default 0).
	Sat(sat float32) HsvkeyBuilder
	// SatExpr set the saturation value (from -1 to 1) (default 0).
	SatExpr(sat expr.Expr) HsvkeyBuilder
	// Val set the value value (from -1 to 1) (default 0).
	Val(val float32) HsvkeyBuilder
	// ValExpr set the value value (from -1 to 1) (default 0).
	ValExpr(val expr.Expr) HsvkeyBuilder
	// Similarity set the hsvkey similarity value (from 1e-05 to 1) (default 0.01).
	Similarity(similarity float32) HsvkeyBuilder
	// SimilarityExpr set the hsvkey similarity value (from 1e-05 to 1) (default 0.01).
	SimilarityExpr(similarity expr.Expr) HsvkeyBuilder
	// Blend set the hsvkey blend value (from 0 to 1) (default 0).
	Blend(blend float32) HsvkeyBuilder
	// BlendExpr set the hsvkey blend value (from 0 to 1) (default 0).
	BlendExpr(blend expr.Expr) HsvkeyBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) HsvkeyBuilder
}

// Hsvkey creates a new HsvkeyBuilder to configure the "hsvkey" filter.
func Hsvkey(opts ...filter.Option) HsvkeyBuilder {
	return &implHsvkeyBuilder{
		f: filter.New("hsvkey", 1, opts...),
	}
}

type implHsvkeyBuilder struct {
	f filter.Filter
}

func (hsvkeyBuilder *implHsvkeyBuilder) String() string {
	return hsvkeyBuilder.f.String()
}

func (hsvkeyBuilder *implHsvkeyBuilder) Outputs() int {
	return hsvkeyBuilder.f.Outputs()
}

func (hsvkeyBuilder *implHsvkeyBuilder) With(key string, value expr.Expr) filter.Filter {
	return hsvkeyBuilder.withOption(key, value)
}

func (hsvkeyBuilder *implHsvkeyBuilder) withOption(key string, value expr.Expr) HsvkeyBuilder {
	bCopy := *hsvkeyBuilder
	bCopy.f = hsvkeyBuilder.f.With(key, value)
	return &bCopy
}

func (hsvkeyBuilder *implHsvkeyBuilder) Hue(hue float32) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("hue", expr.Float(hue))
}

func (hsvkeyBuilder *implHsvkeyBuilder) HueExpr(hue expr.Expr) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("hue", hue)
}

func (hsvkeyBuilder *implHsvkeyBuilder) Sat(sat float32) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("sat", expr.Float(sat))
}

func (hsvkeyBuilder *implHsvkeyBuilder) SatExpr(sat expr.Expr) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("sat", sat)
}

func (hsvkeyBuilder *implHsvkeyBuilder) Val(val float32) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("val", expr.Float(val))
}

func (hsvkeyBuilder *implHsvkeyBuilder) ValExpr(val expr.Expr) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("val", val)
}

func (hsvkeyBuilder *implHsvkeyBuilder) Similarity(similarity float32) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("similarity", expr.Float(similarity))
}

func (hsvkeyBuilder *implHsvkeyBuilder) SimilarityExpr(similarity expr.Expr) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("similarity", similarity)
}

func (hsvkeyBuilder *implHsvkeyBuilder) Blend(blend float32) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("blend", expr.Float(blend))
}

func (hsvkeyBuilder *implHsvkeyBuilder) BlendExpr(blend expr.Expr) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("blend", blend)
}

func (hsvkeyBuilder *implHsvkeyBuilder) Enable(enable expr.Expr) HsvkeyBuilder {
	return hsvkeyBuilder.withOption("enable", enable)
}
