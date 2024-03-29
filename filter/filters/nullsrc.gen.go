// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// NullSourceBuilder corresponds to the "nullsrc" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#nullsrc
type NullSourceBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) NullSourceBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) NullSourceBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) NullSourceBuilder
}

// NullSource creates a new NullSourceBuilder to configure the "nullsrc" filter.
func NullSource(opts ...filter.Option) NullSourceBuilder {
	return &implNullSourceBuilder{
		f: filter.New("nullsrc", 1, opts...),
	}
}

type implNullSourceBuilder struct {
	f filter.Filter
}

func (b *implNullSourceBuilder) String() string {
	return b.f.String()
}

func (b *implNullSourceBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implNullSourceBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implNullSourceBuilder) withOption(key string, value expr.Expr) NullSourceBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implNullSourceBuilder) Size(width, height int) NullSourceBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implNullSourceBuilder) FrameRate(frameRate expr.Expr) NullSourceBuilder {
	return b.withOption("r", frameRate)
}

func (b *implNullSourceBuilder) Duration(duration string) NullSourceBuilder {
	return b.withOption("d", expr.String(duration))
}
