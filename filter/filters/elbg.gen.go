// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ElbgBuilder Apply posterize effect, using the ELBG algorithm.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#elbg
type ElbgBuilder interface {
	filter.Filter
	// CodebookLength set codebook length (from 1 to INT_MAX) (default 256).
	CodebookLength(codebookLength int) ElbgBuilder
	// L set codebook length (from 1 to INT_MAX) (default 256).
	L(l int) ElbgBuilder
	// NbSteps set max number of steps used to compute the mapping (from 1 to INT_MAX) (default 1).
	NbSteps(nbSteps int) ElbgBuilder
	// N set max number of steps used to compute the mapping (from 1 to INT_MAX) (default 1).
	N(n int) ElbgBuilder
	// Seed set the random seed (from -1 to UINT32_MAX) (default -1).
	Seed(seed int64) ElbgBuilder
	// S set the random seed (from -1 to UINT32_MAX) (default -1).
	S(s int64) ElbgBuilder
	// Pal8 set the pal8 output (default false).
	Pal8(pal8 bool) ElbgBuilder
	// UseAlpha use alpha channel for mapping (default false).
	UseAlpha(useAlpha bool) ElbgBuilder
}

// Elbg creates a new ElbgBuilder to configure the "elbg" filter.
func Elbg(opts ...filter.Option) ElbgBuilder {
	return &implElbgBuilder{
		f: filter.New("elbg", 1, opts...),
	}
}

type implElbgBuilder struct {
	f filter.Filter
}

func (elbgBuilder *implElbgBuilder) String() string {
	return elbgBuilder.f.String()
}

func (elbgBuilder *implElbgBuilder) Outputs() int {
	return elbgBuilder.f.Outputs()
}

func (elbgBuilder *implElbgBuilder) With(key string, value expr.Expr) filter.Filter {
	return elbgBuilder.withOption(key, value)
}

func (elbgBuilder *implElbgBuilder) withOption(key string, value expr.Expr) ElbgBuilder {
	bCopy := *elbgBuilder
	bCopy.f = elbgBuilder.f.With(key, value)
	return &bCopy
}

func (elbgBuilder *implElbgBuilder) CodebookLength(codebookLength int) ElbgBuilder {
	return elbgBuilder.withOption("codebook_length", expr.Int(codebookLength))
}

func (elbgBuilder *implElbgBuilder) L(l int) ElbgBuilder {
	return elbgBuilder.withOption("l", expr.Int(l))
}

func (elbgBuilder *implElbgBuilder) NbSteps(nbSteps int) ElbgBuilder {
	return elbgBuilder.withOption("nb_steps", expr.Int(nbSteps))
}

func (elbgBuilder *implElbgBuilder) N(n int) ElbgBuilder {
	return elbgBuilder.withOption("n", expr.Int(n))
}

func (elbgBuilder *implElbgBuilder) Seed(seed int64) ElbgBuilder {
	return elbgBuilder.withOption("seed", expr.Int64(seed))
}

func (elbgBuilder *implElbgBuilder) S(s int64) ElbgBuilder {
	return elbgBuilder.withOption("s", expr.Int64(s))
}

func (elbgBuilder *implElbgBuilder) Pal8(pal8 bool) ElbgBuilder {
	return elbgBuilder.withOption("pal8", expr.Bool(pal8))
}

func (elbgBuilder *implElbgBuilder) UseAlpha(useAlpha bool) ElbgBuilder {
	return elbgBuilder.withOption("use_alpha", expr.Bool(useAlpha))
}