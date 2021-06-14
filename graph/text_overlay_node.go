package graph

import (
	"fmt"
	"strings"

	"github.com/spiretechnology/spireav"
)

type TextOverlayNode struct {
	TransformNode
	Text string
}

func (node TextOverlayNode) GetOutputTypes() ([]spireav.StreamType, error) {
	return []spireav.StreamType{spireav.StreamTypeVideo}, nil
}

func (node TextOverlayNode) FilterString(inputs []spireav.StreamType) string {
	config := make(map[string]string)
	config["text"] = node.Text
	configParts := []string{}
	for k, v := range config {
		configParts = append(configParts, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("drawtext=%s", strings.Join(configParts, ":"))
}
