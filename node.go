package spireav

// Node is a node in a graph.
type Node interface{}

// NodeReadable is a node in a graph that has readable streams. These streams can be
// piped into subsequent nodes in the graph.
type NodeReadable interface {
	Node
	// Stream returns the stream at the specified index.
	Stream(index int) Stream
	// Pipe pipes the first output stream from this node to an input slot on another node.
	Pipe(to Node, toIndex int)
}

// Stream is a single stream of data that can be piped into a subsequent node in the graph.
// A stream is usually video or audio data. Nodes can input and output multiple streams.
type Stream interface {
	Node() Node
	Index() int
	Label() string
	MapLabel() string
	Pipe(to Node, toIndex int)
}
