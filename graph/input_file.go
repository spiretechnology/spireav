package graph

type FileInputNode struct {
	InputNode
	Filename string
}

func (in *FileInputNode) GetFilename() string {
	return in.Filename
}
