// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// Tlut2Builder Compute and apply a lookup table from two successive frames.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#tlut2
type Tlut2Builder interface {
	filter.Filter
	// C0 set component #0 expression (default "x").
	C0(c0 expr.Expr) Tlut2Builder
	// C1 set component #1 expression (default "x").
	C1(c1 expr.Expr) Tlut2Builder
	// C2 set component #2 expression (default "x").
	C2(c2 expr.Expr) Tlut2Builder
	// C3 set component #3 expression (default "x").
	C3(c3 expr.Expr) Tlut2Builder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) Tlut2Builder
}

// Tlut2 creates a new Tlut2Builder to configure the "tlut2" filter.
func Tlut2(opts ...filter.Option) Tlut2Builder {
	return &implTlut2Builder{
		f: filter.New("tlut2", 1, opts...),
	}
}

type implTlut2Builder struct {
	f filter.Filter
}

func (tlut2Builder *implTlut2Builder) String() string {
	return tlut2Builder.f.String()
}

func (tlut2Builder *implTlut2Builder) Outputs() int {
	return tlut2Builder.f.Outputs()
}

func (tlut2Builder *implTlut2Builder) With(key string, value expr.Expr) filter.Filter {
	return tlut2Builder.withOption(key, value)
}

func (tlut2Builder *implTlut2Builder) withOption(key string, value expr.Expr) Tlut2Builder {
	bCopy := *tlut2Builder
	bCopy.f = tlut2Builder.f.With(key, value)
	return &bCopy
}

func (tlut2Builder *implTlut2Builder) C0(c0 expr.Expr) Tlut2Builder {
	return tlut2Builder.withOption("c0", c0)
}

func (tlut2Builder *implTlut2Builder) C1(c1 expr.Expr) Tlut2Builder {
	return tlut2Builder.withOption("c1", c1)
}

func (tlut2Builder *implTlut2Builder) C2(c2 expr.Expr) Tlut2Builder {
	return tlut2Builder.withOption("c2", c2)
}

func (tlut2Builder *implTlut2Builder) C3(c3 expr.Expr) Tlut2Builder {
	return tlut2Builder.withOption("c3", c3)
}

func (tlut2Builder *implTlut2Builder) Enable(enable expr.Expr) Tlut2Builder {
	return tlut2Builder.withOption("enable", enable)
}
