// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DynaudnormBuilder Dynamic Audio Normalizer.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#dynaudnorm
type DynaudnormBuilder interface {
	filter.Filter
	// Framelen set the frame length in msec (from 10 to 8000) (default 500).
	Framelen(framelen int) DynaudnormBuilder
	// FramelenExpr set the frame length in msec (from 10 to 8000) (default 500).
	FramelenExpr(framelen expr.Expr) DynaudnormBuilder
	// F set the frame length in msec (from 10 to 8000) (default 500).
	F(f int) DynaudnormBuilder
	// FExpr set the frame length in msec (from 10 to 8000) (default 500).
	FExpr(f expr.Expr) DynaudnormBuilder
	// Gausssize set the filter size (from 3 to 301) (default 31).
	Gausssize(gausssize int) DynaudnormBuilder
	// GausssizeExpr set the filter size (from 3 to 301) (default 31).
	GausssizeExpr(gausssize expr.Expr) DynaudnormBuilder
	// G set the filter size (from 3 to 301) (default 31).
	G(g int) DynaudnormBuilder
	// GExpr set the filter size (from 3 to 301) (default 31).
	GExpr(g expr.Expr) DynaudnormBuilder
	// Peak set the peak value (from 0 to 1) (default 0.95).
	Peak(peak float64) DynaudnormBuilder
	// PeakExpr set the peak value (from 0 to 1) (default 0.95).
	PeakExpr(peak expr.Expr) DynaudnormBuilder
	// P set the peak value (from 0 to 1) (default 0.95).
	P(p float64) DynaudnormBuilder
	// PExpr set the peak value (from 0 to 1) (default 0.95).
	PExpr(p expr.Expr) DynaudnormBuilder
	// Maxgain set the max amplification (from 1 to 100) (default 10).
	Maxgain(maxgain float64) DynaudnormBuilder
	// MaxgainExpr set the max amplification (from 1 to 100) (default 10).
	MaxgainExpr(maxgain expr.Expr) DynaudnormBuilder
	// M set the max amplification (from 1 to 100) (default 10).
	M(m float64) DynaudnormBuilder
	// MExpr set the max amplification (from 1 to 100) (default 10).
	MExpr(m expr.Expr) DynaudnormBuilder
	// Targetrms set the target RMS (from 0 to 1) (default 0).
	Targetrms(targetrms float64) DynaudnormBuilder
	// TargetrmsExpr set the target RMS (from 0 to 1) (default 0).
	TargetrmsExpr(targetrms expr.Expr) DynaudnormBuilder
	// R set the target RMS (from 0 to 1) (default 0).
	R(r float64) DynaudnormBuilder
	// RExpr set the target RMS (from 0 to 1) (default 0).
	RExpr(r expr.Expr) DynaudnormBuilder
	// Coupling set channel coupling (default true).
	Coupling(coupling bool) DynaudnormBuilder
	// CouplingExpr set channel coupling (default true).
	CouplingExpr(coupling expr.Expr) DynaudnormBuilder
	// N set channel coupling (default true).
	N(n bool) DynaudnormBuilder
	// NExpr set channel coupling (default true).
	NExpr(n expr.Expr) DynaudnormBuilder
	// Correctdc set DC correction (default false).
	Correctdc(correctdc bool) DynaudnormBuilder
	// CorrectdcExpr set DC correction (default false).
	CorrectdcExpr(correctdc expr.Expr) DynaudnormBuilder
	// C set DC correction (default false).
	C(c bool) DynaudnormBuilder
	// CExpr set DC correction (default false).
	CExpr(c expr.Expr) DynaudnormBuilder
	// Altboundary set alternative boundary mode (default false).
	Altboundary(altboundary bool) DynaudnormBuilder
	// AltboundaryExpr set alternative boundary mode (default false).
	AltboundaryExpr(altboundary expr.Expr) DynaudnormBuilder
	// B set alternative boundary mode (default false).
	B(b bool) DynaudnormBuilder
	// BExpr set alternative boundary mode (default false).
	BExpr(b expr.Expr) DynaudnormBuilder
	// Compress set the compress factor (from 0 to 30) (default 0).
	Compress(compress float64) DynaudnormBuilder
	// CompressExpr set the compress factor (from 0 to 30) (default 0).
	CompressExpr(compress expr.Expr) DynaudnormBuilder
	// S set the compress factor (from 0 to 30) (default 0).
	S(s float64) DynaudnormBuilder
	// SExpr set the compress factor (from 0 to 30) (default 0).
	SExpr(s expr.Expr) DynaudnormBuilder
	// Threshold set the threshold value (from 0 to 1) (default 0).
	Threshold(threshold float64) DynaudnormBuilder
	// ThresholdExpr set the threshold value (from 0 to 1) (default 0).
	ThresholdExpr(threshold expr.Expr) DynaudnormBuilder
	// T set the threshold value (from 0 to 1) (default 0).
	T(t float64) DynaudnormBuilder
	// TExpr set the threshold value (from 0 to 1) (default 0).
	TExpr(t expr.Expr) DynaudnormBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) DynaudnormBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) DynaudnormBuilder
	// H set channels to filter (default "all").
	H(h string) DynaudnormBuilder
	// HExpr set channels to filter (default "all").
	HExpr(h expr.Expr) DynaudnormBuilder
	// Overlap set the frame overlap (from 0 to 1) (default 0).
	Overlap(overlap float64) DynaudnormBuilder
	// OverlapExpr set the frame overlap (from 0 to 1) (default 0).
	OverlapExpr(overlap expr.Expr) DynaudnormBuilder
	// O set the frame overlap (from 0 to 1) (default 0).
	O(o float64) DynaudnormBuilder
	// OExpr set the frame overlap (from 0 to 1) (default 0).
	OExpr(o expr.Expr) DynaudnormBuilder
	// Curve set the custom peak mapping curve.
	Curve(curve string) DynaudnormBuilder
	// CurveExpr set the custom peak mapping curve.
	CurveExpr(curve expr.Expr) DynaudnormBuilder
	// V set the custom peak mapping curve.
	V(v string) DynaudnormBuilder
	// VExpr set the custom peak mapping curve.
	VExpr(v expr.Expr) DynaudnormBuilder
}

