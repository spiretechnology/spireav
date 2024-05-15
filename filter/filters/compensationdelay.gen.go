// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CompensationdelayBuilder Audio Compensation Delay Line.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#compensationdelay
type CompensationdelayBuilder interface {
	filter.Filter
	// Mm set mm distance (from 0 to 10) (default 0).
	Mm(mm int) CompensationdelayBuilder
	// MmExpr set mm distance (from 0 to 10) (default 0).
	MmExpr(mm expr.Expr) CompensationdelayBuilder
	// Cm set cm distance (from 0 to 100) (default 0).
	Cm(cm int) CompensationdelayBuilder
	// CmExpr set cm distance (from 0 to 100) (default 0).
	CmExpr(cm expr.Expr) CompensationdelayBuilder
	// M set meter distance (from 0 to 100) (default 0).
	M(m int) CompensationdelayBuilder
	// MExpr set meter distance (from 0 to 100) (default 0).
	MExpr(m expr.Expr) CompensationdelayBuilder
	// Dry set dry amount (from 0 to 1) (default 0).
	Dry(dry float64) CompensationdelayBuilder
	// DryExpr set dry amount (from 0 to 1) (default 0).
	DryExpr(dry expr.Expr) CompensationdelayBuilder
	// Wet set wet amount (from 0 to 1) (default 1).
	Wet(wet float64) CompensationdelayBuilder
	// WetExpr set wet amount (from 0 to 1) (default 1).
	WetExpr(wet expr.Expr) CompensationdelayBuilder
	// Temp set temperature °C (from -50 to 50) (default 20).
	Temp(temp int) CompensationdelayBuilder
	// TempExpr set temperature °C (from -50 to 50) (default 20).
	TempExpr(temp expr.Expr) CompensationdelayBuilder
}

// Compensationdelay creates a new CompensationdelayBuilder to configure the "compensationdelay" filter.
func Compensationdelay(opts ...filter.Option) CompensationdelayBuilder {
	return &implCompensationdelayBuilder{
		f: filter.New("compensationdelay", 1, opts...),
	}
}

type implCompensationdelayBuilder struct {
	f filter.Filter
}

func (compensationdelayBuilder *implCompensationdelayBuilder) String() string {
	return compensationdelayBuilder.f.String()
}

func (compensationdelayBuilder *implCompensationdelayBuilder) Outputs() int {
	return compensationdelayBuilder.f.Outputs()
}

func (compensationdelayBuilder *implCompensationdelayBuilder) With(key string, value expr.Expr) filter.Filter {
	return compensationdelayBuilder.withOption(key, value)
}

func (compensationdelayBuilder *implCompensationdelayBuilder) withOption(key string, value expr.Expr) CompensationdelayBuilder {
	bCopy := *compensationdelayBuilder
	bCopy.f = compensationdelayBuilder.f.With(key, value)
	return &bCopy
}

func (compensationdelayBuilder *implCompensationdelayBuilder) Mm(mm int) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("mm", expr.Int(mm))
}

func (compensationdelayBuilder *implCompensationdelayBuilder) MmExpr(mm expr.Expr) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("mm", mm)
}

func (compensationdelayBuilder *implCompensationdelayBuilder) Cm(cm int) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("cm", expr.Int(cm))
}

func (compensationdelayBuilder *implCompensationdelayBuilder) CmExpr(cm expr.Expr) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("cm", cm)
}

func (compensationdelayBuilder *implCompensationdelayBuilder) M(m int) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("m", expr.Int(m))
}

func (compensationdelayBuilder *implCompensationdelayBuilder) MExpr(m expr.Expr) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("m", m)
}

func (compensationdelayBuilder *implCompensationdelayBuilder) Dry(dry float64) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("dry", expr.Double(dry))
}

func (compensationdelayBuilder *implCompensationdelayBuilder) DryExpr(dry expr.Expr) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("dry", dry)
}

func (compensationdelayBuilder *implCompensationdelayBuilder) Wet(wet float64) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("wet", expr.Double(wet))
}

func (compensationdelayBuilder *implCompensationdelayBuilder) WetExpr(wet expr.Expr) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("wet", wet)
}

func (compensationdelayBuilder *implCompensationdelayBuilder) Temp(temp int) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("temp", expr.Int(temp))
}

func (compensationdelayBuilder *implCompensationdelayBuilder) TempExpr(temp expr.Expr) CompensationdelayBuilder {
	return compensationdelayBuilder.withOption("temp", temp)
}
