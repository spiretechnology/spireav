// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// BilateralBuilder Apply Bilateral filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#bilateral
type BilateralBuilder interface {
	filter.Filter
	// SigmaS set spatial sigma (from 0 to 512) (default 0.1).
	SigmaS(sigmaS float32) BilateralBuilder
	// SigmaSExpr set spatial sigma (from 0 to 512) (default 0.1).
	SigmaSExpr(sigmaS expr.Expr) BilateralBuilder
	// SigmaR set range sigma (from 0 to 1) (default 0.1).
	SigmaR(sigmaR float32) BilateralBuilder
	// SigmaRExpr set range sigma (from 0 to 1) (default 0.1).
	SigmaRExpr(sigmaR expr.Expr) BilateralBuilder
	// Planes set planes to filter (from 0 to 15) (default 1).
	Planes(planes int) BilateralBuilder
	// PlanesExpr set planes to filter (from 0 to 15) (default 1).
	PlanesExpr(planes expr.Expr) BilateralBuilder
}

// Bilateral creates a new BilateralBuilder to configure the "bilateral" filter.
func Bilateral(opts ...filter.Option) BilateralBuilder {
	return &implBilateralBuilder{
		f: filter.New("bilateral", 1, opts...),
	}
}

type implBilateralBuilder struct {
	f filter.Filter
}

func (bilateralBuilder *implBilateralBuilder) String() string {
	return bilateralBuilder.f.String()
}

func (bilateralBuilder *implBilateralBuilder) Outputs() int {
	return bilateralBuilder.f.Outputs()
}

func (bilateralBuilder *implBilateralBuilder) With(key string, value expr.Expr) filter.Filter {
	return bilateralBuilder.withOption(key, value)
}

func (bilateralBuilder *implBilateralBuilder) withOption(key string, value expr.Expr) BilateralBuilder {
	bCopy := *bilateralBuilder
	bCopy.f = bilateralBuilder.f.With(key, value)
	return &bCopy
}

func (bilateralBuilder *implBilateralBuilder) SigmaS(sigmaS float32) BilateralBuilder {
	return bilateralBuilder.withOption("sigmaS", expr.Float(sigmaS))
}

func (bilateralBuilder *implBilateralBuilder) SigmaSExpr(sigmaS expr.Expr) BilateralBuilder {
	return bilateralBuilder.withOption("sigmaS", sigmaS)
}

func (bilateralBuilder *implBilateralBuilder) SigmaR(sigmaR float32) BilateralBuilder {
	return bilateralBuilder.withOption("sigmaR", expr.Float(sigmaR))
}

func (bilateralBuilder *implBilateralBuilder) SigmaRExpr(sigmaR expr.Expr) BilateralBuilder {
	return bilateralBuilder.withOption("sigmaR", sigmaR)
}

func (bilateralBuilder *implBilateralBuilder) Planes(planes int) BilateralBuilder {
	return bilateralBuilder.withOption("planes", expr.Int(planes))
}

func (bilateralBuilder *implBilateralBuilder) PlanesExpr(planes expr.Expr) BilateralBuilder {
	return bilateralBuilder.withOption("planes", planes)
}
