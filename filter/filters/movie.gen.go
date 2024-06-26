// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// MovieBuilder Read from a movie source.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#movie
type MovieBuilder interface {
	filter.Filter
	// FormatName set format name.
	FormatName(formatName string) MovieBuilder
	// F set format name.
	F(f string) MovieBuilder
	// StreamIndex set stream index (from -1 to INT_MAX) (default -1).
	StreamIndex(streamIndex int) MovieBuilder
	// Si set stream index (from -1 to INT_MAX) (default -1).
	Si(si int) MovieBuilder
	// SeekPoint set seekpoint (seconds) (from 0 to 9.22337e+12) (default 0).
	SeekPoint(seekPoint float64) MovieBuilder
	// Sp set seekpoint (seconds) (from 0 to 9.22337e+12) (default 0).
	Sp(sp float64) MovieBuilder
	// Streams set streams.
	Streams(streams string) MovieBuilder
	// S set streams.
	S(s string) MovieBuilder
	// Loop set loop count (from 0 to INT_MAX) (default 1).
	Loop(loop int) MovieBuilder
	// Discontinuity set discontinuity threshold (default 0).
	Discontinuity(discontinuity time.Duration) MovieBuilder
	// DecThreads set the number of threads for decoding (from 0 to INT_MAX) (default 0).
	DecThreads(decThreads int) MovieBuilder
	// FormatOpts set format options for the opened file.
	FormatOpts(formatOpts expr.Dictionary) MovieBuilder
}

// Movie creates a new MovieBuilder to configure the "movie" filter.
func Movie(outputs int, opts ...filter.Option) MovieBuilder {
	f := filter.New("movie", outputs, opts...)
	f = f.With("outputs", expr.Int(outputs))
	return &implMovieBuilder{f: f}
}

type implMovieBuilder struct {
	f filter.Filter
}

func (movieBuilder *implMovieBuilder) String() string {
	return movieBuilder.f.String()
}

func (movieBuilder *implMovieBuilder) Outputs() int {
	return movieBuilder.f.Outputs()
}

func (movieBuilder *implMovieBuilder) With(key string, value expr.Expr) filter.Filter {
	return movieBuilder.withOption(key, value)
}

func (movieBuilder *implMovieBuilder) withOption(key string, value expr.Expr) MovieBuilder {
	bCopy := *movieBuilder
	bCopy.f = movieBuilder.f.With(key, value)
	return &bCopy
}

func (movieBuilder *implMovieBuilder) FormatName(formatName string) MovieBuilder {
	return movieBuilder.withOption("format_name", expr.String(formatName))
}

func (movieBuilder *implMovieBuilder) F(f string) MovieBuilder {
	return movieBuilder.withOption("f", expr.String(f))
}

func (movieBuilder *implMovieBuilder) StreamIndex(streamIndex int) MovieBuilder {
	return movieBuilder.withOption("stream_index", expr.Int(streamIndex))
}

func (movieBuilder *implMovieBuilder) Si(si int) MovieBuilder {
	return movieBuilder.withOption("si", expr.Int(si))
}

func (movieBuilder *implMovieBuilder) SeekPoint(seekPoint float64) MovieBuilder {
	return movieBuilder.withOption("seek_point", expr.Double(seekPoint))
}

func (movieBuilder *implMovieBuilder) Sp(sp float64) MovieBuilder {
	return movieBuilder.withOption("sp", expr.Double(sp))
}

func (movieBuilder *implMovieBuilder) Streams(streams string) MovieBuilder {
	return movieBuilder.withOption("streams", expr.String(streams))
}

func (movieBuilder *implMovieBuilder) S(s string) MovieBuilder {
	return movieBuilder.withOption("s", expr.String(s))
}

func (movieBuilder *implMovieBuilder) Loop(loop int) MovieBuilder {
	return movieBuilder.withOption("loop", expr.Int(loop))
}

func (movieBuilder *implMovieBuilder) Discontinuity(discontinuity time.Duration) MovieBuilder {
	return movieBuilder.withOption("discontinuity", expr.Duration(discontinuity))
}

func (movieBuilder *implMovieBuilder) DecThreads(decThreads int) MovieBuilder {
	return movieBuilder.withOption("dec_threads", expr.Int(decThreads))
}

func (movieBuilder *implMovieBuilder) FormatOpts(formatOpts expr.Dictionary) MovieBuilder {
	return movieBuilder.withOption("format_opts", formatOpts)
}
