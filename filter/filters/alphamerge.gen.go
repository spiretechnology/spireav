// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AlphamergeBuilder Copy the luma value of the second input into the alpha channel of the first input.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#alphamerge
type AlphamergeBuilder interface {
	filter.Filter
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) AlphamergeBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) AlphamergeBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) AlphamergeBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) AlphamergeBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AlphamergeBuilder
}

// Alphamerge creates a new AlphamergeBuilder to configure the "alphamerge" filter.
func Alphamerge(opts ...filter.Option) AlphamergeBuilder {
	return &implAlphamergeBuilder{
		f: filter.New("alphamerge", 1, opts...),
	}
}

type implAlphamergeBuilder struct {
	f filter.Filter
}

func (alphamergeBuilder *implAlphamergeBuilder) String() string {
	return alphamergeBuilder.f.String()
}

func (alphamergeBuilder *implAlphamergeBuilder) Outputs() int {
	return alphamergeBuilder.f.Outputs()
}

func (alphamergeBuilder *implAlphamergeBuilder) With(key string, value expr.Expr) filter.Filter {
	return alphamergeBuilder.withOption(key, value)
}

func (alphamergeBuilder *implAlphamergeBuilder) withOption(key string, value expr.Expr) AlphamergeBuilder {
	bCopy := *alphamergeBuilder
	bCopy.f = alphamergeBuilder.f.With(key, value)
	return &bCopy
}

func (alphamergeBuilder *implAlphamergeBuilder) EofAction(eofAction int) AlphamergeBuilder {
	return alphamergeBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (alphamergeBuilder *implAlphamergeBuilder) Shortest(shortest bool) AlphamergeBuilder {
	return alphamergeBuilder.withOption("shortest", expr.Bool(shortest))
}

func (alphamergeBuilder *implAlphamergeBuilder) Repeatlast(repeatlast bool) AlphamergeBuilder {
	return alphamergeBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (alphamergeBuilder *implAlphamergeBuilder) TsSyncMode(tsSyncMode int) AlphamergeBuilder {
	return alphamergeBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (alphamergeBuilder *implAlphamergeBuilder) Enable(enable expr.Expr) AlphamergeBuilder {
	return alphamergeBuilder.withOption("enable", enable)
}
