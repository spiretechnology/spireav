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
	opts := []string{}
	return opts
}
