// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AfadeBuilder Fade in/out input audio.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#afade
type AfadeBuilder interface {
	filter.Filter
	// Type set the fade direction (from 0 to 1) (default in).
	Type(typ int) AfadeBuilder
	// TypeExpr set the fade direction (from 0 to 1) (default in).
	TypeExpr(typ expr.Expr) AfadeBuilder
	// T set the fade direction (from 0 to 1) (default in).
	T(t int) AfadeBuilder
	// TExpr set the fade direction (from 0 to 1) (default in).
	TExpr(t expr.Expr) AfadeBuilder
	// StartSample set number of first sample to start fading (from 0 to I64_MAX) (default 0).
	StartSample(startSample int64) AfadeBuilder
	// StartSampleExpr set number of first sample to start fading (from 0 to I64_MAX) (default 0).
	StartSampleExpr(startSample expr.Expr) AfadeBuilder
	// Ss set number of first sample to start fading (from 0 to I64_MAX) (default 0).
	Ss(ss int64) AfadeBuilder
	// SsExpr set number of first sample to start fading (from 0 to I64_MAX) (default 0).
	SsExpr(ss expr.Expr) AfadeBuilder
	// NbSamples set number of samples for fade duration (from 1 to I64_MAX) (default 44100).
	NbSamples(nbSamples int64) AfadeBuilder
	// NbSamplesExpr set number of samples for fade duration (from 1 to I64_MAX) (default 44100).
	NbSamplesExpr(nbSamples expr.Expr) AfadeBuilder
	// Ns set number of samples for fade duration (from 1 to I64_MAX) (default 44100).
	Ns(ns int64) AfadeBuilder
	// NsExpr set number of samples for fade duration (from 1 to I64_MAX) (default 44100).
	NsExpr(ns expr.Expr) AfadeBuilder
	// StartTime set time to start fading (default 0).
	StartTime(startTime time.Duration) AfadeBuilder
	// StartTimeExpr set time to start fading (default 0).
	StartTimeExpr(startTime expr.Expr) AfadeBuilder
	// St set time to start fading (default 0).
	St(st time.Duration) AfadeBuilder
	// StExpr set time to start fading (default 0).
	StExpr(st expr.Expr) AfadeBuilder
	// Duration set fade duration (default 0).
	Duration(duration time.Duration) AfadeBuilder
	// DurationExpr set fade duration (default 0).
	DurationExpr(duration expr.Expr) AfadeBuilder
	// D set fade duration (default 0).
	D(d time.Duration) AfadeBuilder
	// DExpr set fade duration (default 0).
	DExpr(d expr.Expr) AfadeBuilder
	// Curve set fade curve type (from -1 to 22) (default tri).
	Curve(curve int) AfadeBuilder
	// CurveExpr set fade curve type (from -1 to 22) (default tri).
	CurveExpr(curve expr.Expr) AfadeBuilder
	// C set fade curve type (from -1 to 22) (default tri).
	C(c int) AfadeBuilder
	// CExpr set fade curve type (from -1 to 22) (default tri).
	CExpr(c expr.Expr) AfadeBuilder
	// Silence set the silence gain (from 0 to 1) (default 0).
	Silence(silence float64) AfadeBuilder
	// SilenceExpr set the silence gain (from 0 to 1) (default 0).
	SilenceExpr(silence expr.Expr) AfadeBuilder
	// Unity set the unity gain (from 0 to 1) (default 1).
	Unity(unity float64) AfadeBuilder
	// UnityExpr set the unity gain (from 0 to 1) (default 1).
	UnityExpr(unity expr.Expr) AfadeBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AfadeBuilder
}

// Afade creates a new AfadeBuilder to configure the "afade" filter.
func Afade(opts ...filter.Option) AfadeBuilder {
	return &implAfadeBuilder{
		f: filter.New("afade", 1, opts...),
	}
}

type implAfadeBuilder struct {
	f filter.Filter
}

func (afadeBuilder *implAfadeBuilder) String() string {
	return afadeBuilder.f.String()
}

func (afadeBuilder *implAfadeBuilder) Outputs() int {
	return afadeBuilder.f.Outputs()
}

func (afadeBuilder *implAfadeBuilder) With(key string, value expr.Expr) filter.Filter {
	return afadeBuilder.withOption(key, value)
}

