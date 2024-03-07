// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CropBuilder corresponds to the "crop" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#crop
type CropBuilder interface {
	filter.Filter
	// PosX sets the "x" option on the filter.
	PosX(posX expr.Expr) CropBuilder
	// PosXInt sets the "x" option on the filter.
	PosXInt(posX int) CropBuilder
	// PosY sets the "y" option on the filter.
	PosY(posY expr.Expr) CropBuilder
	// PosYInt sets the "y" option on the filter.
	PosYInt(posY int) CropBuilder
	// Width sets the "w" option on the filter.
	Width(width expr.Expr) CropBuilder
	// WidthInt sets the "w" option on the filter.
	WidthInt(width int) CropBuilder
	// Height sets the "h" option on the filter.
	Height(height expr.Expr) CropBuilder
	// HeightInt sets the "h" option on the filter.
	HeightInt(height int) CropBuilder
	// KeepAspectRatio sets the "keep_aspect" option on the filter.
	KeepAspectRatio(keepAspectRatio bool) CropBuilder
	// Exact sets the "exact" option on the filter.
	Exact(exact bool) CropBuilder
}

// Crop creates a new CropBuilder to configure the "crop" filter.
func Crop(opts ...filter.Option) CropBuilder {
	return &implCropBuilder{
		f: filter.New("crop", 1, opts...),
	}
}

type implCropBuilder struct {
	f filter.Filter
}

func (b *implCropBuilder) String() string {
	return b.f.String()
}

func (b *implCropBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implCropBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implCropBuilder) withOption(key string, value expr.Expr) CropBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implCropBuilder) PosX(posX expr.Expr) CropBuilder {
	return b.withOption("x", posX)
}

func (b *implCropBuilder) PosXInt(posX int) CropBuilder {
	return b.withOption("x", expr.Int(posX))
}

func (b *implCropBuilder) PosY(posY expr.Expr) CropBuilder {
	return b.withOption("y", posY)
}

func (b *implCropBuilder) PosYInt(posY int) CropBuilder {
	return b.withOption("y", expr.Int(posY))
}

func (b *implCropBuilder) Width(width expr.Expr) CropBuilder {
	return b.withOption("w", width)
}

func (b *implCropBuilder) WidthInt(width int) CropBuilder {
	return b.withOption("w", expr.Int(width))
}

func (b *implCropBuilder) Height(height expr.Expr) CropBuilder {
	return b.withOption("h", height)
}

func (b *implCropBuilder) HeightInt(height int) CropBuilder {
	return b.withOption("h", expr.Int(height))
}

func (b *implCropBuilder) KeepAspectRatio(keepAspectRatio bool) CropBuilder {
	return b.withOption("keep_aspect", expr.Bool(keepAspectRatio))
}

func (b *implCropBuilder) Exact(exact bool) CropBuilder {
	return b.withOption("exact", expr.Bool(exact))
}
