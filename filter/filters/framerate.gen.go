// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FramerateBuilder Upsamples or downsamples progressive source between specified frame rates.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#framerate
type FramerateBuilder interface {
	filter.Filter
	// Fps required output frames per second rate (default "50").
	Fps(fps expr.Rational) FramerateBuilder
	// InterpStart point to start linear interpolation (from 0 to 255) (default 15).
	InterpStart(interpStart int) FramerateBuilder
	// InterpEnd point to end linear interpolation (from 0 to 255) (default 240).
	InterpEnd(interpEnd int) FramerateBuilder
	// Scene scene change level (from 0 to 100) (default 8.2).
	Scene(scene float64) FramerateBuilder
	// Flags set flags (default scene_change_detect+scd).
	Flags(flags ...string) FramerateBuilder
}

// Framerate creates a new FramerateBuilder to configure the "framerate" filter.
func Framerate(opts ...filter.Option) FramerateBuilder {
	return &implFramerateBuilder{
		f: filter.New("framerate", 1, opts...),
	}
}

type implFramerateBuilder struct {
	f filter.Filter
}

func (framerateBuilder *implFramerateBuilder) String() string {
	return framerateBuilder.f.String()
}

func (framerateBuilder *implFramerateBuilder) Outputs() int {
	return framerateBuilder.f.Outputs()
}

func (framerateBuilder *implFramerateBuilder) With(key string, value expr.Expr) filter.Filter {
	return framerateBuilder.withOption(key, value)
}

func (framerateBuilder *implFramerateBuilder) withOption(key string, value expr.Expr) FramerateBuilder {
	bCopy := *framerateBuilder
	bCopy.f = framerateBuilder.f.With(key, value)
	return &bCopy
}

func (framerateBuilder *implFramerateBuilder) Fps(fps expr.Rational) FramerateBuilder {
	return framerateBuilder.withOption("fps", fps)
}

func (framerateBuilder *implFramerateBuilder) InterpStart(interpStart int) FramerateBuilder {
	return framerateBuilder.withOption("interp_start", expr.Int(interpStart))
}

func (framerateBuilder *implFramerateBuilder) InterpEnd(interpEnd int) FramerateBuilder {
	return framerateBuilder.withOption("interp_end", expr.Int(interpEnd))
}

func (framerateBuilder *implFramerateBuilder) Scene(scene float64) FramerateBuilder {
	return framerateBuilder.withOption("scene", expr.Double(scene))
}

func (framerateBuilder *implFramerateBuilder) Flags(flags ...string) FramerateBuilder {
	return framerateBuilder.withOption("flags", expr.Flags(flags))
}
