// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// InterleaveBuilder Temporally interleave video inputs.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#interleave
type InterleaveBuilder interface {
	filter.Filter
	// NbInputs set number of inputs (from 1 to INT_MAX) (default 2).
	NbInputs(nbInputs int) InterleaveBuilder
	// N set number of inputs (from 1 to INT_MAX) (default 2).
	N(n int) InterleaveBuilder
	// Duration how to determine the end-of-stream (from 0 to 2) (default longest).
	Duration(duration int) InterleaveBuilder
}

// Interleave creates a new InterleaveBuilder to configure the "interleave" filter.
func Interleave(opts ...filter.Option) InterleaveBuilder {
	return &implInterleaveBuilder{
		f: filter.New("interleave", 1, opts...),
	}
}

type implInterleaveBuilder struct {
	f filter.Filter
}

func (interleaveBuilder *implInterleaveBuilder) String() string {
	return interleaveBuilder.f.String()
}

func (interleaveBuilder *implInterleaveBuilder) Outputs() int {
	return interleaveBuilder.f.Outputs()
}

func (interleaveBuilder *implInterleaveBuilder) With(key string, value expr.Expr) filter.Filter {
	return interleaveBuilder.withOption(key, value)
}

func (interleaveBuilder *implInterleaveBuilder) withOption(key string, value expr.Expr) InterleaveBuilder {
	bCopy := *interleaveBuilder
	bCopy.f = interleaveBuilder.f.With(key, value)
	return &bCopy
}

func (interleaveBuilder *implInterleaveBuilder) NbInputs(nbInputs int) InterleaveBuilder {
	return interleaveBuilder.withOption("nb_inputs", expr.Int(nbInputs))
}

func (interleaveBuilder *implInterleaveBuilder) N(n int) InterleaveBuilder {
	return interleaveBuilder.withOption("n", expr.Int(n))
}

func (interleaveBuilder *implInterleaveBuilder) Duration(duration int) InterleaveBuilder {
	return interleaveBuilder.withOption("duration", expr.Int(duration))
}
