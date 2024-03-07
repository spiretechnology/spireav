package spireav

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/output"
)

// Graph is a directed acyclic graph of nodes that transforms media inputs (audio and video files) into media outputs (audio or video files).
// The most common use of Graph is to transcode video, overlay timecode or watermarks, adjust resolutions, etc.
type Graph interface {
	FfmpegArgs
	Input(filename string) NodeReadable
	Filter(filter filter.Filter) NodeReadable
	Output(filename string, opts ...output.Option) Node
}

func New() Graph {
	return &implGraph{}
}
