// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PixelizeBuilder Pixelize video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#pixelize
type PixelizeBuilder interface {
	filter.Filter
	// Width set block width (from 1 to 1024) (default 16).
	Width(width int) PixelizeBuilder
	// WidthExpr set block width (from 1 to 1024) (default 16).
	WidthExpr(width expr.Expr) PixelizeBuilder
	// W set block width (from 1 to 1024) (default 16).
	W(w int) PixelizeBuilder
	// WExpr set block width (from 1 to 1024) (default 16).
	WExpr(w expr.Expr) PixelizeBuilder
	// Height set block height (from 1 to 1024) (default 16).
	Height(height int) PixelizeBuilder
	// HeightExpr set block height (from 1 to 1024) (default 16).
	HeightExpr(height expr.Expr) PixelizeBuilder
	// H set block height (from 1 to 1024) (default 16).
	H(h int) PixelizeBuilder
	// HExpr set block height (from 1 to 1024) (default 16).
	HExpr(h expr.Expr) PixelizeBuilder
	// Mode set the pixelize mode (from 0 to 2) (default avg).
	Mode(mode int) PixelizeBuilder
	// ModeExpr set the pixelize mode (from 0 to 2) (default avg).
	ModeExpr(mode expr.Expr) PixelizeBuilder
	// M set the pixelize mode (from 0 to 2) (default avg).
	M(m int) PixelizeBuilder
	// MExpr set the pixelize mode (from 0 to 2) (default avg).
	MExpr(m expr.Expr) PixelizeBuilder
	// Planes set what planes to filter (default F).
	Planes(planes ...string) PixelizeBuilder
	// PlanesExpr set what planes to filter (default F).
	PlanesExpr(planes expr.Expr) PixelizeBuilder
	// P set what planes to filter (default F).
	P(p ...string) PixelizeBuilder
	// PExpr set what planes to filter (default F).
	PExpr(p expr.Expr) PixelizeBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) PixelizeBuilder
}

// Pixelize creates a new PixelizeBuilder to configure the "pixelize" filter.
func Pixelize(opts ...filter.Option) PixelizeBuilder {
	return &implPixelizeBuilder{
		f: filter.New("pixelize", 1, opts...),
	}
}

type implPixelizeBuilder struct {
	f filter.Filter
}

func (pixelizeBuilder *implPixelizeBuilder) String() string {
	return pixelizeBuilder.f.String()
}

func (pixelizeBuilder *implPixelizeBuilder) Outputs() int {
	return pixelizeBuilder.f.Outputs()
}

func (pixelizeBuilder *implPixelizeBuilder) With(key string, value expr.Expr) filter.Filter {
	return pixelizeBuilder.withOption(key, value)
}

func (pixelizeBuilder *implPixelizeBuilder) withOption(key string, value expr.Expr) PixelizeBuilder {
	bCopy := *pixelizeBuilder
	bCopy.f = pixelizeBuilder.f.With(key, value)
	return &bCopy
}

func (pixelizeBuilder *implPixelizeBuilder) Width(width int) PixelizeBuilder {
	return pixelizeBuilder.withOption("width", expr.Int(width))
}

func (pixelizeBuilder *implPixelizeBuilder) WidthExpr(width expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("width", width)
}

func (pixelizeBuilder *implPixelizeBuilder) W(w int) PixelizeBuilder {
	return pixelizeBuilder.withOption("w", expr.Int(w))
}

func (pixelizeBuilder *implPixelizeBuilder) WExpr(w expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("w", w)
}

func (pixelizeBuilder *implPixelizeBuilder) Height(height int) PixelizeBuilder {
	return pixelizeBuilder.withOption("height", expr.Int(height))
}

func (pixelizeBuilder *implPixelizeBuilder) HeightExpr(height expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("height", height)
}

func (pixelizeBuilder *implPixelizeBuilder) H(h int) PixelizeBuilder {
	return pixelizeBuilder.withOption("h", expr.Int(h))
}

func (pixelizeBuilder *implPixelizeBuilder) HExpr(h expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("h", h)
}

func (pixelizeBuilder *implPixelizeBuilder) Mode(mode int) PixelizeBuilder {
	return pixelizeBuilder.withOption("mode", expr.Int(mode))
}

func (pixelizeBuilder *implPixelizeBuilder) ModeExpr(mode expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("mode", mode)
}

func (pixelizeBuilder *implPixelizeBuilder) M(m int) PixelizeBuilder {
	return pixelizeBuilder.withOption("m", expr.Int(m))
}

func (pixelizeBuilder *implPixelizeBuilder) MExpr(m expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("m", m)
}

func (pixelizeBuilder *implPixelizeBuilder) Planes(planes ...string) PixelizeBuilder {
	return pixelizeBuilder.withOption("planes", expr.Flags(planes))
}

func (pixelizeBuilder *implPixelizeBuilder) PlanesExpr(planes expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("planes", planes)
}

func (pixelizeBuilder *implPixelizeBuilder) P(p ...string) PixelizeBuilder {
	return pixelizeBuilder.withOption("p", expr.Flags(p))
}

func (pixelizeBuilder *implPixelizeBuilder) PExpr(p expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("p", p)
}

func (pixelizeBuilder *implPixelizeBuilder) Enable(enable expr.Expr) PixelizeBuilder {
	return pixelizeBuilder.withOption("enable", enable)
}
