// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PanBuilder Remix channels with coefficients (panning).
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#pan
type PanBuilder interface {
	filter.Filter
}

// Pan creates a new PanBuilder to configure the "pan" filter.
func Pan(opts ...filter.Option) PanBuilder {
	return &implPanBuilder{
		f: filter.New("pan", 1, opts...),
	}
}

type implPanBuilder struct {
	f filter.Filter
}

func (panBuilder *implPanBuilder) String() string {
	return panBuilder.f.String()
}

func (panBuilder *implPanBuilder) Outputs() int {
	return panBuilder.f.Outputs()
}

func (panBuilder *implPanBuilder) With(key string, value expr.Expr) filter.Filter {
	return panBuilder.withOption(key, value)
}

func (panBuilder *implPanBuilder) withOption(key string, value expr.Expr) PanBuilder {
	bCopy := *panBuilder
	bCopy.f = panBuilder.f.With(key, value)
	return &bCopy
}
