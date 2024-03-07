package graph

import (
	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/graph/filter"
	"github.com/spiretechnology/spireav/graph/output"
)

// Graph is a directed acyclic graph of nodes that transforms media inputs (audio and video files) into media outputs (audio or video files).
// The most common use of Graph is to transcode video, overlay timecode or watermarks, adjust resolutions, etc.
type Graph interface {
	spireav.FfmpegArger
	NewInput(filename string) NodeReadable
	NewFilter(filter filter.Filter) NodeReadable
	NewOutput(filename string, opts ...output.Option) Node
}

func New() Graph {
	return &implGraph{}
}
