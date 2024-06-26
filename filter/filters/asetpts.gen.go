// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AsetptsBuilder Set PTS for the output audio frame.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#asetpts
type AsetptsBuilder interface {
	filter.Filter
	// Expr Expression determining the frame timestamp (default "PTS").
	Expr(expression string) AsetptsBuilder
	// ExprExpr Expression determining the frame timestamp (default "PTS").
	ExprExpr(expression expr.Expr) AsetptsBuilder
}

// Asetpts creates a new AsetptsBuilder to configure the "asetpts" filter.
func Asetpts(opts ...filter.Option) AsetptsBuilder {
	return &implAsetptsBuilder{
		f: filter.New("asetpts", 1, opts...),
	}
}

type implAsetptsBuilder struct {
	f filter.Filter
}

func (asetptsBuilder *implAsetptsBuilder) String() string {
	return asetptsBuilder.f.String()
}

func (asetptsBuilder *implAsetptsBuilder) Outputs() int {
	return asetptsBuilder.f.Outputs()
}

func (asetptsBuilder *implAsetptsBuilder) With(key string, value expr.Expr) filter.Filter {
	return asetptsBuilder.withOption(key, value)
}

func (asetptsBuilder *implAsetptsBuilder) withOption(key string, value expr.Expr) AsetptsBuilder {
	bCopy := *asetptsBuilder
	bCopy.f = asetptsBuilder.f.With(key, value)
	return &bCopy
}

func (asetptsBuilder *implAsetptsBuilder) Expr(expression string) AsetptsBuilder {
	return asetptsBuilder.withOption("expr", expr.String(expression))
}

func (asetptsBuilder *implAsetptsBuilder) ExprExpr(expression expr.Expr) AsetptsBuilder {
	return asetptsBuilder.withOption("expr", expression)
}
