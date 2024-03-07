package graph

type Node interface{}

type NodeReadable interface {
	Node
	Stream(index int) Stream
}

type Stream interface {
	Node() Node
	Index() int
	Label() string
	MapLabel() string
	Pipe(to Node, toIndex int)
}
