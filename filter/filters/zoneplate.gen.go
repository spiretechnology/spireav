// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ZoneplateBuilder Generate zone-plate.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#zoneplate
type ZoneplateBuilder interface {
	filter.Filter
	// Size set video size (default "320x240").
	Size(size expr.Size) ZoneplateBuilder
	// S set video size (default "320x240").
	S(s expr.Size) ZoneplateBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) ZoneplateBuilder
	// R set video rate (default "25").
	R(r expr.Rational) ZoneplateBuilder
	// Duration set video duration (default -0.000001).
	Duration(duration time.Duration) ZoneplateBuilder
	// D set video duration (default -0.000001).
	D(d time.Duration) ZoneplateBuilder
	// Sar set video sample aspect ratio (from 0 to INT_MAX) (default 1/1).
	Sar(sar expr.Rational) ZoneplateBuilder
	// Precision set LUT precision (from 4 to 16) (default 10).
	Precision(precision int) ZoneplateBuilder
	// Xo set X-axis offset (from INT_MIN to INT_MAX) (default 0).
	Xo(xo int) ZoneplateBuilder
	// XoExpr set X-axis offset (from INT_MIN to INT_MAX) (default 0).
	XoExpr(xo expr.Expr) ZoneplateBuilder
	// Yo set Y-axis offset (from INT_MIN to INT_MAX) (default 0).
	Yo(yo int) ZoneplateBuilder
	// YoExpr set Y-axis offset (from INT_MIN to INT_MAX) (default 0).
	YoExpr(yo expr.Expr) ZoneplateBuilder
	// To set T-axis offset (from INT_MIN to INT_MAX) (default 0).
	To(to int) ZoneplateBuilder
	// ToExpr set T-axis offset (from INT_MIN to INT_MAX) (default 0).
	ToExpr(to expr.Expr) ZoneplateBuilder
	// K0 set 0-order phase (from INT_MIN to INT_MAX) (default 0).
	K0(k0 int) ZoneplateBuilder
	// K0Expr set 0-order phase (from INT_MIN to INT_MAX) (default 0).
	K0Expr(k0 expr.Expr) ZoneplateBuilder
	// Kx set 1-order X-axis phase (from INT_MIN to INT_MAX) (default 0).
	Kx(kx int) ZoneplateBuilder
	// KxExpr set 1-order X-axis phase (from INT_MIN to INT_MAX) (default 0).
	KxExpr(kx expr.Expr) ZoneplateBuilder
	// Ky set 1-order Y-axis phase (from INT_MIN to INT_MAX) (default 0).
	Ky(ky int) ZoneplateBuilder
	// KyExpr set 1-order Y-axis phase (from INT_MIN to INT_MAX) (default 0).
	KyExpr(ky expr.Expr) ZoneplateBuilder
	// Kt set 1-order T-axis phase (from INT_MIN to INT_MAX) (default 0).
	Kt(kt int) ZoneplateBuilder
	// KtExpr set 1-order T-axis phase (from INT_MIN to INT_MAX) (default 0).
	KtExpr(kt expr.Expr) ZoneplateBuilder
	// Kxt set X-axis*T-axis product phase (from INT_MIN to INT_MAX) (default 0).
	Kxt(kxt int) ZoneplateBuilder
	// KxtExpr set X-axis*T-axis product phase (from INT_MIN to INT_MAX) (default 0).
	KxtExpr(kxt expr.Expr) ZoneplateBuilder
	// Kyt set Y-axis*T-axis product phase (from INT_MIN to INT_MAX) (default 0).
	Kyt(kyt int) ZoneplateBuilder
	// KytExpr set Y-axis*T-axis product phase (from INT_MIN to INT_MAX) (default 0).
	KytExpr(kyt expr.Expr) ZoneplateBuilder
	// Kxy set X-axis*Y-axis product phase (from INT_MIN to INT_MAX) (default 0).
	Kxy(kxy int) ZoneplateBuilder
	// KxyExpr set X-axis*Y-axis product phase (from INT_MIN to INT_MAX) (default 0).
	KxyExpr(kxy expr.Expr) ZoneplateBuilder
	// Kx2 set 2-order X-axis phase (from INT_MIN to INT_MAX) (default 0).
	Kx2(kx2 int) ZoneplateBuilder
	// Kx2Expr set 2-order X-axis phase (from INT_MIN to INT_MAX) (default 0).
	Kx2Expr(kx2 expr.Expr) ZoneplateBuilder
	// Ky2 set 2-order Y-axis phase (from INT_MIN to INT_MAX) (default 0).
	Ky2(ky2 int) ZoneplateBuilder
	// Ky2Expr set 2-order Y-axis phase (from INT_MIN to INT_MAX) (default 0).
	Ky2Expr(ky2 expr.Expr) ZoneplateBuilder
	// Kt2 set 2-order T-axis phase (from INT_MIN to INT_MAX) (default 0).
	Kt2(kt2 int) ZoneplateBuilder
	// Kt2Expr set 2-order T-axis phase (from INT_MIN to INT_MAX) (default 0).
	Kt2Expr(kt2 expr.Expr) ZoneplateBuilder
	// Ku set 0-order U-color phase (from INT_MIN to INT_MAX) (default 0).
	Ku(ku int) ZoneplateBuilder
	// KuExpr set 0-order U-color phase (from INT_MIN to INT_MAX) (default 0).
	KuExpr(ku expr.Expr) ZoneplateBuilder
	// Kv set 0-order V-color phase (from INT_MIN to INT_MAX) (default 0).
	Kv(kv int) ZoneplateBuilder
	// KvExpr set 0-order V-color phase (from INT_MIN to INT_MAX) (default 0).
	KvExpr(kv expr.Expr) ZoneplateBuilder
}

