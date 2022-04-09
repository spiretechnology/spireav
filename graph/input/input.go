package input

// Input is a node that is a source of one or more streams of data. InputNode is usually a file. In ffmpeg
// terms, this is a buffersrc
type Input interface {
	Filename() string
}

// New creates a new graph input with the given filename.
func New(filename string) Input {
	return &inputWithFilename{
		filename,
	}
}

type inputWithFilename struct {
	filename string
}

func (i *inputWithFilename) Filename() string {
	return i.filename
}
