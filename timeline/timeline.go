package timeline

import "github.com/spiretechnology/spireav"

// Timeline represents a timeline
type Timeline struct {
	children       []ChildNode
	positionFrames int64
	durationFrames int64
}

// GetFramePosition gets the temporal position of the node's first frame on the timeline, expressed in frames. This
// value is relative to the parent node
func (tl *Timeline) GetFramePosition() int64 {
	return tl.positionFrames
}

// SetFramePosition sets the temporal position of the node relative to its parent node
func (tl *Timeline) SetFramePosition(frame int64) {
	tl.positionFrames = frame
}

// GetFrameDuration gets the duration of the node on the timeline, expressed in frames. This value is relative to
// the parent node
func (tl *Timeline) GetFrameDuration() int64 {
	return tl.durationFrames
}

// SetFrameDuration sets the duration of the node within its parent node
func (tl *Timeline) SetFrameDuration(frames int64) {
	tl.durationFrames = frames
}

// GetChildren gets the children nodes of this node
func (tl *Timeline) GetChildren() []ChildNode {
	return tl.children
}

// AddChild adds a child node to this node
func (tl *Timeline) AddChild(child ChildNode) {
	tl.children = append(tl.children, child)
}

// Source gets the input source for the node
func (tl *Timeline) Source() *spireav.FormatContext {
	return nil
}
