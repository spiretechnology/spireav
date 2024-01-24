package spireav

// FfmpegArger defines the interface for an object that creates FFmpeg command line arguments.
type FfmpegArger interface {
	FfmpegArgs() ([]string, error)
}

// FfmpegArgs is a simple type that implements the FfmpegArger interface.
type FfmpegArgs []string

func (args FfmpegArgs) FfmpegArgs() ([]string, error) {
	return args, nil
}
