// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// HaldclutBuilder Adjust colors using a Hald CLUT.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#haldclut
type HaldclutBuilder interface {
	filter.Filter
	// Clut when to process CLUT (from 0 to 1) (default all).
	Clut(clut int) HaldclutBuilder
	// ClutExpr when to process CLUT (from 0 to 1) (default all).
	ClutExpr(clut expr.Expr) HaldclutBuilder
	// Interp select interpolation mode (from 0 to 4) (default tetrahedral).
	Interp(interp int) HaldclutBuilder
	// InterpExpr select interpolation mode (from 0 to 4) (default tetrahedral).
	InterpExpr(interp expr.Expr) HaldclutBuilder
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) HaldclutBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) HaldclutBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) HaldclutBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) HaldclutBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) HaldclutBuilder
}

// Haldclut creates a new HaldclutBuilder to configure the "haldclut" filter.
func Haldclut(opts ...filter.Option) HaldclutBuilder {
	return &implHaldclutBuilder{
		f: filter.New("haldclut", 1, opts...),
	}
}

type implHaldclutBuilder struct {
	f filter.Filter
}

func (haldclutBuilder *implHaldclutBuilder) String() string {
	return haldclutBuilder.f.String()
}

func (haldclutBuilder *implHaldclutBuilder) Outputs() int {
	return haldclutBuilder.f.Outputs()
}

func (haldclutBuilder *implHaldclutBuilder) With(key string, value expr.Expr) filter.Filter {
	return haldclutBuilder.withOption(key, value)
}

func (haldclutBuilder *implHaldclutBuilder) withOption(key string, value expr.Expr) HaldclutBuilder {
	bCopy := *haldclutBuilder
	bCopy.f = haldclutBuilder.f.With(key, value)
	return &bCopy
}

func (haldclutBuilder *implHaldclutBuilder) Clut(clut int) HaldclutBuilder {
	return haldclutBuilder.withOption("clut", expr.Int(clut))
}

func (haldclutBuilder *implHaldclutBuilder) ClutExpr(clut expr.Expr) HaldclutBuilder {
	return haldclutBuilder.withOption("clut", clut)
}

func (haldclutBuilder *implHaldclutBuilder) Interp(interp int) HaldclutBuilder {
	return haldclutBuilder.withOption("interp", expr.Int(interp))
}

func (haldclutBuilder *implHaldclutBuilder) InterpExpr(interp expr.Expr) HaldclutBuilder {
	return haldclutBuilder.withOption("interp", interp)
}

func (haldclutBuilder *implHaldclutBuilder) EofAction(eofAction int) HaldclutBuilder {
	return haldclutBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (haldclutBuilder *implHaldclutBuilder) Shortest(shortest bool) HaldclutBuilder {
	return haldclutBuilder.withOption("shortest", expr.Bool(shortest))
}

func (haldclutBuilder *implHaldclutBuilder) Repeatlast(repeatlast bool) HaldclutBuilder {
	return haldclutBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (haldclutBuilder *implHaldclutBuilder) TsSyncMode(tsSyncMode int) HaldclutBuilder {
	return haldclutBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (haldclutBuilder *implHaldclutBuilder) Enable(enable expr.Expr) HaldclutBuilder {
	return haldclutBuilder.withOption("enable", enable)
}
