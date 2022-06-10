package output

import "strconv"

func WithFormat(format string) Option {
	return withOptions("-f", format)
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
	return withOptions(
		"-movflags", "use_metadata_tags",
		"-vcodec", "h264",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		"-tune", "fastdecode",
	)
}

func WithFrameRate(frameRate string) Option {
	return withOptions("-r", frameRate)
}

func WithTimecode(timecode string) Option {
	return withOptions("-timecode", timecode)
}

func WithConstantRateFactor(factor int) Option {
	return withOptions("-crf", strconv.Itoa(factor))
}
