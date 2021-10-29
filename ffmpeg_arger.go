package spireav

// FfmpegArger defines the interface for an object that creates FFmpeg command line arguments.
type FfmpegArger interface {
	FfmpegArgs() ([]string, error)
}
