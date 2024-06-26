// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AphasemeterBuilder Convert input audio to phase meter video output.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#aphasemeter
type AphasemeterBuilder interface {
	filter.Filter
	// Rate set video rate (default "25").
	Rate(rate expr.Rational) AphasemeterBuilder
	// R set video rate (default "25").
	R(r expr.Rational) AphasemeterBuilder
	// Size set video size (default "800x400").
	Size(size expr.Size) AphasemeterBuilder
	// S set video size (default "800x400").
	S(s expr.Size) AphasemeterBuilder
	// Rc set red contrast (from 0 to 255) (default 2).
	Rc(rc int) AphasemeterBuilder
	// Gc set green contrast (from 0 to 255) (default 7).
	Gc(gc int) AphasemeterBuilder
	// Bc set blue contrast (from 0 to 255) (default 1).
	Bc(bc int) AphasemeterBuilder
	// Mpc set median phase color (default "none").
	Mpc(mpc string) AphasemeterBuilder
	// Video set video output (default true).
	Video(video bool) AphasemeterBuilder
	// Phasing set mono and out-of-phase detection output (default false).
	Phasing(phasing bool) AphasemeterBuilder
	// Tolerance set phase tolerance for mono detection (from 0 to 1) (default 0).
	Tolerance(tolerance float32) AphasemeterBuilder
	// T set phase tolerance for mono detection (from 0 to 1) (default 0).
	T(t float32) AphasemeterBuilder
	// Angle set angle threshold for out-of-phase detection (from 90 to 180) (default 170).
	Angle(angle float32) AphasemeterBuilder
	// A set angle threshold for out-of-phase detection (from 90 to 180) (default 170).
	A(a float32) AphasemeterBuilder
	// Duration set minimum mono or out-of-phase duration in seconds (default 2).
	Duration(duration time.Duration) AphasemeterBuilder
	// D set minimum mono or out-of-phase duration in seconds (default 2).
	D(d time.Duration) AphasemeterBuilder
}

// Aphasemeter creates a new AphasemeterBuilder to configure the "aphasemeter" filter.
func Aphasemeter(outputs int, opts ...filter.Option) AphasemeterBuilder {
	f := filter.New("aphasemeter", outputs, opts...)
	f = f.With("outputs", expr.Int(outputs))
	return &implAphasemeterBuilder{f: f}
}

type implAphasemeterBuilder struct {
	f filter.Filter
}

func (aphasemeterBuilder *implAphasemeterBuilder) String() string {
	return aphasemeterBuilder.f.String()
}

func (aphasemeterBuilder *implAphasemeterBuilder) Outputs() int {
	return aphasemeterBuilder.f.Outputs()
}

func (aphasemeterBuilder *implAphasemeterBuilder) With(key string, value expr.Expr) filter.Filter {
	return aphasemeterBuilder.withOption(key, value)
}

func (aphasemeterBuilder *implAphasemeterBuilder) withOption(key string, value expr.Expr) AphasemeterBuilder {
	bCopy := *aphasemeterBuilder
	bCopy.f = aphasemeterBuilder.f.With(key, value)
	return &bCopy
}

func (aphasemeterBuilder *implAphasemeterBuilder) Rate(rate expr.Rational) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("rate", rate)
}

func (aphasemeterBuilder *implAphasemeterBuilder) R(r expr.Rational) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("r", r)
}

func (aphasemeterBuilder *implAphasemeterBuilder) Size(size expr.Size) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("size", size)
}

func (aphasemeterBuilder *implAphasemeterBuilder) S(s expr.Size) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("s", s)
}

func (aphasemeterBuilder *implAphasemeterBuilder) Rc(rc int) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("rc", expr.Int(rc))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Gc(gc int) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("gc", expr.Int(gc))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Bc(bc int) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("bc", expr.Int(bc))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Mpc(mpc string) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("mpc", expr.String(mpc))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Video(video bool) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("video", expr.Bool(video))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Phasing(phasing bool) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("phasing", expr.Bool(phasing))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Tolerance(tolerance float32) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("tolerance", expr.Float(tolerance))
}

func (aphasemeterBuilder *implAphasemeterBuilder) T(t float32) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("t", expr.Float(t))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Angle(angle float32) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("angle", expr.Float(angle))
}

func (aphasemeterBuilder *implAphasemeterBuilder) A(a float32) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("a", expr.Float(a))
}

func (aphasemeterBuilder *implAphasemeterBuilder) Duration(duration time.Duration) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("duration", expr.Duration(duration))
}

func (aphasemeterBuilder *implAphasemeterBuilder) D(d time.Duration) AphasemeterBuilder {
	return aphasemeterBuilder.withOption("d", expr.Duration(d))
}
