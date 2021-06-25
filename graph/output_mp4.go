package graph

type OutputNodeMP4 struct {
	Filename string
	Width    int
	Height   int
}

func (o *OutputNodeMP4) GetFilename() string {
	return o.Filename
}

func (o *OutputNodeMP4) GetOptions() []string {
	opts := []string{
		"-movflags", "use_metadata_tags",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		"-f", "mp4",
	}
	return opts
}
