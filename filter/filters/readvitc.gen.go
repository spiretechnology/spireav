// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// ReadvitcBuilder Read vertical interval timecode and write it to frame metadata.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#readvitc
type ReadvitcBuilder interface {
	filter.Filter
	// ScanMax maximum line numbers to scan for VITC data (from -1 to INT_MAX) (default 45).
	ScanMax(scanMax int) ReadvitcBuilder
	// ThrB black color threshold (from 0 to 1) (default 0.2).
	ThrB(thrB float64) ReadvitcBuilder
	// ThrW white color threshold (from 0 to 1) (default 0.6).
	ThrW(thrW float64) ReadvitcBuilder
}

// Readvitc creates a new ReadvitcBuilder to configure the "readvitc" filter.
func Readvitc(opts ...filter.Option) ReadvitcBuilder {
	return &implReadvitcBuilder{
		f: filter.New("readvitc", 1, opts...),
	}
}

type implReadvitcBuilder struct {
	f filter.Filter
}

func (readvitcBuilder *implReadvitcBuilder) String() string {
	return readvitcBuilder.f.String()
}

func (readvitcBuilder *implReadvitcBuilder) Outputs() int {
	return readvitcBuilder.f.Outputs()
}

func (readvitcBuilder *implReadvitcBuilder) With(key string, value expr.Expr) filter.Filter {
	return readvitcBuilder.withOption(key, value)
}

func (readvitcBuilder *implReadvitcBuilder) withOption(key string, value expr.Expr) ReadvitcBuilder {
	bCopy := *readvitcBuilder
	bCopy.f = readvitcBuilder.f.With(key, value)
	return &bCopy
}

func (readvitcBuilder *implReadvitcBuilder) ScanMax(scanMax int) ReadvitcBuilder {
	return readvitcBuilder.withOption("scan_max", expr.Int(scanMax))
}

func (readvitcBuilder *implReadvitcBuilder) ThrB(thrB float64) ReadvitcBuilder {
	return readvitcBuilder.withOption("thr_b", expr.Double(thrB))
}

func (readvitcBuilder *implReadvitcBuilder) ThrW(thrW float64) ReadvitcBuilder {
	return readvitcBuilder.withOption("thr_w", expr.Double(thrW))
}
