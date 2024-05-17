// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// IdentityBuilder Calculate the Identity between two video streams.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#identity
type IdentityBuilder interface {
	filter.Filter
	// EofAction Action to take when encountering EOF from secondary input  (from 0 to 2) (default repeat).
	EofAction(eofAction int) IdentityBuilder
	// Shortest force termination when the shortest input terminates (default false).
	Shortest(shortest bool) IdentityBuilder
	// Repeatlast extend last frame of secondary streams beyond EOF (default true).
	Repeatlast(repeatlast bool) IdentityBuilder
	// TsSyncMode How strictly to sync streams based on secondary input timestamps (from 0 to 1) (default default).
	TsSyncMode(tsSyncMode int) IdentityBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) IdentityBuilder
}

// Identity creates a new IdentityBuilder to configure the "identity" filter.
func Identity(opts ...filter.Option) IdentityBuilder {
	return &implIdentityBuilder{
		f: filter.New("identity", 1, opts...),
	}
}

type implIdentityBuilder struct {
	f filter.Filter
}

func (identityBuilder *implIdentityBuilder) String() string {
	return identityBuilder.f.String()
}

func (identityBuilder *implIdentityBuilder) Outputs() int {
	return identityBuilder.f.Outputs()
}

func (identityBuilder *implIdentityBuilder) With(key string, value expr.Expr) filter.Filter {
	return identityBuilder.withOption(key, value)
}

func (identityBuilder *implIdentityBuilder) withOption(key string, value expr.Expr) IdentityBuilder {
	bCopy := *identityBuilder
	bCopy.f = identityBuilder.f.With(key, value)
	return &bCopy
}

func (identityBuilder *implIdentityBuilder) EofAction(eofAction int) IdentityBuilder {
	return identityBuilder.withOption("eof_action", expr.Int(eofAction))
}

func (identityBuilder *implIdentityBuilder) Shortest(shortest bool) IdentityBuilder {
	return identityBuilder.withOption("shortest", expr.Bool(shortest))
}

func (identityBuilder *implIdentityBuilder) Repeatlast(repeatlast bool) IdentityBuilder {
	return identityBuilder.withOption("repeatlast", expr.Bool(repeatlast))
}

func (identityBuilder *implIdentityBuilder) TsSyncMode(tsSyncMode int) IdentityBuilder {
	return identityBuilder.withOption("ts_sync_mode", expr.Int(tsSyncMode))
}

func (identityBuilder *implIdentityBuilder) Enable(enable expr.Expr) IdentityBuilder {
	return identityBuilder.withOption("enable", enable)
}