func (afadeBuilder *implAfadeBuilder) withOption(key string, value expr.Expr) AfadeBuilder {
	bCopy := *afadeBuilder
	bCopy.f = afadeBuilder.f.With(key, value)
	return &bCopy
}

func (afadeBuilder *implAfadeBuilder) Type(typ int) AfadeBuilder {
	return afadeBuilder.withOption("type", expr.Int(typ))
}

func (afadeBuilder *implAfadeBuilder) TypeExpr(typ expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("type", typ)
}

func (afadeBuilder *implAfadeBuilder) T(t int) AfadeBuilder {
	return afadeBuilder.withOption("t", expr.Int(t))
}

func (afadeBuilder *implAfadeBuilder) TExpr(t expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("t", t)
}

func (afadeBuilder *implAfadeBuilder) StartSample(startSample int64) AfadeBuilder {
	return afadeBuilder.withOption("start_sample", expr.Int64(startSample))
}

func (afadeBuilder *implAfadeBuilder) StartSampleExpr(startSample expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("start_sample", startSample)
}

func (afadeBuilder *implAfadeBuilder) Ss(ss int64) AfadeBuilder {
	return afadeBuilder.withOption("ss", expr.Int64(ss))
}

func (afadeBuilder *implAfadeBuilder) SsExpr(ss expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("ss", ss)
}

func (afadeBuilder *implAfadeBuilder) NbSamples(nbSamples int64) AfadeBuilder {
	return afadeBuilder.withOption("nb_samples", expr.Int64(nbSamples))
}

func (afadeBuilder *implAfadeBuilder) NbSamplesExpr(nbSamples expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("nb_samples", nbSamples)
}

func (afadeBuilder *implAfadeBuilder) Ns(ns int64) AfadeBuilder {
	return afadeBuilder.withOption("ns", expr.Int64(ns))
}

func (afadeBuilder *implAfadeBuilder) NsExpr(ns expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("ns", ns)
}

func (afadeBuilder *implAfadeBuilder) StartTime(startTime time.Duration) AfadeBuilder {
	return afadeBuilder.withOption("start_time", expr.Duration(startTime))
}

func (afadeBuilder *implAfadeBuilder) StartTimeExpr(startTime expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("start_time", startTime)
}

func (afadeBuilder *implAfadeBuilder) St(st time.Duration) AfadeBuilder {
	return afadeBuilder.withOption("st", expr.Duration(st))
}

func (afadeBuilder *implAfadeBuilder) StExpr(st expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("st", st)
}

func (afadeBuilder *implAfadeBuilder) Duration(duration time.Duration) AfadeBuilder {
	return afadeBuilder.withOption("duration", expr.Duration(duration))
}

func (afadeBuilder *implAfadeBuilder) DurationExpr(duration expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("duration", duration)
}

func (afadeBuilder *implAfadeBuilder) D(d time.Duration) AfadeBuilder {
	return afadeBuilder.withOption("d", expr.Duration(d))
}

func (afadeBuilder *implAfadeBuilder) DExpr(d expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("d", d)
}

func (afadeBuilder *implAfadeBuilder) Curve(curve int) AfadeBuilder {
	return afadeBuilder.withOption("curve", expr.Int(curve))
}

func (afadeBuilder *implAfadeBuilder) CurveExpr(curve expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("curve", curve)
}

func (afadeBuilder *implAfadeBuilder) C(c int) AfadeBuilder {
	return afadeBuilder.withOption("c", expr.Int(c))
}

func (afadeBuilder *implAfadeBuilder) CExpr(c expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("c", c)
}

func (afadeBuilder *implAfadeBuilder) Silence(silence float64) AfadeBuilder {
	return afadeBuilder.withOption("silence", expr.Double(silence))
}

func (afadeBuilder *implAfadeBuilder) SilenceExpr(silence expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("silence", silence)
}

func (afadeBuilder *implAfadeBuilder) Unity(unity float64) AfadeBuilder {
	return afadeBuilder.withOption("unity", expr.Double(unity))
}

func (afadeBuilder *implAfadeBuilder) UnityExpr(unity expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("unity", unity)
}

func (afadeBuilder *implAfadeBuilder) Enable(enable expr.Expr) AfadeBuilder {
	return afadeBuilder.withOption("enable", enable)
}
