package transform

import "fmt"

// Opacity is a transform that takes in a picture stream (video or photo) and outputs a stream
// with the opacity multiplied by the give floating point value
type Opacity struct {
	Opacity float64
}

func (node *Opacity) GetOutputsCount() int {
	return 1
}

func (node *Opacity) FilterString() string {
	return fmt.Sprintf("geq=r='r(X,Y)':a='%0.3f*alpha(X,Y)'", node.Opacity)
}
