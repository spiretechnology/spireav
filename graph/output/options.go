package output

func WithFormat(format string) Option {
	return withOptions("-f", format)
}

func WithFormatMP4() Option {
	return withMulti(
		WithFormat("mp4"),
		withDefaultMP4Options(),
	)
}

func withDefaultMP4Options() Option {
	return withOptions(
		"-movflags", "use_metadata_tags",
		"-vcodec", "h264",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
	)
}

func WithFrameRate(frameRate string) Option {
	return withOptions("-r", frameRate)
}
