// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// GrayworldBuilder Adjust white balance using LAB gray world algorithm.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#grayworld
type GrayworldBuilder interface {
	filter.Filter
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) GrayworldBuilder
}

// Grayworld creates a new GrayworldBuilder to configure the "grayworld" filter.
func Grayworld(opts ...filter.Option) GrayworldBuilder {
	return &implGrayworldBuilder{
		f: filter.New("grayworld", 1, opts...),
	}
}

type implGrayworldBuilder struct {
	f filter.Filter
}

func (grayworldBuilder *implGrayworldBuilder) String() string {
	return grayworldBuilder.f.String()
}

func (grayworldBuilder *implGrayworldBuilder) Outputs() int {
	return grayworldBuilder.f.Outputs()
}

func (grayworldBuilder *implGrayworldBuilder) With(key string, value expr.Expr) filter.Filter {
	return grayworldBuilder.withOption(key, value)
}

func (grayworldBuilder *implGrayworldBuilder) withOption(key string, value expr.Expr) GrayworldBuilder {
	bCopy := *grayworldBuilder
	bCopy.f = grayworldBuilder.f.With(key, value)
	return &bCopy
}

func (grayworldBuilder *implGrayworldBuilder) Enable(enable expr.Expr) GrayworldBuilder {
	return grayworldBuilder.withOption("enable", enable)
}