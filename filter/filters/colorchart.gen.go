// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColorChartBuilder corresponds to the "colorchart" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colorchart
type ColorChartBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) ColorChartBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) ColorChartBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) ColorChartBuilder
}

// ColorChart creates a new ColorChartBuilder to configure the "colorchart" filter.
func ColorChart(opts ...filter.Option) ColorChartBuilder {
	return &implColorChartBuilder{
		f: filter.New("colorchart", 1, opts...),
	}
}

type implColorChartBuilder struct {
	f filter.Filter
}

func (b *implColorChartBuilder) String() string {
	return b.f.String()
}

func (b *implColorChartBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implColorChartBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implColorChartBuilder) withOption(key string, value expr.Expr) ColorChartBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implColorChartBuilder) Size(width, height int) ColorChartBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implColorChartBuilder) FrameRate(frameRate expr.Expr) ColorChartBuilder {
	return b.withOption("r", frameRate)
}

func (b *implColorChartBuilder) Duration(duration string) ColorChartBuilder {
	return b.withOption("d", expr.String(duration))
}
