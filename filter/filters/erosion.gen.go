// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ErosionBuilder Apply erosion effect.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#erosion
type ErosionBuilder interface {
	filter.Filter
	// Coordinates set coordinates (from 0 to 255) (default 255).
	Coordinates(coordinates int) ErosionBuilder
	// CoordinatesExpr set coordinates (from 0 to 255) (default 255).
	CoordinatesExpr(coordinates expr.Expr) ErosionBuilder
	// Threshold0 set threshold for 1st plane (from 0 to 65535) (default 65535).
	Threshold0(threshold0 int) ErosionBuilder
	// Threshold0Expr set threshold for 1st plane (from 0 to 65535) (default 65535).
	Threshold0Expr(threshold0 expr.Expr) ErosionBuilder
	// Threshold1 set threshold for 2nd plane (from 0 to 65535) (default 65535).
	Threshold1(threshold1 int) ErosionBuilder
	// Threshold1Expr set threshold for 2nd plane (from 0 to 65535) (default 65535).
	Threshold1Expr(threshold1 expr.Expr) ErosionBuilder
	// Threshold2 set threshold for 3rd plane (from 0 to 65535) (default 65535).
	Threshold2(threshold2 int) ErosionBuilder
	// Threshold2Expr set threshold for 3rd plane (from 0 to 65535) (default 65535).
	Threshold2Expr(threshold2 expr.Expr) ErosionBuilder
	// Threshold3 set threshold for 4th plane (from 0 to 65535) (default 65535).
	Threshold3(threshold3 int) ErosionBuilder
	// Threshold3Expr set threshold for 4th plane (from 0 to 65535) (default 65535).
	Threshold3Expr(threshold3 expr.Expr) ErosionBuilder
}

// Erosion creates a new ErosionBuilder to configure the "erosion" filter.
func Erosion(opts ...filter.Option) ErosionBuilder {
	return &implErosionBuilder{
		f: filter.New("erosion", 1, opts...),
	}
}

type implErosionBuilder struct {
	f filter.Filter
}

func (erosionBuilder *implErosionBuilder) String() string {
	return erosionBuilder.f.String()
}

func (erosionBuilder *implErosionBuilder) Outputs() int {
	return erosionBuilder.f.Outputs()
}

func (erosionBuilder *implErosionBuilder) With(key string, value expr.Expr) filter.Filter {
	return erosionBuilder.withOption(key, value)
}

func (erosionBuilder *implErosionBuilder) withOption(key string, value expr.Expr) ErosionBuilder {
	bCopy := *erosionBuilder
	bCopy.f = erosionBuilder.f.With(key, value)
	return &bCopy
}

func (erosionBuilder *implErosionBuilder) Coordinates(coordinates int) ErosionBuilder {
	return erosionBuilder.withOption("coordinates", expr.Int(coordinates))
}

func (erosionBuilder *implErosionBuilder) CoordinatesExpr(coordinates expr.Expr) ErosionBuilder {
	return erosionBuilder.withOption("coordinates", coordinates)
}

func (erosionBuilder *implErosionBuilder) Threshold0(threshold0 int) ErosionBuilder {
	return erosionBuilder.withOption("threshold0", expr.Int(threshold0))
}

func (erosionBuilder *implErosionBuilder) Threshold0Expr(threshold0 expr.Expr) ErosionBuilder {
	return erosionBuilder.withOption("threshold0", threshold0)
}

func (erosionBuilder *implErosionBuilder) Threshold1(threshold1 int) ErosionBuilder {
	return erosionBuilder.withOption("threshold1", expr.Int(threshold1))
}

func (erosionBuilder *implErosionBuilder) Threshold1Expr(threshold1 expr.Expr) ErosionBuilder {
	return erosionBuilder.withOption("threshold1", threshold1)
}

func (erosionBuilder *implErosionBuilder) Threshold2(threshold2 int) ErosionBuilder {
	return erosionBuilder.withOption("threshold2", expr.Int(threshold2))
}

func (erosionBuilder *implErosionBuilder) Threshold2Expr(threshold2 expr.Expr) ErosionBuilder {
	return erosionBuilder.withOption("threshold2", threshold2)
}

func (erosionBuilder *implErosionBuilder) Threshold3(threshold3 int) ErosionBuilder {
	return erosionBuilder.withOption("threshold3", expr.Int(threshold3))
}

func (erosionBuilder *implErosionBuilder) Threshold3Expr(threshold3 expr.Expr) ErosionBuilder {
	return erosionBuilder.withOption("threshold3", threshold3)
}
