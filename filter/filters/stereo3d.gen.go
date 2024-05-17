// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// Stereo3dBuilder Convert video stereoscopic 3D view.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#stereo3d
type Stereo3dBuilder interface {
	filter.Filter
	// In set input format (from 16 to 32) (default sbsl).
	In(in int) Stereo3dBuilder
	// Out set output format (from 0 to 32) (default arcd).
	Out(out int) Stereo3dBuilder
}

// Stereo3d creates a new Stereo3dBuilder to configure the "stereo3d" filter.
func Stereo3d(opts ...filter.Option) Stereo3dBuilder {
	return &implStereo3dBuilder{
		f: filter.New("stereo3d", 1, opts...),
	}
}

type implStereo3dBuilder struct {
	f filter.Filter
}

func (stereo3dBuilder *implStereo3dBuilder) String() string {
	return stereo3dBuilder.f.String()
}

func (stereo3dBuilder *implStereo3dBuilder) Outputs() int {
	return stereo3dBuilder.f.Outputs()
}

func (stereo3dBuilder *implStereo3dBuilder) With(key string, value expr.Expr) filter.Filter {
	return stereo3dBuilder.withOption(key, value)
}

func (stereo3dBuilder *implStereo3dBuilder) withOption(key string, value expr.Expr) Stereo3dBuilder {
	bCopy := *stereo3dBuilder
	bCopy.f = stereo3dBuilder.f.With(key, value)
	return &bCopy
}

func (stereo3dBuilder *implStereo3dBuilder) In(in int) Stereo3dBuilder {
	return stereo3dBuilder.withOption("in", expr.Int(in))
}

func (stereo3dBuilder *implStereo3dBuilder) Out(out int) Stereo3dBuilder {
	return stereo3dBuilder.withOption("out", expr.Int(out))
}