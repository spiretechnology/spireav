// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// GradientsBuilder Draw a gradients.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#gradients
type GradientsBuilder interface {
	filter.Filter
	// Size set frame size (default "640x480").
	Size(size expr.Size) GradientsBuilder
	// S set frame size (default "640x480").
	S(s expr.Size) GradientsBuilder
	// Rate set frame rate (default "25").
	Rate(rate expr.Rational) GradientsBuilder
	// R set frame rate (default "25").
	R(r expr.Rational) GradientsBuilder
	// C0 set 1st color (default "random").
	C0(c0 expr.Color) GradientsBuilder
	// C1 set 2nd color (default "random").
	C1(c1 expr.Color) GradientsBuilder
	// C2 set 3rd color (default "random").
	C2(c2 expr.Color) GradientsBuilder
	// C3 set 4th color (default "random").
	C3(c3 expr.Color) GradientsBuilder
	// C4 set 5th color (default "random").
	C4(c4 expr.Color) GradientsBuilder
	// C5 set 6th color (default "random").
	C5(c5 expr.Color) GradientsBuilder
	// C6 set 7th color (default "random").
	C6(c6 expr.Color) GradientsBuilder
	// C7 set 8th color (default "random").
	C7(c7 expr.Color) GradientsBuilder
	// X0 set gradient line source x0 (from -1 to INT_MAX) (default -1).
	X0(x0 int) GradientsBuilder
	// Y0 set gradient line source y0 (from -1 to INT_MAX) (default -1).
	Y0(y0 int) GradientsBuilder
	// X1 set gradient line destination x1 (from -1 to INT_MAX) (default -1).
	X1(x1 int) GradientsBuilder
	// Y1 set gradient line destination y1 (from -1 to INT_MAX) (default -1).
	Y1(y1 int) GradientsBuilder
	// NbColors set the number of colors (from 2 to 8) (default 2).
	NbColors(nbColors int) GradientsBuilder
	// N set the number of colors (from 2 to 8) (default 2).
	N(n int) GradientsBuilder
	// Seed set the seed (from -1 to UINT32_MAX) (default -1).
	Seed(seed int64) GradientsBuilder
	// Duration set video duration (default -0.000001).
	Duration(duration time.Duration) GradientsBuilder
	// D set video duration (default -0.000001).
	D(d time.Duration) GradientsBuilder
	// Speed set gradients rotation speed (from 0 to 1) (default 0.01).
	Speed(speed float32) GradientsBuilder
	// SpeedExpr set gradients rotation speed (from 0 to 1) (default 0.01).
	SpeedExpr(speed expr.Expr) GradientsBuilder
	// Type set gradient type (from 0 to 4) (default linear).
	Type(typ int) GradientsBuilder
	// TypeExpr set gradient type (from 0 to 4) (default linear).
	TypeExpr(typ expr.Expr) GradientsBuilder
	// T set gradient type (from 0 to 4) (default linear).
	T(t int) GradientsBuilder
	// TExpr set gradient type (from 0 to 4) (default linear).
	TExpr(t expr.Expr) GradientsBuilder
}

// Gradients creates a new GradientsBuilder to configure the "gradients" filter.
func Gradients(opts ...filter.Option) GradientsBuilder {
	return &implGradientsBuilder{
		f: filter.New("gradients", 1, opts...),
	}
}

type implGradientsBuilder struct {
	f filter.Filter
}

func (gradientsBuilder *implGradientsBuilder) String() string {
	return gradientsBuilder.f.String()
}

func (gradientsBuilder *implGradientsBuilder) Outputs() int {
	return gradientsBuilder.f.Outputs()
}

func (gradientsBuilder *implGradientsBuilder) With(key string, value expr.Expr) filter.Filter {
	return gradientsBuilder.withOption(key, value)
}

func (gradientsBuilder *implGradientsBuilder) withOption(key string, value expr.Expr) GradientsBuilder {
	bCopy := *gradientsBuilder
	bCopy.f = gradientsBuilder.f.With(key, value)
	return &bCopy
}

