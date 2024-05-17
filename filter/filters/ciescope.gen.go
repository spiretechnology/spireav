// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CiescopeBuilder Video CIE scope.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#ciescope
type CiescopeBuilder interface {
	filter.Filter
	// System set color system (from 0 to 9) (default hdtv).
	System(system int) CiescopeBuilder
	// Cie set cie system (from 0 to 2) (default xyy).
	Cie(cie int) CiescopeBuilder
	// Gamuts set what gamuts to draw (default 0).
	Gamuts(gamuts ...string) CiescopeBuilder
	// Size set ciescope size (from 256 to 8192) (default 512).
	Size(size int) CiescopeBuilder
	// S set ciescope size (from 256 to 8192) (default 512).
	S(s int) CiescopeBuilder
	// Intensity set ciescope intensity (from 0 to 1) (default 0.001).
	Intensity(intensity float32) CiescopeBuilder
	// I set ciescope intensity (from 0 to 1) (default 0.001).
	I(i float32) CiescopeBuilder
	// Contrast (from 0 to 1) (default 0.75).
	Contrast(contrast float32) CiescopeBuilder
	// Corrgamma (default true).
	Corrgamma(corrgamma bool) CiescopeBuilder
	// Showwhite (default false).
	Showwhite(showwhite bool) CiescopeBuilder
	// Gamma (from 0.1 to 6) (default 2.6).
	Gamma(gamma float64) CiescopeBuilder
	// Fill fill with CIE colors (default true).
	Fill(fill bool) CiescopeBuilder
}

// Ciescope creates a new CiescopeBuilder to configure the "ciescope" filter.
func Ciescope(opts ...filter.Option) CiescopeBuilder {
	return &implCiescopeBuilder{
		f: filter.New("ciescope", 1, opts...),
	}
}

type implCiescopeBuilder struct {
	f filter.Filter
}

func (ciescopeBuilder *implCiescopeBuilder) String() string {
	return ciescopeBuilder.f.String()
}

func (ciescopeBuilder *implCiescopeBuilder) Outputs() int {
	return ciescopeBuilder.f.Outputs()
}

func (ciescopeBuilder *implCiescopeBuilder) With(key string, value expr.Expr) filter.Filter {
	return ciescopeBuilder.withOption(key, value)
}

func (ciescopeBuilder *implCiescopeBuilder) withOption(key string, value expr.Expr) CiescopeBuilder {
	bCopy := *ciescopeBuilder
	bCopy.f = ciescopeBuilder.f.With(key, value)
	return &bCopy
}

func (ciescopeBuilder *implCiescopeBuilder) System(system int) CiescopeBuilder {
	return ciescopeBuilder.withOption("system", expr.Int(system))
}

func (ciescopeBuilder *implCiescopeBuilder) Cie(cie int) CiescopeBuilder {
	return ciescopeBuilder.withOption("cie", expr.Int(cie))
}

func (ciescopeBuilder *implCiescopeBuilder) Gamuts(gamuts ...string) CiescopeBuilder {
	return ciescopeBuilder.withOption("gamuts", expr.Flags(gamuts))
}

func (ciescopeBuilder *implCiescopeBuilder) Size(size int) CiescopeBuilder {
	return ciescopeBuilder.withOption("size", expr.Int(size))
}

func (ciescopeBuilder *implCiescopeBuilder) S(s int) CiescopeBuilder {
	return ciescopeBuilder.withOption("s", expr.Int(s))
}

func (ciescopeBuilder *implCiescopeBuilder) Intensity(intensity float32) CiescopeBuilder {
	return ciescopeBuilder.withOption("intensity", expr.Float(intensity))
}

func (ciescopeBuilder *implCiescopeBuilder) I(i float32) CiescopeBuilder {
	return ciescopeBuilder.withOption("i", expr.Float(i))
}

func (ciescopeBuilder *implCiescopeBuilder) Contrast(contrast float32) CiescopeBuilder {
	return ciescopeBuilder.withOption("contrast", expr.Float(contrast))
}

func (ciescopeBuilder *implCiescopeBuilder) Corrgamma(corrgamma bool) CiescopeBuilder {
	return ciescopeBuilder.withOption("corrgamma", expr.Bool(corrgamma))
}

func (ciescopeBuilder *implCiescopeBuilder) Showwhite(showwhite bool) CiescopeBuilder {
	return ciescopeBuilder.withOption("showwhite", expr.Bool(showwhite))
}

func (ciescopeBuilder *implCiescopeBuilder) Gamma(gamma float64) CiescopeBuilder {
	return ciescopeBuilder.withOption("gamma", expr.Double(gamma))
}

func (ciescopeBuilder *implCiescopeBuilder) Fill(fill bool) CiescopeBuilder {
	return ciescopeBuilder.withOption("fill", expr.Bool(fill))
}
