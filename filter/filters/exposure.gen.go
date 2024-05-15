// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ExposureBuilder Adjust exposure of the video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#exposure
type ExposureBuilder interface {
	filter.Filter
	// Exposure set the exposure correction (from -3 to 3) (default 0).
	Exposure(exposure float32) ExposureBuilder
	// ExposureExpr set the exposure correction (from -3 to 3) (default 0).
	ExposureExpr(exposure expr.Expr) ExposureBuilder
	// Black set the black level correction (from -1 to 1) (default 0).
	Black(black float32) ExposureBuilder
	// BlackExpr set the black level correction (from -1 to 1) (default 0).
	BlackExpr(black expr.Expr) ExposureBuilder
}

// Exposure creates a new ExposureBuilder to configure the "exposure" filter.
func Exposure(opts ...filter.Option) ExposureBuilder {
	return &implExposureBuilder{
		f: filter.New("exposure", 1, opts...),
	}
}

type implExposureBuilder struct {
	f filter.Filter
}

func (exposureBuilder *implExposureBuilder) String() string {
	return exposureBuilder.f.String()
}

func (exposureBuilder *implExposureBuilder) Outputs() int {
	return exposureBuilder.f.Outputs()
}

func (exposureBuilder *implExposureBuilder) With(key string, value expr.Expr) filter.Filter {
	return exposureBuilder.withOption(key, value)
}

func (exposureBuilder *implExposureBuilder) withOption(key string, value expr.Expr) ExposureBuilder {
	bCopy := *exposureBuilder
	bCopy.f = exposureBuilder.f.With(key, value)
	return &bCopy
}

func (exposureBuilder *implExposureBuilder) Exposure(exposure float32) ExposureBuilder {
	return exposureBuilder.withOption("exposure", expr.Float(exposure))
}

func (exposureBuilder *implExposureBuilder) ExposureExpr(exposure expr.Expr) ExposureBuilder {
	return exposureBuilder.withOption("exposure", exposure)
}

func (exposureBuilder *implExposureBuilder) Black(black float32) ExposureBuilder {
	return exposureBuilder.withOption("black", expr.Float(black))
}

func (exposureBuilder *implExposureBuilder) BlackExpr(black expr.Expr) ExposureBuilder {
	return exposureBuilder.withOption("black", black)
}
