// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// NullBuilder Pass the source unchanged to the output.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#null
type NullBuilder interface {
	filter.Filter
}

// Null creates a new NullBuilder to configure the "null" filter.
func Null(opts ...filter.Option) NullBuilder {
	return &implNullBuilder{
		f: filter.New("null", 1, opts...),
	}
}

type implNullBuilder struct {
	f filter.Filter
}

func (nullBuilder *implNullBuilder) String() string {
	return nullBuilder.f.String()
}

func (nullBuilder *implNullBuilder) Outputs() int {
	return nullBuilder.f.Outputs()
}

func (nullBuilder *implNullBuilder) With(key string, value expr.Expr) filter.Filter {
	return nullBuilder.withOption(key, value)
}

func (nullBuilder *implNullBuilder) withOption(key string, value expr.Expr) NullBuilder {
	bCopy := *nullBuilder
	bCopy.f = nullBuilder.f.With(key, value)
	return &bCopy
}
