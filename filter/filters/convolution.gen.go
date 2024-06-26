// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ConvolutionBuilder Apply convolution filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#convolution
type ConvolutionBuilder interface {
	filter.Filter
	// With0m set matrix for 1st plane (default "0 0 0 0 1 0 0 0 0").
	With0m(val0m string) ConvolutionBuilder
	// With0mExpr set matrix for 1st plane (default "0 0 0 0 1 0 0 0 0").
	With0mExpr(val0m expr.Expr) ConvolutionBuilder
	// With1m set matrix for 2nd plane (default "0 0 0 0 1 0 0 0 0").
	With1m(val1m string) ConvolutionBuilder
	// With1mExpr set matrix for 2nd plane (default "0 0 0 0 1 0 0 0 0").
	With1mExpr(val1m expr.Expr) ConvolutionBuilder
	// With2m set matrix for 3rd plane (default "0 0 0 0 1 0 0 0 0").
	With2m(val2m string) ConvolutionBuilder
	// With2mExpr set matrix for 3rd plane (default "0 0 0 0 1 0 0 0 0").
	With2mExpr(val2m expr.Expr) ConvolutionBuilder
	// With3m set matrix for 4th plane (default "0 0 0 0 1 0 0 0 0").
	With3m(val3m string) ConvolutionBuilder
	// With3mExpr set matrix for 4th plane (default "0 0 0 0 1 0 0 0 0").
	With3mExpr(val3m expr.Expr) ConvolutionBuilder
	// With0rdiv set rdiv for 1st plane (from 0 to INT_MAX) (default 0).
	With0rdiv(val0rdiv float32) ConvolutionBuilder
	// With0rdivExpr set rdiv for 1st plane (from 0 to INT_MAX) (default 0).
	With0rdivExpr(val0rdiv expr.Expr) ConvolutionBuilder
	// With1rdiv set rdiv for 2nd plane (from 0 to INT_MAX) (default 0).
	With1rdiv(val1rdiv float32) ConvolutionBuilder
	// With1rdivExpr set rdiv for 2nd plane (from 0 to INT_MAX) (default 0).
	With1rdivExpr(val1rdiv expr.Expr) ConvolutionBuilder
	// With2rdiv set rdiv for 3rd plane (from 0 to INT_MAX) (default 0).
	With2rdiv(val2rdiv float32) ConvolutionBuilder
	// With2rdivExpr set rdiv for 3rd plane (from 0 to INT_MAX) (default 0).
	With2rdivExpr(val2rdiv expr.Expr) ConvolutionBuilder
	// With3rdiv set rdiv for 4th plane (from 0 to INT_MAX) (default 0).
	With3rdiv(val3rdiv float32) ConvolutionBuilder
	// With3rdivExpr set rdiv for 4th plane (from 0 to INT_MAX) (default 0).
	With3rdivExpr(val3rdiv expr.Expr) ConvolutionBuilder
	// With0bias set bias for 1st plane (from 0 to INT_MAX) (default 0).
	With0bias(val0bias float32) ConvolutionBuilder
	// With0biasExpr set bias for 1st plane (from 0 to INT_MAX) (default 0).
	With0biasExpr(val0bias expr.Expr) ConvolutionBuilder
	// With1bias set bias for 2nd plane (from 0 to INT_MAX) (default 0).
	With1bias(val1bias float32) ConvolutionBuilder
	// With1biasExpr set bias for 2nd plane (from 0 to INT_MAX) (default 0).
	With1biasExpr(val1bias expr.Expr) ConvolutionBuilder
	// With2bias set bias for 3rd plane (from 0 to INT_MAX) (default 0).
	With2bias(val2bias float32) ConvolutionBuilder
	// With2biasExpr set bias for 3rd plane (from 0 to INT_MAX) (default 0).
	With2biasExpr(val2bias expr.Expr) ConvolutionBuilder
	// With3bias set bias for 4th plane (from 0 to INT_MAX) (default 0).
	With3bias(val3bias float32) ConvolutionBuilder
	// With3biasExpr set bias for 4th plane (from 0 to INT_MAX) (default 0).
	With3biasExpr(val3bias expr.Expr) ConvolutionBuilder
	// With0mode set matrix mode for 1st plane (from 0 to 2) (default square).
	With0mode(val0mode int) ConvolutionBuilder
	// With0modeExpr set matrix mode for 1st plane (from 0 to 2) (default square).
	With0modeExpr(val0mode expr.Expr) ConvolutionBuilder
	// With1mode set matrix mode for 2nd plane (from 0 to 2) (default square).
	With1mode(val1mode int) ConvolutionBuilder
	// With1modeExpr set matrix mode for 2nd plane (from 0 to 2) (default square).
	With1modeExpr(val1mode expr.Expr) ConvolutionBuilder
	// With2mode set matrix mode for 3rd plane (from 0 to 2) (default square).
	With2mode(val2mode int) ConvolutionBuilder
	// With2modeExpr set matrix mode for 3rd plane (from 0 to 2) (default square).
	With2modeExpr(val2mode expr.Expr) ConvolutionBuilder
	// With3mode set matrix mode for 4th plane (from 0 to 2) (default square).
	With3mode(val3mode int) ConvolutionBuilder
	// With3modeExpr set matrix mode for 4th plane (from 0 to 2) (default square).
	With3modeExpr(val3mode expr.Expr) ConvolutionBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ConvolutionBuilder
}

// Convolution creates a new ConvolutionBuilder to configure the "convolution" filter.
func Convolution(opts ...filter.Option) ConvolutionBuilder {
	return &implConvolutionBuilder{
		f: filter.New("convolution", 1, opts...),
	}
}