func (gradientsBuilder *implGradientsBuilder) Size(size expr.Size) GradientsBuilder {
	return gradientsBuilder.withOption("size", size)
}

func (gradientsBuilder *implGradientsBuilder) S(s expr.Size) GradientsBuilder {
	return gradientsBuilder.withOption("s", s)
}

func (gradientsBuilder *implGradientsBuilder) Rate(rate expr.Rational) GradientsBuilder {
	return gradientsBuilder.withOption("rate", rate)
}

func (gradientsBuilder *implGradientsBuilder) R(r expr.Rational) GradientsBuilder {
	return gradientsBuilder.withOption("r", r)
}

func (gradientsBuilder *implGradientsBuilder) C0(c0 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c0", c0)
}

func (gradientsBuilder *implGradientsBuilder) C1(c1 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c1", c1)
}

func (gradientsBuilder *implGradientsBuilder) C2(c2 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c2", c2)
}

func (gradientsBuilder *implGradientsBuilder) C3(c3 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c3", c3)
}

func (gradientsBuilder *implGradientsBuilder) C4(c4 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c4", c4)
}

func (gradientsBuilder *implGradientsBuilder) C5(c5 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c5", c5)
}

func (gradientsBuilder *implGradientsBuilder) C6(c6 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c6", c6)
}

func (gradientsBuilder *implGradientsBuilder) C7(c7 expr.Color) GradientsBuilder {
	return gradientsBuilder.withOption("c7", c7)
}

func (gradientsBuilder *implGradientsBuilder) X0(x0 int) GradientsBuilder {
	return gradientsBuilder.withOption("x0", expr.Int(x0))
}

func (gradientsBuilder *implGradientsBuilder) Y0(y0 int) GradientsBuilder {
	return gradientsBuilder.withOption("y0", expr.Int(y0))
}

func (gradientsBuilder *implGradientsBuilder) X1(x1 int) GradientsBuilder {
	return gradientsBuilder.withOption("x1", expr.Int(x1))
}

func (gradientsBuilder *implGradientsBuilder) Y1(y1 int) GradientsBuilder {
	return gradientsBuilder.withOption("y1", expr.Int(y1))
}

func (gradientsBuilder *implGradientsBuilder) NbColors(nbColors int) GradientsBuilder {
	return gradientsBuilder.withOption("nb_colors", expr.Int(nbColors))
}

func (gradientsBuilder *implGradientsBuilder) N(n int) GradientsBuilder {
	return gradientsBuilder.withOption("n", expr.Int(n))
}

func (gradientsBuilder *implGradientsBuilder) Seed(seed int64) GradientsBuilder {
	return gradientsBuilder.withOption("seed", expr.Int64(seed))
}

func (gradientsBuilder *implGradientsBuilder) Duration(duration time.Duration) GradientsBuilder {
	return gradientsBuilder.withOption("duration", expr.Duration(duration))
}

func (gradientsBuilder *implGradientsBuilder) D(d time.Duration) GradientsBuilder {
	return gradientsBuilder.withOption("d", expr.Duration(d))
}

func (gradientsBuilder *implGradientsBuilder) Speed(speed float32) GradientsBuilder {
	return gradientsBuilder.withOption("speed", expr.Float(speed))
}

func (gradientsBuilder *implGradientsBuilder) SpeedExpr(speed expr.Expr) GradientsBuilder {
	return gradientsBuilder.withOption("speed", speed)
}

func (gradientsBuilder *implGradientsBuilder) Type(typ int) GradientsBuilder {
	return gradientsBuilder.withOption("type", expr.Int(typ))
}

func (gradientsBuilder *implGradientsBuilder) TypeExpr(typ expr.Expr) GradientsBuilder {
	return gradientsBuilder.withOption("type", typ)
}

func (gradientsBuilder *implGradientsBuilder) T(t int) GradientsBuilder {
	return gradientsBuilder.withOption("t", expr.Int(t))
}

func (gradientsBuilder *implGradientsBuilder) TExpr(t expr.Expr) GradientsBuilder {
	return gradientsBuilder.withOption("t", t)
}