package graph

import (
	"fmt"

	"github.com/spiretechnology/spireav/graph/filter"
)

type filterNode struct {
	graph    *implGraph
	uniqueID string
	filter   filter.Filter
}

func (fn *filterNode) Stream(index int) Stream {
	return &filterNodeStream{
		node:  fn,
		index: index,
	}
}

type filterNodeStream struct {
	node  *filterNode
	index int
}

func (fns *filterNodeStream) Node() Node {
	return fns.node
}

func (fns *filterNodeStream) Index() int {
	return fns.index
}

func (fns *filterNodeStream) Label() string {
	return fmt.Sprintf("%s_%d", fns.node.uniqueID, fns.index)
}

func (fns *filterNodeStream) MapLabel() string {
	return fmt.Sprintf("[%s_%d]", fns.node.uniqueID, fns.index)
}

func (fns *filterNodeStream) Pipe(to Node, toIndex int) {
	fns.node.graph.addLink(fns, to, toIndex)
}
