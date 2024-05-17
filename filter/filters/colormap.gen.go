// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ColormapBuilder Apply custom Color Maps to video stream.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#colormap
type ColormapBuilder interface {
	filter.Filter
	// PatchSize set patch size (default "64x64").
	PatchSize(patchSize expr.Size) ColormapBuilder
	// PatchSizeExpr set patch size (default "64x64").
	PatchSizeExpr(patchSize expr.Expr) ColormapBuilder
	// NbPatches set number of patches (from 0 to 64) (default 0).
	NbPatches(nbPatches int) ColormapBuilder
	// NbPatchesExpr set number of patches (from 0 to 64) (default 0).
	NbPatchesExpr(nbPatches expr.Expr) ColormapBuilder
	// Type set the target type used (from 0 to 1) (default absolute).
	Type(typ int) ColormapBuilder
	// TypeExpr set the target type used (from 0 to 1) (default absolute).
	TypeExpr(typ expr.Expr) ColormapBuilder
	// Kernel set the kernel used for measuring color difference (from 0 to 1) (default euclidean).
	Kernel(kernel int) ColormapBuilder
	// KernelExpr set the kernel used for measuring color difference (from 0 to 1) (default euclidean).
	KernelExpr(kernel expr.Expr) ColormapBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) ColormapBuilder
}

// Colormap creates a new ColormapBuilder to configure the "colormap" filter.
func Colormap(opts ...filter.Option) ColormapBuilder {
	return &implColormapBuilder{
		f: filter.New("colormap", 1, opts...),
	}
}

type implColormapBuilder struct {
	f filter.Filter
}

func (colormapBuilder *implColormapBuilder) String() string {
	return colormapBuilder.f.String()
}

func (colormapBuilder *implColormapBuilder) Outputs() int {
	return colormapBuilder.f.Outputs()
}

func (colormapBuilder *implColormapBuilder) With(key string, value expr.Expr) filter.Filter {
	return colormapBuilder.withOption(key, value)
}

func (colormapBuilder *implColormapBuilder) withOption(key string, value expr.Expr) ColormapBuilder {
	bCopy := *colormapBuilder
	bCopy.f = colormapBuilder.f.With(key, value)
	return &bCopy
}

func (colormapBuilder *implColormapBuilder) PatchSize(patchSize expr.Size) ColormapBuilder {
	return colormapBuilder.withOption("patch_size", patchSize)
}

func (colormapBuilder *implColormapBuilder) PatchSizeExpr(patchSize expr.Expr) ColormapBuilder {
	return colormapBuilder.withOption("patch_size", patchSize)
}

func (colormapBuilder *implColormapBuilder) NbPatches(nbPatches int) ColormapBuilder {
	return colormapBuilder.withOption("nb_patches", expr.Int(nbPatches))
}

func (colormapBuilder *implColormapBuilder) NbPatchesExpr(nbPatches expr.Expr) ColormapBuilder {
	return colormapBuilder.withOption("nb_patches", nbPatches)
}

func (colormapBuilder *implColormapBuilder) Type(typ int) ColormapBuilder {
	return colormapBuilder.withOption("type", expr.Int(typ))
}

func (colormapBuilder *implColormapBuilder) TypeExpr(typ expr.Expr) ColormapBuilder {
	return colormapBuilder.withOption("type", typ)
}

func (colormapBuilder *implColormapBuilder) Kernel(kernel int) ColormapBuilder {
	return colormapBuilder.withOption("kernel", expr.Int(kernel))
}

func (colormapBuilder *implColormapBuilder) KernelExpr(kernel expr.Expr) ColormapBuilder {
	return colormapBuilder.withOption("kernel", kernel)
}

func (colormapBuilder *implColormapBuilder) Enable(enable expr.Expr) ColormapBuilder {
	return colormapBuilder.withOption("enable", enable)
}
