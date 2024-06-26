// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DrawtextBuilder Draw text on top of video frames using libfreetype library.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#drawtext
type DrawtextBuilder interface {
	filter.Filter
	// Fontfile set font file.
	Fontfile(fontfile string) DrawtextBuilder
	// Text set text.
	Text(text string) DrawtextBuilder
	// TextExpr set text.
	TextExpr(text expr.Expr) DrawtextBuilder
	// Textfile set text file.
	Textfile(textfile string) DrawtextBuilder
	// Fontcolor set foreground color (default "black").
	Fontcolor(fontcolor expr.Color) DrawtextBuilder
	// FontcolorExpr set foreground color expression (default "").
	FontcolorExpr(fontcolorExpr expr.Expr) DrawtextBuilder
	// Boxcolor set box color (default "white").
	Boxcolor(boxcolor expr.Color) DrawtextBuilder
	// BoxcolorExpr set box color (default "white").
	BoxcolorExpr(boxcolor expr.Expr) DrawtextBuilder
	// Bordercolor set border color (default "black").
	Bordercolor(bordercolor expr.Color) DrawtextBuilder
	// BordercolorExpr set border color (default "black").
	BordercolorExpr(bordercolor expr.Expr) DrawtextBuilder
	// Shadowcolor set shadow color (default "black").
	Shadowcolor(shadowcolor expr.Color) DrawtextBuilder
	// ShadowcolorExpr set shadow color (default "black").
	ShadowcolorExpr(shadowcolor expr.Expr) DrawtextBuilder
	// Box set box (default false).
	Box(box bool) DrawtextBuilder
	// BoxExpr set box (default false).
	BoxExpr(box expr.Expr) DrawtextBuilder
	// Boxborderw set box borders width (default "0").
	Boxborderw(boxborderw string) DrawtextBuilder
	// BoxborderwExpr set box borders width (default "0").
	BoxborderwExpr(boxborderw expr.Expr) DrawtextBuilder
	// LineSpacing set line spacing in pixels (from INT_MIN to INT_MAX) (default 0).
	LineSpacing(lineSpacing int) DrawtextBuilder
	// LineSpacingExpr set line spacing in pixels (from INT_MIN to INT_MAX) (default 0).
	LineSpacingExpr(lineSpacing expr.Expr) DrawtextBuilder
	// Fontsize set font size.
	Fontsize(fontsize string) DrawtextBuilder
	// FontsizeExpr set font size.
	FontsizeExpr(fontsize expr.Expr) DrawtextBuilder
	// TextAlign set text alignment (default 0).
	TextAlign(textAlign ...string) DrawtextBuilder
	// TextAlignExpr set text alignment (default 0).
	TextAlignExpr(textAlign expr.Expr) DrawtextBuilder
	// X set x expression (default "0").
	X(x int) DrawtextBuilder
	// XExpr set x expression (default "0").
	XExpr(x expr.Expr) DrawtextBuilder
	// Y set y expression (default "0").
	Y(y int) DrawtextBuilder
	// YExpr set y expression (default "0").
	YExpr(y expr.Expr) DrawtextBuilder
	// Boxw set box width (from 0 to INT_MAX) (default 0).
	Boxw(boxw int) DrawtextBuilder
	// BoxwExpr set box width (from 0 to INT_MAX) (default 0).
	BoxwExpr(boxw expr.Expr) DrawtextBuilder
	// Boxh set box height (from 0 to INT_MAX) (default 0).
	Boxh(boxh int) DrawtextBuilder
	// BoxhExpr set box height (from 0 to INT_MAX) (default 0).
	BoxhExpr(boxh expr.Expr) DrawtextBuilder
	// Shadowx set shadow x offset (from INT_MIN to INT_MAX) (default 0).
	Shadowx(shadowx int) DrawtextBuilder
	// ShadowxExpr set shadow x offset (from INT_MIN to INT_MAX) (default 0).
	ShadowxExpr(shadowx expr.Expr) DrawtextBuilder
	// Shadowy set shadow y offset (from INT_MIN to INT_MAX) (default 0).
	Shadowy(shadowy int) DrawtextBuilder
	// ShadowyExpr set shadow y offset (from INT_MIN to INT_MAX) (default 0).
	ShadowyExpr(shadowy expr.Expr) DrawtextBuilder
	// Borderw set border width (from INT_MIN to INT_MAX) (default 0).
	Borderw(borderw int) DrawtextBuilder
	// BorderwExpr set border width (from INT_MIN to INT_MAX) (default 0).
	BorderwExpr(borderw expr.Expr) DrawtextBuilder
	// Tabsize set tab size (from 0 to INT_MAX) (default 4).
	Tabsize(tabsize int) DrawtextBuilder
	// TabsizeExpr set tab size (from 0 to INT_MAX) (default 4).
	TabsizeExpr(tabsize expr.Expr) DrawtextBuilder
	// Basetime set base time (from I64_MIN to I64_MAX) (default I64_MIN).
	Basetime(basetime int64) DrawtextBuilder
	// Font Font name (default "Sans").
	Font(font string) DrawtextBuilder
	// Expansion set the expansion mode (from 0 to 2) (default normal).
	Expansion(expansion int) DrawtextBuilder
	// YAlign set the y alignment (from 0 to 2) (default text).
	YAlign(yAlign int) DrawtextBuilder
	// YAlignExpr set the y alignment (from 0 to 2) (default text).
	YAlignExpr(yAlign expr.Expr) DrawtextBuilder
	// Timecode set initial timecode.
	Timecode(timecode string) DrawtextBuilder
	// Tc24hmax set 24 hours max (timecode only) (default false).
	Tc24hmax(tc24hmax bool) DrawtextBuilder
	// TimecodeRate set rate (timecode only) (from 0 to INT_MAX) (default 0/1).
	TimecodeRate(timecodeRate expr.Rational) DrawtextBuilder
	// R set rate (timecode only) (from 0 to INT_MAX) (default 0/1).
	R(r expr.Rational) DrawtextBuilder
	// Rate set rate (timecode only) (from 0 to INT_MAX) (default 0/1).
	Rate(rate expr.Rational) DrawtextBuilder
	// Reload reload text file at specified frame interval (from 0 to INT_MAX) (default 0).
	Reload(reload int) DrawtextBuilder
	// Alpha apply alpha while rendering (default "1").
	Alpha(alpha string) DrawtextBuilder
	// AlphaExpr apply alpha while rendering (default "1").
	AlphaExpr(alpha expr.Expr) DrawtextBuilder
	// FixBounds check and fix text coords to avoid clipping (default false).
	FixBounds(fixBounds bool) DrawtextBuilder
	// StartNumber start frame number for n/frame_num variable (from 0 to INT_MAX) (default 0).
	StartNumber(startNumber int) DrawtextBuilder
	// TextSource the source of text.
	TextSource(textSource string) DrawtextBuilder
	// FtLoadFlags set font loading flags for libfreetype (default 0).
	FtLoadFlags(ftLoadFlags ...string) DrawtextBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) DrawtextBuilder
}

