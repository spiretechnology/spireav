// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SrBuilder Apply DNN-based image super resolution to the input.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#sr
type SrBuilder interface {
	filter.Filter
	// DnnBackend DNN backend used for model execution (from 0 to 1) (default 1).
	DnnBackend(dnnBackend int) SrBuilder
	// ScaleFactor scale factor for SRCNN model (from 2 to 4) (default 2).
	ScaleFactor(scaleFactor int) SrBuilder
	// Model path to model file specifying network architecture and its parameters.
	Model(model string) SrBuilder
	// Input input name of the model (default "x").
	Input(input string) SrBuilder
	// Output output name of the model (default "y").
	Output(output string) SrBuilder
}

// Sr creates a new SrBuilder to configure the "sr" filter.
func Sr(opts ...filter.Option) SrBuilder {
	return &implSrBuilder{
		f: filter.New("sr", 1, opts...),
	}
}

type implSrBuilder struct {
	f filter.Filter
}

func (srBuilder *implSrBuilder) String() string {
	return srBuilder.f.String()
}

func (srBuilder *implSrBuilder) Outputs() int {
	return srBuilder.f.Outputs()
}

func (srBuilder *implSrBuilder) With(key string, value expr.Expr) filter.Filter {
	return srBuilder.withOption(key, value)
}

func (srBuilder *implSrBuilder) withOption(key string, value expr.Expr) SrBuilder {
	bCopy := *srBuilder
	bCopy.f = srBuilder.f.With(key, value)
	return &bCopy
}

func (srBuilder *implSrBuilder) DnnBackend(dnnBackend int) SrBuilder {
	return srBuilder.withOption("dnn_backend", expr.Int(dnnBackend))
}

func (srBuilder *implSrBuilder) ScaleFactor(scaleFactor int) SrBuilder {
	return srBuilder.withOption("scale_factor", expr.Int(scaleFactor))
}

func (srBuilder *implSrBuilder) Model(model string) SrBuilder {
	return srBuilder.withOption("model", expr.String(model))
}

func (srBuilder *implSrBuilder) Input(input string) SrBuilder {
	return srBuilder.withOption("input", expr.String(input))
}

func (srBuilder *implSrBuilder) Output(output string) SrBuilder {
	return srBuilder.withOption("output", expr.String(output))
}
