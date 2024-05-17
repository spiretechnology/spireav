// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PerspectiveBuilder Correct the perspective of video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#perspective
type PerspectiveBuilder interface {
	filter.Filter
	// X0 set top left x coordinate (default "0").
	X0(x0 string) PerspectiveBuilder
	// Y0 set top left y coordinate (default "0").
	Y0(y0 string) PerspectiveBuilder
	// X1 set top right x coordinate (default "W").
	X1(x1 string) PerspectiveBuilder
	// Y1 set top right y coordinate (default "0").
	Y1(y1 string) PerspectiveBuilder
	// X2 set bottom left x coordinate (default "0").
	X2(x2 string) PerspectiveBuilder
	// Y2 set bottom left y coordinate (default "H").
	Y2(y2 string) PerspectiveBuilder
	// X3 set bottom right x coordinate (default "W").
	X3(x3 string) PerspectiveBuilder
	// Y3 set bottom right y coordinate (default "H").
	Y3(y3 string) PerspectiveBuilder
	// Interpolation set interpolation (from 0 to 1) (default linear).
	Interpolation(interpolation int) PerspectiveBuilder
	// Sense specify the sense of the coordinates (from 0 to 1) (default source).
	Sense(sense int) PerspectiveBuilder
	// Eval specify when to evaluate expressions (from 0 to 1) (default init).
	Eval(eval int) PerspectiveBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) PerspectiveBuilder
}

// Perspective creates a new PerspectiveBuilder to configure the "perspective" filter.
func Perspective(opts ...filter.Option) PerspectiveBuilder {
	return &implPerspectiveBuilder{
		f: filter.New("perspective", 1, opts...),
	}
}

type implPerspectiveBuilder struct {
	f filter.Filter
}

func (perspectiveBuilder *implPerspectiveBuilder) String() string {
	return perspectiveBuilder.f.String()
}

func (perspectiveBuilder *implPerspectiveBuilder) Outputs() int {
	return perspectiveBuilder.f.Outputs()
}

func (perspectiveBuilder *implPerspectiveBuilder) With(key string, value expr.Expr) filter.Filter {
	return perspectiveBuilder.withOption(key, value)
}

func (perspectiveBuilder *implPerspectiveBuilder) withOption(key string, value expr.Expr) PerspectiveBuilder {
	bCopy := *perspectiveBuilder
	bCopy.f = perspectiveBuilder.f.With(key, value)
	return &bCopy
}

func (perspectiveBuilder *implPerspectiveBuilder) X0(x0 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("x0", expr.String(x0))
}

func (perspectiveBuilder *implPerspectiveBuilder) Y0(y0 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("y0", expr.String(y0))
}

func (perspectiveBuilder *implPerspectiveBuilder) X1(x1 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("x1", expr.String(x1))
}

func (perspectiveBuilder *implPerspectiveBuilder) Y1(y1 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("y1", expr.String(y1))
}

func (perspectiveBuilder *implPerspectiveBuilder) X2(x2 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("x2", expr.String(x2))
}

func (perspectiveBuilder *implPerspectiveBuilder) Y2(y2 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("y2", expr.String(y2))
}

func (perspectiveBuilder *implPerspectiveBuilder) X3(x3 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("x3", expr.String(x3))
}

func (perspectiveBuilder *implPerspectiveBuilder) Y3(y3 string) PerspectiveBuilder {
	return perspectiveBuilder.withOption("y3", expr.String(y3))
}

func (perspectiveBuilder *implPerspectiveBuilder) Interpolation(interpolation int) PerspectiveBuilder {
	return perspectiveBuilder.withOption("interpolation", expr.Int(interpolation))
}

func (perspectiveBuilder *implPerspectiveBuilder) Sense(sense int) PerspectiveBuilder {
	return perspectiveBuilder.withOption("sense", expr.Int(sense))
}

func (perspectiveBuilder *implPerspectiveBuilder) Eval(eval int) PerspectiveBuilder {
	return perspectiveBuilder.withOption("eval", expr.Int(eval))
}

func (perspectiveBuilder *implPerspectiveBuilder) Enable(enable expr.Expr) PerspectiveBuilder {
	return perspectiveBuilder.withOption("enable", enable)
}
