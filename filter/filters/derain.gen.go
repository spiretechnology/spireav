// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DerainBuilder Apply derain filter to the input.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#derain
type DerainBuilder interface {
	filter.Filter
	// FilterType filter type(derain/dehaze) (from 0 to 1) (default derain).
	FilterType(filterType int) DerainBuilder
	// DnnBackend DNN backend (from 0 to 1) (default 1).
	DnnBackend(dnnBackend int) DerainBuilder
	// Model path to model file.
	Model(model string) DerainBuilder
	// Input input name of the model (default "x").
	Input(input string) DerainBuilder
	// Output output name of the model (default "y").
	Output(output string) DerainBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) DerainBuilder
}

// Derain creates a new DerainBuilder to configure the "derain" filter.
func Derain(opts ...filter.Option) DerainBuilder {
	return &implDerainBuilder{
		f: filter.New("derain", 1, opts...),
	}
}

type implDerainBuilder struct {
	f filter.Filter
}

func (derainBuilder *implDerainBuilder) String() string {
	return derainBuilder.f.String()
}

func (derainBuilder *implDerainBuilder) Outputs() int {
	return derainBuilder.f.Outputs()
}

func (derainBuilder *implDerainBuilder) With(key string, value expr.Expr) filter.Filter {
	return derainBuilder.withOption(key, value)
}

func (derainBuilder *implDerainBuilder) withOption(key string, value expr.Expr) DerainBuilder {
	bCopy := *derainBuilder
	bCopy.f = derainBuilder.f.With(key, value)
	return &bCopy
}

func (derainBuilder *implDerainBuilder) FilterType(filterType int) DerainBuilder {
	return derainBuilder.withOption("filter_type", expr.Int(filterType))
}

func (derainBuilder *implDerainBuilder) DnnBackend(dnnBackend int) DerainBuilder {
	return derainBuilder.withOption("dnn_backend", expr.Int(dnnBackend))
}

func (derainBuilder *implDerainBuilder) Model(model string) DerainBuilder {
	return derainBuilder.withOption("model", expr.String(model))
}

func (derainBuilder *implDerainBuilder) Input(input string) DerainBuilder {
	return derainBuilder.withOption("input", expr.String(input))
}

func (derainBuilder *implDerainBuilder) Output(output string) DerainBuilder {
	return derainBuilder.withOption("output", expr.String(output))
}

func (derainBuilder *implDerainBuilder) Enable(enable expr.Expr) DerainBuilder {
	return derainBuilder.withOption("enable", enable)
}
