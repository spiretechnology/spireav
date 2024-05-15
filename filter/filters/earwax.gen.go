// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// EarwaxBuilder Widen the stereo image.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#earwax
type EarwaxBuilder interface {
	filter.Filter
}

// Earwax creates a new EarwaxBuilder to configure the "earwax" filter.
func Earwax(opts ...filter.Option) EarwaxBuilder {
	return &implEarwaxBuilder{
		f: filter.New("earwax", 1, opts...),
	}
}

type implEarwaxBuilder struct {
	f filter.Filter
}

func (earwaxBuilder *implEarwaxBuilder) String() string {
	return earwaxBuilder.f.String()
}

func (earwaxBuilder *implEarwaxBuilder) Outputs() int {
	return earwaxBuilder.f.Outputs()
}

func (earwaxBuilder *implEarwaxBuilder) With(key string, value expr.Expr) filter.Filter {
	return earwaxBuilder.withOption(key, value)
}

func (earwaxBuilder *implEarwaxBuilder) withOption(key string, value expr.Expr) EarwaxBuilder {
	bCopy := *earwaxBuilder
	bCopy.f = earwaxBuilder.f.With(key, value)
	return &bCopy
}
