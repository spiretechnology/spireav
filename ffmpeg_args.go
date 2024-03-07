package spireav

// FfmpegArgs defines the interface for an object that creates FFmpeg command line arguments.
type FfmpegArgs interface {
	FfmpegArgs() ([]string, error)
}

// FfmpegArgsSlice is a simple type that implements the FfmpegArgs interface.
type FfmpegArgsSlice []string

func (args FfmpegArgsSlice) FfmpegArgs() ([]string, error) {
	return args, nil
}
