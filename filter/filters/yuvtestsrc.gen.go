// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// YuvtestsrcBuilder Generate YUV test pattern.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#yuvtestsrc
type YuvtestsrcBuilder interface {
	filter.Filter
	// Size set video size (default "320x240").
	Size(size expr.Size) YuvtestsrcBuilder
	// S set video size (default "320x240").
	S(s expr.Size) YuvtestsrcBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) YuvtestsrcBuilder
	// R set video rate (default "25").
	R(r expr.Rational) YuvtestsrcBuilder
	// Duration set video duration (default -0.000001).
	Duration(duration time.Duration) YuvtestsrcBuilder
	// D set video duration (default -0.000001).
	D(d time.Duration) YuvtestsrcBuilder
	// Sar set video sample aspect ratio (from 0 to INT_MAX) (default 1/1).
	Sar(sar expr.Rational) YuvtestsrcBuilder
}

// Yuvtestsrc creates a new YuvtestsrcBuilder to configure the "yuvtestsrc" filter.
func Yuvtestsrc(opts ...filter.Option) YuvtestsrcBuilder {
	return &implYuvtestsrcBuilder{
		f: filter.New("yuvtestsrc", 1, opts...),
	}
}

type implYuvtestsrcBuilder struct {
	f filter.Filter
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) String() string {
	return yuvtestsrcBuilder.f.String()
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) Outputs() int {
	return yuvtestsrcBuilder.f.Outputs()
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) With(key string, value expr.Expr) filter.Filter {
	return yuvtestsrcBuilder.withOption(key, value)
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) withOption(key string, value expr.Expr) YuvtestsrcBuilder {
	bCopy := *yuvtestsrcBuilder
	bCopy.f = yuvtestsrcBuilder.f.With(key, value)
	return &bCopy
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) Size(size expr.Size) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("size", size)
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) S(s expr.Size) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("s", s)
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) Rate(rate expr.Rational) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("rate", rate)
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) R(r expr.Rational) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("r", r)
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) Duration(duration time.Duration) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("duration", expr.Duration(duration))
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) D(d time.Duration) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("d", expr.Duration(d))
}

func (yuvtestsrcBuilder *implYuvtestsrcBuilder) Sar(sar expr.Rational) YuvtestsrcBuilder {
	return yuvtestsrcBuilder.withOption("sar", sar)
}
