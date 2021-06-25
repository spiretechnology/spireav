package graph

import (
	"fmt"
	"strings"
)

type TextOverlayNode struct {
	TransformNode
	Text string
}

func (node *TextOverlayNode) GetOutputsCount() int {
	return 1
}

func (node *TextOverlayNode) FilterString() string {
	config := make(map[string]string)
	config["text"] = node.Text
	configParts := []string{}
	for k, v := range config {
		configParts = append(configParts, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("drawtext=%s", strings.Join(configParts, ":"))
}
