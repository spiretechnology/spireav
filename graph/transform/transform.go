package transform

import "github.com/spiretechnology/spireav/graph/transform/expr"

// Transform is a node that applies a filter to its inputs, producing one or more outputs.
type Transform interface {
	FilterString() string
	OutputsCount() int
}

// OptionValue defines a single option for a transform node. This includes only a single
// key-value pair.
type OptionValue struct {
	Key string
	Val expr.Expr
}

type transformWithOpts struct {
	name        string
	opts        []OptionValue
	outputCount int
}

func (t *transformWithOpts) FilterString() string {
	opts := make(map[string]string)
	for _, optValue := range t.opts {
		opts[optValue.Key] = optValue.Val.String()
	}
	return FormatTransform(t.name, opts)
}

func (t *transformWithOpts) OutputsCount() int {
	return t.outputCount
}

// New creates a new transform with the given name and some options.
func New(name string, outputs int, opts ...Option) Transform {
	return &transformWithOpts{
		name:        name,
		outputCount: outputs,
		opts:        withMulti(opts...)(),
	}
}

// Option defines an output option middleware that populates some option string
// into the slice of options for the output.
type Option func() []OptionValue

func withOption(key string, value expr.Expr) Option {
	return func() []OptionValue {
		return []OptionValue{
			{
				Key: key,
				Val: value,
			},
		}
	}
}

func withMulti(opts ...Option) Option {
	return func() []OptionValue {
		var options []OptionValue
		for _, opt := range opts {
			options = append(options, opt()...)
		}
		return options
	}
}
