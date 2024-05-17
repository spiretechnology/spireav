// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// EdgedetectBuilder Detect and draw edge.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#edgedetect
type EdgedetectBuilder interface {
	filter.Filter
	// High set high threshold (from 0 to 1) (default 0.196078).
	High(high float64) EdgedetectBuilder
	// Low set low threshold (from 0 to 1) (default 0.0784314).
	Low(low float64) EdgedetectBuilder
	// Mode set mode (from 0 to 2) (default wires).
	Mode(mode int) EdgedetectBuilder
	// Planes set planes to filter (default y+u+v+r+g+b).
	Planes(planes ...string) EdgedetectBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) EdgedetectBuilder
}

// Edgedetect creates a new EdgedetectBuilder to configure the "edgedetect" filter.
func Edgedetect(opts ...filter.Option) EdgedetectBuilder {
	return &implEdgedetectBuilder{
		f: filter.New("edgedetect", 1, opts...),
	}
}

type implEdgedetectBuilder struct {
	f filter.Filter
}

func (edgedetectBuilder *implEdgedetectBuilder) String() string {
	return edgedetectBuilder.f.String()
}

func (edgedetectBuilder *implEdgedetectBuilder) Outputs() int {
	return edgedetectBuilder.f.Outputs()
}

func (edgedetectBuilder *implEdgedetectBuilder) With(key string, value expr.Expr) filter.Filter {
	return edgedetectBuilder.withOption(key, value)
}

func (edgedetectBuilder *implEdgedetectBuilder) withOption(key string, value expr.Expr) EdgedetectBuilder {
	bCopy := *edgedetectBuilder
	bCopy.f = edgedetectBuilder.f.With(key, value)
	return &bCopy
}

func (edgedetectBuilder *implEdgedetectBuilder) High(high float64) EdgedetectBuilder {
	return edgedetectBuilder.withOption("high", expr.Double(high))
}

func (edgedetectBuilder *implEdgedetectBuilder) Low(low float64) EdgedetectBuilder {
	return edgedetectBuilder.withOption("low", expr.Double(low))
}

func (edgedetectBuilder *implEdgedetectBuilder) Mode(mode int) EdgedetectBuilder {
	return edgedetectBuilder.withOption("mode", expr.Int(mode))
}

func (edgedetectBuilder *implEdgedetectBuilder) Planes(planes ...string) EdgedetectBuilder {
	return edgedetectBuilder.withOption("planes", expr.Flags(planes))
}

func (edgedetectBuilder *implEdgedetectBuilder) Enable(enable expr.Expr) EdgedetectBuilder {
	return edgedetectBuilder.withOption("enable", enable)
}