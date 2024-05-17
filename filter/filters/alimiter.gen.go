// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AlimiterBuilder Audio lookahead limiter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#alimiter
type AlimiterBuilder interface {
	filter.Filter
	// LevelIn set input level (from 0.015625 to 64) (default 1).
	LevelIn(levelIn float64) AlimiterBuilder
	// LevelInExpr set input level (from 0.015625 to 64) (default 1).
	LevelInExpr(levelIn expr.Expr) AlimiterBuilder
	// LevelOut set output level (from 0.015625 to 64) (default 1).
	LevelOut(levelOut float64) AlimiterBuilder
	// LevelOutExpr set output level (from 0.015625 to 64) (default 1).
	LevelOutExpr(levelOut expr.Expr) AlimiterBuilder
	// Limit set limit (from 0.0625 to 1) (default 1).
	Limit(limit float64) AlimiterBuilder
	// LimitExpr set limit (from 0.0625 to 1) (default 1).
	LimitExpr(limit expr.Expr) AlimiterBuilder
	// Attack set attack (from 0.1 to 80) (default 5).
	Attack(attack float64) AlimiterBuilder
	// AttackExpr set attack (from 0.1 to 80) (default 5).
	AttackExpr(attack expr.Expr) AlimiterBuilder
	// Release set release (from 1 to 8000) (default 50).
	Release(release float64) AlimiterBuilder
	// ReleaseExpr set release (from 1 to 8000) (default 50).
	ReleaseExpr(release expr.Expr) AlimiterBuilder
	// Asc enable asc (default false).
	Asc(asc bool) AlimiterBuilder
	// AscExpr enable asc (default false).
	AscExpr(asc expr.Expr) AlimiterBuilder
	// AscLevel set asc level (from 0 to 1) (default 0.5).
	AscLevel(ascLevel float64) AlimiterBuilder
	// AscLevelExpr set asc level (from 0 to 1) (default 0.5).
	AscLevelExpr(ascLevel expr.Expr) AlimiterBuilder
	// Level auto level (default true).
	Level(level bool) AlimiterBuilder
	// LevelExpr auto level (default true).
	LevelExpr(level expr.Expr) AlimiterBuilder
	// Latency compensate delay (default false).
	Latency(latency bool) AlimiterBuilder
	// LatencyExpr compensate delay (default false).
	LatencyExpr(latency expr.Expr) AlimiterBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AlimiterBuilder
}

// Alimiter creates a new AlimiterBuilder to configure the "alimiter" filter.
func Alimiter(opts ...filter.Option) AlimiterBuilder {
	return &implAlimiterBuilder{
		f: filter.New("alimiter", 1, opts...),
	}
}

type implAlimiterBuilder struct {
	f filter.Filter
}

func (alimiterBuilder *implAlimiterBuilder) String() string {
	return alimiterBuilder.f.String()
}

func (alimiterBuilder *implAlimiterBuilder) Outputs() int {
	return alimiterBuilder.f.Outputs()
}

func (alimiterBuilder *implAlimiterBuilder) With(key string, value expr.Expr) filter.Filter {
	return alimiterBuilder.withOption(key, value)
}

func (alimiterBuilder *implAlimiterBuilder) withOption(key string, value expr.Expr) AlimiterBuilder {
	bCopy := *alimiterBuilder
	bCopy.f = alimiterBuilder.f.With(key, value)
	return &bCopy
}

func (alimiterBuilder *implAlimiterBuilder) LevelIn(levelIn float64) AlimiterBuilder {
	return alimiterBuilder.withOption("level_in", expr.Double(levelIn))
}

func (alimiterBuilder *implAlimiterBuilder) LevelInExpr(levelIn expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("level_in", levelIn)
}

func (alimiterBuilder *implAlimiterBuilder) LevelOut(levelOut float64) AlimiterBuilder {
	return alimiterBuilder.withOption("level_out", expr.Double(levelOut))
}

func (alimiterBuilder *implAlimiterBuilder) LevelOutExpr(levelOut expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("level_out", levelOut)
}

func (alimiterBuilder *implAlimiterBuilder) Limit(limit float64) AlimiterBuilder {
	return alimiterBuilder.withOption("limit", expr.Double(limit))
}

func (alimiterBuilder *implAlimiterBuilder) LimitExpr(limit expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("limit", limit)
}

func (alimiterBuilder *implAlimiterBuilder) Attack(attack float64) AlimiterBuilder {
	return alimiterBuilder.withOption("attack", expr.Double(attack))
}

func (alimiterBuilder *implAlimiterBuilder) AttackExpr(attack expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("attack", attack)
}

func (alimiterBuilder *implAlimiterBuilder) Release(release float64) AlimiterBuilder {
	return alimiterBuilder.withOption("release", expr.Double(release))
}

func (alimiterBuilder *implAlimiterBuilder) ReleaseExpr(release expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("release", release)
}

func (alimiterBuilder *implAlimiterBuilder) Asc(asc bool) AlimiterBuilder {
	return alimiterBuilder.withOption("asc", expr.Bool(asc))
}

func (alimiterBuilder *implAlimiterBuilder) AscExpr(asc expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("asc", asc)
}

func (alimiterBuilder *implAlimiterBuilder) AscLevel(ascLevel float64) AlimiterBuilder {
	return alimiterBuilder.withOption("asc_level", expr.Double(ascLevel))
}

func (alimiterBuilder *implAlimiterBuilder) AscLevelExpr(ascLevel expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("asc_level", ascLevel)
}

func (alimiterBuilder *implAlimiterBuilder) Level(level bool) AlimiterBuilder {
	return alimiterBuilder.withOption("level", expr.Bool(level))
}

func (alimiterBuilder *implAlimiterBuilder) LevelExpr(level expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("level", level)
}

func (alimiterBuilder *implAlimiterBuilder) Latency(latency bool) AlimiterBuilder {
	return alimiterBuilder.withOption("latency", expr.Bool(latency))
}

func (alimiterBuilder *implAlimiterBuilder) LatencyExpr(latency expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("latency", latency)
}

func (alimiterBuilder *implAlimiterBuilder) Enable(enable expr.Expr) AlimiterBuilder {
	return alimiterBuilder.withOption("enable", enable)
}
