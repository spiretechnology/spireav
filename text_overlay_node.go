package spireav

import (
	"fmt"
	"strings"
)

type TextOverlayNode struct {
	Text string
}

func (node TextOverlayNode) GetOutputTypes() ([]StreamType, error) {
	return []StreamType{StreamTypeVideo}, nil
}

func (node TextOverlayNode) FilterString(inputs []StreamType) string {
	config := make(map[string]string)
	config["text"] = node.Text
	configParts := []string{}
	for k, v := range config {
		configParts = append(configParts, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("drawtext=%s", strings.Join(configParts, ":"))
}

func (node TextOverlayNode) Type() NodeType {
	return NodeTypeTransform
}
