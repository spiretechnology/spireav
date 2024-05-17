// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// RemovelogoBuilder Remove a TV logo based on a mask image.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#removelogo
type RemovelogoBuilder interface {
	filter.Filter
	// Filename set bitmap filename.
	Filename(filename string) RemovelogoBuilder
	// F set bitmap filename.
	F(f string) RemovelogoBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) RemovelogoBuilder
}

// Removelogo creates a new RemovelogoBuilder to configure the "removelogo" filter.
func Removelogo(opts ...filter.Option) RemovelogoBuilder {
	return &implRemovelogoBuilder{
		f: filter.New("removelogo", 1, opts...),
	}
}

type implRemovelogoBuilder struct {
	f filter.Filter
}

func (removelogoBuilder *implRemovelogoBuilder) String() string {
	return removelogoBuilder.f.String()
}

func (removelogoBuilder *implRemovelogoBuilder) Outputs() int {
	return removelogoBuilder.f.Outputs()
}

func (removelogoBuilder *implRemovelogoBuilder) With(key string, value expr.Expr) filter.Filter {
	return removelogoBuilder.withOption(key, value)
}

func (removelogoBuilder *implRemovelogoBuilder) withOption(key string, value expr.Expr) RemovelogoBuilder {
	bCopy := *removelogoBuilder
	bCopy.f = removelogoBuilder.f.With(key, value)
	return &bCopy
}

func (removelogoBuilder *implRemovelogoBuilder) Filename(filename string) RemovelogoBuilder {
	return removelogoBuilder.withOption("filename", expr.String(filename))
}

func (removelogoBuilder *implRemovelogoBuilder) F(f string) RemovelogoBuilder {
	return removelogoBuilder.withOption("f", expr.String(f))
}

func (removelogoBuilder *implRemovelogoBuilder) Enable(enable expr.Expr) RemovelogoBuilder {
	return removelogoBuilder.withOption("enable", enable)
}