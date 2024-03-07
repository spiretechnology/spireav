package filter

import (
	"fmt"
	"strings"

	"github.com/spiretechnology/spireav/graph/filter/expr"
)

// Filter is a node that applies a filter to its inputs, producing one or more outputs.
type Filter interface {
	String() string
	Outputs() int
}

// Option defines an output option middleware that populates some option string
// into the slice of options for the output.
type Option func() []OptionValue

// New creates a new transform with the given name and some options.
func New(name string, outputs int, opts ...Option) Filter {
	return &filterWithOpts{
		name:    name,
		outputs: outputs,
		opts:    WithMulti(opts...)(),
	}
}

// WithOption creates a new option middleware with a key and value.
func WithOption(key string, value expr.Expr) Option {
	return func() []OptionValue {
		return []OptionValue{
			{
				Key: key,
				Val: value,
			},
		}
	}
}

// WithMulti creates a new option middleware with multiple options.
func WithMulti(opts ...Option) Option {
	return func() []OptionValue {
		var options []OptionValue
		for _, opt := range opts {
			options = append(options, opt()...)
		}
		return options
	}
}

// OptionValue defines a single option for a transform node. This includes only a single
// key-value pair.
type OptionValue struct {
	Key string
	Val expr.Expr
}

type filterWithOpts struct {
	name    string
	outputs int
	opts    []OptionValue
}

func (t *filterWithOpts) String() string {
	opts := make(map[string]string)
	for _, optValue := range t.opts {
		opts[optValue.Key] = optValue.Val.String()
	}
	return formatTransform(t.name, opts)
}

func (t *filterWithOpts) Outputs() int {
	return t.outputs
}

// formatTransform formats a filter transform so that it conforms to the FFMpeg format
func formatTransform(filterName string, opts map[string]string) string {
	optsStr := formatTransformOptions(opts)
	if len(optsStr) == 0 {
		return filterName
	}
	return fmt.Sprintf("%s=%s", filterName, optsStr)
}

// formatTransformOptions converts a map of keys and values into a string of options
func formatTransformOptions(opts map[string]string) string {
	var keyvals []string
	for k, v := range opts {
		pair := fmt.Sprintf("%s=%s", k, formatTransformOptionValue(v))
		keyvals = append(keyvals, pair)
	}
	return strings.Join(keyvals, ":")
}

// formatTransformOptionValue formats a single value to ensure it's properly escaped and quote-wrapped
func formatTransformOptionValue(value string) string {
	if strings.Contains(value, ":") && !strings.HasPrefix(value, "'") {
		return fmt.Sprintf("'%s'", value)
	}
	return value
}