// Dynaudnorm creates a new DynaudnormBuilder to configure the "dynaudnorm" filter.
func Dynaudnorm(opts ...filter.Option) DynaudnormBuilder {
	return &implDynaudnormBuilder{
		f: filter.New("dynaudnorm", 1, opts...),
	}
}

type implDynaudnormBuilder struct {
	f filter.Filter
}

func (dynaudnormBuilder *implDynaudnormBuilder) String() string {
	return dynaudnormBuilder.f.String()
}

func (dynaudnormBuilder *implDynaudnormBuilder) Outputs() int {
	return dynaudnormBuilder.f.Outputs()
}

func (dynaudnormBuilder *implDynaudnormBuilder) With(key string, value expr.Expr) filter.Filter {
	return dynaudnormBuilder.withOption(key, value)
}

func (dynaudnormBuilder *implDynaudnormBuilder) withOption(key string, value expr.Expr) DynaudnormBuilder {
	bCopy := *dynaudnormBuilder
	bCopy.f = dynaudnormBuilder.f.With(key, value)
	return &bCopy
}

func (dynaudnormBuilder *implDynaudnormBuilder) Framelen(framelen int) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("framelen", expr.Int(framelen))
}

func (dynaudnormBuilder *implDynaudnormBuilder) FramelenExpr(framelen expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("framelen", framelen)
}

func (dynaudnormBuilder *implDynaudnormBuilder) F(f int) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("f", expr.Int(f))
}

func (dynaudnormBuilder *implDynaudnormBuilder) FExpr(f expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("f", f)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Gausssize(gausssize int) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("gausssize", expr.Int(gausssize))
}

func (dynaudnormBuilder *implDynaudnormBuilder) GausssizeExpr(gausssize expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("gausssize", gausssize)
}

func (dynaudnormBuilder *implDynaudnormBuilder) G(g int) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("g", expr.Int(g))
}

