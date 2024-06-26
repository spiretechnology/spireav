// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PrewittBuilder Apply prewitt operator.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#prewitt
type PrewittBuilder interface {
	filter.Filter
	// Planes set planes to filter (from 0 to 15) (default 15).
	Planes(planes int) PrewittBuilder
	// PlanesExpr set planes to filter (from 0 to 15) (default 15).
	PlanesExpr(planes expr.Expr) PrewittBuilder
	// Scale set scale (from 0 to 65535) (default 1).
	Scale(scale float32) PrewittBuilder
	// ScaleExpr set scale (from 0 to 65535) (default 1).
	ScaleExpr(scale expr.Expr) PrewittBuilder
	// Delta set delta (from -65535 to 65535) (default 0).
	Delta(delta float32) PrewittBuilder
	// DeltaExpr set delta (from -65535 to 65535) (default 0).
	DeltaExpr(delta expr.Expr) PrewittBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) PrewittBuilder
}

// Prewitt creates a new PrewittBuilder to configure the "prewitt" filter.
func Prewitt(opts ...filter.Option) PrewittBuilder {
	return &implPrewittBuilder{
		f: filter.New("prewitt", 1, opts...),
	}
}

type implPrewittBuilder struct {
	f filter.Filter
}

func (prewittBuilder *implPrewittBuilder) String() string {
	return prewittBuilder.f.String()
}

func (prewittBuilder *implPrewittBuilder) Outputs() int {
	return prewittBuilder.f.Outputs()
}

func (prewittBuilder *implPrewittBuilder) With(key string, value expr.Expr) filter.Filter {
	return prewittBuilder.withOption(key, value)
}

func (prewittBuilder *implPrewittBuilder) withOption(key string, value expr.Expr) PrewittBuilder {
	bCopy := *prewittBuilder
	bCopy.f = prewittBuilder.f.With(key, value)
	return &bCopy
}

func (prewittBuilder *implPrewittBuilder) Planes(planes int) PrewittBuilder {
	return prewittBuilder.withOption("planes", expr.Int(planes))
}

func (prewittBuilder *implPrewittBuilder) PlanesExpr(planes expr.Expr) PrewittBuilder {
	return prewittBuilder.withOption("planes", planes)
}

func (prewittBuilder *implPrewittBuilder) Scale(scale float32) PrewittBuilder {
	return prewittBuilder.withOption("scale", expr.Float(scale))
}

func (prewittBuilder *implPrewittBuilder) ScaleExpr(scale expr.Expr) PrewittBuilder {
	return prewittBuilder.withOption("scale", scale)
}

func (prewittBuilder *implPrewittBuilder) Delta(delta float32) PrewittBuilder {
	return prewittBuilder.withOption("delta", expr.Float(delta))
}

func (prewittBuilder *implPrewittBuilder) DeltaExpr(delta expr.Expr) PrewittBuilder {
	return prewittBuilder.withOption("delta", delta)
}

func (prewittBuilder *implPrewittBuilder) Enable(enable expr.Expr) PrewittBuilder {
	return prewittBuilder.withOption("enable", enable)
}
