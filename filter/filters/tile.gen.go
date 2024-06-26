// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// TileBuilder Tile several successive frames together.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#tile
type TileBuilder interface {
	filter.Filter
	// Layout set grid size (default "6x5").
	Layout(layout expr.Size) TileBuilder
	// NbFrames set maximum number of frame to render (from 0 to INT_MAX) (default 0).
	NbFrames(nbFrames int) TileBuilder
	// Margin set outer border margin in pixels (from 0 to 1024) (default 0).
	Margin(margin int) TileBuilder
	// Padding set inner border thickness in pixels (from 0 to 1024) (default 0).
	Padding(padding int) TileBuilder
	// Color set the color of the unused area (default "black").
	Color(color expr.Color) TileBuilder
	// Overlap set how many frames to overlap for each render (from 0 to INT_MAX) (default 0).
	Overlap(overlap int) TileBuilder
	// InitPadding set how many frames to initially pad (from 0 to INT_MAX) (default 0).
	InitPadding(initPadding int) TileBuilder
}

// Tile creates a new TileBuilder to configure the "tile" filter.
func Tile(opts ...filter.Option) TileBuilder {
	return &implTileBuilder{
		f: filter.New("tile", 1, opts...),
	}
}

type implTileBuilder struct {
	f filter.Filter
}

func (tileBuilder *implTileBuilder) String() string {
	return tileBuilder.f.String()
}

func (tileBuilder *implTileBuilder) Outputs() int {
	return tileBuilder.f.Outputs()
}

func (tileBuilder *implTileBuilder) With(key string, value expr.Expr) filter.Filter {
	return tileBuilder.withOption(key, value)
}

func (tileBuilder *implTileBuilder) withOption(key string, value expr.Expr) TileBuilder {
	bCopy := *tileBuilder
	bCopy.f = tileBuilder.f.With(key, value)
	return &bCopy
}

func (tileBuilder *implTileBuilder) Layout(layout expr.Size) TileBuilder {
	return tileBuilder.withOption("layout", layout)
}

func (tileBuilder *implTileBuilder) NbFrames(nbFrames int) TileBuilder {
	return tileBuilder.withOption("nb_frames", expr.Int(nbFrames))
}

func (tileBuilder *implTileBuilder) Margin(margin int) TileBuilder {
	return tileBuilder.withOption("margin", expr.Int(margin))
}

func (tileBuilder *implTileBuilder) Padding(padding int) TileBuilder {
	return tileBuilder.withOption("padding", expr.Int(padding))
}

func (tileBuilder *implTileBuilder) Color(color expr.Color) TileBuilder {
	return tileBuilder.withOption("color", color)
}

func (tileBuilder *implTileBuilder) Overlap(overlap int) TileBuilder {
	return tileBuilder.withOption("overlap", expr.Int(overlap))
}

func (tileBuilder *implTileBuilder) InitPadding(initPadding int) TileBuilder {
	return tileBuilder.withOption("init_padding", expr.Int(initPadding))
}
