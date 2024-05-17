// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// McdeintBuilder Apply motion compensating deinterlacing.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#mcdeint
type McdeintBuilder interface {
	filter.Filter
	// Mode set mode (from 0 to 3) (default fast).
	Mode(mode int) McdeintBuilder
	// Parity set the assumed picture field parity (from -1 to 1) (default bff).
	Parity(parity int) McdeintBuilder
	// Qp set qp (from INT_MIN to INT_MAX) (default 1).
	Qp(qp int) McdeintBuilder
}

// Mcdeint creates a new McdeintBuilder to configure the "mcdeint" filter.
func Mcdeint(opts ...filter.Option) McdeintBuilder {
	return &implMcdeintBuilder{
		f: filter.New("mcdeint", 1, opts...),
	}
}

type implMcdeintBuilder struct {
	f filter.Filter
}

func (mcdeintBuilder *implMcdeintBuilder) String() string {
	return mcdeintBuilder.f.String()
}

func (mcdeintBuilder *implMcdeintBuilder) Outputs() int {
	return mcdeintBuilder.f.Outputs()
}

func (mcdeintBuilder *implMcdeintBuilder) With(key string, value expr.Expr) filter.Filter {
	return mcdeintBuilder.withOption(key, value)
}

func (mcdeintBuilder *implMcdeintBuilder) withOption(key string, value expr.Expr) McdeintBuilder {
	bCopy := *mcdeintBuilder
	bCopy.f = mcdeintBuilder.f.With(key, value)
	return &bCopy
}

func (mcdeintBuilder *implMcdeintBuilder) Mode(mode int) McdeintBuilder {
	return mcdeintBuilder.withOption("mode", expr.Int(mode))
}

func (mcdeintBuilder *implMcdeintBuilder) Parity(parity int) McdeintBuilder {
	return mcdeintBuilder.withOption("parity", expr.Int(parity))
}

func (mcdeintBuilder *implMcdeintBuilder) Qp(qp int) McdeintBuilder {
	return mcdeintBuilder.withOption("qp", expr.Int(qp))
}