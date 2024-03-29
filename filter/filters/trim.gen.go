// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TrimBuilder corresponds to the "trim" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#trim
type TrimBuilder interface {
	filter.Filter
	// StartSeconds sets the "start" option on the filter.
	StartSeconds(startSeconds expr.Expr) TrimBuilder
	// StartSecondsString sets the "start" option on the filter.
	StartSecondsString(startSeconds string) TrimBuilder
	// EndSeconds sets the "end" option on the filter.
	EndSeconds(endSeconds expr.Expr) TrimBuilder
	// EndSecondsString sets the "end" option on the filter.
	EndSecondsString(endSeconds string) TrimBuilder
	// StartFrame sets the "start_frame" option on the filter.
	StartFrame(startFrame int) TrimBuilder
	// EndFrame sets the "end_frame" option on the filter.
	EndFrame(endFrame int) TrimBuilder
}

// Trim creates a new TrimBuilder to configure the "trim" filter.
func Trim(opts ...filter.Option) TrimBuilder {
	return &implTrimBuilder{
		f: filter.New("trim", 1, opts...),
	}
}

type implTrimBuilder struct {
	f filter.Filter
}

func (b *implTrimBuilder) String() string {
	return b.f.String()
}

func (b *implTrimBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implTrimBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implTrimBuilder) withOption(key string, value expr.Expr) TrimBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implTrimBuilder) StartSeconds(startSeconds expr.Expr) TrimBuilder {
	return b.withOption("start", startSeconds)
}

func (b *implTrimBuilder) StartSecondsString(startSeconds string) TrimBuilder {
	return b.withOption("start", expr.String(startSeconds))
}

func (b *implTrimBuilder) EndSeconds(endSeconds expr.Expr) TrimBuilder {
	return b.withOption("end", endSeconds)
}

func (b *implTrimBuilder) EndSecondsString(endSeconds string) TrimBuilder {
	return b.withOption("end", expr.String(endSeconds))
}

func (b *implTrimBuilder) StartFrame(startFrame int) TrimBuilder {
	return b.withOption("start_frame", expr.Int(startFrame))
}

func (b *implTrimBuilder) EndFrame(endFrame int) TrimBuilder {
	return b.withOption("end_frame", expr.Int(endFrame))
}
