package output

// Output is a node to which a graph outputs final stream data. Output is usually a file. In ffmpeg terms,
// this is a buffersink.
type Output interface {
	Options() []string
}

// Option defines an output option middleware that populates some option string
// into the slice of options for the output.
type Option func() []string

// New creates a new output node with some slice of options.
func New(filename string, opts ...Option) Output {
	options := withMulti(opts...)()
	options = append(options, filename)
	return &outputWithArgs{options}
}

type outputWithArgs struct {
	opts []string
}

func (o *outputWithArgs) Options() []string {
	return o.opts
}

func WithOptions(opts ...string) Option {
	return func() []string {
		return opts
	}
}

func withMulti(opts ...Option) Option {
	return func() []string {
		var options []string
		for _, opt := range opts {
			options = append(options, opt()...)
		}
		return options
	}
}