type implConvolutionBuilder struct {
	f filter.Filter
}

func (convolutionBuilder *implConvolutionBuilder) String() string {
	return convolutionBuilder.f.String()
}

func (convolutionBuilder *implConvolutionBuilder) Outputs() int {
	return convolutionBuilder.f.Outputs()
}

func (convolutionBuilder *implConvolutionBuilder) With(key string, value expr.Expr) filter.Filter {
	return convolutionBuilder.withOption(key, value)
}

func (convolutionBuilder *implConvolutionBuilder) withOption(key string, value expr.Expr) ConvolutionBuilder {
	bCopy := *convolutionBuilder
	bCopy.f = convolutionBuilder.f.With(key, value)
	return &bCopy
}

func (convolutionBuilder *implConvolutionBuilder) With0m(val0m string) ConvolutionBuilder {
	return convolutionBuilder.withOption("0m", expr.String(val0m))
}

func (convolutionBuilder *implConvolutionBuilder) With0mExpr(val0m expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("0m", val0m)
}

func (convolutionBuilder *implConvolutionBuilder) With1m(val1m string) ConvolutionBuilder {
	return convolutionBuilder.withOption("1m", expr.String(val1m))
}

func (convolutionBuilder *implConvolutionBuilder) With1mExpr(val1m expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("1m", val1m)
}

func (convolutionBuilder *implConvolutionBuilder) With2m(val2m string) ConvolutionBuilder {
	return convolutionBuilder.withOption("2m", expr.String(val2m))
}

func (convolutionBuilder *implConvolutionBuilder) With2mExpr(val2m expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("2m", val2m)
}

func (convolutionBuilder *implConvolutionBuilder) With3m(val3m string) ConvolutionBuilder {
	return convolutionBuilder.withOption("3m", expr.String(val3m))
}

func (convolutionBuilder *implConvolutionBuilder) With3mExpr(val3m expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("3m", val3m)
}

func (convolutionBuilder *implConvolutionBuilder) With0rdiv(val0rdiv float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("0rdiv", expr.Float(val0rdiv))
}

func (convolutionBuilder *implConvolutionBuilder) With0rdivExpr(val0rdiv expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("0rdiv", val0rdiv)
}

func (convolutionBuilder *implConvolutionBuilder) With1rdiv(val1rdiv float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("1rdiv", expr.Float(val1rdiv))
}

func (convolutionBuilder *implConvolutionBuilder) With1rdivExpr(val1rdiv expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("1rdiv", val1rdiv)
}

func (convolutionBuilder *implConvolutionBuilder) With2rdiv(val2rdiv float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("2rdiv", expr.Float(val2rdiv))
}

func (convolutionBuilder *implConvolutionBuilder) With2rdivExpr(val2rdiv expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("2rdiv", val2rdiv)
}

func (convolutionBuilder *implConvolutionBuilder) With3rdiv(val3rdiv float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("3rdiv", expr.Float(val3rdiv))
}

func (convolutionBuilder *implConvolutionBuilder) With3rdivExpr(val3rdiv expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("3rdiv", val3rdiv)
}

func (convolutionBuilder *implConvolutionBuilder) With0bias(val0bias float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("0bias", expr.Float(val0bias))
}

func (convolutionBuilder *implConvolutionBuilder) With0biasExpr(val0bias expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("0bias", val0bias)
}

func (convolutionBuilder *implConvolutionBuilder) With1bias(val1bias float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("1bias", expr.Float(val1bias))
}

func (convolutionBuilder *implConvolutionBuilder) With1biasExpr(val1bias expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("1bias", val1bias)
}

func (convolutionBuilder *implConvolutionBuilder) With2bias(val2bias float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("2bias", expr.Float(val2bias))
}

func (convolutionBuilder *implConvolutionBuilder) With2biasExpr(val2bias expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("2bias", val2bias)
}

func (convolutionBuilder *implConvolutionBuilder) With3bias(val3bias float32) ConvolutionBuilder {
	return convolutionBuilder.withOption("3bias", expr.Float(val3bias))
}

func (convolutionBuilder *implConvolutionBuilder) With3biasExpr(val3bias expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("3bias", val3bias)
}

func (convolutionBuilder *implConvolutionBuilder) With0mode(val0mode int) ConvolutionBuilder {
	return convolutionBuilder.withOption("0mode", expr.Int(val0mode))
}

func (convolutionBuilder *implConvolutionBuilder) With0modeExpr(val0mode expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("0mode", val0mode)
}

func (convolutionBuilder *implConvolutionBuilder) With1mode(val1mode int) ConvolutionBuilder {
	return convolutionBuilder.withOption("1mode", expr.Int(val1mode))
}

func (convolutionBuilder *implConvolutionBuilder) With1modeExpr(val1mode expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("1mode", val1mode)
}

func (convolutionBuilder *implConvolutionBuilder) With2mode(val2mode int) ConvolutionBuilder {
	return convolutionBuilder.withOption("2mode", expr.Int(val2mode))
}

func (convolutionBuilder *implConvolutionBuilder) With2modeExpr(val2mode expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("2mode", val2mode)
}

func (convolutionBuilder *implConvolutionBuilder) With3mode(val3mode int) ConvolutionBuilder {
	return convolutionBuilder.withOption("3mode", expr.Int(val3mode))
}

func (convolutionBuilder *implConvolutionBuilder) With3modeExpr(val3mode expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("3mode", val3mode)
}

func (convolutionBuilder *implConvolutionBuilder) Enable(enable expr.Expr) ConvolutionBuilder {
	return convolutionBuilder.withOption("enable", enable)
}
