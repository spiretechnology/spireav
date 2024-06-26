// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DeconvolveBuilder Deconvolve first video stream with second video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#deconvolve
type DeconvolveBuilder interface {
	filter.Filter
	// Planes set planes to deconvolve (from 0 to 15) (default 7).
	Planes(planes int) DeconvolveBuilder
	// Impulse when to process impulses (from 0 to 1) (default all).
	Impulse(impulse int) DeconvolveBuilder
	// Noise set noise (from 0 to 1) (default 1e-07).
	Noise(noise float32) DeconvolveBuilder
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) DeconvolveBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) DeconvolveBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) DeconvolveBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) DeconvolveBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) DeconvolveBuilder
}

// Deconvolve creates a new DeconvolveBuilder to configure the "deconvolve" filter.
func Deconvolve(opts ...filter.Option) DeconvolveBuilder {
	return &implDeconvolveBuilder{
		f: filter.New("deconvolve", 1, opts...),
	}
}

type implDeconvolveBuilder struct {
	f filter.Filter
}

func (deconvolveBuilder *implDeconvolveBuilder) String() string {
	return deconvolveBuilder.f.String()
}

func (deconvolveBuilder *implDeconvolveBuilder) Outputs() int {
	return deconvolveBuilder.f.Outputs()
}

func (deconvolveBuilder *implDeconvolveBuilder) With(key string, value expr.Expr) filter.Filter {
	return deconvolveBuilder.withOption(key, value)
}

func (deconvolveBuilder *implDeconvolveBuilder) withOption(key string, value expr.Expr) DeconvolveBuilder {
	bCopy := *deconvolveBuilder
	bCopy.f = deconvolveBuilder.f.With(key, value)
	return &bCopy
}

func (deconvolveBuilder *implDeconvolveBuilder) Planes(planes int) DeconvolveBuilder {
	return deconvolveBuilder.withOption("planes", expr.Int(planes))
}

func (deconvolveBuilder *implDeconvolveBuilder) Impulse(impulse int) DeconvolveBuilder {
	return deconvolveBuilder.withOption("impulse", expr.Int(impulse))
}

func (deconvolveBuilder *implDeconvolveBuilder) Noise(noise float32) DeconvolveBuilder {
	return deconvolveBuilder.withOption("noise", expr.Float(noise))
}

func (deconvolveBuilder *implDeconvolveBuilder) EofAction(eofAction int) DeconvolveBuilder {
	return deconvolveBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (deconvolveBuilder *implDeconvolveBuilder) Shortest(shortest bool) DeconvolveBuilder {
	return deconvolveBuilder.withOption("shortest", expr.Bool(shortest))
}

func (deconvolveBuilder *implDeconvolveBuilder) Repeatlast(repeatlast bool) DeconvolveBuilder {
	return deconvolveBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (deconvolveBuilder *implDeconvolveBuilder) TsSyncMode(tsSyncMode int) DeconvolveBuilder {
	return deconvolveBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (deconvolveBuilder *implDeconvolveBuilder) Enable(enable expr.Expr) DeconvolveBuilder {
	return deconvolveBuilder.withOption("enable", enable)
}
