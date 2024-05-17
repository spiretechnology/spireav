// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FieldorderBuilder Set the field order.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#fieldorder
type FieldorderBuilder interface {
	filter.Filter
	// Order output field order (from 0 to 1) (default tff).
	Order(order int) FieldorderBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) FieldorderBuilder
}

// Fieldorder creates a new FieldorderBuilder to configure the "fieldorder" filter.
func Fieldorder(opts ...filter.Option) FieldorderBuilder {
	return &implFieldorderBuilder{
		f: filter.New("fieldorder", 1, opts...),
	}
}

type implFieldorderBuilder struct {
	f filter.Filter
}

func (fieldorderBuilder *implFieldorderBuilder) String() string {
	return fieldorderBuilder.f.String()
}

func (fieldorderBuilder *implFieldorderBuilder) Outputs() int {
	return fieldorderBuilder.f.Outputs()
}

func (fieldorderBuilder *implFieldorderBuilder) With(key string, value expr.Expr) filter.Filter {
	return fieldorderBuilder.withOption(key, value)
}

func (fieldorderBuilder *implFieldorderBuilder) withOption(key string, value expr.Expr) FieldorderBuilder {
	bCopy := *fieldorderBuilder
	bCopy.f = fieldorderBuilder.f.With(key, value)
	return &bCopy
}

func (fieldorderBuilder *implFieldorderBuilder) Order(order int) FieldorderBuilder {
	return fieldorderBuilder.withOption("order", expr.Int(order))
}

func (fieldorderBuilder *implFieldorderBuilder) Enable(enable expr.Expr) FieldorderBuilder {
	return fieldorderBuilder.withOption("enable", enable)
}