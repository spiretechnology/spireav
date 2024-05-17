// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SuperequalizerBuilder Apply 18 band equalization filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#superequalizer
type SuperequalizerBuilder interface {
	filter.Filter
	// With1b set 65Hz band gain (from 0 to 20) (default 1).
	With1b(val1b float32) SuperequalizerBuilder
	// With2b set 92Hz band gain (from 0 to 20) (default 1).
	With2b(val2b float32) SuperequalizerBuilder
	// With3b set 131Hz band gain (from 0 to 20) (default 1).
	With3b(val3b float32) SuperequalizerBuilder
	// With4b set 185Hz band gain (from 0 to 20) (default 1).
	With4b(val4b float32) SuperequalizerBuilder
	// With5b set 262Hz band gain (from 0 to 20) (default 1).
	With5b(val5b float32) SuperequalizerBuilder
	// With6b set 370Hz band gain (from 0 to 20) (default 1).
	With6b(val6b float32) SuperequalizerBuilder
	// With7b set 523Hz band gain (from 0 to 20) (default 1).
	With7b(val7b float32) SuperequalizerBuilder
	// With8b set 740Hz band gain (from 0 to 20) (default 1).
	With8b(val8b float32) SuperequalizerBuilder
	// With9b set 1047Hz band gain (from 0 to 20) (default 1).
	With9b(val9b float32) SuperequalizerBuilder
	// With10b set 1480Hz band gain (from 0 to 20) (default 1).
	With10b(val10b float32) SuperequalizerBuilder
	// With11b set 2093Hz band gain (from 0 to 20) (default 1).
	With11b(val11b float32) SuperequalizerBuilder
	// With12b set 2960Hz band gain (from 0 to 20) (default 1).
	With12b(val12b float32) SuperequalizerBuilder
	// With13b set 4186Hz band gain (from 0 to 20) (default 1).
	With13b(val13b float32) SuperequalizerBuilder
	// With14b set 5920Hz band gain (from 0 to 20) (default 1).
	With14b(val14b float32) SuperequalizerBuilder
	// With15b set 8372Hz band gain (from 0 to 20) (default 1).
	With15b(val15b float32) SuperequalizerBuilder
	// With16b set 11840Hz band gain (from 0 to 20) (default 1).
	With16b(val16b float32) SuperequalizerBuilder
	// With17b set 16744Hz band gain (from 0 to 20) (default 1).
	With17b(val17b float32) SuperequalizerBuilder
	// With18b set 20000Hz band gain (from 0 to 20) (default 1).
	With18b(val18b float32) SuperequalizerBuilder
}

// Superequalizer creates a new SuperequalizerBuilder to configure the "superequalizer" filter.
func Superequalizer(opts ...filter.Option) SuperequalizerBuilder {
	return &implSuperequalizerBuilder{
		f: filter.New("superequalizer", 1, opts...),
	}
}

type implSuperequalizerBuilder struct {
	f filter.Filter
}

func (superequalizerBuilder *implSuperequalizerBuilder) String() string {
	return superequalizerBuilder.f.String()
}

func (superequalizerBuilder *implSuperequalizerBuilder) Outputs() int {
	return superequalizerBuilder.f.Outputs()
}

func (superequalizerBuilder *implSuperequalizerBuilder) With(key string, value expr.Expr) filter.Filter {
	return superequalizerBuilder.withOption(key, value)
}

func (superequalizerBuilder *implSuperequalizerBuilder) withOption(key string, value expr.Expr) SuperequalizerBuilder {
	bCopy := *superequalizerBuilder
	bCopy.f = superequalizerBuilder.f.With(key, value)
	return &bCopy
}

func (superequalizerBuilder *implSuperequalizerBuilder) With1b(val1b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("1b", expr.Float(val1b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With2b(val2b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("2b", expr.Float(val2b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With3b(val3b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("3b", expr.Float(val3b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With4b(val4b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("4b", expr.Float(val4b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With5b(val5b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("5b", expr.Float(val5b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With6b(val6b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("6b", expr.Float(val6b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With7b(val7b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("7b", expr.Float(val7b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With8b(val8b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("8b", expr.Float(val8b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With9b(val9b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("9b", expr.Float(val9b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With10b(val10b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("10b", expr.Float(val10b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With11b(val11b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("11b", expr.Float(val11b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With12b(val12b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("12b", expr.Float(val12b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With13b(val13b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("13b", expr.Float(val13b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With14b(val14b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("14b", expr.Float(val14b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With15b(val15b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("15b", expr.Float(val15b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With16b(val16b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("16b", expr.Float(val16b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With17b(val17b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("17b", expr.Float(val17b))
}

func (superequalizerBuilder *implSuperequalizerBuilder) With18b(val18b float32) SuperequalizerBuilder {
	return superequalizerBuilder.withOption("18b", expr.Float(val18b))
}