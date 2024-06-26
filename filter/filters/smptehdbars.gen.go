// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SmptehdbarsBuilder Generate SMPTE HD color bars.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#smptehdbars
type SmptehdbarsBuilder interface {
	filter.Filter
	// Size set video size (default "320x240").
	Size(size expr.Size) SmptehdbarsBuilder
	// S set video size (default "320x240").
	S(s expr.Size) SmptehdbarsBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) SmptehdbarsBuilder
	// R set video rate (default "25").
	R(r expr.Rational) SmptehdbarsBuilder
	// Duration set video duration (default -0.000001).
	Duration(duration time.Duration) SmptehdbarsBuilder
	// D set video duration (default -0.000001).
	D(d time.Duration) SmptehdbarsBuilder
	// Sar set video sample aspect ratio (from 0 to INT_MAX) (default 1/1).
	Sar(sar expr.Rational) SmptehdbarsBuilder
}

// Smptehdbars creates a new SmptehdbarsBuilder to configure the "smptehdbars" filter.
func Smptehdbars(opts ...filter.Option) SmptehdbarsBuilder {
	return &implSmptehdbarsBuilder{
		f: filter.New("smptehdbars", 1, opts...),
	}
}

type implSmptehdbarsBuilder struct {
	f filter.Filter
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) String() string {
	return smptehdbarsBuilder.f.String()
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) Outputs() int {
	return smptehdbarsBuilder.f.Outputs()
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) With(key string, value expr.Expr) filter.Filter {
	return smptehdbarsBuilder.withOption(key, value)
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) withOption(key string, value expr.Expr) SmptehdbarsBuilder {
	bCopy := *smptehdbarsBuilder
	bCopy.f = smptehdbarsBuilder.f.With(key, value)
	return &bCopy
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) Size(size expr.Size) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("size", size)
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) S(s expr.Size) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("s", s)
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) Rate(rate expr.Rational) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("rate", rate)
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) R(r expr.Rational) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("r", r)
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) Duration(duration time.Duration) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("duration", expr.Duration(duration))
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) D(d time.Duration) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("d", expr.Duration(d))
}

func (smptehdbarsBuilder *implSmptehdbarsBuilder) Sar(sar expr.Rational) SmptehdbarsBuilder {
	return smptehdbarsBuilder.withOption("sar", sar)
}
