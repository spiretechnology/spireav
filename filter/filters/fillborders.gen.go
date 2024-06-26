// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FillbordersBuilder Fill borders of the input video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#fillborders
type FillbordersBuilder interface {
	filter.Filter
	// Left set the left fill border (from 0 to INT_MAX) (default 0).
	Left(left int) FillbordersBuilder
	// LeftExpr set the left fill border (from 0 to INT_MAX) (default 0).
	LeftExpr(left expr.Expr) FillbordersBuilder
	// Right set the right fill border (from 0 to INT_MAX) (default 0).
	Right(right int) FillbordersBuilder
	// RightExpr set the right fill border (from 0 to INT_MAX) (default 0).
	RightExpr(right expr.Expr) FillbordersBuilder
	// Top set the top fill border (from 0 to INT_MAX) (default 0).
	Top(top int) FillbordersBuilder
	// TopExpr set the top fill border (from 0 to INT_MAX) (default 0).
	TopExpr(top expr.Expr) FillbordersBuilder
	// Bottom set the bottom fill border (from 0 to INT_MAX) (default 0).
	Bottom(bottom int) FillbordersBuilder
	// BottomExpr set the bottom fill border (from 0 to INT_MAX) (default 0).
	BottomExpr(bottom expr.Expr) FillbordersBuilder
	// Mode set the fill borders mode (from 0 to 6) (default smear).
	Mode(mode int) FillbordersBuilder
	// ModeExpr set the fill borders mode (from 0 to 6) (default smear).
	ModeExpr(mode expr.Expr) FillbordersBuilder
	// Color set the color for the fixed/fade mode (default "black").
	Color(color expr.Color) FillbordersBuilder
	// ColorExpr set the color for the fixed/fade mode (default "black").
	ColorExpr(color expr.Expr) FillbordersBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) FillbordersBuilder
}

// Fillborders creates a new FillbordersBuilder to configure the "fillborders" filter.
func Fillborders(opts ...filter.Option) FillbordersBuilder {
	return &implFillbordersBuilder{
		f: filter.New("fillborders", 1, opts...),
	}
}

type implFillbordersBuilder struct {
	f filter.Filter
}

func (fillbordersBuilder *implFillbordersBuilder) String() string {
	return fillbordersBuilder.f.String()
}

func (fillbordersBuilder *implFillbordersBuilder) Outputs() int {
	return fillbordersBuilder.f.Outputs()
}

func (fillbordersBuilder *implFillbordersBuilder) With(key string, value expr.Expr) filter.Filter {
	return fillbordersBuilder.withOption(key, value)
}

func (fillbordersBuilder *implFillbordersBuilder) withOption(key string, value expr.Expr) FillbordersBuilder {
	bCopy := *fillbordersBuilder
	bCopy.f = fillbordersBuilder.f.With(key, value)
	return &bCopy
}

func (fillbordersBuilder *implFillbordersBuilder) Left(left int) FillbordersBuilder {
	return fillbordersBuilder.withOption("left", expr.Int(left))
}

func (fillbordersBuilder *implFillbordersBuilder) LeftExpr(left expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("left", left)
}

func (fillbordersBuilder *implFillbordersBuilder) Right(right int) FillbordersBuilder {
	return fillbordersBuilder.withOption("right", expr.Int(right))
}

func (fillbordersBuilder *implFillbordersBuilder) RightExpr(right expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("right", right)
}

func (fillbordersBuilder *implFillbordersBuilder) Top(top int) FillbordersBuilder {
	return fillbordersBuilder.withOption("top", expr.Int(top))
}

func (fillbordersBuilder *implFillbordersBuilder) TopExpr(top expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("top", top)
}

func (fillbordersBuilder *implFillbordersBuilder) Bottom(bottom int) FillbordersBuilder {
	return fillbordersBuilder.withOption("bottom", expr.Int(bottom))
}

func (fillbordersBuilder *implFillbordersBuilder) BottomExpr(bottom expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("bottom", bottom)
}

func (fillbordersBuilder *implFillbordersBuilder) Mode(mode int) FillbordersBuilder {
	return fillbordersBuilder.withOption("mode", expr.Int(mode))
}

func (fillbordersBuilder *implFillbordersBuilder) ModeExpr(mode expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("mode", mode)
}

func (fillbordersBuilder *implFillbordersBuilder) Color(color expr.Color) FillbordersBuilder {
	return fillbordersBuilder.withOption("color", color)
}

func (fillbordersBuilder *implFillbordersBuilder) ColorExpr(color expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("color", color)
}

func (fillbordersBuilder *implFillbordersBuilder) Enable(enable expr.Expr) FillbordersBuilder {
	return fillbordersBuilder.withOption("enable", enable)
}