// Zoneplate creates a new ZoneplateBuilder to configure the "zoneplate" filter.
func Zoneplate(opts ...filter.Option) ZoneplateBuilder {
	return &implZoneplateBuilder{
		f: filter.New("zoneplate", 1, opts...),
	}
}

type implZoneplateBuilder struct {
	f filter.Filter
}

func (zoneplateBuilder *implZoneplateBuilder) String() string {
	return zoneplateBuilder.f.String()
}

func (zoneplateBuilder *implZoneplateBuilder) Outputs() int {
	return zoneplateBuilder.f.Outputs()
}

func (zoneplateBuilder *implZoneplateBuilder) With(key string, value expr.Expr) filter.Filter {
	return zoneplateBuilder.withOption(key, value)
}

func (zoneplateBuilder *implZoneplateBuilder) withOption(key string, value expr.Expr) ZoneplateBuilder {
	bCopy := *zoneplateBuilder
	bCopy.f = zoneplateBuilder.f.With(key, value)
	return &bCopy
}

func (zoneplateBuilder *implZoneplateBuilder) Size(size expr.Size) ZoneplateBuilder {
	return zoneplateBuilder.withOption("size", size)
}

func (zoneplateBuilder *implZoneplateBuilder) S(s expr.Size) ZoneplateBuilder {
	return zoneplateBuilder.withOption("s", s)
}

func (zoneplateBuilder *implZoneplateBuilder) Rate(rate expr.Rational) ZoneplateBuilder {
	return zoneplateBuilder.withOption("rate", rate)
}

func (zoneplateBuilder *implZoneplateBuilder) R(r expr.Rational) ZoneplateBuilder {
	return zoneplateBuilder.withOption("r", r)
}

func (zoneplateBuilder *implZoneplateBuilder) Duration(duration time.Duration) ZoneplateBuilder {
	return zoneplateBuilder.withOption("duration", expr.Duration(duration))
}

func (zoneplateBuilder *implZoneplateBuilder) D(d time.Duration) ZoneplateBuilder {
	return zoneplateBuilder.withOption("d", expr.Duration(d))
}

