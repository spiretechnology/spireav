// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColorbalanceBuilder Adjust the color balance.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colorbalance
type ColorbalanceBuilder interface {
	filter.Filter
	// Rs set red shadows (from -1 to 1) (default 0).
	Rs(rs float32) ColorbalanceBuilder
	// RsExpr set red shadows (from -1 to 1) (default 0).
	RsExpr(rs expr.Expr) ColorbalanceBuilder
	// Gs set green shadows (from -1 to 1) (default 0).
	Gs(gs float32) ColorbalanceBuilder
	// GsExpr set green shadows (from -1 to 1) (default 0).
	GsExpr(gs expr.Expr) ColorbalanceBuilder
	// Bs set blue shadows (from -1 to 1) (default 0).
	Bs(bs float32) ColorbalanceBuilder
	// BsExpr set blue shadows (from -1 to 1) (default 0).
	BsExpr(bs expr.Expr) ColorbalanceBuilder
	// Rm set red midtones (from -1 to 1) (default 0).
	Rm(rm float32) ColorbalanceBuilder
	// RmExpr set red midtones (from -1 to 1) (default 0).
	RmExpr(rm expr.Expr) ColorbalanceBuilder
	// Gm set green midtones (from -1 to 1) (default 0).
	Gm(gm float32) ColorbalanceBuilder
	// GmExpr set green midtones (from -1 to 1) (default 0).
	GmExpr(gm expr.Expr) ColorbalanceBuilder
	// Bm set blue midtones (from -1 to 1) (default 0).
	Bm(bm float32) ColorbalanceBuilder
	// BmExpr set blue midtones (from -1 to 1) (default 0).
	BmExpr(bm expr.Expr) ColorbalanceBuilder
	// Rh set red highlights (from -1 to 1) (default 0).
	Rh(rh float32) ColorbalanceBuilder
	// RhExpr set red highlights (from -1 to 1) (default 0).
	RhExpr(rh expr.Expr) ColorbalanceBuilder
	// Gh set green highlights (from -1 to 1) (default 0).
	Gh(gh float32) ColorbalanceBuilder
	// GhExpr set green highlights (from -1 to 1) (default 0).
	GhExpr(gh expr.Expr) ColorbalanceBuilder
	// Bh set blue highlights (from -1 to 1) (default 0).
	Bh(bh float32) ColorbalanceBuilder
	// BhExpr set blue highlights (from -1 to 1) (default 0).
	BhExpr(bh expr.Expr) ColorbalanceBuilder
	// Pl preserve lightness (default false).
	Pl(pl bool) ColorbalanceBuilder
	// PlExpr preserve lightness (default false).
	PlExpr(pl expr.Expr) ColorbalanceBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ColorbalanceBuilder
}

// Colorbalance creates a new ColorbalanceBuilder to configure the "colorbalance" filter.
func Colorbalance(opts ...filter.Option) ColorbalanceBuilder {
	return &implColorbalanceBuilder{
		f: filter.New("colorbalance", 1, opts...),
	}
}

type implColorbalanceBuilder struct {
	f filter.Filter
}

func (colorbalanceBuilder *implColorbalanceBuilder) String() string {
	return colorbalanceBuilder.f.String()
}

func (colorbalanceBuilder *implColorbalanceBuilder) Outputs() int {
	return colorbalanceBuilder.f.Outputs()
}

func (colorbalanceBuilder *implColorbalanceBuilder) With(key string, value expr.Expr) filter.Filter {
	return colorbalanceBuilder.withOption(key, value)
}

func (colorbalanceBuilder *implColorbalanceBuilder) withOption(key string, value expr.Expr) ColorbalanceBuilder {
	bCopy := *colorbalanceBuilder
	bCopy.f = colorbalanceBuilder.f.With(key, value)
	return &bCopy
}

func (colorbalanceBuilder *implColorbalanceBuilder) Rs(rs float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("rs", expr.Float(rs))
}

func (colorbalanceBuilder *implColorbalanceBuilder) RsExpr(rs expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("rs", rs)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Gs(gs float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("gs", expr.Float(gs))
}

func (colorbalanceBuilder *implColorbalanceBuilder) GsExpr(gs expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("gs", gs)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Bs(bs float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("bs", expr.Float(bs))
}

func (colorbalanceBuilder *implColorbalanceBuilder) BsExpr(bs expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("bs", bs)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Rm(rm float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("rm", expr.Float(rm))
}

func (colorbalanceBuilder *implColorbalanceBuilder) RmExpr(rm expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("rm", rm)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Gm(gm float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("gm", expr.Float(gm))
}

func (colorbalanceBuilder *implColorbalanceBuilder) GmExpr(gm expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("gm", gm)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Bm(bm float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("bm", expr.Float(bm))
}

func (colorbalanceBuilder *implColorbalanceBuilder) BmExpr(bm expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("bm", bm)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Rh(rh float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("rh", expr.Float(rh))
}

func (colorbalanceBuilder *implColorbalanceBuilder) RhExpr(rh expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("rh", rh)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Gh(gh float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("gh", expr.Float(gh))
}

func (colorbalanceBuilder *implColorbalanceBuilder) GhExpr(gh expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("gh", gh)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Bh(bh float32) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("bh", expr.Float(bh))
}

func (colorbalanceBuilder *implColorbalanceBuilder) BhExpr(bh expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("bh", bh)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Pl(pl bool) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("pl", expr.Bool(pl))
}

func (colorbalanceBuilder *implColorbalanceBuilder) PlExpr(pl expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("pl", pl)
}

func (colorbalanceBuilder *implColorbalanceBuilder) Enable(enable expr.Expr) ColorbalanceBuilder {
	return colorbalanceBuilder.withOption("enable", enable)
}