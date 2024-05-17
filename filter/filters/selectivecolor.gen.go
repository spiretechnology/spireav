// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SelectivecolorBuilder Apply CMYK adjustments to specific color ranges.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#selectivecolor
type SelectivecolorBuilder interface {
	filter.Filter
	// CorrectionMethod select correction method (from 0 to 1) (default absolute).
	CorrectionMethod(correctionMethod int) SelectivecolorBuilder
	// Reds adjust red regions.
	Reds(reds string) SelectivecolorBuilder
	// Yellows adjust yellow regions.
	Yellows(yellows string) SelectivecolorBuilder
	// Greens adjust green regions.
	Greens(greens string) SelectivecolorBuilder
	// Cyans adjust cyan regions.
	Cyans(cyans string) SelectivecolorBuilder
	// Blues adjust blue regions.
	Blues(blues string) SelectivecolorBuilder
	// Magentas adjust magenta regions.
	Magentas(magentas string) SelectivecolorBuilder
	// Whites adjust white regions.
	Whites(whites string) SelectivecolorBuilder
	// Neutrals adjust neutral regions.
	Neutrals(neutrals string) SelectivecolorBuilder
	// Blacks adjust black regions.
	Blacks(blacks string) SelectivecolorBuilder
	// Psfile set Photoshop selectivecolor file name.
	Psfile(psfile string) SelectivecolorBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) SelectivecolorBuilder
}

// Selectivecolor creates a new SelectivecolorBuilder to configure the "selectivecolor" filter.
func Selectivecolor(opts ...filter.Option) SelectivecolorBuilder {
	return &implSelectivecolorBuilder{
		f: filter.New("selectivecolor", 1, opts...),
	}
}

type implSelectivecolorBuilder struct {
	f filter.Filter
}

func (selectivecolorBuilder *implSelectivecolorBuilder) String() string {
	return selectivecolorBuilder.f.String()
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Outputs() int {
	return selectivecolorBuilder.f.Outputs()
}

func (selectivecolorBuilder *implSelectivecolorBuilder) With(key string, value expr.Expr) filter.Filter {
	return selectivecolorBuilder.withOption(key, value)
}

func (selectivecolorBuilder *implSelectivecolorBuilder) withOption(key string, value expr.Expr) SelectivecolorBuilder {
	bCopy := *selectivecolorBuilder
	bCopy.f = selectivecolorBuilder.f.With(key, value)
	return &bCopy
}

func (selectivecolorBuilder *implSelectivecolorBuilder) CorrectionMethod(correctionMethod int) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("correction_method", expr.Int(correctionMethod))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Reds(reds string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("reds", expr.String(reds))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Yellows(yellows string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("yellows", expr.String(yellows))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Greens(greens string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("greens", expr.String(greens))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Cyans(cyans string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("cyans", expr.String(cyans))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Blues(blues string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("blues", expr.String(blues))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Magentas(magentas string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("magentas", expr.String(magentas))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Whites(whites string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("whites", expr.String(whites))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Neutrals(neutrals string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("neutrals", expr.String(neutrals))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Blacks(blacks string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("blacks", expr.String(blacks))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Psfile(psfile string) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("psfile", expr.String(psfile))
}

func (selectivecolorBuilder *implSelectivecolorBuilder) Enable(enable expr.Expr) SelectivecolorBuilder {
	return selectivecolorBuilder.withOption("enable", enable)
}