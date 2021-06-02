package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
//#include <libavfilter/buffersink.h>
//#include <libavfilter/buffersrc.h>
//#include <libavutil/pixdesc.h>
import "C"
import (
	"fmt"
	"sort"
	"strings"
	"unsafe"

	"github.com/spiretechnology/spireav"
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

func (graph *Graph) getInputTypes(node *GraphNode) ([]spireav.StreamType, error) {

	// Get the links into this node
	linksIn := graph.getLinksToNode(node)

	// Get the input types
	var inputTypes []spireav.StreamType
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

func (graph *Graph) formatNodeOutputNameClean(nodeid uint, nodeOutputIndex uint) string {
	return fmt.Sprintf("node%d_%d", nodeid, nodeOutputIndex)
}

func (graph *Graph) formatNodeOutputName(nodeid uint, nodeOutputIndex uint) string {
	return fmt.Sprintf("[%s]", graph.formatNodeOutputNameClean(nodeid, nodeOutputIndex))
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
		inputNames += graph.formatNodeOutputName(link.fromNode.nodeid, link.fromOutputIndex)
	}

	// Get the output types
	outputTypes, err := node.node.GetOutputTypes()
	if err != nil {
		return "", err
	}

	// Create the output names
	outputNames := ""
	for i := range outputTypes {
		outputNames += graph.formatNodeOutputName(node.nodeid, uint(i))
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

func (graph *Graph) generateOverallFiltersString() (string, error) {

	// Create the slice of filters strings
	filtersStrs := []string{}

	// Loop through the output nodes
	for _, node := range graph.nodes {
		if _, ok := node.node.(OutputNode); ok {
			filtersStr, err := graph.generateFiltersStringForOutput(node)
			if err != nil {
				return "", err
			}
			filtersStrs = append(filtersStrs, filtersStr)
		}
	}

	// Join the filters string
	return strings.Join(filtersStrs, ";"), nil

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

func (graph *Graph) getStreamTypesIntoOutput(outputNode *GraphNode) []spireav.StreamType {

	// Slice of stream types in the output
	outputStreamTypes := []spireav.StreamType{}

	// Get the links leading into this node
	linksIntoOutput := graph.getLinksToNode(outputNode)

	// Loop through the links into the node
	for _, link := range linksIntoOutput {

		// Get the output types on the FROM node
		fromTypes, err := link.fromNode.node.GetOutputTypes()
		if err != nil {
			fmt.Println("Error getting node output types: ", err)
			continue
		}

		// Get the type at the index
		fromType := fromTypes[link.fromOutputIndex]

		// Add the type to the slice
		outputStreamTypes = append(outputStreamTypes, fromType)

	}

	// Return the slice of types
	return outputStreamTypes

}

// getInputNodes gets the input nodes to this transcoding graph
func (graph *Graph) getInputNodes() []InputNode {
	inputNodes := []InputNode{}
	for _, node := range graph.nodes {
		inputNode, ok := node.node.(InputNode)
		if ok {
			inputNodes = append(inputNodes, inputNode)
		}
	}
	return inputNodes
}

// getOutputNodes gets the output nodes from this transcoding graph
func (graph *Graph) getOutputNodes() []OutputNode {
	outputNodes := []OutputNode{}
	for _, node := range graph.nodes {
		outputNode, ok := node.node.(OutputNode)
		if ok {
			outputNodes = append(outputNodes, outputNode)
		}
	}
	return outputNodes
}

func (graph *Graph) NewJob() *GraphJob {
	return &GraphJob{
		graph: graph,
	}
}

func CopyStreamSettings(from *InputStream, to *OutputStream) {

	// If it's a video
	if from.decCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
		to.encCtx.width = from.decCtx.width
		to.encCtx.height = from.decCtx.height
		to.encCtx.sample_aspect_ratio = from.decCtx.sample_aspect_ratio
		if to.encoder.avCodec.pix_fmts != nil {
			to.encCtx.pix_fmt = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(to.encoder.avCodec.pix_fmts))))
		} else {
			to.encCtx.pix_fmt = from.decCtx.pix_fmt
		}
		to.encCtx.framerate = from.decCtx.framerate
		to.encCtx.time_base = C.av_inv_q(to.encCtx.framerate)
		to.avStream.time_base = to.encCtx.time_base
	} else {
		to.encCtx.sample_rate = from.decCtx.sample_rate
		to.encCtx.channel_layout = from.decCtx.channel_layout
		to.encCtx.channels = C.av_get_channel_layout_nb_channels(to.encCtx.channel_layout)

		// encCtx.sample_fmt = encoder.sample_fmts[0]
		to.encCtx.sample_fmt = from.decCtx.sample_fmt
		to.encCtx.time_base = C.struct_AVRational{
			num: 1,
			den: to.encCtx.sample_rate,
		}
	}

}
