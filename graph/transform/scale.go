package transform

import "fmt"

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
