package timeline

import (
	"fmt"

	"github.com/spiretechnology/spireav"
)

// SourceNode defines a node that includes a media source
type SourceNode struct {
	FormatContext  *spireav.FormatContext
	positionFrames int64
	durationFrames int64
}

// GetFramePosition gets the temporal position of the node's first frame on the timeline, expressed in frames. This
// value is relative to the parent node
func (sn *SourceNode) GetFramePosition() int64 {
	return sn.positionFrames
}

// SetFramePosition sets the temporal position of the node relative to its parent node
func (sn *SourceNode) SetFramePosition(frame int64) {
	sn.positionFrames = frame
}

// GetFrameDuration gets the duration of the node on the timeline, expressed in frames. This value is relative to
// the parent node
func (sn *SourceNode) GetFrameDuration() int64 {
	return sn.durationFrames
}

// SetFrameDuration sets the duration of the node within its parent node
func (sn *SourceNode) SetFrameDuration(frames int64) {
	sn.durationFrames = frames
}

// GetChildren gets the children nodes of this node
func (sn *SourceNode) GetChildren() []ChildNode {
	return nil
}

// AddChild adds a child node to this node
func (sn *SourceNode) AddChild(child ChildNode) {
	fmt.Println("WARN: cannot add child to a source node")
}

// Source gets the input source for the node
func (sn *SourceNode) Source() *spireav.FormatContext {
	return sn.FormatContext
}
