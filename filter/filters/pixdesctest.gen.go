// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PixdesctestBuilder Test pixel format definitions.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#pixdesctest
type PixdesctestBuilder interface {
	filter.Filter
}

// Pixdesctest creates a new PixdesctestBuilder to configure the "pixdesctest" filter.
func Pixdesctest(opts ...filter.Option) PixdesctestBuilder {
	return &implPixdesctestBuilder{
		f: filter.New("pixdesctest", 1, opts...),
	}
}

type implPixdesctestBuilder struct {
	f filter.Filter
}

func (pixdesctestBuilder *implPixdesctestBuilder) String() string {
	return pixdesctestBuilder.f.String()
}

func (pixdesctestBuilder *implPixdesctestBuilder) Outputs() int {
	return pixdesctestBuilder.f.Outputs()
}

func (pixdesctestBuilder *implPixdesctestBuilder) With(key string, value expr.Expr) filter.Filter {
	return pixdesctestBuilder.withOption(key, value)
}

func (pixdesctestBuilder *implPixdesctestBuilder) withOption(key string, value expr.Expr) PixdesctestBuilder {
	bCopy := *pixdesctestBuilder
	bCopy.f = pixdesctestBuilder.f.With(key, value)
	return &bCopy
}