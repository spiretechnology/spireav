// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AcueBuilder Delay filtering to match a cue.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#acue
type AcueBuilder interface {
	filter.Filter
	// Cue cue unix timestamp in microseconds (from 0 to I64_MAX) (default 0).
	Cue(cue int64) AcueBuilder
	// Preroll preroll duration in seconds (default 0).
	Preroll(preroll time.Duration) AcueBuilder
	// Buffer buffer duration in seconds (default 0).
	Buffer(buffer time.Duration) AcueBuilder
}

// Acue creates a new AcueBuilder to configure the "acue" filter.
func Acue(opts ...filter.Option) AcueBuilder {
	return &implAcueBuilder{
		f: filter.New("acue", 1, opts...),
	}
}

type implAcueBuilder struct {
	f filter.Filter
}

func (acueBuilder *implAcueBuilder) String() string {
	return acueBuilder.f.String()
}

func (acueBuilder *implAcueBuilder) Outputs() int {
	return acueBuilder.f.Outputs()
}

func (acueBuilder *implAcueBuilder) With(key string, value expr.Expr) filter.Filter {
	return acueBuilder.withOption(key, value)
}

func (acueBuilder *implAcueBuilder) withOption(key string, value expr.Expr) AcueBuilder {
	bCopy := *acueBuilder
	bCopy.f = acueBuilder.f.With(key, value)
	return &bCopy
}

func (acueBuilder *implAcueBuilder) Cue(cue int64) AcueBuilder {
	return acueBuilder.withOption("cue", expr.Int64(cue))
}

func (acueBuilder *implAcueBuilder) Preroll(preroll time.Duration) AcueBuilder {
	return acueBuilder.withOption("preroll", expr.Duration(preroll))
}

func (acueBuilder *implAcueBuilder) Buffer(buffer time.Duration) AcueBuilder {
	return acueBuilder.withOption("buffer", expr.Duration(buffer))
}
