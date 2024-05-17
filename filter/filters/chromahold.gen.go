// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ChromaholdBuilder Turns a certain color range into gray.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#chromahold
type ChromaholdBuilder interface {
	filter.Filter
	// Color set the chromahold key color (default "black").
	Color(color expr.Color) ChromaholdBuilder
	// ColorExpr set the chromahold key color (default "black").
	ColorExpr(color expr.Expr) ChromaholdBuilder
	// Similarity set the chromahold similarity value (from 1e-05 to 1) (default 0.01).
	Similarity(similarity float32) ChromaholdBuilder
	// SimilarityExpr set the chromahold similarity value (from 1e-05 to 1) (default 0.01).
	SimilarityExpr(similarity expr.Expr) ChromaholdBuilder
	// Blend set the chromahold blend value (from 0 to 1) (default 0).
	Blend(blend float32) ChromaholdBuilder
	// BlendExpr set the chromahold blend value (from 0 to 1) (default 0).
	BlendExpr(blend expr.Expr) ChromaholdBuilder
	// Yuv color parameter is in yuv instead of rgb (default false).
	Yuv(yuv bool) ChromaholdBuilder
	// YuvExpr color parameter is in yuv instead of rgb (default false).
	YuvExpr(yuv expr.Expr) ChromaholdBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ChromaholdBuilder
}

// Chromahold creates a new ChromaholdBuilder to configure the "chromahold" filter.
func Chromahold(opts ...filter.Option) ChromaholdBuilder {
	return &implChromaholdBuilder{
		f: filter.New("chromahold", 1, opts...),
	}
}

type implChromaholdBuilder struct {
	f filter.Filter
}

func (chromaholdBuilder *implChromaholdBuilder) String() string {
	return chromaholdBuilder.f.String()
}

func (chromaholdBuilder *implChromaholdBuilder) Outputs() int {
	return chromaholdBuilder.f.Outputs()
}

func (chromaholdBuilder *implChromaholdBuilder) With(key string, value expr.Expr) filter.Filter {
	return chromaholdBuilder.withOption(key, value)
}

func (chromaholdBuilder *implChromaholdBuilder) withOption(key string, value expr.Expr) ChromaholdBuilder {
	bCopy := *chromaholdBuilder
	bCopy.f = chromaholdBuilder.f.With(key, value)
	return &bCopy
}

func (chromaholdBuilder *implChromaholdBuilder) Color(color expr.Color) ChromaholdBuilder {
	return chromaholdBuilder.withOption("color", color)
}

func (chromaholdBuilder *implChromaholdBuilder) ColorExpr(color expr.Expr) ChromaholdBuilder {
	return chromaholdBuilder.withOption("color", color)
}

func (chromaholdBuilder *implChromaholdBuilder) Similarity(similarity float32) ChromaholdBuilder {
	return chromaholdBuilder.withOption("similarity", expr.Float(similarity))
}

func (chromaholdBuilder *implChromaholdBuilder) SimilarityExpr(similarity expr.Expr) ChromaholdBuilder {
	return chromaholdBuilder.withOption("similarity", similarity)
}

func (chromaholdBuilder *implChromaholdBuilder) Blend(blend float32) ChromaholdBuilder {
	return chromaholdBuilder.withOption("blend", expr.Float(blend))
}

func (chromaholdBuilder *implChromaholdBuilder) BlendExpr(blend expr.Expr) ChromaholdBuilder {
	return chromaholdBuilder.withOption("blend", blend)
}

func (chromaholdBuilder *implChromaholdBuilder) Yuv(yuv bool) ChromaholdBuilder {
	return chromaholdBuilder.withOption("yuv", expr.Bool(yuv))
}

func (chromaholdBuilder *implChromaholdBuilder) YuvExpr(yuv expr.Expr) ChromaholdBuilder {
	return chromaholdBuilder.withOption("yuv", yuv)
}

func (chromaholdBuilder *implChromaholdBuilder) Enable(enable expr.Expr) ChromaholdBuilder {
	return chromaholdBuilder.withOption("enable", enable)
}