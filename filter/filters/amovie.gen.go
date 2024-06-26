// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AmovieBuilder Read audio from a movie source.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#amovie
type AmovieBuilder interface {
	filter.Filter
	// FormatName set format name.
	FormatName(formatName string) AmovieBuilder
	// F set format name.
	F(f string) AmovieBuilder
	// StreamIndex set stream index (from -1 to INT_MAX) (default -1).
	StreamIndex(streamIndex int) AmovieBuilder
	// Si set stream index (from -1 to INT_MAX) (default -1).
	Si(si int) AmovieBuilder
	// SeekPoint set seekpoint (seconds) (from 0 to 9.22337e+12) (default 0).
	SeekPoint(seekPoint float64) AmovieBuilder
	// Sp set seekpoint (seconds) (from 0 to 9.22337e+12) (default 0).
	Sp(sp float64) AmovieBuilder
	// Streams set streams.
	Streams(streams string) AmovieBuilder
	// S set streams.
	S(s string) AmovieBuilder
	// Loop set loop count (from 0 to INT_MAX) (default 1).
	Loop(loop int) AmovieBuilder
	// Discontinuity set discontinuity threshold (default 0).
	Discontinuity(discontinuity time.Duration) AmovieBuilder
	// DecThreads set the number of threads for decoding (from 0 to INT_MAX) (default 0).
	DecThreads(decThreads int) AmovieBuilder
	// FormatOpts set format options for the opened file.
	FormatOpts(formatOpts expr.Dictionary) AmovieBuilder
}

// Amovie creates a new AmovieBuilder to configure the "amovie" filter.
func Amovie(outputs int, opts ...filter.Option) AmovieBuilder {
	f := filter.New("amovie", outputs, opts...)
	f = f.With("outputs", expr.Int(outputs))
	return &implAmovieBuilder{f: f}
}

type implAmovieBuilder struct {
	f filter.Filter
}

func (amovieBuilder *implAmovieBuilder) String() string {
	return amovieBuilder.f.String()
}

func (amovieBuilder *implAmovieBuilder) Outputs() int {
	return amovieBuilder.f.Outputs()
}

func (amovieBuilder *implAmovieBuilder) With(key string, value expr.Expr) filter.Filter {
	return amovieBuilder.withOption(key, value)
}

func (amovieBuilder *implAmovieBuilder) withOption(key string, value expr.Expr) AmovieBuilder {
	bCopy := *amovieBuilder
	bCopy.f = amovieBuilder.f.With(key, value)
	return &bCopy
}

func (amovieBuilder *implAmovieBuilder) FormatName(formatName string) AmovieBuilder {
	return amovieBuilder.withOption("format_name", expr.String(formatName))
}

func (amovieBuilder *implAmovieBuilder) F(f string) AmovieBuilder {
	return amovieBuilder.withOption("f", expr.String(f))
}

func (amovieBuilder *implAmovieBuilder) StreamIndex(streamIndex int) AmovieBuilder {
	return amovieBuilder.withOption("stream_index", expr.Int(streamIndex))
}

func (amovieBuilder *implAmovieBuilder) Si(si int) AmovieBuilder {
	return amovieBuilder.withOption("si", expr.Int(si))
}

func (amovieBuilder *implAmovieBuilder) SeekPoint(seekPoint float64) AmovieBuilder {
	return amovieBuilder.withOption("seek_point", expr.Double(seekPoint))
}

func (amovieBuilder *implAmovieBuilder) Sp(sp float64) AmovieBuilder {
	return amovieBuilder.withOption("sp", expr.Double(sp))
}

func (amovieBuilder *implAmovieBuilder) Streams(streams string) AmovieBuilder {
	return amovieBuilder.withOption("streams", expr.String(streams))
}

func (amovieBuilder *implAmovieBuilder) S(s string) AmovieBuilder {
	return amovieBuilder.withOption("s", expr.String(s))
}

func (amovieBuilder *implAmovieBuilder) Loop(loop int) AmovieBuilder {
	return amovieBuilder.withOption("loop", expr.Int(loop))
}

func (amovieBuilder *implAmovieBuilder) Discontinuity(discontinuity time.Duration) AmovieBuilder {
	return amovieBuilder.withOption("discontinuity", expr.Duration(discontinuity))
}

func (amovieBuilder *implAmovieBuilder) DecThreads(decThreads int) AmovieBuilder {
	return amovieBuilder.withOption("dec_threads", expr.Int(decThreads))
}

func (amovieBuilder *implAmovieBuilder) FormatOpts(formatOpts expr.Dictionary) AmovieBuilder {
	return amovieBuilder.withOption("format_opts", formatOpts)
}
