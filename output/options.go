package output

import (
	"runtime"
	"strconv"
)

func WithFormat(format string) Option {
	return WithOptions("-f", format)
}

func WithFormatMP4() Option {
	return withMulti(
		WithFormat("mp4"),
		withDefaultMP4Options(),
	)
}

func WithFormatMOV() Option {
	return withMulti(
		WithFormat("mov"),
		withDefaultMP4Options(),
	)
}

func withDefaultMP4Options() Option {
	codec := "h264"
	if runtime.GOOS == "darwin" {
		codec = "h264_videotoolbox"
	}
	return WithOptions(
		"-movflags", "use_metadata_tags",
		"-vcodec", codec,
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		"-tune", "fastdecode",
		"-g", "12", // 12-frame GOP size
	)
}

func WithFrameRate(frameRate string) Option {
	return WithOptions("-r", frameRate)
}

func WithTimecode(timecode string) Option {
	return WithOptions("-timecode", timecode)
}

func WithConstantRateFactor(factor int) Option {
	return WithOptions("-crf", strconv.Itoa(factor))
}

func WithVideoCodec(codec string) Option {
	return WithOptions("-vcodec", codec)
}

func WithAudioCodec(codec string) Option {
	return WithOptions("-acodec", codec)
}
