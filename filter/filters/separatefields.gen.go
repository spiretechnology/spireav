// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SeparatefieldsBuilder Split input video frames into fields.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#separatefields
type SeparatefieldsBuilder interface {
	filter.Filter
}

// Separatefields creates a new SeparatefieldsBuilder to configure the "separatefields" filter.
func Separatefields(opts ...filter.Option) SeparatefieldsBuilder {
	return &implSeparatefieldsBuilder{
		f: filter.New("separatefields", 1, opts...),
	}
}

type implSeparatefieldsBuilder struct {
	f filter.Filter
}

func (separatefieldsBuilder *implSeparatefieldsBuilder) String() string {
	return separatefieldsBuilder.f.String()
}

func (separatefieldsBuilder *implSeparatefieldsBuilder) Outputs() int {
	return separatefieldsBuilder.f.Outputs()
}

func (separatefieldsBuilder *implSeparatefieldsBuilder) With(key string, value expr.Expr) filter.Filter {
	return separatefieldsBuilder.withOption(key, value)
}

func (separatefieldsBuilder *implSeparatefieldsBuilder) withOption(key string, value expr.Expr) SeparatefieldsBuilder {
	bCopy := *separatefieldsBuilder
	bCopy.f = separatefieldsBuilder.f.With(key, value)
	return &bCopy
}
