// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CorrBuilder Calculate the correlation between two video streams.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#corr
type CorrBuilder interface {
	filter.Filter
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) CorrBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) CorrBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) CorrBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) CorrBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) CorrBuilder
}

// Corr creates a new CorrBuilder to configure the "corr" filter.
func Corr(opts ...filter.Option) CorrBuilder {
	return &implCorrBuilder{
		f: filter.New("corr", 1, opts...),
	}
}

type implCorrBuilder struct {
	f filter.Filter
}

func (corrBuilder *implCorrBuilder) String() string {
	return corrBuilder.f.String()
}

func (corrBuilder *implCorrBuilder) Outputs() int {
	return corrBuilder.f.Outputs()
}

func (corrBuilder *implCorrBuilder) With(key string, value expr.Expr) filter.Filter {
	return corrBuilder.withOption(key, value)
}

func (corrBuilder *implCorrBuilder) withOption(key string, value expr.Expr) CorrBuilder {
	bCopy := *corrBuilder
	bCopy.f = corrBuilder.f.With(key, value)
	return &bCopy
}

func (corrBuilder *implCorrBuilder) EofAction(eofAction int) CorrBuilder {
	return corrBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (corrBuilder *implCorrBuilder) Shortest(shortest bool) CorrBuilder {
	return corrBuilder.withOption("shortest", expr.Bool(shortest))
}

func (corrBuilder *implCorrBuilder) Repeatlast(repeatlast bool) CorrBuilder {
	return corrBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (corrBuilder *implCorrBuilder) TsSyncMode(tsSyncMode int) CorrBuilder {
	return corrBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (corrBuilder *implCorrBuilder) Enable(enable expr.Expr) CorrBuilder {
	return corrBuilder.withOption("enable", enable)
}
