// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AdecorrelateBuilder Apply decorrelation to input audio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#adecorrelate
type AdecorrelateBuilder interface {
	filter.Filter
	// Stages set filtering stages (from 1 to 16) (default 6).
	Stages(stages int) AdecorrelateBuilder
	// Seed set random seed (from -1 to UINT32_MAX) (default -1).
	Seed(seed int64) AdecorrelateBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AdecorrelateBuilder
}

// Adecorrelate creates a new AdecorrelateBuilder to configure the "adecorrelate" filter.
func Adecorrelate(opts ...filter.Option) AdecorrelateBuilder {
	return &implAdecorrelateBuilder{
		f: filter.New("adecorrelate", 1, opts...),
	}
}

type implAdecorrelateBuilder struct {
	f filter.Filter
}

func (adecorrelateBuilder *implAdecorrelateBuilder) String() string {
	return adecorrelateBuilder.f.String()
}

func (adecorrelateBuilder *implAdecorrelateBuilder) Outputs() int {
	return adecorrelateBuilder.f.Outputs()
}

func (adecorrelateBuilder *implAdecorrelateBuilder) With(key string, value expr.Expr) filter.Filter {
	return adecorrelateBuilder.withOption(key, value)
}

func (adecorrelateBuilder *implAdecorrelateBuilder) withOption(key string, value expr.Expr) AdecorrelateBuilder {
	bCopy := *adecorrelateBuilder
	bCopy.f = adecorrelateBuilder.f.With(key, value)
	return &bCopy
}

func (adecorrelateBuilder *implAdecorrelateBuilder) Stages(stages int) AdecorrelateBuilder {
	return adecorrelateBuilder.withOption("stages", expr.Int(stages))
}

func (adecorrelateBuilder *implAdecorrelateBuilder) Seed(seed int64) AdecorrelateBuilder {
	return adecorrelateBuilder.withOption("seed", expr.Int64(seed))
}

func (adecorrelateBuilder *implAdecorrelateBuilder) Enable(enable expr.Expr) AdecorrelateBuilder {
	return adecorrelateBuilder.withOption("enable", enable)
}
