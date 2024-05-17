// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// PermsBuilder Set permissions for the output video frame.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#perms
type PermsBuilder interface {
	filter.Filter
	// Mode select permissions mode (from 0 to 4) (default none).
	Mode(mode int) PermsBuilder
	// ModeExpr select permissions mode (from 0 to 4) (default none).
	ModeExpr(mode expr.Expr) PermsBuilder
	// Seed set the seed for the random mode (from -1 to UINT32_MAX) (default -1).
	Seed(seed int64) PermsBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) PermsBuilder
}

// Perms creates a new PermsBuilder to configure the "perms" filter.
func Perms(opts ...filter.Option) PermsBuilder {
	return &implPermsBuilder{
		f: filter.New("perms", 1, opts...),
	}
}

type implPermsBuilder struct {
	f filter.Filter
}

func (permsBuilder *implPermsBuilder) String() string {
	return permsBuilder.f.String()
}

func (permsBuilder *implPermsBuilder) Outputs() int {
	return permsBuilder.f.Outputs()
}

func (permsBuilder *implPermsBuilder) With(key string, value expr.Expr) filter.Filter {
	return permsBuilder.withOption(key, value)
}

func (permsBuilder *implPermsBuilder) withOption(key string, value expr.Expr) PermsBuilder {
	bCopy := *permsBuilder
	bCopy.f = permsBuilder.f.With(key, value)
	return &bCopy
}

func (permsBuilder *implPermsBuilder) Mode(mode int) PermsBuilder {
	return permsBuilder.withOption("mode", expr.Int(mode))
}

func (permsBuilder *implPermsBuilder) ModeExpr(mode expr.Expr) PermsBuilder {
	return permsBuilder.withOption("mode", mode)
}

func (permsBuilder *implPermsBuilder) Seed(seed int64) PermsBuilder {
	return permsBuilder.withOption("seed", expr.Int64(seed))
}

func (permsBuilder *implPermsBuilder) Enable(enable expr.Expr) PermsBuilder {
	return permsBuilder.withOption("enable", enable)
}