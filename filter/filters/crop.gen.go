// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// CropBuilder Crop the input video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#crop
type CropBuilder interface {
	filter.Filter
	// OutW set the width crop area expression (default "iw").
	OutW(outW string) CropBuilder
	// OutWExpr set the width crop area expression (default "iw").
	OutWExpr(outW expr.Expr) CropBuilder
	// W set the width crop area expression (default "iw").
	W(w string) CropBuilder
	// WExpr set the width crop area expression (default "iw").
	WExpr(w expr.Expr) CropBuilder
	// OutH set the height crop area expression (default "ih").
	OutH(outH string) CropBuilder
	// OutHExpr set the height crop area expression (default "ih").
	OutHExpr(outH expr.Expr) CropBuilder
	// H set the height crop area expression (default "ih").
	H(h string) CropBuilder
	// HExpr set the height crop area expression (default "ih").
	HExpr(h expr.Expr) CropBuilder
	// X set the x crop area expression (default "(in_w-out_w)/2").
	X(x string) CropBuilder
	// XExpr set the x crop area expression (default "(in_w-out_w)/2").
	XExpr(x expr.Expr) CropBuilder
	// Y set the y crop area expression (default "(in_h-out_h)/2").
	Y(y string) CropBuilder
	// YExpr set the y crop area expression (default "(in_h-out_h)/2").
	YExpr(y expr.Expr) CropBuilder
	// KeepAspect keep aspect ratio (default false).
	KeepAspect(keepAspect bool) CropBuilder
	// Exact do exact cropping (default false).
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

func (cropBuilder *implCropBuilder) String() string {
	return cropBuilder.f.String()
}

func (cropBuilder *implCropBuilder) Outputs() int {
	return cropBuilder.f.Outputs()
}

func (cropBuilder *implCropBuilder) With(key string, value expr.Expr) filter.Filter {
	return cropBuilder.withOption(key, value)
}

func (cropBuilder *implCropBuilder) withOption(key string, value expr.Expr) CropBuilder {
	bCopy := *cropBuilder
	bCopy.f = cropBuilder.f.With(key, value)
	return &bCopy
}

func (cropBuilder *implCropBuilder) OutW(outW string) CropBuilder {
	return cropBuilder.withOption("out_w", expr.String(outW))
}

func (cropBuilder *implCropBuilder) OutWExpr(outW expr.Expr) CropBuilder {
	return cropBuilder.withOption("out_w", outW)
}

func (cropBuilder *implCropBuilder) W(w string) CropBuilder {
	return cropBuilder.withOption("w", expr.String(w))
}

func (cropBuilder *implCropBuilder) WExpr(w expr.Expr) CropBuilder {
	return cropBuilder.withOption("w", w)
}

func (cropBuilder *implCropBuilder) OutH(outH string) CropBuilder {
	return cropBuilder.withOption("out_h", expr.String(outH))
}

func (cropBuilder *implCropBuilder) OutHExpr(outH expr.Expr) CropBuilder {
	return cropBuilder.withOption("out_h", outH)
}

func (cropBuilder *implCropBuilder) H(h string) CropBuilder {
	return cropBuilder.withOption("h", expr.String(h))
}

func (cropBuilder *implCropBuilder) HExpr(h expr.Expr) CropBuilder {
	return cropBuilder.withOption("h", h)
}

func (cropBuilder *implCropBuilder) X(x string) CropBuilder {
	return cropBuilder.withOption("x", expr.String(x))
}

func (cropBuilder *implCropBuilder) XExpr(x expr.Expr) CropBuilder {
	return cropBuilder.withOption("x", x)
}

func (cropBuilder *implCropBuilder) Y(y string) CropBuilder {
	return cropBuilder.withOption("y", expr.String(y))
}

func (cropBuilder *implCropBuilder) YExpr(y expr.Expr) CropBuilder {
	return cropBuilder.withOption("y", y)
}

func (cropBuilder *implCropBuilder) KeepAspect(keepAspect bool) CropBuilder {
	return cropBuilder.withOption("keep_aspect", expr.Bool(keepAspect))
}

func (cropBuilder *implCropBuilder) Exact(exact bool) CropBuilder {
	return cropBuilder.withOption("exact", expr.Bool(exact))
}
