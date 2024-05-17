package spireav

// Pipeline creates a new pipeline of nodes. The pipeline returned can, itself, be treated as
// a node in the graph and piped to other nodes.
func Pipeline(nodes ...Pipeable) NodeReadable {
	if len(nodes) == 0 {
		return nil
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Pipe(nodes[i+1], 0)
	}
	return &implPipeline{
		left:  nodes[0],
		right: nodes[len(nodes)-1],
	}
}

type implPipeline struct {
	left, right Pipeable
}

func (p *implPipeline) Pipe(to Node, toIndex int) {
	p.right.Pipe(to, toIndex)
}

func (p *implPipeline) Stream(index int) Stream {
	switch t := p.right.(type) {
	case NodeReadable:
		return t.Stream(index)
	case Stream:
		return t
	default:
		return nil
	}
}
