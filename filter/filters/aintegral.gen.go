// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AintegralBuilder Compute integral of input audio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#aintegral
type AintegralBuilder interface {
	filter.Filter
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AintegralBuilder
}

// Aintegral creates a new AintegralBuilder to configure the "aintegral" filter.
func Aintegral(opts ...filter.Option) AintegralBuilder {
	return &implAintegralBuilder{
		f: filter.New("aintegral", 1, opts...),
	}
}

type implAintegralBuilder struct {
	f filter.Filter
}

func (aintegralBuilder *implAintegralBuilder) String() string {
	return aintegralBuilder.f.String()
}

func (aintegralBuilder *implAintegralBuilder) Outputs() int {
	return aintegralBuilder.f.Outputs()
}

func (aintegralBuilder *implAintegralBuilder) With(key string, value expr.Expr) filter.Filter {
	return aintegralBuilder.withOption(key, value)
}

func (aintegralBuilder *implAintegralBuilder) withOption(key string, value expr.Expr) AintegralBuilder {
	bCopy := *aintegralBuilder
	bCopy.f = aintegralBuilder.f.With(key, value)
	return &bCopy
}

func (aintegralBuilder *implAintegralBuilder) Enable(enable expr.Expr) AintegralBuilder {
	return aintegralBuilder.withOption("enable", enable)
}
