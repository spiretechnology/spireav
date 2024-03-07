package spireav

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/output"
)

// Graph is a directed acyclic graph of nodes that transforms media inputs (audio and video files) into media outputs (audio or video files).
// The most common use of Graph is to transcode video, overlay timecode or watermarks, adjust resolutions, etc.
type Graph interface {
	FfmpegArgs
	Input(filename string) NodeReadable
	Filter(filter filter.Filter) NodeReadable
	Output(filename string, opts ...output.Option) Node
}

func New() Graph {
	return &implGraph{}
}

// graphLink is a connection between one node's output to another node's input
type graphLink struct {
	from         Stream
	toNode       Node
	toInputIndex int
}

type implGraph struct {
	inputs  []*inputNode
	filters []*filterNode
	outputs []output.Output
	links   []graphLink
}

func (g *implGraph) Input(filename string) NodeReadable {
	node := &inputNode{
		graph:      g,
		inputIndex: len(g.inputs),
		filename:   filename,
	}
	g.inputs = append(g.inputs, node)
	return node
}

func (g *implGraph) Filter(filter filter.Filter) NodeReadable {
	node := &filterNode{
		graph:    g,
		uniqueID: fmt.Sprintf("f%d", len(g.filters)),
		filter:   filter,
	}
	g.filters = append(g.filters, node)
	return node
}

func (g *implGraph) Output(filename string, opts ...output.Option) Node {
	out := output.New(filename, opts...)
	g.outputs = append(g.outputs, out)
	return out
}

// AddLink adds a link between nodes in the graph
func (g *implGraph) addLink(from Stream, toNode Node, toInputIndex int) {
	g.links = append(g.links, graphLink{
		from,
		toNode,
		toInputIndex,
	})
}

func (g *implGraph) getOutputMappings(node Node) []string {
	// Get all the links into the output node
	links := g.getLinksToNode(node)

	// Collect the output names
	var outputNames []string
	for _, link := range links {
		outputNames = append(outputNames, link.from.MapLabel())
	}
	return outputNames
}

func (g *implGraph) getLinksToNode(node Node) []graphLink {
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

func (g *implGraph) generateFiltersString(node *filterNode) string {
	// Get the links into the node
	linksIn := g.getLinksToNode(node)

	// Create the input names
	inputNames := []string{}
	for _, link := range linksIn {
		inputNames = append(inputNames, fmt.Sprintf("[%s]", link.from.Label()))
	}

	// Create the output names
	outputNames := []string{}
	for i := 0; i < node.filter.Outputs(); i++ {
		outputNames = append(outputNames, fmt.Sprintf("[%s]", node.Stream(i).Label()))
	}

	// Create the full filters string
	return strings.Join(inputNames, "") + node.filter.String() + strings.Join(outputNames, "")
}

func (g *implGraph) FilterString() string {
	// Slice for all of the individual filter strings
	var filterStrs []string

	// Loop through all of the nodes
	for _, node := range g.filters {
		// Get the node filter string
		filterStrs = append(filterStrs, g.generateFiltersString(node))
	}

	// Join all the filter strings
	return strings.Join(filterStrs, ";")
}

// FfmpegArgs implements the interface FfmpegArger and produces the slice of all arguments
// to be passed into the FFmpeg command to execute this graph
func (g *implGraph) FfmpegArgs() ([]string, error) {
	// A slice to hold all the arguments
	var args []string

	// Loop through all the inputs
	for _, in := range g.inputs {
		args = append(args, "-i", in.filename)
	}

	// Add the -y argument to skip interactive prompts
	args = append(args, "-y")

	// Add the filters string
	args = append(args, "-filter_complex", g.FilterString())

	// Loop through all the outputs. Apply the stream mappings, then all other output options
	for _, out := range g.outputs {
		for _, mapping := range g.getOutputMappings(out) {
			args = append(args, "-map", mapping)
		}
		args = append(args, out.Options()...)
	}

	// Return the slice of arguments
	return args, nil
}
