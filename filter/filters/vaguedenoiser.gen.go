// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// VaguedenoiserBuilder Apply a Wavelet based Denoiser.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#vaguedenoiser
type VaguedenoiserBuilder interface {
	filter.Filter
	// Threshold set filtering strength (from 0 to DBL_MAX) (default 2).
	Threshold(threshold float32) VaguedenoiserBuilder
	// Method set filtering method (from 0 to 2) (default garrote).
	Method(method int) VaguedenoiserBuilder
	// Nsteps set number of steps (from 1 to 32) (default 6).
	Nsteps(nsteps int) VaguedenoiserBuilder
	// Percent set percent of full denoising (from 0 to 100) (default 85).
	Percent(percent float32) VaguedenoiserBuilder
	// Planes set planes to filter (from 0 to 15) (default 15).
	Planes(planes int) VaguedenoiserBuilder
	// Type set threshold type (from 0 to 1) (default universal).
	Type(typ int) VaguedenoiserBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) VaguedenoiserBuilder
}

// Vaguedenoiser creates a new VaguedenoiserBuilder to configure the "vaguedenoiser" filter.
func Vaguedenoiser(opts ...filter.Option) VaguedenoiserBuilder {
	return &implVaguedenoiserBuilder{
		f: filter.New("vaguedenoiser", 1, opts...),
	}
}

type implVaguedenoiserBuilder struct {
	f filter.Filter
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) String() string {
	return vaguedenoiserBuilder.f.String()
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Outputs() int {
	return vaguedenoiserBuilder.f.Outputs()
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) With(key string, value expr.Expr) filter.Filter {
	return vaguedenoiserBuilder.withOption(key, value)
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) withOption(key string, value expr.Expr) VaguedenoiserBuilder {
	bCopy := *vaguedenoiserBuilder
	bCopy.f = vaguedenoiserBuilder.f.With(key, value)
	return &bCopy
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Threshold(threshold float32) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("threshold", expr.Float(threshold))
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Method(method int) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("method", expr.Int(method))
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Nsteps(nsteps int) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("nsteps", expr.Int(nsteps))
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Percent(percent float32) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("percent", expr.Float(percent))
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Planes(planes int) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("planes", expr.Int(planes))
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Type(typ int) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("type", expr.Int(typ))
}

func (vaguedenoiserBuilder *implVaguedenoiserBuilder) Enable(enable expr.Expr) VaguedenoiserBuilder {
	return vaguedenoiserBuilder.withOption("enable", enable)
}
