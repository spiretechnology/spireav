package render

import (
	"fmt"

	"github.com/spiretechnology/spireav/timeline"
)

// Renderer is a type for generating renders of timelines to a flat output
type Renderer struct {
	// Timeline the timeline to render
	Timeline *timeline.Timeline
	// Destination the destination to which to write the output
	Destination Destination
}

// Render renders the timeline to its output destination
func (r *Renderer) Render() error {

	// Construct a context for the timeline
	ctx, err := NewContextFromTimeline(r.Timeline)
	if err != nil {
		return err
	}

	//
	fmt.Println("Sources: ", ctx.Sources)

	// Return no error
	return nil

}
