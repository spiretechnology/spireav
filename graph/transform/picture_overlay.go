package transform

import "fmt"

type PictureOverlay struct {
	X          string
	Y          string
	EnableExpr string
}

func (node *PictureOverlay) GetOutputsCount() int {
	return 1
}

func (node *PictureOverlay) FilterString() string {

	// Create the basic filter string
	filter := fmt.Sprintf("overlay=%s:%s", node.X, node.Y)

	// If there is an enable expression
	if len(node.EnableExpr) > 0 {
		filter += fmt.Sprintf(":enable=%s", formatTransformOptionValue(node.EnableExpr))
	}

	// Return the filter string
	return filter

}