// Drawtext creates a new DrawtextBuilder to configure the "drawtext" filter.
func Drawtext(opts ...filter.Option) DrawtextBuilder {
	return &implDrawtextBuilder{
		f: filter.New("drawtext", 1, opts...),
	}
}

type implDrawtextBuilder struct {
	f filter.Filter
}

func (drawtextBuilder *implDrawtextBuilder) String() string {
	return drawtextBuilder.f.String()
}

func (drawtextBuilder *implDrawtextBuilder) Outputs() int {
	return drawtextBuilder.f.Outputs()
}

func (drawtextBuilder *implDrawtextBuilder) With(key string, value expr.Expr) filter.Filter {
	return drawtextBuilder.withOption(key, value)
}

func (drawtextBuilder *implDrawtextBuilder) withOption(key string, value expr.Expr) DrawtextBuilder {
	bCopy := *drawtextBuilder
	bCopy.f = drawtextBuilder.f.With(key, value)
	return &bCopy
}

func (drawtextBuilder *implDrawtextBuilder) Fontfile(fontfile string) DrawtextBuilder {
	return drawtextBuilder.withOption("fontfile", expr.String(fontfile))
}

func (drawtextBuilder *implDrawtextBuilder) Text(text string) DrawtextBuilder {
	return drawtextBuilder.withOption("text", expr.String(text))
}

func (drawtextBuilder *implDrawtextBuilder) TextExpr(text expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("text", text)
}

func (drawtextBuilder *implDrawtextBuilder) Textfile(textfile string) DrawtextBuilder {
	return drawtextBuilder.withOption("textfile", expr.String(textfile))
}

