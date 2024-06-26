// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CopyBuilder Copy the input video unchanged to the output.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#copy
type CopyBuilder interface {
	filter.Filter
}

// Copy creates a new CopyBuilder to configure the "copy" filter.
func Copy(opts ...filter.Option) CopyBuilder {
	return &implCopyBuilder{
		f: filter.New("copy", 1, opts...),
	}
}

type implCopyBuilder struct {
	f filter.Filter
}

func (copyBuilder *implCopyBuilder) String() string {
	return copyBuilder.f.String()
}

func (copyBuilder *implCopyBuilder) Outputs() int {
	return copyBuilder.f.Outputs()
}

func (copyBuilder *implCopyBuilder) With(key string, value expr.Expr) filter.Filter {
	return copyBuilder.withOption(key, value)
}

func (copyBuilder *implCopyBuilder) withOption(key string, value expr.Expr) CopyBuilder {
	bCopy := *copyBuilder
	bCopy.f = copyBuilder.f.With(key, value)
	return &bCopy
}
