// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TestSourceBuilder corresponds to the "testsrc" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#testsrc
type TestSourceBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) TestSourceBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) TestSourceBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) TestSourceBuilder
}

// TestSource creates a new TestSourceBuilder to configure the "testsrc" filter.
func TestSource(opts ...filter.Option) TestSourceBuilder {
	return &implTestSourceBuilder{
		f: filter.New("testsrc", 1, opts...),
	}
}

type implTestSourceBuilder struct {
	f filter.Filter
}

func (b *implTestSourceBuilder) String() string {
	return b.f.String()
}

func (b *implTestSourceBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implTestSourceBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implTestSourceBuilder) withOption(key string, value expr.Expr) TestSourceBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implTestSourceBuilder) Size(width, height int) TestSourceBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implTestSourceBuilder) FrameRate(frameRate expr.Expr) TestSourceBuilder {
	return b.withOption("r", frameRate)
}

func (b *implTestSourceBuilder) Duration(duration string) TestSourceBuilder {
	return b.withOption("d", expr.String(duration))
}