func (drawtextBuilder *implDrawtextBuilder) Fontcolor(fontcolor expr.Color) DrawtextBuilder {
	return drawtextBuilder.withOption("fontcolor", fontcolor)
}

func (drawtextBuilder *implDrawtextBuilder) FontcolorExpr(fontcolorExpr expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("fontcolor_expr", fontcolorExpr)
}

func (drawtextBuilder *implDrawtextBuilder) Boxcolor(boxcolor expr.Color) DrawtextBuilder {
	return drawtextBuilder.withOption("boxcolor", boxcolor)
}

func (drawtextBuilder *implDrawtextBuilder) BoxcolorExpr(boxcolor expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("boxcolor", boxcolor)
}

func (drawtextBuilder *implDrawtextBuilder) Bordercolor(bordercolor expr.Color) DrawtextBuilder {
	return drawtextBuilder.withOption("bordercolor", bordercolor)
}

func (drawtextBuilder *implDrawtextBuilder) BordercolorExpr(bordercolor expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("bordercolor", bordercolor)
}

func (drawtextBuilder *implDrawtextBuilder) Shadowcolor(shadowcolor expr.Color) DrawtextBuilder {
	return drawtextBuilder.withOption("shadowcolor", shadowcolor)
}

func (drawtextBuilder *implDrawtextBuilder) ShadowcolorExpr(shadowcolor expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("shadowcolor", shadowcolor)
}

func (drawtextBuilder *implDrawtextBuilder) Box(box bool) DrawtextBuilder {
	return drawtextBuilder.withOption("box", expr.Bool(box))
}

func (drawtextBuilder *implDrawtextBuilder) BoxExpr(box expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("box", box)
}

func (drawtextBuilder *implDrawtextBuilder) Boxborderw(boxborderw string) DrawtextBuilder {
	return drawtextBuilder.withOption("boxborderw", expr.String(boxborderw))
}

func (drawtextBuilder *implDrawtextBuilder) BoxborderwExpr(boxborderw expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("boxborderw", boxborderw)
}

func (drawtextBuilder *implDrawtextBuilder) LineSpacing(lineSpacing int) DrawtextBuilder {
	return drawtextBuilder.withOption("line_spacing", expr.Int(lineSpacing))
}

func (drawtextBuilder *implDrawtextBuilder) LineSpacingExpr(lineSpacing expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("line_spacing", lineSpacing)
}

func (drawtextBuilder *implDrawtextBuilder) Fontsize(fontsize string) DrawtextBuilder {
	return drawtextBuilder.withOption("fontsize", expr.String(fontsize))
}

func (drawtextBuilder *implDrawtextBuilder) FontsizeExpr(fontsize expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("fontsize", fontsize)
}

func (drawtextBuilder *implDrawtextBuilder) TextAlign(textAlign ...string) DrawtextBuilder {
	return drawtextBuilder.withOption("text_align", expr.Flags(textAlign))
}

func (drawtextBuilder *implDrawtextBuilder) TextAlignExpr(textAlign expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("text_align", textAlign)
}

func (drawtextBuilder *implDrawtextBuilder) X(x int) DrawtextBuilder {
	return drawtextBuilder.withOption("x", expr.Int(x))
}

func (drawtextBuilder *implDrawtextBuilder) XExpr(x expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("x", x)
}

func (drawtextBuilder *implDrawtextBuilder) Y(y int) DrawtextBuilder {
	return drawtextBuilder.withOption("y", expr.Int(y))
}

func (drawtextBuilder *implDrawtextBuilder) YExpr(y expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("y", y)
}

func (drawtextBuilder *implDrawtextBuilder) Boxw(boxw int) DrawtextBuilder {
	return drawtextBuilder.withOption("boxw", expr.Int(boxw))
}

func (drawtextBuilder *implDrawtextBuilder) BoxwExpr(boxw expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("boxw", boxw)
}

func (drawtextBuilder *implDrawtextBuilder) Boxh(boxh int) DrawtextBuilder {
	return drawtextBuilder.withOption("boxh", expr.Int(boxh))
}

func (drawtextBuilder *implDrawtextBuilder) BoxhExpr(boxh expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("boxh", boxh)
}

func (drawtextBuilder *implDrawtextBuilder) Shadowx(shadowx int) DrawtextBuilder {
	return drawtextBuilder.withOption("shadowx", expr.Int(shadowx))
}

