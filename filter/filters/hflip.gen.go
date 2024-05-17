// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// HflipBuilder Horizontally flip the input video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#hflip
type HflipBuilder interface {
	filter.Filter
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) HflipBuilder
}

// Hflip creates a new HflipBuilder to configure the "hflip" filter.
func Hflip(opts ...filter.Option) HflipBuilder {
	return &implHflipBuilder{
		f: filter.New("hflip", 1, opts...),
	}
}

type implHflipBuilder struct {
	f filter.Filter
}

func (hflipBuilder *implHflipBuilder) String() string {
	return hflipBuilder.f.String()
}

func (hflipBuilder *implHflipBuilder) Outputs() int {
	return hflipBuilder.f.Outputs()
}

func (hflipBuilder *implHflipBuilder) With(key string, value expr.Expr) filter.Filter {
	return hflipBuilder.withOption(key, value)
}

func (hflipBuilder *implHflipBuilder) withOption(key string, value expr.Expr) HflipBuilder {
	bCopy := *hflipBuilder
	bCopy.f = hflipBuilder.f.With(key, value)
	return &bCopy
}

func (hflipBuilder *implHflipBuilder) Enable(enable expr.Expr) HflipBuilder {
	return hflipBuilder.withOption("enable", enable)
}
