// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ConvolveBuilder Convolve first video stream with second video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#convolve
type ConvolveBuilder interface {
	filter.Filter
	// Planes set planes to convolve (from 0 to 15) (default 7).
	Planes(planes int) ConvolveBuilder
	// Impulse when to process impulses (from 0 to 1) (default all).
	Impulse(impulse int) ConvolveBuilder
	// Noise set noise (from 0 to 1) (default 1e-07).
	Noise(noise float32) ConvolveBuilder
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) ConvolveBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) ConvolveBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) ConvolveBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) ConvolveBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ConvolveBuilder
}

// Convolve creates a new ConvolveBuilder to configure the "convolve" filter.
func Convolve(opts ...filter.Option) ConvolveBuilder {
	return &implConvolveBuilder{
		f: filter.New("convolve", 1, opts...),
	}
}

type implConvolveBuilder struct {
	f filter.Filter
}

func (convolveBuilder *implConvolveBuilder) String() string {
	return convolveBuilder.f.String()
}

func (convolveBuilder *implConvolveBuilder) Outputs() int {
	return convolveBuilder.f.Outputs()
}

func (convolveBuilder *implConvolveBuilder) With(key string, value expr.Expr) filter.Filter {
	return convolveBuilder.withOption(key, value)
}

func (convolveBuilder *implConvolveBuilder) withOption(key string, value expr.Expr) ConvolveBuilder {
	bCopy := *convolveBuilder
	bCopy.f = convolveBuilder.f.With(key, value)
	return &bCopy
}

func (convolveBuilder *implConvolveBuilder) Planes(planes int) ConvolveBuilder {
	return convolveBuilder.withOption("planes", expr.Int(planes))
}

func (convolveBuilder *implConvolveBuilder) Impulse(impulse int) ConvolveBuilder {
	return convolveBuilder.withOption("impulse", expr.Int(impulse))
}

func (convolveBuilder *implConvolveBuilder) Noise(noise float32) ConvolveBuilder {
	return convolveBuilder.withOption("noise", expr.Float(noise))
}

func (convolveBuilder *implConvolveBuilder) EofAction(eofAction int) ConvolveBuilder {
	return convolveBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (convolveBuilder *implConvolveBuilder) Shortest(shortest bool) ConvolveBuilder {
	return convolveBuilder.withOption("shortest", expr.Bool(shortest))
}

func (convolveBuilder *implConvolveBuilder) Repeatlast(repeatlast bool) ConvolveBuilder {
	return convolveBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (convolveBuilder *implConvolveBuilder) TsSyncMode(tsSyncMode int) ConvolveBuilder {
	return convolveBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (convolveBuilder *implConvolveBuilder) Enable(enable expr.Expr) ConvolveBuilder {
	return convolveBuilder.withOption("enable", enable)
}