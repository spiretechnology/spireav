// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FormatBuilder Convert the input video to one of the specified pixel formats.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#format
type FormatBuilder interface {
	filter.Filter
	// PixFmts A '|'-separated list of pixel formats.
	PixFmts(pixFmts string) FormatBuilder
	// ColorSpaces A '|'-separated list of color spaces.
	ColorSpaces(colorSpaces string) FormatBuilder
	// ColorRanges A '|'-separated list of color ranges.
	ColorRanges(colorRanges string) FormatBuilder
}

// Format creates a new FormatBuilder to configure the "format" filter.
func Format(opts ...filter.Option) FormatBuilder {
	return &implFormatBuilder{
		f: filter.New("format", 1, opts...),
	}
}

type implFormatBuilder struct {
	f filter.Filter
}

func (formatBuilder *implFormatBuilder) String() string {
	return formatBuilder.f.String()
}

func (formatBuilder *implFormatBuilder) Outputs() int {
	return formatBuilder.f.Outputs()
}

func (formatBuilder *implFormatBuilder) With(key string, value expr.Expr) filter.Filter {
	return formatBuilder.withOption(key, value)
}

func (formatBuilder *implFormatBuilder) withOption(key string, value expr.Expr) FormatBuilder {
	bCopy := *formatBuilder
	bCopy.f = formatBuilder.f.With(key, value)
	return &bCopy
}

func (formatBuilder *implFormatBuilder) PixFmts(pixFmts string) FormatBuilder {
	return formatBuilder.withOption("pix_fmts", expr.String(pixFmts))
}

func (formatBuilder *implFormatBuilder) ColorSpaces(colorSpaces string) FormatBuilder {
	return formatBuilder.withOption("color_spaces", expr.String(colorSpaces))
}

func (formatBuilder *implFormatBuilder) ColorRanges(colorRanges string) FormatBuilder {
	return formatBuilder.withOption("color_ranges", expr.String(colorRanges))
}