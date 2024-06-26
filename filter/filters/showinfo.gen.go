// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ShowinfoBuilder Show textual information for each video frame.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#showinfo
type ShowinfoBuilder interface {
	filter.Filter
	// Checksum calculate checksums (default true).
	Checksum(checksum bool) ShowinfoBuilder
	// UduSeiAsAscii try to print user data unregistered SEI as ascii character when possible (default false).
	UduSeiAsAscii(uduSeiAsAscii bool) ShowinfoBuilder
}

// Showinfo creates a new ShowinfoBuilder to configure the "showinfo" filter.
func Showinfo(opts ...filter.Option) ShowinfoBuilder {
	return &implShowinfoBuilder{
		f: filter.New("showinfo", 1, opts...),
	}
}

type implShowinfoBuilder struct {
	f filter.Filter
}

func (showinfoBuilder *implShowinfoBuilder) String() string {
	return showinfoBuilder.f.String()
}

func (showinfoBuilder *implShowinfoBuilder) Outputs() int {
	return showinfoBuilder.f.Outputs()
}

func (showinfoBuilder *implShowinfoBuilder) With(key string, value expr.Expr) filter.Filter {
	return showinfoBuilder.withOption(key, value)
}

func (showinfoBuilder *implShowinfoBuilder) withOption(key string, value expr.Expr) ShowinfoBuilder {
	bCopy := *showinfoBuilder
	bCopy.f = showinfoBuilder.f.With(key, value)
	return &bCopy
}

func (showinfoBuilder *implShowinfoBuilder) Checksum(checksum bool) ShowinfoBuilder {
	return showinfoBuilder.withOption("checksum", expr.Bool(checksum))
}

func (showinfoBuilder *implShowinfoBuilder) UduSeiAsAscii(uduSeiAsAscii bool) ShowinfoBuilder {
	return showinfoBuilder.withOption("udu_sei_as_ascii", expr.Bool(uduSeiAsAscii))
}
