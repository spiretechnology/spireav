// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ExtrastereoBuilder Increase difference between stereo audio channels.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#extrastereo
type ExtrastereoBuilder interface {
	filter.Filter
	// M set the difference coefficient (from -10 to 10) (default 2.5).
	M(m float32) ExtrastereoBuilder
	// MExpr set the difference coefficient (from -10 to 10) (default 2.5).
	MExpr(m expr.Expr) ExtrastereoBuilder
	// C enable clipping (default true).
	C(c bool) ExtrastereoBuilder
	// CExpr enable clipping (default true).
	CExpr(c expr.Expr) ExtrastereoBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ExtrastereoBuilder
}

// Extrastereo creates a new ExtrastereoBuilder to configure the "extrastereo" filter.
func Extrastereo(opts ...filter.Option) ExtrastereoBuilder {
	return &implExtrastereoBuilder{
		f: filter.New("extrastereo", 1, opts...),
	}
}

type implExtrastereoBuilder struct {
	f filter.Filter
}

func (extrastereoBuilder *implExtrastereoBuilder) String() string {
	return extrastereoBuilder.f.String()
}

func (extrastereoBuilder *implExtrastereoBuilder) Outputs() int {
	return extrastereoBuilder.f.Outputs()
}

func (extrastereoBuilder *implExtrastereoBuilder) With(key string, value expr.Expr) filter.Filter {
	return extrastereoBuilder.withOption(key, value)
}

func (extrastereoBuilder *implExtrastereoBuilder) withOption(key string, value expr.Expr) ExtrastereoBuilder {
	bCopy := *extrastereoBuilder
	bCopy.f = extrastereoBuilder.f.With(key, value)
	return &bCopy
}

func (extrastereoBuilder *implExtrastereoBuilder) M(m float32) ExtrastereoBuilder {
	return extrastereoBuilder.withOption("m", expr.Float(m))
}

func (extrastereoBuilder *implExtrastereoBuilder) MExpr(m expr.Expr) ExtrastereoBuilder {
	return extrastereoBuilder.withOption("m", m)
}

func (extrastereoBuilder *implExtrastereoBuilder) C(c bool) ExtrastereoBuilder {
	return extrastereoBuilder.withOption("c", expr.Bool(c))
}

func (extrastereoBuilder *implExtrastereoBuilder) CExpr(c expr.Expr) ExtrastereoBuilder {
	return extrastereoBuilder.withOption("c", c)
}

func (extrastereoBuilder *implExtrastereoBuilder) Enable(enable expr.Expr) ExtrastereoBuilder {
	return extrastereoBuilder.withOption("enable", enable)
}
