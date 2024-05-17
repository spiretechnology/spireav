// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// LatencyBuilder Report video filtering latency.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#latency
type LatencyBuilder interface {
	filter.Filter
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) LatencyBuilder
}

// Latency creates a new LatencyBuilder to configure the "latency" filter.
func Latency(opts ...filter.Option) LatencyBuilder {
	return &implLatencyBuilder{
		f: filter.New("latency", 1, opts...),
	}
}

type implLatencyBuilder struct {
	f filter.Filter
}

func (latencyBuilder *implLatencyBuilder) String() string {
	return latencyBuilder.f.String()
}

func (latencyBuilder *implLatencyBuilder) Outputs() int {
	return latencyBuilder.f.Outputs()
}

func (latencyBuilder *implLatencyBuilder) With(key string, value expr.Expr) filter.Filter {
	return latencyBuilder.withOption(key, value)
}

func (latencyBuilder *implLatencyBuilder) withOption(key string, value expr.Expr) LatencyBuilder {
	bCopy := *latencyBuilder
	bCopy.f = latencyBuilder.f.With(key, value)
	return &bCopy
}

func (latencyBuilder *implLatencyBuilder) Enable(enable expr.Expr) LatencyBuilder {
	return latencyBuilder.withOption("enable", enable)
}