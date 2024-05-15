// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SetsarBuilder Set the pixel sample aspect ratio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#setsar
type SetsarBuilder interface {
	filter.Filter
	// Sar set sample (pixel) aspect ratio (default "0").
	Sar(sar string) SetsarBuilder
	// Ratio set sample (pixel) aspect ratio (default "0").
	Ratio(ratio string) SetsarBuilder
	// R set sample (pixel) aspect ratio (default "0").
	R(r string) SetsarBuilder
	// Max set max value for nominator or denominator in the ratio (from 1 to INT_MAX) (default 100).
	Max(max int) SetsarBuilder
}

// Setsar creates a new SetsarBuilder to configure the "setsar" filter.
func Setsar(opts ...filter.Option) SetsarBuilder {
	return &implSetsarBuilder{
		f: filter.New("setsar", 1, opts...),
	}
}

type implSetsarBuilder struct {
	f filter.Filter
}

func (setsarBuilder *implSetsarBuilder) String() string {
	return setsarBuilder.f.String()
}

func (setsarBuilder *implSetsarBuilder) Outputs() int {
	return setsarBuilder.f.Outputs()
}

func (setsarBuilder *implSetsarBuilder) With(key string, value expr.Expr) filter.Filter {
	return setsarBuilder.withOption(key, value)
}

func (setsarBuilder *implSetsarBuilder) withOption(key string, value expr.Expr) SetsarBuilder {
	bCopy := *setsarBuilder
	bCopy.f = setsarBuilder.f.With(key, value)
	return &bCopy
}

func (setsarBuilder *implSetsarBuilder) Sar(sar string) SetsarBuilder {
	return setsarBuilder.withOption("sar", expr.String(sar))
}

func (setsarBuilder *implSetsarBuilder) Ratio(ratio string) SetsarBuilder {
	return setsarBuilder.withOption("ratio", expr.String(ratio))
}

func (setsarBuilder *implSetsarBuilder) R(r string) SetsarBuilder {
	return setsarBuilder.withOption("r", expr.String(r))
}

func (setsarBuilder *implSetsarBuilder) Max(max int) SetsarBuilder {
	return setsarBuilder.withOption("max", expr.Int(max))
}
