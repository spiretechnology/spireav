// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PpBuilder Filter video using libpostproc.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#pp
type PpBuilder interface {
	filter.Filter
	// Subfilters set postprocess subfilters (default "de").
	Subfilters(subfilters string) PpBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) PpBuilder
}

// Pp creates a new PpBuilder to configure the "pp" filter.
func Pp(opts ...filter.Option) PpBuilder {
	return &implPpBuilder{
		f: filter.New("pp", 1, opts...),
	}
}

type implPpBuilder struct {
	f filter.Filter
}

func (ppBuilder *implPpBuilder) String() string {
	return ppBuilder.f.String()
}

func (ppBuilder *implPpBuilder) Outputs() int {
	return ppBuilder.f.Outputs()
}

func (ppBuilder *implPpBuilder) With(key string, value expr.Expr) filter.Filter {
	return ppBuilder.withOption(key, value)
}

func (ppBuilder *implPpBuilder) withOption(key string, value expr.Expr) PpBuilder {
	bCopy := *ppBuilder
	bCopy.f = ppBuilder.f.With(key, value)
	return &bCopy
}

func (ppBuilder *implPpBuilder) Subfilters(subfilters string) PpBuilder {
	return ppBuilder.withOption("subfilters", expr.String(subfilters))
}

func (ppBuilder *implPpBuilder) Enable(enable expr.Expr) PpBuilder {
	return ppBuilder.withOption("enable", enable)
}
