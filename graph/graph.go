package graph

import (
	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/graph/input"
	"github.com/spiretechnology/spireav/graph/output"
	"github.com/spiretechnology/spireav/graph/transform"
)

// nodeNamePrefix is a prefix string added to the generated FFmpeg command to make filter outputs unique
const nodeNamePrefix = "spire"

// Node represents an input, output, or transform node in the graph.
type Node any

// Graph is a directed acyclic graph of nodes that transforms media inputs (audio and video files) into media outputs (audio or video files).
// The most common use of Graph is to transcode video, overlay timecode or watermarks, adjust resolutions, etc.
type Graph interface {
	spireav.FfmpegArger
	NewInput(filename string) input.Input
	AddInput(node input.Input) input.Input
	AddTransform(node transform.Transform) transform.Transform
	NewOutput(filename string, opts ...output.Option) output.Output
	AddOutput(node output.Output) output.Output
	AddLink(fromNode Node, fromOutputIndex int, toNode Node, toInputIndex int)
}

// New creates a new transcoding graph.
func New() Graph {
	return &graphImpl{}
}
