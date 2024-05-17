// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColorlevelsBuilder Adjust the color levels.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colorlevels
type ColorlevelsBuilder interface {
	filter.Filter
	// Rimin set input red black point (from -1 to 1) (default 0).
	Rimin(rimin float64) ColorlevelsBuilder
	// RiminExpr set input red black point (from -1 to 1) (default 0).
	RiminExpr(rimin expr.Expr) ColorlevelsBuilder
	// Gimin set input green black point (from -1 to 1) (default 0).
	Gimin(gimin float64) ColorlevelsBuilder
	// GiminExpr set input green black point (from -1 to 1) (default 0).
	GiminExpr(gimin expr.Expr) ColorlevelsBuilder
	// Bimin set input blue black point (from -1 to 1) (default 0).
	Bimin(bimin float64) ColorlevelsBuilder
	// BiminExpr set input blue black point (from -1 to 1) (default 0).
	BiminExpr(bimin expr.Expr) ColorlevelsBuilder
	// Aimin set input alpha black point (from -1 to 1) (default 0).
	Aimin(aimin float64) ColorlevelsBuilder
	// AiminExpr set input alpha black point (from -1 to 1) (default 0).
	AiminExpr(aimin expr.Expr) ColorlevelsBuilder
	// Rimax set input red white point (from -1 to 1) (default 1).
	Rimax(rimax float64) ColorlevelsBuilder
	// RimaxExpr set input red white point (from -1 to 1) (default 1).
	RimaxExpr(rimax expr.Expr) ColorlevelsBuilder
	// Gimax set input green white point (from -1 to 1) (default 1).
	Gimax(gimax float64) ColorlevelsBuilder
	// GimaxExpr set input green white point (from -1 to 1) (default 1).
	GimaxExpr(gimax expr.Expr) ColorlevelsBuilder
	// Bimax set input blue white point (from -1 to 1) (default 1).
	Bimax(bimax float64) ColorlevelsBuilder
	// BimaxExpr set input blue white point (from -1 to 1) (default 1).
	BimaxExpr(bimax expr.Expr) ColorlevelsBuilder
	// Aimax set input alpha white point (from -1 to 1) (default 1).
	Aimax(aimax float64) ColorlevelsBuilder
	// AimaxExpr set input alpha white point (from -1 to 1) (default 1).
	AimaxExpr(aimax expr.Expr) ColorlevelsBuilder
	// Romin set output red black point (from 0 to 1) (default 0).
	Romin(romin float64) ColorlevelsBuilder
	// RominExpr set output red black point (from 0 to 1) (default 0).
	RominExpr(romin expr.Expr) ColorlevelsBuilder
	// Gomin set output green black point (from 0 to 1) (default 0).
	Gomin(gomin float64) ColorlevelsBuilder
	// GominExpr set output green black point (from 0 to 1) (default 0).
	GominExpr(gomin expr.Expr) ColorlevelsBuilder
	// Bomin set output blue black point (from 0 to 1) (default 0).
	Bomin(bomin float64) ColorlevelsBuilder
	// BominExpr set output blue black point (from 0 to 1) (default 0).
	BominExpr(bomin expr.Expr) ColorlevelsBuilder
	// Aomin set output alpha black point (from 0 to 1) (default 0).
	Aomin(aomin float64) ColorlevelsBuilder
	// AominExpr set output alpha black point (from 0 to 1) (default 0).
	AominExpr(aomin expr.Expr) ColorlevelsBuilder
	// Romax set output red white point (from 0 to 1) (default 1).
	Romax(romax float64) ColorlevelsBuilder
	// RomaxExpr set output red white point (from 0 to 1) (default 1).
	RomaxExpr(romax expr.Expr) ColorlevelsBuilder
	// Gomax set output green white point (from 0 to 1) (default 1).
	Gomax(gomax float64) ColorlevelsBuilder
	// GomaxExpr set output green white point (from 0 to 1) (default 1).
	GomaxExpr(gomax expr.Expr) ColorlevelsBuilder
	// Bomax set output blue white point (from 0 to 1) (default 1).
	Bomax(bomax float64) ColorlevelsBuilder
	// BomaxExpr set output blue white point (from 0 to 1) (default 1).
	BomaxExpr(bomax expr.Expr) ColorlevelsBuilder
	// Aomax set output alpha white point (from 0 to 1) (default 1).
	Aomax(aomax float64) ColorlevelsBuilder
	// AomaxExpr set output alpha white point (from 0 to 1) (default 1).
	AomaxExpr(aomax expr.Expr) ColorlevelsBuilder
	// Preserve set preserve color mode (from 0 to 6) (default none).
	Preserve(preserve int) ColorlevelsBuilder
	// PreserveExpr set preserve color mode (from 0 to 6) (default none).
	PreserveExpr(preserve expr.Expr) ColorlevelsBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ColorlevelsBuilder
}

