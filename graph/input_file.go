package graph

// FileInputNode is an InputNode that uses a file as its data source
type FileInputNode struct {
	InputNode
	Filename string
}

// GetFilename returns the filename of the input file
func (in *FileInputNode) GetFilename() string {
	return in.Filename
}
