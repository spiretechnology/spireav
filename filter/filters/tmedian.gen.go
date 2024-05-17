// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TmedianBuilder Pick median pixels from successive frames.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#tmedian
type TmedianBuilder interface {
	filter.Filter
	// Radius set median filter radius (from 1 to 127) (default 1).
	Radius(radius int) TmedianBuilder
	// Planes set planes to filter (from 0 to 15) (default 15).
	Planes(planes int) TmedianBuilder
	// PlanesExpr set planes to filter (from 0 to 15) (default 15).
	PlanesExpr(planes expr.Expr) TmedianBuilder
	// Percentile set percentile (from 0 to 1) (default 0.5).
	Percentile(percentile float32) TmedianBuilder
	// PercentileExpr set percentile (from 0 to 1) (default 0.5).
	PercentileExpr(percentile expr.Expr) TmedianBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) TmedianBuilder
}

// Tmedian creates a new TmedianBuilder to configure the "tmedian" filter.
func Tmedian(opts ...filter.Option) TmedianBuilder {
	return &implTmedianBuilder{
		f: filter.New("tmedian", 1, opts...),
	}
}

type implTmedianBuilder struct {
	f filter.Filter
}

func (tmedianBuilder *implTmedianBuilder) String() string {
	return tmedianBuilder.f.String()
}

func (tmedianBuilder *implTmedianBuilder) Outputs() int {
	return tmedianBuilder.f.Outputs()
}

func (tmedianBuilder *implTmedianBuilder) With(key string, value expr.Expr) filter.Filter {
	return tmedianBuilder.withOption(key, value)
}

func (tmedianBuilder *implTmedianBuilder) withOption(key string, value expr.Expr) TmedianBuilder {
	bCopy := *tmedianBuilder
	bCopy.f = tmedianBuilder.f.With(key, value)
	return &bCopy
}

func (tmedianBuilder *implTmedianBuilder) Radius(radius int) TmedianBuilder {
	return tmedianBuilder.withOption("radius", expr.Int(radius))
}

func (tmedianBuilder *implTmedianBuilder) Planes(planes int) TmedianBuilder {
	return tmedianBuilder.withOption("planes", expr.Int(planes))
}

func (tmedianBuilder *implTmedianBuilder) PlanesExpr(planes expr.Expr) TmedianBuilder {
	return tmedianBuilder.withOption("planes", planes)
}

func (tmedianBuilder *implTmedianBuilder) Percentile(percentile float32) TmedianBuilder {
	return tmedianBuilder.withOption("percentile", expr.Float(percentile))
}

func (tmedianBuilder *implTmedianBuilder) PercentileExpr(percentile expr.Expr) TmedianBuilder {
	return tmedianBuilder.withOption("percentile", percentile)
}

func (tmedianBuilder *implTmedianBuilder) Enable(enable expr.Expr) TmedianBuilder {
	return tmedianBuilder.withOption("enable", enable)
}