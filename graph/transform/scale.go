package transform

import "fmt"

// Scale is a transform that takes in a picture stream (video or photo) and outputs a scaled version
// of the stream.
type Scale struct {
	Width  int
	Height int
}

func (node *Scale) GetOutputsCount() int {
	return 1
}

func (node *Scale) FilterString() string {
	return fmt.Sprintf("scale=%d:%d", node.Width, node.Height)
}
