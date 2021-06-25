package transform

import "fmt"

type AudioMerge struct {
	Inputs int
}

func (node *AudioMerge) GetOutputsCount() int {
	return 1
}

func (node *AudioMerge) FilterString() string {
	return FormatTransform("amerge", map[string]string{
		"inputs": fmt.Sprintf("%d", node.Inputs),
	})
}
