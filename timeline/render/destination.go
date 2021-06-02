package render

import "github.com/spiretechnology/spireav"

// Destination defines the interface for a timeline render destination
type Destination interface {
	// GetFormatContext gets the format context to which the render will be written
	GetFormatContext() *spireav.FormatContext
}
