// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColorSpectrumBuilder corresponds to the "colorspectrum" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colorspectrum
type ColorSpectrumBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) ColorSpectrumBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) ColorSpectrumBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) ColorSpectrumBuilder
}

// ColorSpectrum creates a new ColorSpectrumBuilder to configure the "colorspectrum" filter.
func ColorSpectrum(opts ...filter.Option) ColorSpectrumBuilder {
	return &implColorSpectrumBuilder{
		f: filter.New("colorspectrum", 1, opts...),
	}
}

type implColorSpectrumBuilder struct {
	f filter.Filter
}

func (b *implColorSpectrumBuilder) String() string {
	return b.f.String()
}

func (b *implColorSpectrumBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implColorSpectrumBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implColorSpectrumBuilder) withOption(key string, value expr.Expr) ColorSpectrumBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implColorSpectrumBuilder) Size(width, height int) ColorSpectrumBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implColorSpectrumBuilder) FrameRate(frameRate expr.Expr) ColorSpectrumBuilder {
	return b.withOption("r", frameRate)
}

func (b *implColorSpectrumBuilder) Duration(duration string) ColorSpectrumBuilder {
	return b.withOption("d", expr.String(duration))
}