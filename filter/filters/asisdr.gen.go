// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AsisdrBuilder Measure Audio Scale-Invariant Signal-to-Distortion Ratio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#asisdr
type AsisdrBuilder interface {
	filter.Filter
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AsisdrBuilder
}

// Asisdr creates a new AsisdrBuilder to configure the "asisdr" filter.
func Asisdr(opts ...filter.Option) AsisdrBuilder {
	return &implAsisdrBuilder{
		f: filter.New("asisdr", 1, opts...),
	}
}

type implAsisdrBuilder struct {
	f filter.Filter
}

func (asisdrBuilder *implAsisdrBuilder) String() string {
	return asisdrBuilder.f.String()
}

func (asisdrBuilder *implAsisdrBuilder) Outputs() int {
	return asisdrBuilder.f.Outputs()
}

func (asisdrBuilder *implAsisdrBuilder) With(key string, value expr.Expr) filter.Filter {
	return asisdrBuilder.withOption(key, value)
}

func (asisdrBuilder *implAsisdrBuilder) withOption(key string, value expr.Expr) AsisdrBuilder {
	bCopy := *asisdrBuilder
	bCopy.f = asisdrBuilder.f.With(key, value)
	return &bCopy
}

func (asisdrBuilder *implAsisdrBuilder) Enable(enable expr.Expr) AsisdrBuilder {
	return asisdrBuilder.withOption("enable", enable)
}