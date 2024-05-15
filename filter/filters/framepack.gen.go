// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FramepackBuilder Generate a frame packed stereoscopic video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#framepack
type FramepackBuilder interface {
	filter.Filter
	// Format Frame pack output format (from 0 to INT_MAX) (default sbs).
	Format(format int) FramepackBuilder
}

// Framepack creates a new FramepackBuilder to configure the "framepack" filter.
func Framepack(opts ...filter.Option) FramepackBuilder {
	return &implFramepackBuilder{
		f: filter.New("framepack", 1, opts...),
	}
}

type implFramepackBuilder struct {
	f filter.Filter
}

func (framepackBuilder *implFramepackBuilder) String() string {
	return framepackBuilder.f.String()
}

func (framepackBuilder *implFramepackBuilder) Outputs() int {
	return framepackBuilder.f.Outputs()
}

func (framepackBuilder *implFramepackBuilder) With(key string, value expr.Expr) filter.Filter {
	return framepackBuilder.withOption(key, value)
}

func (framepackBuilder *implFramepackBuilder) withOption(key string, value expr.Expr) FramepackBuilder {
	bCopy := *framepackBuilder
	bCopy.f = framepackBuilder.f.With(key, value)
	return &bCopy
}

func (framepackBuilder *implFramepackBuilder) Format(format int) FramepackBuilder {
	return framepackBuilder.withOption("format", expr.Int(format))
}
