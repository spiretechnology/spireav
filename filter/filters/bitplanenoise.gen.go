// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// BitplanenoiseBuilder Measure bit plane noise.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#bitplanenoise
type BitplanenoiseBuilder interface {
	filter.Filter
	// Bitplane set bit plane to use for measuring noise (from 1 to 16) (default 1).
	Bitplane(bitplane int) BitplanenoiseBuilder
	// Filter show noisy pixels (default false).
	Filter(filter bool) BitplanenoiseBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) BitplanenoiseBuilder
}

// Bitplanenoise creates a new BitplanenoiseBuilder to configure the "bitplanenoise" filter.
func Bitplanenoise(opts ...filter.Option) BitplanenoiseBuilder {
	return &implBitplanenoiseBuilder{
		f: filter.New("bitplanenoise", 1, opts...),
	}
}

type implBitplanenoiseBuilder struct {
	f filter.Filter
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) String() string {
	return bitplanenoiseBuilder.f.String()
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) Outputs() int {
	return bitplanenoiseBuilder.f.Outputs()
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) With(key string, value expr.Expr) filter.Filter {
	return bitplanenoiseBuilder.withOption(key, value)
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) withOption(key string, value expr.Expr) BitplanenoiseBuilder {
	bCopy := *bitplanenoiseBuilder
	bCopy.f = bitplanenoiseBuilder.f.With(key, value)
	return &bCopy
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) Bitplane(bitplane int) BitplanenoiseBuilder {
	return bitplanenoiseBuilder.withOption("bitplane", expr.Int(bitplane))
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) Filter(filter bool) BitplanenoiseBuilder {
	return bitplanenoiseBuilder.withOption("filter", expr.Bool(filter))
}

func (bitplanenoiseBuilder *implBitplanenoiseBuilder) Enable(enable expr.Expr) BitplanenoiseBuilder {
	return bitplanenoiseBuilder.withOption("enable", enable)
}
