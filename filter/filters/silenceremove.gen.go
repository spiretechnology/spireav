// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"time"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SilenceremoveBuilder Remove silence.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#silenceremove
type SilenceremoveBuilder interface {
	filter.Filter
	// StartPeriods set periods of silence parts to skip from start (from 0 to 9000) (default 0).
	StartPeriods(startPeriods int) SilenceremoveBuilder
	// StartDuration set start duration of non-silence part (default 0).
	StartDuration(startDuration time.Duration) SilenceremoveBuilder
	// StartThreshold set threshold for start silence detection (from 0 to DBL_MAX) (default 0).
	StartThreshold(startThreshold float64) SilenceremoveBuilder
	// StartThresholdExpr set threshold for start silence detection (from 0 to DBL_MAX) (default 0).
	StartThresholdExpr(startThreshold expr.Expr) SilenceremoveBuilder
	// StartSilence set start duration of silence part to keep (default 0).
	StartSilence(startSilence time.Duration) SilenceremoveBuilder
	// StartMode set which channel will trigger trimming from start (from 0 to 1) (default any).
	StartMode(startMode int) SilenceremoveBuilder
	// StartModeExpr set which channel will trigger trimming from start (from 0 to 1) (default any).
	StartModeExpr(startMode expr.Expr) SilenceremoveBuilder
	// StopPeriods set periods of silence parts to skip from end (from -9000 to 9000) (default 0).
	StopPeriods(stopPeriods int) SilenceremoveBuilder
	// StopDuration set stop duration of silence part (default 0).
	StopDuration(stopDuration time.Duration) SilenceremoveBuilder
	// StopThreshold set threshold for stop silence detection (from 0 to DBL_MAX) (default 0).
	StopThreshold(stopThreshold float64) SilenceremoveBuilder
	// StopThresholdExpr set threshold for stop silence detection (from 0 to DBL_MAX) (default 0).
	StopThresholdExpr(stopThreshold expr.Expr) SilenceremoveBuilder
	// StopSilence set stop duration of silence part to keep (default 0).
	StopSilence(stopSilence time.Duration) SilenceremoveBuilder
	// StopMode set which channel will trigger trimming from end (from 0 to 1) (default all).
	StopMode(stopMode int) SilenceremoveBuilder
	// StopModeExpr set which channel will trigger trimming from end (from 0 to 1) (default all).
	StopModeExpr(stopMode expr.Expr) SilenceremoveBuilder
	// Detection set how silence is detected (from 0 to 5) (default rms).
	Detection(detection int) SilenceremoveBuilder
	// Window set duration of window for silence detection (default 0.02).
	Window(window time.Duration) SilenceremoveBuilder
	// Timestamp set how every output frame timestamp is processed (from 0 to 1) (default write).
	Timestamp(timestamp int) SilenceremoveBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) SilenceremoveBuilder
}

// Silenceremove creates a new SilenceremoveBuilder to configure the "silenceremove" filter.
func Silenceremove(opts ...filter.Option) SilenceremoveBuilder {
	return &implSilenceremoveBuilder{
		f: filter.New("silenceremove", 1, opts...),
	}
}

type implSilenceremoveBuilder struct {
	f filter.Filter
}

func (silenceremoveBuilder *implSilenceremoveBuilder) String() string {
	return silenceremoveBuilder.f.String()
}

func (silenceremoveBuilder *implSilenceremoveBuilder) Outputs() int {
	return silenceremoveBuilder.f.Outputs()
}

func (silenceremoveBuilder *implSilenceremoveBuilder) With(key string, value expr.Expr) filter.Filter {
	return silenceremoveBuilder.withOption(key, value)
}

func (silenceremoveBuilder *implSilenceremoveBuilder) withOption(key string, value expr.Expr) SilenceremoveBuilder {
	bCopy := *silenceremoveBuilder
	bCopy.f = silenceremoveBuilder.f.With(key, value)
	return &bCopy
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartPeriods(startPeriods int) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_periods", expr.Int(startPeriods))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartDuration(startDuration time.Duration) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_duration", expr.Duration(startDuration))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartThreshold(startThreshold float64) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_threshold", expr.Double(startThreshold))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartThresholdExpr(startThreshold expr.Expr) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_threshold", startThreshold)
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartSilence(startSilence time.Duration) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_silence", expr.Duration(startSilence))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartMode(startMode int) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_mode", expr.Int(startMode))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StartModeExpr(startMode expr.Expr) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("start_mode", startMode)
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopPeriods(stopPeriods int) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_periods", expr.Int(stopPeriods))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopDuration(stopDuration time.Duration) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_duration", expr.Duration(stopDuration))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopThreshold(stopThreshold float64) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_threshold", expr.Double(stopThreshold))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopThresholdExpr(stopThreshold expr.Expr) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_threshold", stopThreshold)
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopSilence(stopSilence time.Duration) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_silence", expr.Duration(stopSilence))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopMode(stopMode int) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_mode", expr.Int(stopMode))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) StopModeExpr(stopMode expr.Expr) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("stop_mode", stopMode)
}

func (silenceremoveBuilder *implSilenceremoveBuilder) Detection(detection int) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("detection", expr.Int(detection))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) Window(window time.Duration) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("window", expr.Duration(window))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) Timestamp(timestamp int) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("timestamp", expr.Int(timestamp))
}

func (silenceremoveBuilder *implSilenceremoveBuilder) Enable(enable expr.Expr) SilenceremoveBuilder {
	return silenceremoveBuilder.withOption("enable", enable)
}
