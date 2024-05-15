// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// MetadataBuilder Manipulate video frame metadata.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#metadata
type MetadataBuilder interface {
	filter.Filter
	// Mode set a mode of operation (from 0 to 4) (default select).
	Mode(mode int) MetadataBuilder
	// Key set metadata key.
	Key(key string) MetadataBuilder
	// Value set metadata value.
	Value(value string) MetadataBuilder
	// Function function for comparing values (from 0 to 6) (default same_str).
	Function(function int) MetadataBuilder
	// Expr set expression for expr function.
	Expr(expression string) MetadataBuilder
	// File set file where to print metadata information.
	File(file string) MetadataBuilder
	// Direct reduce buffering when printing to user-set file or pipe (default false).
	Direct(direct bool) MetadataBuilder
}

// Metadata creates a new MetadataBuilder to configure the "metadata" filter.
func Metadata(opts ...filter.Option) MetadataBuilder {
	return &implMetadataBuilder{
		f: filter.New("metadata", 1, opts...),
	}
}

type implMetadataBuilder struct {
	f filter.Filter
}

func (metadataBuilder *implMetadataBuilder) String() string {
	return metadataBuilder.f.String()
}

func (metadataBuilder *implMetadataBuilder) Outputs() int {
	return metadataBuilder.f.Outputs()
}

func (metadataBuilder *implMetadataBuilder) With(key string, value expr.Expr) filter.Filter {
	return metadataBuilder.withOption(key, value)
}

func (metadataBuilder *implMetadataBuilder) withOption(key string, value expr.Expr) MetadataBuilder {
	bCopy := *metadataBuilder
	bCopy.f = metadataBuilder.f.With(key, value)
	return &bCopy
}

func (metadataBuilder *implMetadataBuilder) Mode(mode int) MetadataBuilder {
	return metadataBuilder.withOption("mode", expr.Int(mode))
}

func (metadataBuilder *implMetadataBuilder) Key(key string) MetadataBuilder {
	return metadataBuilder.withOption("key", expr.String(key))
}

func (metadataBuilder *implMetadataBuilder) Value(value string) MetadataBuilder {
	return metadataBuilder.withOption("value", expr.String(value))
}

func (metadataBuilder *implMetadataBuilder) Function(function int) MetadataBuilder {
	return metadataBuilder.withOption("function", expr.Int(function))
}

func (metadataBuilder *implMetadataBuilder) Expr(expression string) MetadataBuilder {
	return metadataBuilder.withOption("expr", expr.String(expression))
}

func (metadataBuilder *implMetadataBuilder) File(file string) MetadataBuilder {
	return metadataBuilder.withOption("file", expr.String(file))
}

func (metadataBuilder *implMetadataBuilder) Direct(direct bool) MetadataBuilder {
	return metadataBuilder.withOption("direct", expr.Bool(direct))
}
