// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CurvesBuilder Adjust components curves.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#curves
type CurvesBuilder interface {
	filter.Filter
	// Preset select a color curves preset (from 0 to 10) (default none).
	Preset(preset int) CurvesBuilder
	// PresetExpr select a color curves preset (from 0 to 10) (default none).
	PresetExpr(preset expr.Expr) CurvesBuilder
	// Master set master points coordinates.
	Master(master string) CurvesBuilder
	// MasterExpr set master points coordinates.
	MasterExpr(master expr.Expr) CurvesBuilder
	// M set master points coordinates.
	M(m string) CurvesBuilder
	// MExpr set master points coordinates.
	MExpr(m expr.Expr) CurvesBuilder
	// Red set red points coordinates.
	Red(red string) CurvesBuilder
	// RedExpr set red points coordinates.
	RedExpr(red expr.Expr) CurvesBuilder
	// R set red points coordinates.
	R(r string) CurvesBuilder
	// RExpr set red points coordinates.
	RExpr(r expr.Expr) CurvesBuilder
	// Green set green points coordinates.
	Green(green string) CurvesBuilder
	// GreenExpr set green points coordinates.
	GreenExpr(green expr.Expr) CurvesBuilder
	// G set green points coordinates.
	G(g string) CurvesBuilder
	// GExpr set green points coordinates.
	GExpr(g expr.Expr) CurvesBuilder
	// Blue set blue points coordinates.
	Blue(blue string) CurvesBuilder
	// BlueExpr set blue points coordinates.
	BlueExpr(blue expr.Expr) CurvesBuilder
	// B set blue points coordinates.
	B(b string) CurvesBuilder
	// BExpr set blue points coordinates.
	BExpr(b expr.Expr) CurvesBuilder
	// All set points coordinates for all components.
	All(all string) CurvesBuilder
	// AllExpr set points coordinates for all components.
	AllExpr(all expr.Expr) CurvesBuilder
	// Psfile set Photoshop curves file name.
	Psfile(psfile string) CurvesBuilder
	// PsfileExpr set Photoshop curves file name.
	PsfileExpr(psfile expr.Expr) CurvesBuilder
	// Plot save Gnuplot script of the curves in specified file.
	Plot(plot string) CurvesBuilder
	// PlotExpr save Gnuplot script of the curves in specified file.
	PlotExpr(plot expr.Expr) CurvesBuilder
	// Interp specify the kind of interpolation (from 0 to 1) (default natural).
	Interp(interp int) CurvesBuilder
	// InterpExpr specify the kind of interpolation (from 0 to 1) (default natural).
	InterpExpr(interp expr.Expr) CurvesBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) CurvesBuilder
}

// Curves creates a new CurvesBuilder to configure the "curves" filter.
func Curves(opts ...filter.Option) CurvesBuilder {
	return &implCurvesBuilder{
		f: filter.New("curves", 1, opts...),
	}
}

type implCurvesBuilder struct {
	f filter.Filter
}

func (curvesBuilder *implCurvesBuilder) String() string {
	return curvesBuilder.f.String()
}

func (curvesBuilder *implCurvesBuilder) Outputs() int {
	return curvesBuilder.f.Outputs()
}

func (curvesBuilder *implCurvesBuilder) With(key string, value expr.Expr) filter.Filter {
	return curvesBuilder.withOption(key, value)
}

func (curvesBuilder *implCurvesBuilder) withOption(key string, value expr.Expr) CurvesBuilder {
	bCopy := *curvesBuilder
	bCopy.f = curvesBuilder.f.With(key, value)
	return &bCopy
}

func (curvesBuilder *implCurvesBuilder) Preset(preset int) CurvesBuilder {
	return curvesBuilder.withOption("preset", expr.Int(preset))
}

func (curvesBuilder *implCurvesBuilder) PresetExpr(preset expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("preset", preset)
}

func (curvesBuilder *implCurvesBuilder) Master(master string) CurvesBuilder {
	return curvesBuilder.withOption("master", expr.String(master))
}

func (curvesBuilder *implCurvesBuilder) MasterExpr(master expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("master", master)
}

func (curvesBuilder *implCurvesBuilder) M(m string) CurvesBuilder {
	return curvesBuilder.withOption("m", expr.String(m))
}

func (curvesBuilder *implCurvesBuilder) MExpr(m expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("m", m)
}

func (curvesBuilder *implCurvesBuilder) Red(red string) CurvesBuilder {
	return curvesBuilder.withOption("red", expr.String(red))
}

func (curvesBuilder *implCurvesBuilder) RedExpr(red expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("red", red)
}

func (curvesBuilder *implCurvesBuilder) R(r string) CurvesBuilder {
	return curvesBuilder.withOption("r", expr.String(r))
}

func (curvesBuilder *implCurvesBuilder) RExpr(r expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("r", r)
}

func (curvesBuilder *implCurvesBuilder) Green(green string) CurvesBuilder {
	return curvesBuilder.withOption("green", expr.String(green))
}

func (curvesBuilder *implCurvesBuilder) GreenExpr(green expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("green", green)
}

func (curvesBuilder *implCurvesBuilder) G(g string) CurvesBuilder {
	return curvesBuilder.withOption("g", expr.String(g))
}

func (curvesBuilder *implCurvesBuilder) GExpr(g expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("g", g)
}

func (curvesBuilder *implCurvesBuilder) Blue(blue string) CurvesBuilder {
	return curvesBuilder.withOption("blue", expr.String(blue))
}

func (curvesBuilder *implCurvesBuilder) BlueExpr(blue expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("blue", blue)
}

func (curvesBuilder *implCurvesBuilder) B(b string) CurvesBuilder {
	return curvesBuilder.withOption("b", expr.String(b))
}

func (curvesBuilder *implCurvesBuilder) BExpr(b expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("b", b)
}

func (curvesBuilder *implCurvesBuilder) All(all string) CurvesBuilder {
	return curvesBuilder.withOption("all", expr.String(all))
}

func (curvesBuilder *implCurvesBuilder) AllExpr(all expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("all", all)
}

func (curvesBuilder *implCurvesBuilder) Psfile(psfile string) CurvesBuilder {
	return curvesBuilder.withOption("psfile", expr.String(psfile))
}

func (curvesBuilder *implCurvesBuilder) PsfileExpr(psfile expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("psfile", psfile)
}

func (curvesBuilder *implCurvesBuilder) Plot(plot string) CurvesBuilder {
	return curvesBuilder.withOption("plot", expr.String(plot))
}

func (curvesBuilder *implCurvesBuilder) PlotExpr(plot expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("plot", plot)
}

func (curvesBuilder *implCurvesBuilder) Interp(interp int) CurvesBuilder {
	return curvesBuilder.withOption("interp", expr.Int(interp))
}

func (curvesBuilder *implCurvesBuilder) InterpExpr(interp expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("interp", interp)
}

func (curvesBuilder *implCurvesBuilder) Enable(enable expr.Expr) CurvesBuilder {
	return curvesBuilder.withOption("enable", enable)
}
