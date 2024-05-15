// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DisplaceBuilder Displace pixels.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#displace
type DisplaceBuilder interface {
	filter.Filter
	// Edge set edge mode (from 0 to 3) (default smear).
	Edge(edge int) DisplaceBuilder
	// EdgeExpr set edge mode (from 0 to 3) (default smear).
	EdgeExpr(edge expr.Expr) DisplaceBuilder
}

// Displace creates a new DisplaceBuilder to configure the "displace" filter.
func Displace(opts ...filter.Option) DisplaceBuilder {
	return &implDisplaceBuilder{
		f: filter.New("displace", 1, opts...),
	}
}

type implDisplaceBuilder struct {
	f filter.Filter
}

func (displaceBuilder *implDisplaceBuilder) String() string {
	return displaceBuilder.f.String()
}

func (displaceBuilder *implDisplaceBuilder) Outputs() int {
	return displaceBuilder.f.Outputs()
}

func (displaceBuilder *implDisplaceBuilder) With(key string, value expr.Expr) filter.Filter {
	return displaceBuilder.withOption(key, value)
}

func (displaceBuilder *implDisplaceBuilder) withOption(key string, value expr.Expr) DisplaceBuilder {
	bCopy := *displaceBuilder
	bCopy.f = displaceBuilder.f.With(key, value)
	return &bCopy
}

func (displaceBuilder *implDisplaceBuilder) Edge(edge int) DisplaceBuilder {
	return displaceBuilder.withOption("edge", expr.Int(edge))
}

func (displaceBuilder *implDisplaceBuilder) EdgeExpr(edge expr.Expr) DisplaceBuilder {
	return displaceBuilder.withOption("edge", edge)
}
