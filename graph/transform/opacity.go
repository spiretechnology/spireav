package transform

import "fmt"

type Opacity struct {
	Opacity float64
}

func (node *Opacity) GetOutputsCount() int {
	return 1
}

func (node *Opacity) FilterString() string {
	return fmt.Sprintf("geq=r='r(X,Y)':a='%0.3f*alpha(X,Y)'", node.Opacity)
}
