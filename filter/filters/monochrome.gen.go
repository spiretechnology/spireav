// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// MonochromeBuilder Convert video to gray using custom color filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#monochrome
type MonochromeBuilder interface {
	filter.Filter
	// Cb set the chroma blue spot (from -1 to 1) (default 0).
	Cb(cb float32) MonochromeBuilder
	// CbExpr set the chroma blue spot (from -1 to 1) (default 0).
	CbExpr(cb expr.Expr) MonochromeBuilder
	// Cr set the chroma red spot (from -1 to 1) (default 0).
	Cr(cr float32) MonochromeBuilder
	// CrExpr set the chroma red spot (from -1 to 1) (default 0).
	CrExpr(cr expr.Expr) MonochromeBuilder
	// Size set the color filter size (from 0.1 to 10) (default 1).
	Size(size float32) MonochromeBuilder
	// SizeExpr set the color filter size (from 0.1 to 10) (default 1).
	SizeExpr(size expr.Expr) MonochromeBuilder
	// High set the highlights strength (from 0 to 1) (default 0).
	High(high float32) MonochromeBuilder
	// HighExpr set the highlights strength (from 0 to 1) (default 0).
	HighExpr(high expr.Expr) MonochromeBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) MonochromeBuilder
}

// Monochrome creates a new MonochromeBuilder to configure the "monochrome" filter.
func Monochrome(opts ...filter.Option) MonochromeBuilder {
	return &implMonochromeBuilder{
		f: filter.New("monochrome", 1, opts...),
	}
}

type implMonochromeBuilder struct {
	f filter.Filter
}

func (monochromeBuilder *implMonochromeBuilder) String() string {
	return monochromeBuilder.f.String()
}

func (monochromeBuilder *implMonochromeBuilder) Outputs() int {
	return monochromeBuilder.f.Outputs()
}

func (monochromeBuilder *implMonochromeBuilder) With(key string, value expr.Expr) filter.Filter {
	return monochromeBuilder.withOption(key, value)
}

func (monochromeBuilder *implMonochromeBuilder) withOption(key string, value expr.Expr) MonochromeBuilder {
	bCopy := *monochromeBuilder
	bCopy.f = monochromeBuilder.f.With(key, value)
	return &bCopy
}

func (monochromeBuilder *implMonochromeBuilder) Cb(cb float32) MonochromeBuilder {
	return monochromeBuilder.withOption("cb", expr.Float(cb))
}

func (monochromeBuilder *implMonochromeBuilder) CbExpr(cb expr.Expr) MonochromeBuilder {
	return monochromeBuilder.withOption("cb", cb)
}

func (monochromeBuilder *implMonochromeBuilder) Cr(cr float32) MonochromeBuilder {
	return monochromeBuilder.withOption("cr", expr.Float(cr))
}

func (monochromeBuilder *implMonochromeBuilder) CrExpr(cr expr.Expr) MonochromeBuilder {
	return monochromeBuilder.withOption("cr", cr)
}

func (monochromeBuilder *implMonochromeBuilder) Size(size float32) MonochromeBuilder {
	return monochromeBuilder.withOption("size", expr.Float(size))
}

func (monochromeBuilder *implMonochromeBuilder) SizeExpr(size expr.Expr) MonochromeBuilder {
	return monochromeBuilder.withOption("size", size)
}

func (monochromeBuilder *implMonochromeBuilder) High(high float32) MonochromeBuilder {
	return monochromeBuilder.withOption("high", expr.Float(high))
}

func (monochromeBuilder *implMonochromeBuilder) HighExpr(high expr.Expr) MonochromeBuilder {
	return monochromeBuilder.withOption("high", high)
}

func (monochromeBuilder *implMonochromeBuilder) Enable(enable expr.Expr) MonochromeBuilder {
	return monochromeBuilder.withOption("enable", enable)
}
