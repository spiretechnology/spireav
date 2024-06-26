// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// W3fdifBuilder Apply Martin Weston three field deinterlace.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#w3fdif
type W3fdifBuilder interface {
	filter.Filter
	// Filter specify the filter (from 0 to 1) (default complex).
	Filter(filter int) W3fdifBuilder
	// FilterExpr specify the filter (from 0 to 1) (default complex).
	FilterExpr(filter expr.Expr) W3fdifBuilder
	// Mode specify the interlacing mode (from 0 to 1) (default field).
	Mode(mode int) W3fdifBuilder
	// ModeExpr specify the interlacing mode (from 0 to 1) (default field).
	ModeExpr(mode expr.Expr) W3fdifBuilder
	// Parity specify the assumed picture field parity (from -1 to 1) (default auto).
	Parity(parity int) W3fdifBuilder
	// ParityExpr specify the assumed picture field parity (from -1 to 1) (default auto).
	ParityExpr(parity expr.Expr) W3fdifBuilder
	// Deint specify which frames to deinterlace (from 0 to 1) (default all).
	Deint(deint int) W3fdifBuilder
	// DeintExpr specify which frames to deinterlace (from 0 to 1) (default all).
	DeintExpr(deint expr.Expr) W3fdifBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) W3fdifBuilder
}

// W3fdif creates a new W3fdifBuilder to configure the "w3fdif" filter.
func W3fdif(opts ...filter.Option) W3fdifBuilder {
	return &implW3fdifBuilder{
		f: filter.New("w3fdif", 1, opts...),
	}
}

type implW3fdifBuilder struct {
	f filter.Filter
}

func (w3fdifBuilder *implW3fdifBuilder) String() string {
	return w3fdifBuilder.f.String()
}

func (w3fdifBuilder *implW3fdifBuilder) Outputs() int {
	return w3fdifBuilder.f.Outputs()
}

func (w3fdifBuilder *implW3fdifBuilder) With(key string, value expr.Expr) filter.Filter {
	return w3fdifBuilder.withOption(key, value)
}

func (w3fdifBuilder *implW3fdifBuilder) withOption(key string, value expr.Expr) W3fdifBuilder {
	bCopy := *w3fdifBuilder
	bCopy.f = w3fdifBuilder.f.With(key, value)
	return &bCopy
}

func (w3fdifBuilder *implW3fdifBuilder) Filter(filter int) W3fdifBuilder {
	return w3fdifBuilder.withOption("filter", expr.Int(filter))
}

func (w3fdifBuilder *implW3fdifBuilder) FilterExpr(filter expr.Expr) W3fdifBuilder {
	return w3fdifBuilder.withOption("filter", filter)
}

func (w3fdifBuilder *implW3fdifBuilder) Mode(mode int) W3fdifBuilder {
	return w3fdifBuilder.withOption("mode", expr.Int(mode))
}

func (w3fdifBuilder *implW3fdifBuilder) ModeExpr(mode expr.Expr) W3fdifBuilder {
	return w3fdifBuilder.withOption("mode", mode)
}

func (w3fdifBuilder *implW3fdifBuilder) Parity(parity int) W3fdifBuilder {
	return w3fdifBuilder.withOption("parity", expr.Int(parity))
}

func (w3fdifBuilder *implW3fdifBuilder) ParityExpr(parity expr.Expr) W3fdifBuilder {
	return w3fdifBuilder.withOption("parity", parity)
}

func (w3fdifBuilder *implW3fdifBuilder) Deint(deint int) W3fdifBuilder {
	return w3fdifBuilder.withOption("deint", expr.Int(deint))
}

func (w3fdifBuilder *implW3fdifBuilder) DeintExpr(deint expr.Expr) W3fdifBuilder {
	return w3fdifBuilder.withOption("deint", deint)
}

func (w3fdifBuilder *implW3fdifBuilder) Enable(enable expr.Expr) W3fdifBuilder {
	return w3fdifBuilder.withOption("enable", enable)
}
