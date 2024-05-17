// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// StereowidenBuilder Apply stereo widening effect.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#stereowiden
type StereowidenBuilder interface {
	filter.Filter
	// Delay set delay time (from 1 to 100) (default 20).
	Delay(delay float32) StereowidenBuilder
	// Feedback set feedback gain (from 0 to 0.9) (default 0.3).
	Feedback(feedback float32) StereowidenBuilder
	// FeedbackExpr set feedback gain (from 0 to 0.9) (default 0.3).
	FeedbackExpr(feedback expr.Expr) StereowidenBuilder
	// Crossfeed set cross feed (from 0 to 0.8) (default 0.3).
	Crossfeed(crossfeed float32) StereowidenBuilder
	// CrossfeedExpr set cross feed (from 0 to 0.8) (default 0.3).
	CrossfeedExpr(crossfeed expr.Expr) StereowidenBuilder
	// Drymix set dry-mix (from 0 to 1) (default 0.8).
	Drymix(drymix float32) StereowidenBuilder
	// DrymixExpr set dry-mix (from 0 to 1) (default 0.8).
	DrymixExpr(drymix expr.Expr) StereowidenBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) StereowidenBuilder
}

// Stereowiden creates a new StereowidenBuilder to configure the "stereowiden" filter.
func Stereowiden(opts ...filter.Option) StereowidenBuilder {
	return &implStereowidenBuilder{
		f: filter.New("stereowiden", 1, opts...),
	}
}

type implStereowidenBuilder struct {
	f filter.Filter
}

func (stereowidenBuilder *implStereowidenBuilder) String() string {
	return stereowidenBuilder.f.String()
}

func (stereowidenBuilder *implStereowidenBuilder) Outputs() int {
	return stereowidenBuilder.f.Outputs()
}

func (stereowidenBuilder *implStereowidenBuilder) With(key string, value expr.Expr) filter.Filter {
	return stereowidenBuilder.withOption(key, value)
}

func (stereowidenBuilder *implStereowidenBuilder) withOption(key string, value expr.Expr) StereowidenBuilder {
	bCopy := *stereowidenBuilder
	bCopy.f = stereowidenBuilder.f.With(key, value)
	return &bCopy
}

func (stereowidenBuilder *implStereowidenBuilder) Delay(delay float32) StereowidenBuilder {
	return stereowidenBuilder.withOption("delay", expr.Float(delay))
}

func (stereowidenBuilder *implStereowidenBuilder) Feedback(feedback float32) StereowidenBuilder {
	return stereowidenBuilder.withOption("feedback", expr.Float(feedback))
}

func (stereowidenBuilder *implStereowidenBuilder) FeedbackExpr(feedback expr.Expr) StereowidenBuilder {
	return stereowidenBuilder.withOption("feedback", feedback)
}

func (stereowidenBuilder *implStereowidenBuilder) Crossfeed(crossfeed float32) StereowidenBuilder {
	return stereowidenBuilder.withOption("crossfeed", expr.Float(crossfeed))
}

func (stereowidenBuilder *implStereowidenBuilder) CrossfeedExpr(crossfeed expr.Expr) StereowidenBuilder {
	return stereowidenBuilder.withOption("crossfeed", crossfeed)
}

func (stereowidenBuilder *implStereowidenBuilder) Drymix(drymix float32) StereowidenBuilder {
	return stereowidenBuilder.withOption("drymix", expr.Float(drymix))
}

func (stereowidenBuilder *implStereowidenBuilder) DrymixExpr(drymix expr.Expr) StereowidenBuilder {
	return stereowidenBuilder.withOption("drymix", drymix)
}

func (stereowidenBuilder *implStereowidenBuilder) Enable(enable expr.Expr) StereowidenBuilder {
	return stereowidenBuilder.withOption("enable", enable)
}