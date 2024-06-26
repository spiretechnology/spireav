// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// VifBuilder Calculate the VIF between two video streams.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#vif
type VifBuilder interface {
	filter.Filter
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) VifBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) VifBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) VifBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) VifBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) VifBuilder
}

// Vif creates a new VifBuilder to configure the "vif" filter.
func Vif(opts ...filter.Option) VifBuilder {
	return &implVifBuilder{
		f: filter.New("vif", 1, opts...),
	}
}

type implVifBuilder struct {
	f filter.Filter
}

func (vifBuilder *implVifBuilder) String() string {
	return vifBuilder.f.String()
}

func (vifBuilder *implVifBuilder) Outputs() int {
	return vifBuilder.f.Outputs()
}

func (vifBuilder *implVifBuilder) With(key string, value expr.Expr) filter.Filter {
	return vifBuilder.withOption(key, value)
}

func (vifBuilder *implVifBuilder) withOption(key string, value expr.Expr) VifBuilder {
	bCopy := *vifBuilder
	bCopy.f = vifBuilder.f.With(key, value)
	return &bCopy
}

func (vifBuilder *implVifBuilder) EofAction(eofAction int) VifBuilder {
	return vifBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (vifBuilder *implVifBuilder) Shortest(shortest bool) VifBuilder {
	return vifBuilder.withOption("shortest", expr.Bool(shortest))
}

func (vifBuilder *implVifBuilder) Repeatlast(repeatlast bool) VifBuilder {
	return vifBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (vifBuilder *implVifBuilder) TsSyncMode(tsSyncMode int) VifBuilder {
	return vifBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (vifBuilder *implVifBuilder) Enable(enable expr.Expr) VifBuilder {
	return vifBuilder.withOption("enable", enable)
}
