// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PhaseBuilder Phase shift fields.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#phase
type PhaseBuilder interface {
	filter.Filter
	// Mode set phase mode (from 0 to 8) (default A).
	Mode(mode int) PhaseBuilder
	// ModeExpr set phase mode (from 0 to 8) (default A).
	ModeExpr(mode expr.Expr) PhaseBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) PhaseBuilder
}

// Phase creates a new PhaseBuilder to configure the "phase" filter.
func Phase(opts ...filter.Option) PhaseBuilder {
	return &implPhaseBuilder{
		f: filter.New("phase", 1, opts...),
	}
}

type implPhaseBuilder struct {
	f filter.Filter
}

func (phaseBuilder *implPhaseBuilder) String() string {
	return phaseBuilder.f.String()
}

func (phaseBuilder *implPhaseBuilder) Outputs() int {
	return phaseBuilder.f.Outputs()
}

func (phaseBuilder *implPhaseBuilder) With(key string, value expr.Expr) filter.Filter {
	return phaseBuilder.withOption(key, value)
}

func (phaseBuilder *implPhaseBuilder) withOption(key string, value expr.Expr) PhaseBuilder {
	bCopy := *phaseBuilder
	bCopy.f = phaseBuilder.f.With(key, value)
	return &bCopy
}

func (phaseBuilder *implPhaseBuilder) Mode(mode int) PhaseBuilder {
	return phaseBuilder.withOption("mode", expr.Int(mode))
}

func (phaseBuilder *implPhaseBuilder) ModeExpr(mode expr.Expr) PhaseBuilder {
	return phaseBuilder.withOption("mode", mode)
}

func (phaseBuilder *implPhaseBuilder) Enable(enable expr.Expr) PhaseBuilder {
	return phaseBuilder.withOption("enable", enable)
}
