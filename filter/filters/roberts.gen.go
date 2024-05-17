// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// RobertsBuilder Apply roberts cross operator.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#roberts
type RobertsBuilder interface {
	filter.Filter
	// Planes set planes to filter (from 0 to 15) (default 15).
	Planes(planes int) RobertsBuilder
	// PlanesExpr set planes to filter (from 0 to 15) (default 15).
	PlanesExpr(planes expr.Expr) RobertsBuilder
	// Scale set scale (from 0 to 65535) (default 1).
	Scale(scale float32) RobertsBuilder
	// ScaleExpr set scale (from 0 to 65535) (default 1).
	ScaleExpr(scale expr.Expr) RobertsBuilder
	// Delta set delta (from -65535 to 65535) (default 0).
	Delta(delta float32) RobertsBuilder
	// DeltaExpr set delta (from -65535 to 65535) (default 0).
	DeltaExpr(delta expr.Expr) RobertsBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) RobertsBuilder
}

// Roberts creates a new RobertsBuilder to configure the "roberts" filter.
func Roberts(opts ...filter.Option) RobertsBuilder {
	return &implRobertsBuilder{
		f: filter.New("roberts", 1, opts...),
	}
}

type implRobertsBuilder struct {
	f filter.Filter
}

func (robertsBuilder *implRobertsBuilder) String() string {
	return robertsBuilder.f.String()
}

func (robertsBuilder *implRobertsBuilder) Outputs() int {
	return robertsBuilder.f.Outputs()
}

func (robertsBuilder *implRobertsBuilder) With(key string, value expr.Expr) filter.Filter {
	return robertsBuilder.withOption(key, value)
}

func (robertsBuilder *implRobertsBuilder) withOption(key string, value expr.Expr) RobertsBuilder {
	bCopy := *robertsBuilder
	bCopy.f = robertsBuilder.f.With(key, value)
	return &bCopy
}

func (robertsBuilder *implRobertsBuilder) Planes(planes int) RobertsBuilder {
	return robertsBuilder.withOption("planes", expr.Int(planes))
}

func (robertsBuilder *implRobertsBuilder) PlanesExpr(planes expr.Expr) RobertsBuilder {
	return robertsBuilder.withOption("planes", planes)
}

func (robertsBuilder *implRobertsBuilder) Scale(scale float32) RobertsBuilder {
	return robertsBuilder.withOption("scale", expr.Float(scale))
}

func (robertsBuilder *implRobertsBuilder) ScaleExpr(scale expr.Expr) RobertsBuilder {
	return robertsBuilder.withOption("scale", scale)
}

func (robertsBuilder *implRobertsBuilder) Delta(delta float32) RobertsBuilder {
	return robertsBuilder.withOption("delta", expr.Float(delta))
}

func (robertsBuilder *implRobertsBuilder) DeltaExpr(delta expr.Expr) RobertsBuilder {
	return robertsBuilder.withOption("delta", delta)
}

func (robertsBuilder *implRobertsBuilder) Enable(enable expr.Expr) RobertsBuilder {
	return robertsBuilder.withOption("enable", enable)
}