// Colorlevels creates a new ColorlevelsBuilder to configure the "colorlevels" filter.
func Colorlevels(opts ...filter.Option) ColorlevelsBuilder {
	return &implColorlevelsBuilder{
		f: filter.New("colorlevels", 1, opts...),
	}
}

type implColorlevelsBuilder struct {
	f filter.Filter
}

func (colorlevelsBuilder *implColorlevelsBuilder) String() string {
	return colorlevelsBuilder.f.String()
}

func (colorlevelsBuilder *implColorlevelsBuilder) Outputs() int {
	return colorlevelsBuilder.f.Outputs()
}

func (colorlevelsBuilder *implColorlevelsBuilder) With(key string, value expr.Expr) filter.Filter {
	return colorlevelsBuilder.withOption(key, value)
}

func (colorlevelsBuilder *implColorlevelsBuilder) withOption(key string, value expr.Expr) ColorlevelsBuilder {
	bCopy := *colorlevelsBuilder
	bCopy.f = colorlevelsBuilder.f.With(key, value)
	return &bCopy
}

func (colorlevelsBuilder *implColorlevelsBuilder) Rimin(rimin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("rimin", expr.Double(rimin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) RiminExpr(rimin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("rimin", rimin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Gimin(gimin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gimin", expr.Double(gimin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) GiminExpr(gimin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gimin", gimin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Bimin(bimin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bimin", expr.Double(bimin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) BiminExpr(bimin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bimin", bimin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Aimin(aimin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aimin", expr.Double(aimin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) AiminExpr(aimin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aimin", aimin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Rimax(rimax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("rimax", expr.Double(rimax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) RimaxExpr(rimax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("rimax", rimax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Gimax(gimax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gimax", expr.Double(gimax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) GimaxExpr(gimax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gimax", gimax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Bimax(bimax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bimax", expr.Double(bimax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) BimaxExpr(bimax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bimax", bimax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Aimax(aimax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aimax", expr.Double(aimax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) AimaxExpr(aimax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aimax", aimax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Romin(romin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("romin", expr.Double(romin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) RominExpr(romin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("romin", romin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Gomin(gomin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gomin", expr.Double(gomin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) GominExpr(gomin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gomin", gomin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Bomin(bomin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bomin", expr.Double(bomin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) BominExpr(bomin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bomin", bomin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Aomin(aomin float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aomin", expr.Double(aomin))
}

func (colorlevelsBuilder *implColorlevelsBuilder) AominExpr(aomin expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aomin", aomin)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Romax(romax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("romax", expr.Double(romax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) RomaxExpr(romax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("romax", romax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Gomax(gomax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gomax", expr.Double(gomax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) GomaxExpr(gomax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("gomax", gomax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Bomax(bomax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bomax", expr.Double(bomax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) BomaxExpr(bomax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("bomax", bomax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Aomax(aomax float64) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aomax", expr.Double(aomax))
}

func (colorlevelsBuilder *implColorlevelsBuilder) AomaxExpr(aomax expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("aomax", aomax)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Preserve(preserve int) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("preserve", expr.Int(preserve))
}

func (colorlevelsBuilder *implColorlevelsBuilder) PreserveExpr(preserve expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("preserve", preserve)
}

func (colorlevelsBuilder *implColorlevelsBuilder) Enable(enable expr.Expr) ColorlevelsBuilder {
	return colorlevelsBuilder.withOption("enable", enable)
}
