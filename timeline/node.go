package timeline

import "github.com/spiretechnology/spireav"

// ChildNode represents a node that is a child of another node, and all of the data needed to manage that
// embedded state
type ChildNode struct {
	// Node the node that is a child of the context node
	Node Node
	// LayerIndex is a numeric value which determines the z-index of nodes
	LayerIndex int
}

// Node defines the interface for a node on a timeline
type Node interface {

	// GetFramePosition gets the temporal position of the node's first frame on the timeline, expressed in frames. This
	// value is relative to the parent node
	GetFramePosition() int64

	// SetFramePosition sets the temporal position of the node relative to its parent node
	SetFramePosition(frame int64)

	// GetFrameDuration gets the duration of the node on the timeline, expressed in frames. This value is relative to
	// the parent node
	GetFrameDuration() int64

	// SetFrameDuration sets the duration of the node within its parent node
	SetFrameDuration(frames int64)

	// GetChildren gets the children nodes of this node
	GetChildren() []ChildNode

	// AddChild adds a child node to this node
	AddChild(child ChildNode)

	// Source gets the source for the node, if there is one. This is a relevant media source, such as a video file
	Source() *spireav.FormatContext
}
