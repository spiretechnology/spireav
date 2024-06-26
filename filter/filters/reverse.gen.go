// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ReverseBuilder Reverse a clip.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#reverse
type ReverseBuilder interface {
	filter.Filter
}

// Reverse creates a new ReverseBuilder to configure the "reverse" filter.
func Reverse(opts ...filter.Option) ReverseBuilder {
	return &implReverseBuilder{
		f: filter.New("reverse", 1, opts...),
	}
}

type implReverseBuilder struct {
	f filter.Filter
}

func (reverseBuilder *implReverseBuilder) String() string {
	return reverseBuilder.f.String()
}

func (reverseBuilder *implReverseBuilder) Outputs() int {
	return reverseBuilder.f.Outputs()
}

func (reverseBuilder *implReverseBuilder) With(key string, value expr.Expr) filter.Filter {
	return reverseBuilder.withOption(key, value)
}

func (reverseBuilder *implReverseBuilder) withOption(key string, value expr.Expr) ReverseBuilder {
	bCopy := *reverseBuilder
	bCopy.f = reverseBuilder.f.With(key, value)
	return &bCopy
}
