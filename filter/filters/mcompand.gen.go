// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// McompandBuilder Multiband Compress or expand audio dynamic range.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#mcompand
type McompandBuilder interface {
	filter.Filter
	// Args set parameters for each band (default "0.005,0.1 6 -47/-40,-34/-34,-17/-33 100 | 0.003,0.05 6 -47/-40,-34/-34,-17/-33 400 | 0.000625,0.0125 6 -47/-40,-34/-34,-15/-33 1600 | 0.0001,0.025 6 -47/-40,-34/-34,-31/-31,-0/-30 6400 | 0,0.025 6 -38/-31,-28/-28,-0/-25 22000").
	Args(args string) McompandBuilder
}

// Mcompand creates a new McompandBuilder to configure the "mcompand" filter.
func Mcompand(opts ...filter.Option) McompandBuilder {
	return &implMcompandBuilder{
		f: filter.New("mcompand", 1, opts...),
	}
}

type implMcompandBuilder struct {
	f filter.Filter
}

func (mcompandBuilder *implMcompandBuilder) String() string {
	return mcompandBuilder.f.String()
}

func (mcompandBuilder *implMcompandBuilder) Outputs() int {
	return mcompandBuilder.f.Outputs()
}

func (mcompandBuilder *implMcompandBuilder) With(key string, value expr.Expr) filter.Filter {
	return mcompandBuilder.withOption(key, value)
}

func (mcompandBuilder *implMcompandBuilder) withOption(key string, value expr.Expr) McompandBuilder {
	bCopy := *mcompandBuilder
	bCopy.f = mcompandBuilder.f.With(key, value)
	return &bCopy
}

func (mcompandBuilder *implMcompandBuilder) Args(args string) McompandBuilder {
	return mcompandBuilder.withOption("args", expr.String(args))
}
