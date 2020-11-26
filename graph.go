package spireav

import (
	"fmt"
	"sort"
	"strings"
)

type Link struct {
	fromNode        *GraphNode
	fromOutputIndex uint
	toNode          *GraphNode
	toInputIndex    uint
}

type GraphNode struct {
	nodeid      uint
	node        Node
	inputIndex  int
	outputIndex int
}

type Graph struct {
	nodes        []*GraphNode
	links        []Link
	inputsCount  uint
	outputsCount uint
}

// NewGraph allocates a new graph
func NewGraph() Graph {
	return Graph{
		nodes:        []*GraphNode{},
		links:        []Link{},
		inputsCount:  0,
		outputsCount: 0,
	}
}

// AddNode adds a node to the graph
func (graph *Graph) AddNode(node Node) *GraphNode {
	graphNode := &GraphNode{
		nodeid:      uint(len(graph.nodes)),
		node:        node,
		inputIndex:  -1,
		outputIndex: -1,
	}
	nodeType := node.Type()
	if nodeType == NodeTypeInput {
		graphNode.inputIndex = int(graph.inputsCount)
		graph.inputsCount++
	} else if nodeType == NodeTypeOutput {
		graphNode.outputIndex = int(graph.outputsCount)
		graph.outputsCount++
	}
	graph.nodes = append(graph.nodes, graphNode)
	return graphNode
}

// AddLink adds a link between nodes in the graph
func (graph *Graph) AddLink(fromNode *GraphNode, fromOutputIndex uint, toNode *GraphNode, toInputIndex uint) {
	link := Link{
		fromNode,
		fromOutputIndex,
		toNode,
		toInputIndex,
	}
	graph.links = append(graph.links, link)
}

func (graph *Graph) getLinksFromNode(node *GraphNode) []Link {

	// Get the links
	var links []Link
	for i := range graph.links {
		if graph.links[i].fromNode == node {
			links = append(links, graph.links[i])
		}
	}

	// Sort the links by input index
	sort.Slice(links, func(i, j int) bool {
		return links[i].fromOutputIndex < links[j].fromOutputIndex
	})

	// Return the links
	return links

}

func (graph *Graph) getLinksToNode(node *GraphNode) []Link {

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

func (graph *Graph) getInputTypes(node *GraphNode) ([]StreamType, error) {

	// Get the links into this node
	linksIn := graph.getLinksToNode(node)

	// Get the input types
	var inputTypes []StreamType
	for i := range linksIn {
		prevOutputIndex := linksIn[i].fromOutputIndex
		prevOutputTypes, err := linksIn[i].fromNode.node.GetOutputTypes()
		if err != nil {
			return nil, err
		}
		if prevOutputIndex >= uint(len(prevOutputTypes)) {
			return nil, fmt.Errorf("output #%d out of bounds for graph node", prevOutputIndex)
		}
		inputTypes = append(inputTypes, prevOutputTypes[prevOutputIndex])
	}

	// Return the types
	return inputTypes, nil

}

func (graph *Graph) generateFiltersString(node *GraphNode) (string, error) {

	// If the node is an output or an input
	nodeType := node.node.Type()
	if nodeType == NodeTypeInput || nodeType == NodeTypeOutput {
		return "", nil
	}

	// Get the input types
	inputTypes, err := graph.getInputTypes(node)
	if err != nil {
		return "", err
	}

	// Create the filter string for the node
	filterStr := node.node.FilterString(inputTypes)

	// Get the links into the node
	linksIn := graph.getLinksToNode(node)

	// Create the input names
	inputNames := ""
	for i := range linksIn {
		link := linksIn[i]
		inputNames += fmt.Sprintf("[node%d_%d]", link.fromNode.nodeid, link.fromOutputIndex)
	}

	// Get the output types
	outputTypes, err := node.node.GetOutputTypes()
	if err != nil {
		return "", err
	}

	// Create the output names
	outputNames := ""
	for i := range outputTypes {
		outputNames += fmt.Sprintf("[node%d_%d]", node.nodeid, i)
	}

	// Create the full filters string
	return inputNames + filterStr + outputNames, nil

}

func (graph *Graph) recursiveGenerateFiltersString(node *GraphNode, nodeids *[]uint) ([]string, error) {

	// If the node identifier is already in the list
	for i := range *nodeids {
		if (*nodeids)[i] == node.nodeid {
			return []string{}, nil
		}
	}

	// Add the node id to the slice
	*nodeids = append(*nodeids, node.nodeid)

	// Add the string for the output
	filterStr, err := graph.generateFiltersString(node)
	if err != nil {
		return nil, err
	}

	// Create the slice for inputs
	var inputStrs []string

	// Get the links into the node
	linksIn := graph.getLinksToNode(node)
	for i := range linksIn {
		upstreamInputStrs, err := graph.recursiveGenerateFiltersString(linksIn[i].fromNode, nodeids)
		if err != nil {
			return nil, err
		}
		inputStrs = append(inputStrs, upstreamInputStrs...)
	}

	// Add this node to the slice last
	if len(filterStr) > 0 {
		inputStrs = append(inputStrs, filterStr)
	}

	// Return the slice
	return inputStrs, nil

}

func (graph *Graph) generateFiltersStringForOutput(outputNode *GraphNode) (string, error) {

	// Keep track of the node ids already added to the filter string
	var nodeids []uint

	// Generate the filter string
	filtersStrs, err := graph.recursiveGenerateFiltersString(outputNode, &nodeids)
	if err != nil {
		return "", err
	}

	// Create the complete filter string
	return strings.Join(filtersStrs, ";"), nil

}

func (graph *Graph) ProduceOutput(outputNode *GraphNode) error {

	// Create the filters string
	filters, err := graph.generateFiltersStringForOutput(outputNode)
	if err != nil {
		return err
	}

	fmt.Println(filters)

	return nil

}
