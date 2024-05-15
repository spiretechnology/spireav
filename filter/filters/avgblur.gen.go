// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AvgblurBuilder Apply Average Blur filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#avgblur
type AvgblurBuilder interface {
	filter.Filter
	// SizeX set horizontal size (from 1 to 1024) (default 1).
	SizeX(sizeX int) AvgblurBuilder
	// SizeXExpr set horizontal size (from 1 to 1024) (default 1).
	SizeXExpr(sizeX expr.Expr) AvgblurBuilder
	// Planes set planes to filter (from 0 to 15) (default 15).
	Planes(planes int) AvgblurBuilder
	// PlanesExpr set planes to filter (from 0 to 15) (default 15).
	PlanesExpr(planes expr.Expr) AvgblurBuilder
	// SizeY set vertical size (from 0 to 1024) (default 0).
	SizeY(sizeY int) AvgblurBuilder
	// SizeYExpr set vertical size (from 0 to 1024) (default 0).
	SizeYExpr(sizeY expr.Expr) AvgblurBuilder
}

// Avgblur creates a new AvgblurBuilder to configure the "avgblur" filter.
func Avgblur(opts ...filter.Option) AvgblurBuilder {
	return &implAvgblurBuilder{
		f: filter.New("avgblur", 1, opts...),
	}
}

type implAvgblurBuilder struct {
	f filter.Filter
}

func (avgblurBuilder *implAvgblurBuilder) String() string {
	return avgblurBuilder.f.String()
}

func (avgblurBuilder *implAvgblurBuilder) Outputs() int {
	return avgblurBuilder.f.Outputs()
}

func (avgblurBuilder *implAvgblurBuilder) With(key string, value expr.Expr) filter.Filter {
	return avgblurBuilder.withOption(key, value)
}

func (avgblurBuilder *implAvgblurBuilder) withOption(key string, value expr.Expr) AvgblurBuilder {
	bCopy := *avgblurBuilder
	bCopy.f = avgblurBuilder.f.With(key, value)
	return &bCopy
}

func (avgblurBuilder *implAvgblurBuilder) SizeX(sizeX int) AvgblurBuilder {
	return avgblurBuilder.withOption("sizeX", expr.Int(sizeX))
}

func (avgblurBuilder *implAvgblurBuilder) SizeXExpr(sizeX expr.Expr) AvgblurBuilder {
	return avgblurBuilder.withOption("sizeX", sizeX)
}

func (avgblurBuilder *implAvgblurBuilder) Planes(planes int) AvgblurBuilder {
	return avgblurBuilder.withOption("planes", expr.Int(planes))
}

func (avgblurBuilder *implAvgblurBuilder) PlanesExpr(planes expr.Expr) AvgblurBuilder {
	return avgblurBuilder.withOption("planes", planes)
}

func (avgblurBuilder *implAvgblurBuilder) SizeY(sizeY int) AvgblurBuilder {
	return avgblurBuilder.withOption("sizeY", expr.Int(sizeY))
}

func (avgblurBuilder *implAvgblurBuilder) SizeYExpr(sizeY expr.Expr) AvgblurBuilder {
	return avgblurBuilder.withOption("sizeY", sizeY)
}