func (drawtextBuilder *implDrawtextBuilder) ShadowxExpr(shadowx expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("shadowx", shadowx)
}

func (drawtextBuilder *implDrawtextBuilder) Shadowy(shadowy int) DrawtextBuilder {
	return drawtextBuilder.withOption("shadowy", expr.Int(shadowy))
}

func (drawtextBuilder *implDrawtextBuilder) ShadowyExpr(shadowy expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("shadowy", shadowy)
}

func (drawtextBuilder *implDrawtextBuilder) Borderw(borderw int) DrawtextBuilder {
	return drawtextBuilder.withOption("borderw", expr.Int(borderw))
}

func (drawtextBuilder *implDrawtextBuilder) BorderwExpr(borderw expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("borderw", borderw)
}

func (drawtextBuilder *implDrawtextBuilder) Tabsize(tabsize int) DrawtextBuilder {
	return drawtextBuilder.withOption("tabsize", expr.Int(tabsize))
}

func (drawtextBuilder *implDrawtextBuilder) TabsizeExpr(tabsize expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("tabsize", tabsize)
}

func (drawtextBuilder *implDrawtextBuilder) Basetime(basetime int64) DrawtextBuilder {
	return drawtextBuilder.withOption("basetime", expr.Int64(basetime))
}

func (drawtextBuilder *implDrawtextBuilder) Font(font string) DrawtextBuilder {
	return drawtextBuilder.withOption("font", expr.String(font))
}

func (drawtextBuilder *implDrawtextBuilder) Expansion(expansion int) DrawtextBuilder {
	return drawtextBuilder.withOption("expansion", expr.Int(expansion))
}

func (drawtextBuilder *implDrawtextBuilder) YAlign(yAlign int) DrawtextBuilder {
	return drawtextBuilder.withOption("y_align", expr.Int(yAlign))
}

func (drawtextBuilder *implDrawtextBuilder) YAlignExpr(yAlign expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("y_align", yAlign)
}

func (drawtextBuilder *implDrawtextBuilder) Timecode(timecode string) DrawtextBuilder {
	return drawtextBuilder.withOption("timecode", expr.String(timecode))
}

func (drawtextBuilder *implDrawtextBuilder) Tc24hmax(tc24hmax bool) DrawtextBuilder {
	return drawtextBuilder.withOption("tc24hmax", expr.Bool(tc24hmax))
}

func (drawtextBuilder *implDrawtextBuilder) TimecodeRate(timecodeRate expr.Rational) DrawtextBuilder {
	return drawtextBuilder.withOption("timecode_rate", timecodeRate)
}

func (drawtextBuilder *implDrawtextBuilder) R(r expr.Rational) DrawtextBuilder {
	return drawtextBuilder.withOption("r", r)
}

func (drawtextBuilder *implDrawtextBuilder) Rate(rate expr.Rational) DrawtextBuilder {
	return drawtextBuilder.withOption("rate", rate)
}

func (drawtextBuilder *implDrawtextBuilder) Reload(reload int) DrawtextBuilder {
	return drawtextBuilder.withOption("reload", expr.Int(reload))
}

func (drawtextBuilder *implDrawtextBuilder) Alpha(alpha string) DrawtextBuilder {
	return drawtextBuilder.withOption("alpha", expr.String(alpha))
}

func (drawtextBuilder *implDrawtextBuilder) AlphaExpr(alpha expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("alpha", alpha)
}

func (drawtextBuilder *implDrawtextBuilder) FixBounds(fixBounds bool) DrawtextBuilder {
	return drawtextBuilder.withOption("fix_bounds", expr.Bool(fixBounds))
}

func (drawtextBuilder *implDrawtextBuilder) StartNumber(startNumber int) DrawtextBuilder {
	return drawtextBuilder.withOption("start_number", expr.Int(startNumber))
}

func (drawtextBuilder *implDrawtextBuilder) TextSource(textSource string) DrawtextBuilder {
	return drawtextBuilder.withOption("text_source", expr.String(textSource))
}

func (drawtextBuilder *implDrawtextBuilder) FtLoadFlags(ftLoadFlags ...string) DrawtextBuilder {
	return drawtextBuilder.withOption("ft_load_flags", expr.Flags(ftLoadFlags))
}

func (drawtextBuilder *implDrawtextBuilder) Enable(enable expr.Expr) DrawtextBuilder {
	return drawtextBuilder.withOption("enable", enable)
}
