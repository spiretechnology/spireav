// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// NnediBuilder Apply neural network edge directed interpolation intra-only deinterlacer.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#nnedi
type NnediBuilder interface {
	filter.Filter
	// Weights set weights file (default "nnedi3_weights.bin").
	Weights(weights string) NnediBuilder
	// Deint set which frames to deinterlace (from 0 to 1) (default all).
	Deint(deint int) NnediBuilder
	// DeintExpr set which frames to deinterlace (from 0 to 1) (default all).
	DeintExpr(deint expr.Expr) NnediBuilder
	// Field set mode of operation (from -2 to 3) (default a).
	Field(field int) NnediBuilder
	// FieldExpr set mode of operation (from -2 to 3) (default a).
	FieldExpr(field expr.Expr) NnediBuilder
	// Planes set which planes to process (from 0 to 15) (default 7).
	Planes(planes int) NnediBuilder
	// PlanesExpr set which planes to process (from 0 to 15) (default 7).
	PlanesExpr(planes expr.Expr) NnediBuilder
	// Nsize set size of local neighborhood around each pixel, used by the predictor neural network (from 0 to 6) (default s32x4).
	Nsize(nsize int) NnediBuilder
	// NsizeExpr set size of local neighborhood around each pixel, used by the predictor neural network (from 0 to 6) (default s32x4).
	NsizeExpr(nsize expr.Expr) NnediBuilder
	// Nns set number of neurons in predictor neural network (from 0 to 4) (default n32).
	Nns(nns int) NnediBuilder
	// NnsExpr set number of neurons in predictor neural network (from 0 to 4) (default n32).
	NnsExpr(nns expr.Expr) NnediBuilder
	// Qual set quality (from 1 to 2) (default fast).
	Qual(qual int) NnediBuilder
	// QualExpr set quality (from 1 to 2) (default fast).
	QualExpr(qual expr.Expr) NnediBuilder
	// Etype set which set of weights to use in the predictor (from 0 to 1) (default a).
	Etype(etype int) NnediBuilder
	// EtypeExpr set which set of weights to use in the predictor (from 0 to 1) (default a).
	EtypeExpr(etype expr.Expr) NnediBuilder
	// Pscrn set prescreening (from 0 to 4) (default new).
	Pscrn(pscrn int) NnediBuilder
	// PscrnExpr set prescreening (from 0 to 4) (default new).
	PscrnExpr(pscrn expr.Expr) NnediBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) NnediBuilder
}

// Nnedi creates a new NnediBuilder to configure the "nnedi" filter.
func Nnedi(opts ...filter.Option) NnediBuilder {
	return &implNnediBuilder{
		f: filter.New("nnedi", 1, opts...),
	}
}

type implNnediBuilder struct {
	f filter.Filter
}

func (nnediBuilder *implNnediBuilder) String() string {
	return nnediBuilder.f.String()
}

func (nnediBuilder *implNnediBuilder) Outputs() int {
	return nnediBuilder.f.Outputs()
}

func (nnediBuilder *implNnediBuilder) With(key string, value expr.Expr) filter.Filter {
	return nnediBuilder.withOption(key, value)
}

func (nnediBuilder *implNnediBuilder) withOption(key string, value expr.Expr) NnediBuilder {
	bCopy := *nnediBuilder
	bCopy.f = nnediBuilder.f.With(key, value)
	return &bCopy
}

func (nnediBuilder *implNnediBuilder) Weights(weights string) NnediBuilder {
	return nnediBuilder.withOption("weights", expr.String(weights))
}

func (nnediBuilder *implNnediBuilder) Deint(deint int) NnediBuilder {
	return nnediBuilder.withOption("deint", expr.Int(deint))
}

func (nnediBuilder *implNnediBuilder) DeintExpr(deint expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("deint", deint)
}

func (nnediBuilder *implNnediBuilder) Field(field int) NnediBuilder {
	return nnediBuilder.withOption("field", expr.Int(field))
}

func (nnediBuilder *implNnediBuilder) FieldExpr(field expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("field", field)
}

func (nnediBuilder *implNnediBuilder) Planes(planes int) NnediBuilder {
	return nnediBuilder.withOption("planes", expr.Int(planes))
}

func (nnediBuilder *implNnediBuilder) PlanesExpr(planes expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("planes", planes)
}

func (nnediBuilder *implNnediBuilder) Nsize(nsize int) NnediBuilder {
	return nnediBuilder.withOption("nsize", expr.Int(nsize))
}

func (nnediBuilder *implNnediBuilder) NsizeExpr(nsize expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("nsize", nsize)
}

func (nnediBuilder *implNnediBuilder) Nns(nns int) NnediBuilder {
	return nnediBuilder.withOption("nns", expr.Int(nns))
}

func (nnediBuilder *implNnediBuilder) NnsExpr(nns expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("nns", nns)
}

func (nnediBuilder *implNnediBuilder) Qual(qual int) NnediBuilder {
	return nnediBuilder.withOption("qual", expr.Int(qual))
}

func (nnediBuilder *implNnediBuilder) QualExpr(qual expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("qual", qual)
}

func (nnediBuilder *implNnediBuilder) Etype(etype int) NnediBuilder {
	return nnediBuilder.withOption("etype", expr.Int(etype))
}

func (nnediBuilder *implNnediBuilder) EtypeExpr(etype expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("etype", etype)
}

func (nnediBuilder *implNnediBuilder) Pscrn(pscrn int) NnediBuilder {
	return nnediBuilder.withOption("pscrn", expr.Int(pscrn))
}

func (nnediBuilder *implNnediBuilder) PscrnExpr(pscrn expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("pscrn", pscrn)
}

func (nnediBuilder *implNnediBuilder) Enable(enable expr.Expr) NnediBuilder {
	return nnediBuilder.withOption("enable", enable)
}
