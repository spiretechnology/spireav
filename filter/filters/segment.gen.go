// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SegmentBuilder Segment video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#segment
type SegmentBuilder interface {
	filter.Filter
	// Timestamps timestamps of input at which to split input.
	Timestamps(timestamps string) SegmentBuilder
	// Frames frames at which to split input.
	Frames(frames string) SegmentBuilder
}

// Segment creates a new SegmentBuilder to configure the "segment" filter.
func Segment(outputs int, opts ...filter.Option) SegmentBuilder {
	f := filter.New("segment", outputs, opts...)
	f = f.With("outputs", expr.Int(outputs))
	return &implSegmentBuilder{f: f}
}

type implSegmentBuilder struct {
	f filter.Filter
}

func (segmentBuilder *implSegmentBuilder) String() string {
	return segmentBuilder.f.String()
}

func (segmentBuilder *implSegmentBuilder) Outputs() int {
	return segmentBuilder.f.Outputs()
}

func (segmentBuilder *implSegmentBuilder) With(key string, value expr.Expr) filter.Filter {
	return segmentBuilder.withOption(key, value)
}

func (segmentBuilder *implSegmentBuilder) withOption(key string, value expr.Expr) SegmentBuilder {
	bCopy := *segmentBuilder
	bCopy.f = segmentBuilder.f.With(key, value)
	return &bCopy
}

func (segmentBuilder *implSegmentBuilder) Timestamps(timestamps string) SegmentBuilder {
	return segmentBuilder.withOption("timestamps", expr.String(timestamps))
}

func (segmentBuilder *implSegmentBuilder) Frames(frames string) SegmentBuilder {
	return segmentBuilder.withOption("frames", expr.String(frames))
}