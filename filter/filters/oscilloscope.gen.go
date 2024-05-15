// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// OscilloscopeBuilder 2D Video Oscilloscope.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#oscilloscope
type OscilloscopeBuilder interface {
	filter.Filter
	// X set scope x position (from 0 to 1) (default 0.5).
	X(x float32) OscilloscopeBuilder
	// XExpr set scope x position (from 0 to 1) (default 0.5).
	XExpr(x expr.Expr) OscilloscopeBuilder
	// Y set scope y position (from 0 to 1) (default 0.5).
	Y(y float32) OscilloscopeBuilder
	// YExpr set scope y position (from 0 to 1) (default 0.5).
	YExpr(y expr.Expr) OscilloscopeBuilder
	// S set scope size (from 0 to 1) (default 0.8).
	S(s float32) OscilloscopeBuilder
	// SExpr set scope size (from 0 to 1) (default 0.8).
	SExpr(s expr.Expr) OscilloscopeBuilder
	// T set scope tilt (from 0 to 1) (default 0.5).
	T(t float32) OscilloscopeBuilder
	// TExpr set scope tilt (from 0 to 1) (default 0.5).
	TExpr(t expr.Expr) OscilloscopeBuilder
	// O set trace opacity (from 0 to 1) (default 0.8).
	O(o float32) OscilloscopeBuilder
	// OExpr set trace opacity (from 0 to 1) (default 0.8).
	OExpr(o expr.Expr) OscilloscopeBuilder
	// Tx set trace x position (from 0 to 1) (default 0.5).
	Tx(tx float32) OscilloscopeBuilder
	// TxExpr set trace x position (from 0 to 1) (default 0.5).
	TxExpr(tx expr.Expr) OscilloscopeBuilder
	// Ty set trace y position (from 0 to 1) (default 0.9).
	Ty(ty float32) OscilloscopeBuilder
	// TyExpr set trace y position (from 0 to 1) (default 0.9).
	TyExpr(ty expr.Expr) OscilloscopeBuilder
	// Tw set trace width (from 0.1 to 1) (default 0.8).
	Tw(tw float32) OscilloscopeBuilder
	// TwExpr set trace width (from 0.1 to 1) (default 0.8).
	TwExpr(tw expr.Expr) OscilloscopeBuilder
	// Th set trace height (from 0.1 to 1) (default 0.3).
	Th(th float32) OscilloscopeBuilder
	// ThExpr set trace height (from 0.1 to 1) (default 0.3).
	ThExpr(th expr.Expr) OscilloscopeBuilder
	// C set components to trace (from 0 to 15) (default 7).
	C(c int) OscilloscopeBuilder
	// CExpr set components to trace (from 0 to 15) (default 7).
	CExpr(c expr.Expr) OscilloscopeBuilder
	// G draw trace grid (default true).
	G(g bool) OscilloscopeBuilder
	// GExpr draw trace grid (default true).
	GExpr(g expr.Expr) OscilloscopeBuilder
	// St draw statistics (default true).
	St(st bool) OscilloscopeBuilder
	// StExpr draw statistics (default true).
	StExpr(st expr.Expr) OscilloscopeBuilder
	// Sc draw scope (default true).
	Sc(sc bool) OscilloscopeBuilder
	// ScExpr draw scope (default true).
	ScExpr(sc expr.Expr) OscilloscopeBuilder
}

// Oscilloscope creates a new OscilloscopeBuilder to configure the "oscilloscope" filter.
func Oscilloscope(opts ...filter.Option) OscilloscopeBuilder {
	return &implOscilloscopeBuilder{
		f: filter.New("oscilloscope", 1, opts...),
	}
}

type implOscilloscopeBuilder struct {
	f filter.Filter
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) String() string {
	return oscilloscopeBuilder.f.String()
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Outputs() int {
	return oscilloscopeBuilder.f.Outputs()
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) With(key string, value expr.Expr) filter.Filter {
	return oscilloscopeBuilder.withOption(key, value)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) withOption(key string, value expr.Expr) OscilloscopeBuilder {
	bCopy := *oscilloscopeBuilder
	bCopy.f = oscilloscopeBuilder.f.With(key, value)
	return &bCopy
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) X(x float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("x", expr.Float(x))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) XExpr(x expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("x", x)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Y(y float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("y", expr.Float(y))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) YExpr(y expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("y", y)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) S(s float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("s", expr.Float(s))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) SExpr(s expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("s", s)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) T(t float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("t", expr.Float(t))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) TExpr(t expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("t", t)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) O(o float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("o", expr.Float(o))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) OExpr(o expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("o", o)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Tx(tx float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("tx", expr.Float(tx))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) TxExpr(tx expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("tx", tx)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Ty(ty float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("ty", expr.Float(ty))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) TyExpr(ty expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("ty", ty)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Tw(tw float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("tw", expr.Float(tw))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) TwExpr(tw expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("tw", tw)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Th(th float32) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("th", expr.Float(th))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) ThExpr(th expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("th", th)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) C(c int) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("c", expr.Int(c))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) CExpr(c expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("c", c)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) G(g bool) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("g", expr.Bool(g))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) GExpr(g expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("g", g)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) St(st bool) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("st", expr.Bool(st))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) StExpr(st expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("st", st)
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) Sc(sc bool) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("sc", expr.Bool(sc))
}

func (oscilloscopeBuilder *implOscilloscopeBuilder) ScExpr(sc expr.Expr) OscilloscopeBuilder {
	return oscilloscopeBuilder.withOption("sc", sc)
}
