// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// FreezeframesBuilder Freeze video frames.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#freezeframes
type FreezeframesBuilder interface {
	filter.Filter
	// First set first frame to freeze (from 0 to I64_MAX) (default 0).
	First(first int64) FreezeframesBuilder
	// Last set last frame to freeze (from 0 to I64_MAX) (default 0).
	Last(last int64) FreezeframesBuilder
	// Replace set frame to replace (from 0 to I64_MAX) (default 0).
	Replace(replace int64) FreezeframesBuilder
}

// Freezeframes creates a new FreezeframesBuilder to configure the "freezeframes" filter.
func Freezeframes(opts ...filter.Option) FreezeframesBuilder {
	return &implFreezeframesBuilder{
		f: filter.New("freezeframes", 1, opts...),
	}
}

type implFreezeframesBuilder struct {
	f filter.Filter
}

func (freezeframesBuilder *implFreezeframesBuilder) String() string {
	return freezeframesBuilder.f.String()
}

func (freezeframesBuilder *implFreezeframesBuilder) Outputs() int {
	return freezeframesBuilder.f.Outputs()
}

func (freezeframesBuilder *implFreezeframesBuilder) With(key string, value expr.Expr) filter.Filter {
	return freezeframesBuilder.withOption(key, value)
}

func (freezeframesBuilder *implFreezeframesBuilder) withOption(key string, value expr.Expr) FreezeframesBuilder {
	bCopy := *freezeframesBuilder
	bCopy.f = freezeframesBuilder.f.With(key, value)
	return &bCopy
}

func (freezeframesBuilder *implFreezeframesBuilder) First(first int64) FreezeframesBuilder {
	return freezeframesBuilder.withOption("first", expr.Int64(first))
}

func (freezeframesBuilder *implFreezeframesBuilder) Last(last int64) FreezeframesBuilder {
	return freezeframesBuilder.withOption("last", expr.Int64(last))
}

func (freezeframesBuilder *implFreezeframesBuilder) Replace(replace int64) FreezeframesBuilder {
	return freezeframesBuilder.withOption("replace", expr.Int64(replace))
}
