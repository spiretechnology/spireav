package render

import (
	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/timeline"
)

// Context is the context for a timeline render. It accumulates all of the needed information about the timeline
// needed to produce the output
type Context struct {
	// Sources is a slice of all sources for the timeline
	Sources []*spireav.FormatContext
}

func NewContextFromTimeline(timeline *timeline.Timeline) (*Context, error) {

	// If the timeline is nil, for some reason
	if timeline == nil {
		return nil, nil
	}

	// Create the context
	ctx := &Context{
		Sources: []*spireav.FormatContext{},
	}

	// Recursively accumulate the context
	err := accumulateContext(timeline, ctx)
	if err != nil {
		return nil, err
	}

	// Return the context
	return ctx, nil

}

func accumulateContext(node timeline.Node, ctx *Context) error {

	// Loop through the child nodes
	for _, child := range node.GetChildren() {

		// Add the child to the context
		err := addChildToContext(child, ctx)
		if err != nil {
			return err
		}

		// Accumulate the context on the child
		err = accumulateContext(child.Node, ctx)
		if err != nil {
			return err
		}

	}

	// Return no error
	return nil

}

func addChildToContext(child timeline.ChildNode, ctx *Context) error {

	// Get the source for the child node
	source := child.Node.Source()
	if source != nil {
		ctx.Sources = append(ctx.Sources, source)
	}

	// Return no error
	return nil

}
