package spireav

// Node is a node in a graph.
type Node interface{}

// Pipeable is a node in a graph that can be piped to another node.
type Pipeable interface {
	// Pipe pipes the first output stream from this node to an input slot on another node.
	Pipe(to Node, toIndex int)
}

// NodeReadable is a node in a graph that has readable streams. These streams can be
// piped into subsequent nodes in the graph.
type NodeReadable interface {
	Node
	Pipeable
	// Stream returns the stream at the specified index.
	Stream(index int) Stream
}

// InputNodeReadable is a node in a graph that is an input node and has readable streams.
type InputNodeReadable interface {
	NodeReadable
	Video(index int) Stream
	Audio(index int) Stream
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
