// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// BackgroundkeyBuilder Turns a static background into transparency.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#backgroundkey
type BackgroundkeyBuilder interface {
	filter.Filter
	// Threshold set the scene change threshold (from 0 to 1) (default 0.08).
	Threshold(threshold float32) BackgroundkeyBuilder
	// ThresholdExpr set the scene change threshold (from 0 to 1) (default 0.08).
	ThresholdExpr(threshold expr.Expr) BackgroundkeyBuilder
	// Similarity set the similarity (from 0 to 1) (default 0.1).
	Similarity(similarity float32) BackgroundkeyBuilder
	// SimilarityExpr set the similarity (from 0 to 1) (default 0.1).
	SimilarityExpr(similarity expr.Expr) BackgroundkeyBuilder
	// Blend set the blend value (from 0 to 1) (default 0).
	Blend(blend float32) BackgroundkeyBuilder
	// BlendExpr set the blend value (from 0 to 1) (default 0).
	BlendExpr(blend expr.Expr) BackgroundkeyBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) BackgroundkeyBuilder
}

// Backgroundkey creates a new BackgroundkeyBuilder to configure the "backgroundkey" filter.
func Backgroundkey(opts ...filter.Option) BackgroundkeyBuilder {
	return &implBackgroundkeyBuilder{
		f: filter.New("backgroundkey", 1, opts...),
	}
}

type implBackgroundkeyBuilder struct {
	f filter.Filter
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) String() string {
	return backgroundkeyBuilder.f.String()
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) Outputs() int {
	return backgroundkeyBuilder.f.Outputs()
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) With(key string, value expr.Expr) filter.Filter {
	return backgroundkeyBuilder.withOption(key, value)
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) withOption(key string, value expr.Expr) BackgroundkeyBuilder {
	bCopy := *backgroundkeyBuilder
	bCopy.f = backgroundkeyBuilder.f.With(key, value)
	return &bCopy
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) Threshold(threshold float32) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("threshold", expr.Float(threshold))
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) ThresholdExpr(threshold expr.Expr) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("threshold", threshold)
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) Similarity(similarity float32) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("similarity", expr.Float(similarity))
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) SimilarityExpr(similarity expr.Expr) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("similarity", similarity)
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) Blend(blend float32) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("blend", expr.Float(blend))
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) BlendExpr(blend expr.Expr) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("blend", blend)
}

func (backgroundkeyBuilder *implBackgroundkeyBuilder) Enable(enable expr.Expr) BackgroundkeyBuilder {
	return backgroundkeyBuilder.withOption("enable", enable)
}