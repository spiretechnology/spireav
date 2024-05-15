// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CellautoBuilder Create pattern generated by an elementary cellular automaton.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#cellauto
type CellautoBuilder interface {
	filter.Filter
	// Filename read initial pattern from file.
	Filename(filename string) CellautoBuilder
	// F read initial pattern from file.
	F(f string) CellautoBuilder
	// Pattern set initial pattern.
	Pattern(pattern string) CellautoBuilder
	// P set initial pattern.
	P(p string) CellautoBuilder
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) CellautoBuilder
	// R set video rate (default "25").
	R(r expr.Rational) CellautoBuilder
	// Size set video size.
	Size(size expr.Size) CellautoBuilder
	// S set video size.
	S(s expr.Size) CellautoBuilder
	// Rule set rule (from 0 to 255) (default 110).
	Rule(rule int) CellautoBuilder
	// RandomFillRatio set fill ratio for filling initial grid randomly (from 0 to 1) (default 0.618034).
	RandomFillRatio(randomFillRatio float64) CellautoBuilder
	// Ratio set fill ratio for filling initial grid randomly (from 0 to 1) (default 0.618034).
	Ratio(ratio float64) CellautoBuilder
	// RandomSeed set the seed for filling the initial grid randomly (from -1 to UINT32_MAX) (default -1).
	RandomSeed(randomSeed int64) CellautoBuilder
	// Seed set the seed for filling the initial grid randomly (from -1 to UINT32_MAX) (default -1).
	Seed(seed int64) CellautoBuilder
	// Scroll scroll pattern downward (default true).
	Scroll(scroll bool) CellautoBuilder
	// StartFull start filling the whole video (default false).
	StartFull(startFull bool) CellautoBuilder
	// Full start filling the whole video (default true).
	Full(full bool) CellautoBuilder
	// Stitch stitch boundaries (default true).
	Stitch(stitch bool) CellautoBuilder
}

// Cellauto creates a new CellautoBuilder to configure the "cellauto" filter.
func Cellauto(opts ...filter.Option) CellautoBuilder {
	return &implCellautoBuilder{
		f: filter.New("cellauto", 1, opts...),
	}
}

type implCellautoBuilder struct {
	f filter.Filter
}

func (cellautoBuilder *implCellautoBuilder) String() string {
	return cellautoBuilder.f.String()
}

func (cellautoBuilder *implCellautoBuilder) Outputs() int {
	return cellautoBuilder.f.Outputs()
}

func (cellautoBuilder *implCellautoBuilder) With(key string, value expr.Expr) filter.Filter {
	return cellautoBuilder.withOption(key, value)
}

func (cellautoBuilder *implCellautoBuilder) withOption(key string, value expr.Expr) CellautoBuilder {
	bCopy := *cellautoBuilder
	bCopy.f = cellautoBuilder.f.With(key, value)
	return &bCopy
}

func (cellautoBuilder *implCellautoBuilder) Filename(filename string) CellautoBuilder {
	return cellautoBuilder.withOption("filename", expr.String(filename))
}

func (cellautoBuilder *implCellautoBuilder) F(f string) CellautoBuilder {
	return cellautoBuilder.withOption("f", expr.String(f))
}

func (cellautoBuilder *implCellautoBuilder) Pattern(pattern string) CellautoBuilder {
	return cellautoBuilder.withOption("pattern", expr.String(pattern))
}

func (cellautoBuilder *implCellautoBuilder) P(p string) CellautoBuilder {
	return cellautoBuilder.withOption("p", expr.String(p))
}

func (cellautoBuilder *implCellautoBuilder) Rate(rate expr.Rational) CellautoBuilder {
	return cellautoBuilder.withOption("rate", rate)
}

func (cellautoBuilder *implCellautoBuilder) R(r expr.Rational) CellautoBuilder {
	return cellautoBuilder.withOption("r", r)
}

func (cellautoBuilder *implCellautoBuilder) Size(size expr.Size) CellautoBuilder {
	return cellautoBuilder.withOption("size", size)
}

func (cellautoBuilder *implCellautoBuilder) S(s expr.Size) CellautoBuilder {
	return cellautoBuilder.withOption("s", s)
}

func (cellautoBuilder *implCellautoBuilder) Rule(rule int) CellautoBuilder {
	return cellautoBuilder.withOption("rule", expr.Int(rule))
}

func (cellautoBuilder *implCellautoBuilder) RandomFillRatio(randomFillRatio float64) CellautoBuilder {
	return cellautoBuilder.withOption("random_fill_ratio", expr.Double(randomFillRatio))
}

func (cellautoBuilder *implCellautoBuilder) Ratio(ratio float64) CellautoBuilder {
	return cellautoBuilder.withOption("ratio", expr.Double(ratio))
}

func (cellautoBuilder *implCellautoBuilder) RandomSeed(randomSeed int64) CellautoBuilder {
	return cellautoBuilder.withOption("random_seed", expr.Int64(randomSeed))
}

func (cellautoBuilder *implCellautoBuilder) Seed(seed int64) CellautoBuilder {
	return cellautoBuilder.withOption("seed", expr.Int64(seed))
}

func (cellautoBuilder *implCellautoBuilder) Scroll(scroll bool) CellautoBuilder {
	return cellautoBuilder.withOption("scroll", expr.Bool(scroll))
}

func (cellautoBuilder *implCellautoBuilder) StartFull(startFull bool) CellautoBuilder {
	return cellautoBuilder.withOption("start_full", expr.Bool(startFull))
}

func (cellautoBuilder *implCellautoBuilder) Full(full bool) CellautoBuilder {
	return cellautoBuilder.withOption("full", expr.Bool(full))
}

func (cellautoBuilder *implCellautoBuilder) Stitch(stitch bool) CellautoBuilder {
	return cellautoBuilder.withOption("stitch", expr.Bool(stitch))
}
