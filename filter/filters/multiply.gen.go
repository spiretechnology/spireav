// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// MultiplyBuilder Multiply first video stream with second video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#multiply
type MultiplyBuilder interface {
	filter.Filter
	// Scale set scale (from 0 to 9) (default 1).
	Scale(scale float32) MultiplyBuilder
	// ScaleExpr set scale (from 0 to 9) (default 1).
	ScaleExpr(scale expr.Expr) MultiplyBuilder
	// Offset set offset (from -1 to 1) (default 0.5).
	Offset(offset float32) MultiplyBuilder
	// OffsetExpr set offset (from -1 to 1) (default 0.5).
	OffsetExpr(offset expr.Expr) MultiplyBuilder
	// Planes set planes (default F).
	Planes(planes ...string) MultiplyBuilder
	// PlanesExpr set planes (default F).
	PlanesExpr(planes expr.Expr) MultiplyBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) MultiplyBuilder
}

// Multiply creates a new MultiplyBuilder to configure the "multiply" filter.
func Multiply(opts ...filter.Option) MultiplyBuilder {
	return &implMultiplyBuilder{
		f: filter.New("multiply", 1, opts...),
	}
}

type implMultiplyBuilder struct {
	f filter.Filter
}

func (multiplyBuilder *implMultiplyBuilder) String() string {
	return multiplyBuilder.f.String()
}

func (multiplyBuilder *implMultiplyBuilder) Outputs() int {
	return multiplyBuilder.f.Outputs()
}

func (multiplyBuilder *implMultiplyBuilder) With(key string, value expr.Expr) filter.Filter {
	return multiplyBuilder.withOption(key, value)
}

func (multiplyBuilder *implMultiplyBuilder) withOption(key string, value expr.Expr) MultiplyBuilder {
	bCopy := *multiplyBuilder
	bCopy.f = multiplyBuilder.f.With(key, value)
	return &bCopy
}

func (multiplyBuilder *implMultiplyBuilder) Scale(scale float32) MultiplyBuilder {
	return multiplyBuilder.withOption("scale", expr.Float(scale))
}

func (multiplyBuilder *implMultiplyBuilder) ScaleExpr(scale expr.Expr) MultiplyBuilder {
	return multiplyBuilder.withOption("scale", scale)
}

func (multiplyBuilder *implMultiplyBuilder) Offset(offset float32) MultiplyBuilder {
	return multiplyBuilder.withOption("offset", expr.Float(offset))
}

func (multiplyBuilder *implMultiplyBuilder) OffsetExpr(offset expr.Expr) MultiplyBuilder {
	return multiplyBuilder.withOption("offset", offset)
}

func (multiplyBuilder *implMultiplyBuilder) Planes(planes ...string) MultiplyBuilder {
	return multiplyBuilder.withOption("planes", expr.Flags(planes))
}

func (multiplyBuilder *implMultiplyBuilder) PlanesExpr(planes expr.Expr) MultiplyBuilder {
	return multiplyBuilder.withOption("planes", planes)
}

func (multiplyBuilder *implMultiplyBuilder) Enable(enable expr.Expr) MultiplyBuilder {
	return multiplyBuilder.withOption("enable", enable)
}