func (zoneplateBuilder *implZoneplateBuilder) Sar(sar expr.Rational) ZoneplateBuilder {
	return zoneplateBuilder.withOption("sar", sar)
}

func (zoneplateBuilder *implZoneplateBuilder) Precision(precision int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("precision", expr.Int(precision))
}

func (zoneplateBuilder *implZoneplateBuilder) Xo(xo int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("xo", expr.Int(xo))
}

func (zoneplateBuilder *implZoneplateBuilder) XoExpr(xo expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("xo", xo)
}

func (zoneplateBuilder *implZoneplateBuilder) Yo(yo int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("yo", expr.Int(yo))
}

func (zoneplateBuilder *implZoneplateBuilder) YoExpr(yo expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("yo", yo)
}

func (zoneplateBuilder *implZoneplateBuilder) To(to int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("to", expr.Int(to))
}

func (zoneplateBuilder *implZoneplateBuilder) ToExpr(to expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("to", to)
}

func (zoneplateBuilder *implZoneplateBuilder) K0(k0 int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("k0", expr.Int(k0))
}

func (zoneplateBuilder *implZoneplateBuilder) K0Expr(k0 expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("k0", k0)
}

func (zoneplateBuilder *implZoneplateBuilder) Kx(kx int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kx", expr.Int(kx))
}

func (zoneplateBuilder *implZoneplateBuilder) KxExpr(kx expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kx", kx)
}

func (zoneplateBuilder *implZoneplateBuilder) Ky(ky int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("ky", expr.Int(ky))
}

func (zoneplateBuilder *implZoneplateBuilder) KyExpr(ky expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("ky", ky)
}

func (zoneplateBuilder *implZoneplateBuilder) Kt(kt int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kt", expr.Int(kt))
}

func (zoneplateBuilder *implZoneplateBuilder) KtExpr(kt expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kt", kt)
}

func (zoneplateBuilder *implZoneplateBuilder) Kxt(kxt int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kxt", expr.Int(kxt))
}

func (zoneplateBuilder *implZoneplateBuilder) KxtExpr(kxt expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kxt", kxt)
}

func (zoneplateBuilder *implZoneplateBuilder) Kyt(kyt int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kyt", expr.Int(kyt))
}

func (zoneplateBuilder *implZoneplateBuilder) KytExpr(kyt expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kyt", kyt)
}

func (zoneplateBuilder *implZoneplateBuilder) Kxy(kxy int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kxy", expr.Int(kxy))
}

func (zoneplateBuilder *implZoneplateBuilder) KxyExpr(kxy expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kxy", kxy)
}

func (zoneplateBuilder *implZoneplateBuilder) Kx2(kx2 int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kx2", expr.Int(kx2))
}

func (zoneplateBuilder *implZoneplateBuilder) Kx2Expr(kx2 expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kx2", kx2)
}

func (zoneplateBuilder *implZoneplateBuilder) Ky2(ky2 int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("ky2", expr.Int(ky2))
}

func (zoneplateBuilder *implZoneplateBuilder) Ky2Expr(ky2 expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("ky2", ky2)
}

func (zoneplateBuilder *implZoneplateBuilder) Kt2(kt2 int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kt2", expr.Int(kt2))
}

func (zoneplateBuilder *implZoneplateBuilder) Kt2Expr(kt2 expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kt2", kt2)
}

func (zoneplateBuilder *implZoneplateBuilder) Ku(ku int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("ku", expr.Int(ku))
}

func (zoneplateBuilder *implZoneplateBuilder) KuExpr(ku expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("ku", ku)
}

func (zoneplateBuilder *implZoneplateBuilder) Kv(kv int) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kv", expr.Int(kv))
}

func (zoneplateBuilder *implZoneplateBuilder) KvExpr(kv expr.Expr) ZoneplateBuilder {
	return zoneplateBuilder.withOption("kv", kv)
}
