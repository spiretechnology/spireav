// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CrystalizerBuilder Simple audio noise sharpening filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#crystalizer
type CrystalizerBuilder interface {
	filter.Filter
	// I set intensity (from -10 to 10) (default 2).
	I(i float32) CrystalizerBuilder
	// IExpr set intensity (from -10 to 10) (default 2).
	IExpr(i expr.Expr) CrystalizerBuilder
	// C enable clipping (default true).
	C(c bool) CrystalizerBuilder
	// CExpr enable clipping (default true).
	CExpr(c expr.Expr) CrystalizerBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) CrystalizerBuilder
}

// Crystalizer creates a new CrystalizerBuilder to configure the "crystalizer" filter.
func Crystalizer(opts ...filter.Option) CrystalizerBuilder {
	return &implCrystalizerBuilder{
		f: filter.New("crystalizer", 1, opts...),
	}
}

type implCrystalizerBuilder struct {
	f filter.Filter
}

func (crystalizerBuilder *implCrystalizerBuilder) String() string {
	return crystalizerBuilder.f.String()
}

func (crystalizerBuilder *implCrystalizerBuilder) Outputs() int {
	return crystalizerBuilder.f.Outputs()
}

func (crystalizerBuilder *implCrystalizerBuilder) With(key string, value expr.Expr) filter.Filter {
	return crystalizerBuilder.withOption(key, value)
}

func (crystalizerBuilder *implCrystalizerBuilder) withOption(key string, value expr.Expr) CrystalizerBuilder {
	bCopy := *crystalizerBuilder
	bCopy.f = crystalizerBuilder.f.With(key, value)
	return &bCopy
}

func (crystalizerBuilder *implCrystalizerBuilder) I(i float32) CrystalizerBuilder {
	return crystalizerBuilder.withOption("i", expr.Float(i))
}

func (crystalizerBuilder *implCrystalizerBuilder) IExpr(i expr.Expr) CrystalizerBuilder {
	return crystalizerBuilder.withOption("i", i)
}

func (crystalizerBuilder *implCrystalizerBuilder) C(c bool) CrystalizerBuilder {
	return crystalizerBuilder.withOption("c", expr.Bool(c))
}

func (crystalizerBuilder *implCrystalizerBuilder) CExpr(c expr.Expr) CrystalizerBuilder {
	return crystalizerBuilder.withOption("c", c)
}

func (crystalizerBuilder *implCrystalizerBuilder) Enable(enable expr.Expr) CrystalizerBuilder {
	return crystalizerBuilder.withOption("enable", enable)
}