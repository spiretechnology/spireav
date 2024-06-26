// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// Lut2Builder Compute and apply a lookup table from two video inputs.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#lut2
type Lut2Builder interface {
	filter.Filter
	// C0 set component #0 expression (default "x").
	C0(c0 expr.Expr) Lut2Builder
	// C1 set component #1 expression (default "x").
	C1(c1 expr.Expr) Lut2Builder
	// C2 set component #2 expression (default "x").
	C2(c2 expr.Expr) Lut2Builder
	// C3 set component #3 expression (default "x").
	C3(c3 expr.Expr) Lut2Builder
	// D set output depth (from 0 to 16) (default 0).
	D(d int) Lut2Builder
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) Lut2Builder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) Lut2Builder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) Lut2Builder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) Lut2Builder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) Lut2Builder
}

// Lut2 creates a new Lut2Builder to configure the "lut2" filter.
func Lut2(opts ...filter.Option) Lut2Builder {
	return &implLut2Builder{
		f: filter.New("lut2", 1, opts...),
	}
}

type implLut2Builder struct {
	f filter.Filter
}

func (lut2Builder *implLut2Builder) String() string {
	return lut2Builder.f.String()
}

func (lut2Builder *implLut2Builder) Outputs() int {
	return lut2Builder.f.Outputs()
}

func (lut2Builder *implLut2Builder) With(key string, value expr.Expr) filter.Filter {
	return lut2Builder.withOption(key, value)
}

func (lut2Builder *implLut2Builder) withOption(key string, value expr.Expr) Lut2Builder {
	bCopy := *lut2Builder
	bCopy.f = lut2Builder.f.With(key, value)
	return &bCopy
}

func (lut2Builder *implLut2Builder) C0(c0 expr.Expr) Lut2Builder {
	return lut2Builder.withOption("c0", c0)
}

func (lut2Builder *implLut2Builder) C1(c1 expr.Expr) Lut2Builder {
	return lut2Builder.withOption("c1", c1)
}

func (lut2Builder *implLut2Builder) C2(c2 expr.Expr) Lut2Builder {
	return lut2Builder.withOption("c2", c2)
}

func (lut2Builder *implLut2Builder) C3(c3 expr.Expr) Lut2Builder {
	return lut2Builder.withOption("c3", c3)
}

func (lut2Builder *implLut2Builder) D(d int) Lut2Builder {
	return lut2Builder.withOption("d", expr.Int(d))
}

func (lut2Builder *implLut2Builder) EofAction(eofAction int) Lut2Builder {
	return lut2Builder.withOption("eof_action", expr.Int(eofAction))
}

func (lut2Builder *implLut2Builder) Shortest(shortest bool) Lut2Builder {
	return lut2Builder.withOption("shortest", expr.Bool(shortest))
}

func (lut2Builder *implLut2Builder) Repeatlast(repeatlast bool) Lut2Builder {
	return lut2Builder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (lut2Builder *implLut2Builder) TsSyncMode(tsSyncMode int) Lut2Builder {
	return lut2Builder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (lut2Builder *implLut2Builder) Enable(enable expr.Expr) Lut2Builder {
	return lut2Builder.withOption("enable", enable)
}