func (dynaudnormBuilder *implDynaudnormBuilder) GExpr(g expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("g", g)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Peak(peak float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("peak", expr.Double(peak))
}

func (dynaudnormBuilder *implDynaudnormBuilder) PeakExpr(peak expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("peak", peak)
}

func (dynaudnormBuilder *implDynaudnormBuilder) P(p float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("p", expr.Double(p))
}

func (dynaudnormBuilder *implDynaudnormBuilder) PExpr(p expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("p", p)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Maxgain(maxgain float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("maxgain", expr.Double(maxgain))
}

func (dynaudnormBuilder *implDynaudnormBuilder) MaxgainExpr(maxgain expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("maxgain", maxgain)
}

func (dynaudnormBuilder *implDynaudnormBuilder) M(m float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("m", expr.Double(m))
}

func (dynaudnormBuilder *implDynaudnormBuilder) MExpr(m expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("m", m)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Targetrms(targetrms float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("targetrms", expr.Double(targetrms))
}

func (dynaudnormBuilder *implDynaudnormBuilder) TargetrmsExpr(targetrms expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("targetrms", targetrms)
}

func (dynaudnormBuilder *implDynaudnormBuilder) R(r float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("r", expr.Double(r))
}

func (dynaudnormBuilder *implDynaudnormBuilder) RExpr(r expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("r", r)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Coupling(coupling bool) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("coupling", expr.Bool(coupling))
}

func (dynaudnormBuilder *implDynaudnormBuilder) CouplingExpr(coupling expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("coupling", coupling)
}

func (dynaudnormBuilder *implDynaudnormBuilder) N(n bool) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("n", expr.Bool(n))
}

func (dynaudnormBuilder *implDynaudnormBuilder) NExpr(n expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("n", n)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Correctdc(correctdc bool) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("correctdc", expr.Bool(correctdc))
}

func (dynaudnormBuilder *implDynaudnormBuilder) CorrectdcExpr(correctdc expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("correctdc", correctdc)
}

func (dynaudnormBuilder *implDynaudnormBuilder) C(c bool) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("c", expr.Bool(c))
}

func (dynaudnormBuilder *implDynaudnormBuilder) CExpr(c expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("c", c)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Altboundary(altboundary bool) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("altboundary", expr.Bool(altboundary))
}

func (dynaudnormBuilder *implDynaudnormBuilder) AltboundaryExpr(altboundary expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("altboundary", altboundary)
}

func (dynaudnormBuilder *implDynaudnormBuilder) B(b bool) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("b", expr.Bool(b))
}

func (dynaudnormBuilder *implDynaudnormBuilder) BExpr(b expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("b", b)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Compress(compress float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("compress", expr.Double(compress))
}

func (dynaudnormBuilder *implDynaudnormBuilder) CompressExpr(compress expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("compress", compress)
}

func (dynaudnormBuilder *implDynaudnormBuilder) S(s float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("s", expr.Double(s))
}

func (dynaudnormBuilder *implDynaudnormBuilder) SExpr(s expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("s", s)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Threshold(threshold float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("threshold", expr.Double(threshold))
}

func (dynaudnormBuilder *implDynaudnormBuilder) ThresholdExpr(threshold expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("threshold", threshold)
}

func (dynaudnormBuilder *implDynaudnormBuilder) T(t float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("t", expr.Double(t))
}

func (dynaudnormBuilder *implDynaudnormBuilder) TExpr(t expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("t", t)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Channels(channels string) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("channels", expr.String(channels))
}

func (dynaudnormBuilder *implDynaudnormBuilder) ChannelsExpr(channels expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("channels", channels)
}

func (dynaudnormBuilder *implDynaudnormBuilder) H(h string) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("h", expr.String(h))
}

func (dynaudnormBuilder *implDynaudnormBuilder) HExpr(h expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("h", h)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Overlap(overlap float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("overlap", expr.Double(overlap))
}

func (dynaudnormBuilder *implDynaudnormBuilder) OverlapExpr(overlap expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("overlap", overlap)
}

func (dynaudnormBuilder *implDynaudnormBuilder) O(o float64) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("o", expr.Double(o))
}

func (dynaudnormBuilder *implDynaudnormBuilder) OExpr(o expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("o", o)
}

func (dynaudnormBuilder *implDynaudnormBuilder) Curve(curve string) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("curve", expr.String(curve))
}

func (dynaudnormBuilder *implDynaudnormBuilder) CurveExpr(curve expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("curve", curve)
}

func (dynaudnormBuilder *implDynaudnormBuilder) V(v string) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("v", expr.String(v))
}

func (dynaudnormBuilder *implDynaudnormBuilder) VExpr(v expr.Expr) DynaudnormBuilder {
	return dynaudnormBuilder.withOption("v", v)
}
