// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// YadifBuilder Deinterlace the input image.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#yadif
type YadifBuilder interface {
	filter.Filter
	// Mode specify the interlacing mode (from 0 to 3) (default send_frame).
	Mode(mode int) YadifBuilder
	// Parity specify the assumed picture field parity (from -1 to 1) (default auto).
	Parity(parity int) YadifBuilder
	// Deint specify which frames to deinterlace (from 0 to 1) (default all).
	Deint(deint int) YadifBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) YadifBuilder
}

// Yadif creates a new YadifBuilder to configure the "yadif" filter.
func Yadif(opts ...filter.Option) YadifBuilder {
	return &implYadifBuilder{
		f: filter.New("yadif", 1, opts...),
	}
}

type implYadifBuilder struct {
	f filter.Filter
}

func (yadifBuilder *implYadifBuilder) String() string {
	return yadifBuilder.f.String()
}

func (yadifBuilder *implYadifBuilder) Outputs() int {
	return yadifBuilder.f.Outputs()
}

func (yadifBuilder *implYadifBuilder) With(key string, value expr.Expr) filter.Filter {
	return yadifBuilder.withOption(key, value)
}

func (yadifBuilder *implYadifBuilder) withOption(key string, value expr.Expr) YadifBuilder {
	bCopy := *yadifBuilder
	bCopy.f = yadifBuilder.f.With(key, value)
	return &bCopy
}

func (yadifBuilder *implYadifBuilder) Mode(mode int) YadifBuilder {
	return yadifBuilder.withOption("mode", expr.Int(mode))
}

func (yadifBuilder *implYadifBuilder) Parity(parity int) YadifBuilder {
	return yadifBuilder.withOption("parity", expr.Int(parity))
}

func (yadifBuilder *implYadifBuilder) Deint(deint int) YadifBuilder {
	return yadifBuilder.withOption("deint", expr.Int(deint))
}

func (yadifBuilder *implYadifBuilder) Enable(enable expr.Expr) YadifBuilder {
	return yadifBuilder.withOption("enable", enable)
}
