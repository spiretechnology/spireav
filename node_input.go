package spireav

import "fmt"

type inputNode struct {
	graph      *implGraph
	inputIndex int
	filename   string
}

func (in *inputNode) Stream(index int) Stream {
	return &inputNodeStream{
		node:  in,
		index: index,
	}
}

func (in *inputNode) Video(index int) Stream {
	return &inputNodeStream{
		node:      in,
		typeLabel: "v",
		index:     index,
	}
}

func (in *inputNode) Audio(index int) Stream {
	return &inputNodeStream{
		node:      in,
		typeLabel: "a",
		index:     index,
	}
}

func (in *inputNode) Pipe(to Node, toIndex int) {
	in.Stream(0).Pipe(to, toIndex)
}

type inputNodeStream struct {
	node      *inputNode
	index     int
	typeLabel string
}

func (ins *inputNodeStream) Node() Node {
	return ins.node
}

func (ins *inputNodeStream) Index() int {
	return ins.index
}

func (ins *inputNodeStream) Label() string {
	if ins.typeLabel != "" {
		return fmt.Sprintf("%d:%s:%d", ins.node.inputIndex, ins.typeLabel, ins.index)
	}
	return fmt.Sprintf("%d:%d", ins.node.inputIndex, ins.index)
}

func (ins *inputNodeStream) MapLabel() string {
	return ins.Label()
}

func (ins *inputNodeStream) Pipe(to Node, toIndex int) {
	ins.node.graph.addLink(ins, to, toIndex)
}
