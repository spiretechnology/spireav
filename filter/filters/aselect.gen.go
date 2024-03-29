// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AudioSelectBuilder corresponds to the "aselect" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#aselect
type AudioSelectBuilder interface {
	filter.Filter
	// Expression sets the "e" option on the filter.
	Expression(expression expr.Expr) AudioSelectBuilder
	// ExpressionString sets the "e" option on the filter.
	ExpressionString(expression string) AudioSelectBuilder
}

// AudioSelect creates a new AudioSelectBuilder to configure the "aselect" filter.
func AudioSelect(outputs int, opts ...filter.Option) AudioSelectBuilder {
	f := filter.New("aselect", outputs, opts...)
	f = f.With("outputs", expr.Int(outputs))
	return &implAudioSelectBuilder{f: f}
}

type implAudioSelectBuilder struct {
	f filter.Filter
}

func (b *implAudioSelectBuilder) String() string {
	return b.f.String()
}

func (b *implAudioSelectBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implAudioSelectBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implAudioSelectBuilder) withOption(key string, value expr.Expr) AudioSelectBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implAudioSelectBuilder) Expression(expression expr.Expr) AudioSelectBuilder {
	return b.withOption("e", expression)
}

func (b *implAudioSelectBuilder) ExpressionString(expression string) AudioSelectBuilder {
	return b.withOption("e", expr.String(expression))
}
