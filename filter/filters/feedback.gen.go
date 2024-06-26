// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FeedbackBuilder Apply feedback video filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#feedback
type FeedbackBuilder interface {
	filter.Filter
	// X set top left crop position (from 0 to INT_MAX) (default 0).
	X(x int) FeedbackBuilder
	// XExpr set top left crop position (from 0 to INT_MAX) (default 0).
	XExpr(x expr.Expr) FeedbackBuilder
	// Y set top left crop position (from 0 to INT_MAX) (default 0).
	Y(y int) FeedbackBuilder
	// YExpr set top left crop position (from 0 to INT_MAX) (default 0).
	YExpr(y expr.Expr) FeedbackBuilder
	// W set crop size (from 0 to INT_MAX) (default 0).
	W(w int) FeedbackBuilder
	// H set crop size (from 0 to INT_MAX) (default 0).
	H(h int) FeedbackBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) FeedbackBuilder
}

// Feedback creates a new FeedbackBuilder to configure the "feedback" filter.
func Feedback(opts ...filter.Option) FeedbackBuilder {
	return &implFeedbackBuilder{
		f: filter.New("feedback", 2, opts...),
	}
}

type implFeedbackBuilder struct {
	f filter.Filter
}

func (feedbackBuilder *implFeedbackBuilder) String() string {
	return feedbackBuilder.f.String()
}

func (feedbackBuilder *implFeedbackBuilder) Outputs() int {
	return feedbackBuilder.f.Outputs()
}

func (feedbackBuilder *implFeedbackBuilder) With(key string, value expr.Expr) filter.Filter {
	return feedbackBuilder.withOption(key, value)
}

func (feedbackBuilder *implFeedbackBuilder) withOption(key string, value expr.Expr) FeedbackBuilder {
	bCopy := *feedbackBuilder
	bCopy.f = feedbackBuilder.f.With(key, value)
	return &bCopy
}

func (feedbackBuilder *implFeedbackBuilder) X(x int) FeedbackBuilder {
	return feedbackBuilder.withOption("x", expr.Int(x))
}

func (feedbackBuilder *implFeedbackBuilder) XExpr(x expr.Expr) FeedbackBuilder {
	return feedbackBuilder.withOption("x", x)
}

func (feedbackBuilder *implFeedbackBuilder) Y(y int) FeedbackBuilder {
	return feedbackBuilder.withOption("y", expr.Int(y))
}

func (feedbackBuilder *implFeedbackBuilder) YExpr(y expr.Expr) FeedbackBuilder {
	return feedbackBuilder.withOption("y", y)
}

func (feedbackBuilder *implFeedbackBuilder) W(w int) FeedbackBuilder {
	return feedbackBuilder.withOption("w", expr.Int(w))
}

func (feedbackBuilder *implFeedbackBuilder) H(h int) FeedbackBuilder {
	return feedbackBuilder.withOption("h", expr.Int(h))
}

func (feedbackBuilder *implFeedbackBuilder) Enable(enable expr.Expr) FeedbackBuilder {
	return feedbackBuilder.withOption("enable", enable)
}
