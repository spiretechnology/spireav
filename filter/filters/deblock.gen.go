// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DeblockBuilder Deblock video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#deblock
type DeblockBuilder interface {
	filter.Filter
	// Filter set type of filter (from 0 to 1) (default strong).
	Filter(filter int) DeblockBuilder
	// FilterExpr set type of filter (from 0 to 1) (default strong).
	FilterExpr(filter expr.Expr) DeblockBuilder
	// Block set size of block (from 4 to 512) (default 8).
	Block(block int) DeblockBuilder
	// BlockExpr set size of block (from 4 to 512) (default 8).
	BlockExpr(block expr.Expr) DeblockBuilder
	// Alpha set 1st detection threshold (from 0 to 1) (default 0.098).
	Alpha(alpha float32) DeblockBuilder
	// AlphaExpr set 1st detection threshold (from 0 to 1) (default 0.098).
	AlphaExpr(alpha expr.Expr) DeblockBuilder
	// Beta set 2nd detection threshold (from 0 to 1) (default 0.05).
	Beta(beta float32) DeblockBuilder
	// BetaExpr set 2nd detection threshold (from 0 to 1) (default 0.05).
	BetaExpr(beta expr.Expr) DeblockBuilder
	// Gamma set 3rd detection threshold (from 0 to 1) (default 0.05).
	Gamma(gamma float32) DeblockBuilder
	// GammaExpr set 3rd detection threshold (from 0 to 1) (default 0.05).
	GammaExpr(gamma expr.Expr) DeblockBuilder
	// Delta set 4th detection threshold (from 0 to 1) (default 0.05).
	Delta(delta float32) DeblockBuilder
	// DeltaExpr set 4th detection threshold (from 0 to 1) (default 0.05).
	DeltaExpr(delta expr.Expr) DeblockBuilder
	// Planes set planes to filter (from 0 to 15) (default 15).
	Planes(planes int) DeblockBuilder
	// PlanesExpr set planes to filter (from 0 to 15) (default 15).
	PlanesExpr(planes expr.Expr) DeblockBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) DeblockBuilder
}

// Deblock creates a new DeblockBuilder to configure the "deblock" filter.
func Deblock(opts ...filter.Option) DeblockBuilder {
	return &implDeblockBuilder{
		f: filter.New("deblock", 1, opts...),
	}
}

type implDeblockBuilder struct {
	f filter.Filter
}

func (deblockBuilder *implDeblockBuilder) String() string {
	return deblockBuilder.f.String()
}

func (deblockBuilder *implDeblockBuilder) Outputs() int {
	return deblockBuilder.f.Outputs()
}

func (deblockBuilder *implDeblockBuilder) With(key string, value expr.Expr) filter.Filter {
	return deblockBuilder.withOption(key, value)
}

func (deblockBuilder *implDeblockBuilder) withOption(key string, value expr.Expr) DeblockBuilder {
	bCopy := *deblockBuilder
	bCopy.f = deblockBuilder.f.With(key, value)
	return &bCopy
}

func (deblockBuilder *implDeblockBuilder) Filter(filter int) DeblockBuilder {
	return deblockBuilder.withOption("filter", expr.Int(filter))
}

func (deblockBuilder *implDeblockBuilder) FilterExpr(filter expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("filter", filter)
}

func (deblockBuilder *implDeblockBuilder) Block(block int) DeblockBuilder {
	return deblockBuilder.withOption("block", expr.Int(block))
}

func (deblockBuilder *implDeblockBuilder) BlockExpr(block expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("block", block)
}

func (deblockBuilder *implDeblockBuilder) Alpha(alpha float32) DeblockBuilder {
	return deblockBuilder.withOption("alpha", expr.Float(alpha))
}

func (deblockBuilder *implDeblockBuilder) AlphaExpr(alpha expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("alpha", alpha)
}

func (deblockBuilder *implDeblockBuilder) Beta(beta float32) DeblockBuilder {
	return deblockBuilder.withOption("beta", expr.Float(beta))
}

func (deblockBuilder *implDeblockBuilder) BetaExpr(beta expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("beta", beta)
}

func (deblockBuilder *implDeblockBuilder) Gamma(gamma float32) DeblockBuilder {
	return deblockBuilder.withOption("gamma", expr.Float(gamma))
}

func (deblockBuilder *implDeblockBuilder) GammaExpr(gamma expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("gamma", gamma)
}

func (deblockBuilder *implDeblockBuilder) Delta(delta float32) DeblockBuilder {
	return deblockBuilder.withOption("delta", expr.Float(delta))
}

func (deblockBuilder *implDeblockBuilder) DeltaExpr(delta expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("delta", delta)
}

func (deblockBuilder *implDeblockBuilder) Planes(planes int) DeblockBuilder {
	return deblockBuilder.withOption("planes", expr.Int(planes))
}

func (deblockBuilder *implDeblockBuilder) PlanesExpr(planes expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("planes", planes)
}

func (deblockBuilder *implDeblockBuilder) Enable(enable expr.Expr) DeblockBuilder {
	return deblockBuilder.withOption("enable", enable)
}