// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AcontrastBuilder Simple audio dynamic range compression/expansion filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#acontrast
type AcontrastBuilder interface {
	filter.Filter
	// Contrast set contrast (from 0 to 100) (default 33).
	Contrast(contrast float32) AcontrastBuilder
}

// Acontrast creates a new AcontrastBuilder to configure the "acontrast" filter.
func Acontrast(opts ...filter.Option) AcontrastBuilder {
	return &implAcontrastBuilder{
		f: filter.New("acontrast", 1, opts...),
	}
}

type implAcontrastBuilder struct {
	f filter.Filter
}

func (acontrastBuilder *implAcontrastBuilder) String() string {
	return acontrastBuilder.f.String()
}

func (acontrastBuilder *implAcontrastBuilder) Outputs() int {
	return acontrastBuilder.f.Outputs()
}

func (acontrastBuilder *implAcontrastBuilder) With(key string, value expr.Expr) filter.Filter {
	return acontrastBuilder.withOption(key, value)
}

func (acontrastBuilder *implAcontrastBuilder) withOption(key string, value expr.Expr) AcontrastBuilder {
	bCopy := *acontrastBuilder
	bCopy.f = acontrastBuilder.f.With(key, value)
	return &bCopy
}

func (acontrastBuilder *implAcontrastBuilder) Contrast(contrast float32) AcontrastBuilder {
	return acontrastBuilder.withOption("contrast", expr.Float(contrast))
}