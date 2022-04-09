package transform

// Transform is a node that applies a filter to its inputs, producing one or more outputs.
type Transform interface {
	FilterString() string
	OutputsCount() int
}
