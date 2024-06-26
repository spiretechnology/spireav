// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PalettegenBuilder Find the optimal palette for a given stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#palettegen
type PalettegenBuilder interface {
	filter.Filter
	// MaxColors set the maximum number of colors to use in the palette (from 2 to 256) (default 256).
	MaxColors(maxColors int) PalettegenBuilder
	// ReserveTransparent reserve a palette entry for transparency (default true).
	ReserveTransparent(reserveTransparent bool) PalettegenBuilder
	// TransparencyColor set a background color for transparency (default "lime").
	TransparencyColor(transparencyColor expr.Color) PalettegenBuilder
	// StatsMode set statistics mode (from 0 to 2) (default full).
	StatsMode(statsMode int) PalettegenBuilder
}

// Palettegen creates a new PalettegenBuilder to configure the "palettegen" filter.
func Palettegen(opts ...filter.Option) PalettegenBuilder {
	return &implPalettegenBuilder{
		f: filter.New("palettegen", 1, opts...),
	}
}

type implPalettegenBuilder struct {
	f filter.Filter
}

func (palettegenBuilder *implPalettegenBuilder) String() string {
	return palettegenBuilder.f.String()
}

func (palettegenBuilder *implPalettegenBuilder) Outputs() int {
	return palettegenBuilder.f.Outputs()
}

func (palettegenBuilder *implPalettegenBuilder) With(key string, value expr.Expr) filter.Filter {
	return palettegenBuilder.withOption(key, value)
}

func (palettegenBuilder *implPalettegenBuilder) withOption(key string, value expr.Expr) PalettegenBuilder {
	bCopy := *palettegenBuilder
	bCopy.f = palettegenBuilder.f.With(key, value)
	return &bCopy
}

func (palettegenBuilder *implPalettegenBuilder) MaxColors(maxColors int) PalettegenBuilder {
	return palettegenBuilder.withOption("max_colors", expr.Int(maxColors))
}

func (palettegenBuilder *implPalettegenBuilder) ReserveTransparent(reserveTransparent bool) PalettegenBuilder {
	return palettegenBuilder.withOption("reserve_transparent", expr.Bool(reserveTransparent))
}

func (palettegenBuilder *implPalettegenBuilder) TransparencyColor(transparencyColor expr.Color) PalettegenBuilder {
	return palettegenBuilder.withOption("transparency_color", transparencyColor)
}

func (palettegenBuilder *implPalettegenBuilder) StatsMode(statsMode int) PalettegenBuilder {
	return palettegenBuilder.withOption("stats_mode", expr.Int(statsMode))
}
