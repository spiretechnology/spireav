// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PullupBuilder Pullup from field sequence to frames.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#pullup
type PullupBuilder interface {
	filter.Filter
	// Jl set left junk size (from 0 to INT_MAX) (default 1).
	Jl(jl int) PullupBuilder
	// Jr set right junk size (from 0 to INT_MAX) (default 1).
	Jr(jr int) PullupBuilder
	// Jt set top junk size (from 1 to INT_MAX) (default 4).
	Jt(jt int) PullupBuilder
	// Jb set bottom junk size (from 1 to INT_MAX) (default 4).
	Jb(jb int) PullupBuilder
	// Sb set strict breaks (default false).
	Sb(sb bool) PullupBuilder
	// Mp set metric plane (from 0 to 2) (default y).
	Mp(mp int) PullupBuilder
}

// Pullup creates a new PullupBuilder to configure the "pullup" filter.
func Pullup(opts ...filter.Option) PullupBuilder {
	return &implPullupBuilder{
		f: filter.New("pullup", 1, opts...),
	}
}

type implPullupBuilder struct {
	f filter.Filter
}

func (pullupBuilder *implPullupBuilder) String() string {
	return pullupBuilder.f.String()
}

func (pullupBuilder *implPullupBuilder) Outputs() int {
	return pullupBuilder.f.Outputs()
}

func (pullupBuilder *implPullupBuilder) With(key string, value expr.Expr) filter.Filter {
	return pullupBuilder.withOption(key, value)
}

func (pullupBuilder *implPullupBuilder) withOption(key string, value expr.Expr) PullupBuilder {
	bCopy := *pullupBuilder
	bCopy.f = pullupBuilder.f.With(key, value)
	return &bCopy
}

func (pullupBuilder *implPullupBuilder) Jl(jl int) PullupBuilder {
	return pullupBuilder.withOption("jl", expr.Int(jl))
}

func (pullupBuilder *implPullupBuilder) Jr(jr int) PullupBuilder {
	return pullupBuilder.withOption("jr", expr.Int(jr))
}

func (pullupBuilder *implPullupBuilder) Jt(jt int) PullupBuilder {
	return pullupBuilder.withOption("jt", expr.Int(jt))
}

func (pullupBuilder *implPullupBuilder) Jb(jb int) PullupBuilder {
	return pullupBuilder.withOption("jb", expr.Int(jb))
}

func (pullupBuilder *implPullupBuilder) Sb(sb bool) PullupBuilder {
	return pullupBuilder.withOption("sb", expr.Bool(sb))
}

func (pullupBuilder *implPullupBuilder) Mp(mp int) PullupBuilder {
	return pullupBuilder.withOption("mp", expr.Int(mp))
}
