package graph

type OutputNodeMP4 struct {
	Filename  string
	Width     int
	Height    int
	FrameRate string
}

func (o *OutputNodeMP4) GetFilename() string {
	return o.Filename
}

func (o *OutputNodeMP4) GetOptions() []string {
	opts := []string{
		"-movflags", "use_metadata_tags",
		"-vcodec", "h264",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		"-f", "mp4",
	}
	if len(o.FrameRate) > 0 {
		opts = append(opts, "-r", o.FrameRate)
	}
	return opts
}
