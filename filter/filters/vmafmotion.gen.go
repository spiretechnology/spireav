// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// VmafmotionBuilder Calculate the VMAF Motion score.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#vmafmotion
type VmafmotionBuilder interface {
	filter.Filter
	// StatsFile Set file where to store per-frame difference information.
	StatsFile(statsFile string) VmafmotionBuilder
}

// Vmafmotion creates a new VmafmotionBuilder to configure the "vmafmotion" filter.
func Vmafmotion(opts ...filter.Option) VmafmotionBuilder {
	return &implVmafmotionBuilder{
		f: filter.New("vmafmotion", 1, opts...),
	}
}

type implVmafmotionBuilder struct {
	f filter.Filter
}

func (vmafmotionBuilder *implVmafmotionBuilder) String() string {
	return vmafmotionBuilder.f.String()
}

func (vmafmotionBuilder *implVmafmotionBuilder) Outputs() int {
	return vmafmotionBuilder.f.Outputs()
}

func (vmafmotionBuilder *implVmafmotionBuilder) With(key string, value expr.Expr) filter.Filter {
	return vmafmotionBuilder.withOption(key, value)
}

func (vmafmotionBuilder *implVmafmotionBuilder) withOption(key string, value expr.Expr) VmafmotionBuilder {
	bCopy := *vmafmotionBuilder
	bCopy.f = vmafmotionBuilder.f.With(key, value)
	return &bCopy
}

func (vmafmotionBuilder *implVmafmotionBuilder) StatsFile(statsFile string) VmafmotionBuilder {
	return vmafmotionBuilder.withOption("stats_file", expr.String(statsFile))
}
