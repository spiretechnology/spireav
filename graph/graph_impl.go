package graph

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spiretechnology/spireav/graph/input"
	"github.com/spiretechnology/spireav/graph/output"
	"github.com/spiretechnology/spireav/graph/transform"
)

// Ensure that *graphImpl correctly implements Graph.
var _ Graph = (*graphImpl)(nil)

// graphLink is a connection between one node's output to another node's input
type graphLink struct {
	fromNode        Node
	fromOutputIndex int
	toNode          Node
	toInputIndex    int
}

type graphImpl struct {
	inputs     []input.Input
	outputs    []output.Output
	transforms []transform.Transform
	links      []graphLink
}

func (g *graphImpl) getInputIndex(node input.Input) int {
	for i := range g.inputs {
		if g.inputs[i] == node {
			return i
		}
	}
	return -1
}

func (g *graphImpl) getTransformIndex(node transform.Transform) int {
	for i := range g.transforms {
		if g.transforms[i] == node {
			return i
		}
	}
	return -1
}

// NewOutput creates a new input on the g.
func (g *graphImpl) NewInput(filename string) input.Input {
	return g.AddInput(input.New(filename))
}

// AddInput adds an input node to this transcoding g. The node data will be read into the g.
func (g *graphImpl) AddInput(node input.Input) input.Input {
	g.inputs = append(g.inputs, node)
	return node
}

// AddTransform adds a transform node to this transcoding graph
func (g *graphImpl) AddTransform(node transform.Transform) transform.Transform {
	g.transforms = append(g.transforms, node)
	return node
}

// NewOutput creates a new input on the g.
func (g *graphImpl) NewOutput(filename string, opts ...output.Option) output.Output {
	return g.AddOutput(output.New(filename, opts...))
}

// AddOutput adds an output node to this transcoding g. The output will be written to when the graph is executed.
func (g *graphImpl) AddOutput(node output.Output) output.Output {
	g.outputs = append(g.outputs, node)
	return node
}

// AddLink adds a link between nodes in the graph
func (g *graphImpl) AddLink(fromNode Node, fromOutputIndex int, toNode Node, toInputIndex int) {
	link := graphLink{
		fromNode,
		fromOutputIndex,
		toNode,
		toInputIndex,
	}
	g.links = append(g.links, link)
}

func (g *graphImpl) getOutputMappings(node output.Output) []string {

	// Get all the links into the output node
	links := g.getLinksToNode(node)

	// Collect the output names
	outputNames := []string{}
	for _, link := range links {
		outputNames = append(outputNames, g.formatNodeOutputName(link.fromNode, link.fromOutputIndex, true))
	}

	// Return the names
	return outputNames

}

func (g *graphImpl) getLinksToNode(node Node) []graphLink {

	// Get the links
	var links []graphLink
	for i := range g.links {
		if g.links[i].toNode == node {
			links = append(links, g.links[i])
		}
	}

	// Sort the links by input index
	sort.Slice(links, func(i, j int) bool {
		return links[i].toInputIndex < links[j].toInputIndex
	})

	// Return the links
	return links

}

func (g *graphImpl) formatNodeOutputName(
	node Node,
	nodeOutputIndex int,
	mapping bool,
) string {
	if in, ok := node.(input.Input); ok {
		if mapping {
			return fmt.Sprintf("%d:%d", g.getInputIndex(in), nodeOutputIndex)
		} else {
			return fmt.Sprintf("[%d:%d]", g.getInputIndex(in), nodeOutputIndex)
		}
	} else if tn, ok := node.(transform.Transform); ok {
		return fmt.Sprintf("[%s%d_%d]", nodeNamePrefix, g.getTransformIndex(tn), nodeOutputIndex)
	}
	return ""
}

func (g *graphImpl) generateFiltersString(node transform.Transform) string {

	// Create the filter string for the node
	filterStr := node.FilterString()

	// Get the links into the node
	linksIn := g.getLinksToNode(node)

	// Create the input names
	inputNames := []string{}
	for _, link := range linksIn {
		inputNames = append(inputNames, g.formatNodeOutputName(link.fromNode, link.fromOutputIndex, false))
	}

	// Create the output names
	outputNames := []string{}
	for i := 0; i < node.OutputsCount(); i++ {
		outputNames = append(outputNames, g.formatNodeOutputName(node, i, false))
	}

	// Create the full filters string
	return strings.Join(inputNames, "") + filterStr + strings.Join(outputNames, "")

}

// filterGraphString generates the complete filtergraph string to be passed into FFmpeg
func (g *graphImpl) filterGraphString() string {

	// Slice for all of the individual filter strings
	filterStrs := []string{}

	// Loop through all of the nodes
	for _, node := range g.transforms {

		// Get the node filter string
		filterStrs = append(filterStrs, g.generateFiltersString(node))

	}

	// Join all the filter strings
	return strings.Join(filterStrs, ";")

}

// FfmpegArgs implements the interface FfmpegArger and produces the slice of all arguments
// to be passed into the FFmpeg command to execute this graph
func (g *graphImpl) FfmpegArgs() ([]string, error) {

	// A slice to hold all the arguments
	args := []string{}

	// Loop through all the inputs
	for _, in := range g.inputs {
		args = append(args, "-i", in.Filename())
	}

	// Add the -y argument to skip interactive prompts
	args = append(args, "-y")

	// Add the filters string
	args = append(args, "-filter_complex", g.filterGraphString())

	// Loop through all the outputs
	for _, out := range g.outputs {

		// Get the mappings for the output
		for _, mapping := range g.getOutputMappings(out) {
			args = append(args, "-map", mapping)
		}

		// Add all the other options for the output
		args = append(args, out.Options()...)

	}

	// Return the slice of arguments
	return args, nil

}
