package spireav

import (
	"fmt"

	"github.com/spiretechnology/spireav/filter"
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

func (fn *filterNode) Pipe(to Node, toIndex int) {
	fn.Stream(0).Pipe(to, toIndex)
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
