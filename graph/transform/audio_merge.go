package transform

import "fmt"

// AudioMerge is a transform that takes in N number of discrete audio streams and merges them into
// a single output stream with N channels.
type AudioMerge struct {
	Inputs int
}

func (node *AudioMerge) OutputsCount() int {
	return 1
}

func (node *AudioMerge) FilterString() string {
	return FormatTransform("amerge", map[string]string{
		"inputs": fmt.Sprintf("%d", node.Inputs),
	})
}
