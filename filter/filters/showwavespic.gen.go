// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ShowwavespicBuilder Convert input audio to a video output single picture.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#showwavespic
type ShowwavespicBuilder interface {
	filter.Filter
	// Size set video size (default "600x240").
	Size(size expr.Size) ShowwavespicBuilder
	// S set video size (default "600x240").
	S(s expr.Size) ShowwavespicBuilder
	// SplitChannels draw channels separately (default false).
	SplitChannels(splitChannels bool) ShowwavespicBuilder
	// Colors set channels colors (default "red|green|blue|yellow|orange|lime|pink|magenta|brown").
	Colors(colors string) ShowwavespicBuilder
	// Scale set amplitude scale (from 0 to 3) (default lin).
	Scale(scale int) ShowwavespicBuilder
	// Draw set draw mode (from 0 to 1) (default scale).
	Draw(draw int) ShowwavespicBuilder
	// Filter set filter mode (from 0 to 1) (default average).
	Filter(filter int) ShowwavespicBuilder
}

// Showwavespic creates a new ShowwavespicBuilder to configure the "showwavespic" filter.
func Showwavespic(opts ...filter.Option) ShowwavespicBuilder {
	return &implShowwavespicBuilder{
		f: filter.New("showwavespic", 1, opts...),
	}
}

type implShowwavespicBuilder struct {
	f filter.Filter
}

func (showwavespicBuilder *implShowwavespicBuilder) String() string {
	return showwavespicBuilder.f.String()
}

func (showwavespicBuilder *implShowwavespicBuilder) Outputs() int {
	return showwavespicBuilder.f.Outputs()
}

func (showwavespicBuilder *implShowwavespicBuilder) With(key string, value expr.Expr) filter.Filter {
	return showwavespicBuilder.withOption(key, value)
}

func (showwavespicBuilder *implShowwavespicBuilder) withOption(key string, value expr.Expr) ShowwavespicBuilder {
	bCopy := *showwavespicBuilder
	bCopy.f = showwavespicBuilder.f.With(key, value)
	return &bCopy
}

func (showwavespicBuilder *implShowwavespicBuilder) Size(size expr.Size) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("size", size)
}

func (showwavespicBuilder *implShowwavespicBuilder) S(s expr.Size) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("s", s)
}

func (showwavespicBuilder *implShowwavespicBuilder) SplitChannels(splitChannels bool) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("split_channels", expr.Bool(splitChannels))
}

func (showwavespicBuilder *implShowwavespicBuilder) Colors(colors string) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("colors", expr.String(colors))
}

func (showwavespicBuilder *implShowwavespicBuilder) Scale(scale int) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("scale", expr.Int(scale))
}

func (showwavespicBuilder *implShowwavespicBuilder) Draw(draw int) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("draw", expr.Int(draw))
}

func (showwavespicBuilder *implShowwavespicBuilder) Filter(filter int) ShowwavespicBuilder {
	return showwavespicBuilder.withOption("filter", expr.Int(filter))
}