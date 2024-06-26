// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PaletteuseBuilder Use a palette to downsample an input video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#paletteuse
type PaletteuseBuilder interface {
	filter.Filter
	// Dither select dithering mode (from 0 to 8) (default sierra2_4a).
	Dither(dither int) PaletteuseBuilder
	// BayerScale set scale for bayer dithering (from 0 to 5) (default 2).
	BayerScale(bayerScale int) PaletteuseBuilder
	// DiffMode set frame difference mode (from 0 to 1) (default 0).
	DiffMode(diffMode int) PaletteuseBuilder
	// New take new palette for each output frame (default false).
	New(new bool) PaletteuseBuilder
	// AlphaThreshold set the alpha threshold for transparency (from 0 to 255) (default 128).
	AlphaThreshold(alphaThreshold int) PaletteuseBuilder
	// DebugKdtree save Graphviz graph of the kdtree in specified file.
	DebugKdtree(debugKdtree string) PaletteuseBuilder
}

// Paletteuse creates a new PaletteuseBuilder to configure the "paletteuse" filter.
func Paletteuse(opts ...filter.Option) PaletteuseBuilder {
	return &implPaletteuseBuilder{
		f: filter.New("paletteuse", 1, opts...),
	}
}

type implPaletteuseBuilder struct {
	f filter.Filter
}

func (paletteuseBuilder *implPaletteuseBuilder) String() string {
	return paletteuseBuilder.f.String()
}

func (paletteuseBuilder *implPaletteuseBuilder) Outputs() int {
	return paletteuseBuilder.f.Outputs()
}

func (paletteuseBuilder *implPaletteuseBuilder) With(key string, value expr.Expr) filter.Filter {
	return paletteuseBuilder.withOption(key, value)
}

func (paletteuseBuilder *implPaletteuseBuilder) withOption(key string, value expr.Expr) PaletteuseBuilder {
	bCopy := *paletteuseBuilder
	bCopy.f = paletteuseBuilder.f.With(key, value)
	return &bCopy
}

func (paletteuseBuilder *implPaletteuseBuilder) Dither(dither int) PaletteuseBuilder {
	return paletteuseBuilder.withOption("dither", expr.Int(dither))
}

func (paletteuseBuilder *implPaletteuseBuilder) BayerScale(bayerScale int) PaletteuseBuilder {
	return paletteuseBuilder.withOption("bayer_scale", expr.Int(bayerScale))
}

func (paletteuseBuilder *implPaletteuseBuilder) DiffMode(diffMode int) PaletteuseBuilder {
	return paletteuseBuilder.withOption("diff_mode", expr.Int(diffMode))
}

func (paletteuseBuilder *implPaletteuseBuilder) New(new bool) PaletteuseBuilder {
	return paletteuseBuilder.withOption("new", expr.Bool(new))
}

func (paletteuseBuilder *implPaletteuseBuilder) AlphaThreshold(alphaThreshold int) PaletteuseBuilder {
	return paletteuseBuilder.withOption("alpha_threshold", expr.Int(alphaThreshold))
}

func (paletteuseBuilder *implPaletteuseBuilder) DebugKdtree(debugKdtree string) PaletteuseBuilder {
	return paletteuseBuilder.withOption("debug_kdtree", expr.String(debugKdtree))
}
