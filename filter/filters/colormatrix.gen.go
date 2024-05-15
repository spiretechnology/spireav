// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColormatrixBuilder Convert color matrix.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colormatrix
type ColormatrixBuilder interface {
	filter.Filter
	// Src set source color matrix (from -1 to 4) (default -1).
	Src(src int) ColormatrixBuilder
	// Dst set destination color matrix (from -1 to 4) (default -1).
	Dst(dst int) ColormatrixBuilder
}

// Colormatrix creates a new ColormatrixBuilder to configure the "colormatrix" filter.
func Colormatrix(opts ...filter.Option) ColormatrixBuilder {
	return &implColormatrixBuilder{
		f: filter.New("colormatrix", 1, opts...),
	}
}

type implColormatrixBuilder struct {
	f filter.Filter
}

func (colormatrixBuilder *implColormatrixBuilder) String() string {
	return colormatrixBuilder.f.String()
}

func (colormatrixBuilder *implColormatrixBuilder) Outputs() int {
	return colormatrixBuilder.f.Outputs()
}

func (colormatrixBuilder *implColormatrixBuilder) With(key string, value expr.Expr) filter.Filter {
	return colormatrixBuilder.withOption(key, value)
}

func (colormatrixBuilder *implColormatrixBuilder) withOption(key string, value expr.Expr) ColormatrixBuilder {
	bCopy := *colormatrixBuilder
	bCopy.f = colormatrixBuilder.f.With(key, value)
	return &bCopy
}

func (colormatrixBuilder *implColormatrixBuilder) Src(src int) ColormatrixBuilder {
	return colormatrixBuilder.withOption("src", expr.Int(src))
}

func (colormatrixBuilder *implColormatrixBuilder) Dst(dst int) ColormatrixBuilder {
	return colormatrixBuilder.withOption("dst", expr.Int(dst))
}
