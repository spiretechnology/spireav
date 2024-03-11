// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// RGBTestBuilder corresponds to the "rgbtestsrc" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#rgbtestsrc
type RGBTestBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) RGBTestBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) RGBTestBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) RGBTestBuilder
}

// RGBTest creates a new RGBTestBuilder to configure the "rgbtestsrc" filter.
func RGBTest(opts ...filter.Option) RGBTestBuilder {
	return &implRGBTestBuilder{
		f: filter.New("rgbtestsrc", 1, opts...),
	}
}

type implRGBTestBuilder struct {
	f filter.Filter
}

func (b *implRGBTestBuilder) String() string {
	return b.f.String()
}

func (b *implRGBTestBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implRGBTestBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implRGBTestBuilder) withOption(key string, value expr.Expr) RGBTestBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implRGBTestBuilder) Size(width, height int) RGBTestBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implRGBTestBuilder) FrameRate(frameRate expr.Expr) RGBTestBuilder {
	return b.withOption("r", frameRate)
}

func (b *implRGBTestBuilder) Duration(duration string) RGBTestBuilder {
	return b.withOption("d", expr.String(duration))
}