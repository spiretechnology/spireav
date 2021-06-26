package graph

import (
	"fmt"
	"sort"
	"strings"
)

// nodeNamePrefix is a prefix string added to the generated FFmpeg command to make filter outputs unique
const nodeNamePrefix = "spire"

type Node interface{}

type InputNode interface {
	Node
	GetFilename() string
}

type TransformNode interface {
	Node
	FilterString() string
	GetOutputsCount() int
}

type OutputNode interface {
	Node
	GetFilename() string
	GetOptions() []string
}

type Link struct {
	fromNode        Node
	fromOutputIndex int
	toNode          Node
	toInputIndex    int
}

type Graph struct {
	inputs     []InputNode
	outputs    []OutputNode
	transforms []TransformNode
	links      []Link
}

func (graph *Graph) getInputIndex(node InputNode) int {
	for i := range graph.inputs {
		if graph.inputs[i] == node {
			return i
		}
	}
	return -1
}

func (graph *Graph) getTransformIndex(node TransformNode) int {
	for i := range graph.transforms {
		if graph.transforms[i] == node {
			return i
		}
	}
	return -1
}

func (graph *Graph) AddInput(node InputNode) InputNode {
	graph.inputs = append(graph.inputs, node)
	return node
}

func (graph *Graph) AddTransform(node TransformNode) TransformNode {
	graph.transforms = append(graph.transforms, node)
	return node
}

func (graph *Graph) AddOutput(node OutputNode) OutputNode {
	graph.outputs = append(graph.outputs, node)
	return node
}

// AddLink adds a link between nodes in the graph
func (graph *Graph) AddLink(fromNode Node, fromOutputIndex int, toNode Node, toInputIndex int) {
	link := Link{
		fromNode,
		fromOutputIndex,
		toNode,
		toInputIndex,
	}
	graph.links = append(graph.links, link)
}

func (graph *Graph) GetOutputMappings(node OutputNode) []string {

	// Get all the links into the output node
	links := graph.getLinksToNode(node)

	// Collect the output names
	outputNames := []string{}
	for _, link := range links {
		outputNames = append(outputNames, graph.formatNodeOutputName(link.fromNode, link.fromOutputIndex, true))
	}

	// Return the names
	return outputNames

}

func (graph *Graph) getLinksToNode(node Node) []Link {

	// Get the links
	var links []Link
	for i := range graph.links {
		if graph.links[i].toNode == node {
			links = append(links, graph.links[i])
		}
	}

	// Sort the links by input index
	sort.Slice(links, func(i, j int) bool {
		return links[i].toInputIndex < links[j].toInputIndex
	})

	// Return the links
	return links

}

func (graph *Graph) formatNodeOutputName(
	node Node,
	nodeOutputIndex int,
	mapping bool,
) string {
	if in, ok := node.(InputNode); ok {
		if mapping {
			return fmt.Sprintf("%d:%d", graph.getInputIndex(in), nodeOutputIndex)
		} else {
			return fmt.Sprintf("[%d:%d]", graph.getInputIndex(in), nodeOutputIndex)
		}
	} else if tn, ok := node.(TransformNode); ok {
		return fmt.Sprintf("[%s%d_%d]", nodeNamePrefix, graph.getTransformIndex(tn), nodeOutputIndex)
	}
	return ""
}

func (graph *Graph) generateFiltersString(node TransformNode) string {

	// Create the filter string for the node
	filterStr := node.FilterString()

	// Get the links into the node
	linksIn := graph.getLinksToNode(node)

	// Create the input names
	inputNames := []string{}
	for _, link := range linksIn {
		inputNames = append(inputNames, graph.formatNodeOutputName(link.fromNode, link.fromOutputIndex, false))
	}

	// Create the output names
	outputNames := []string{}
	for i := 0; i < node.GetOutputsCount(); i++ {
		outputNames = append(outputNames, graph.formatNodeOutputName(node, i, false))
	}

	// Create the full filters string
	return strings.Join(inputNames, "") + filterStr + strings.Join(outputNames, "")

}

func (graph *Graph) FilterString() string {

	// Slice for all of the individual filter strings
	filterStrs := []string{}

	// Loop through all of the nodes
	for _, node := range graph.transforms {

		// Get the node filter string
		filterStrs = append(filterStrs, graph.generateFiltersString(node))

	}

	// Join all the filter strings
	return strings.Join(filterStrs, ";")

}

// GetInputs gets the input nodes to this transcoding graph
func (graph *Graph) GetInputs() []InputNode {
	return graph.inputs
}

// GetOutputs gets the output nodes from this transcoding graph
func (graph *Graph) GetOutputs() []OutputNode {
	return graph.outputs
}
