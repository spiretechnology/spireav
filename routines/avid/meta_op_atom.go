package avid

import (
	"errors"

	"github.com/spiretechnology/spireav/meta"
)

// AvidMxfMeta is the metadata specific to AVID that can be gleaned from looking at the
// metadata pulled directly from the file.
//
// Reel is the equivalent of Spire's "tape"
// Material Package is the equivalent of Spire's "clip"
type AvidMxfMeta struct {
	ReelName             string
	ReelUmid             string
	MaterialPackageName  string
	MaterialPackageUmid  string
	MediaComposerVersion string
	CameraManufacturer   string
	CameraModel          string
	CameraSerialNum      string
	EssenceStream        *meta.StreamMeta
}

// ParseAvidMxfOpAtomMeta parses a single op-atom file's metadata and returns the Avid-specific identifiers
// and names for the tape and clip (reel and material package)
func ParseAvidMxfOpAtomMeta(metadata *meta.Meta) (*AvidMxfMeta, error) {

	// Get the stream that actually contains essence data, and thusly is represented by this Op-Atom file
	stream := findOpAtomStreamWithEssenceData(metadata)
	if stream == nil {
		return nil, errors.New("op-atom MXF file contains no stream with essence")
	}

	// Create the starting point for the metadata
	avidMeta := AvidMxfMeta{
		ReelName:             stream.Tags["reel_name"],
		ReelUmid:             stream.Tags["reel_umid"],
		MaterialPackageName:  metadata.Format.Tags["material_package_name"],
		MaterialPackageUmid:  metadata.Format.Tags["material_package_umid"],
		MediaComposerVersion: metadata.Format.Tags["product_name"],
		CameraManufacturer:   metadata.Format.Tags["comment_manufacturer"],
		CameraModel:          metadata.Format.Tags["comment_modelName"],
		CameraSerialNum:      metadata.Format.Tags["comment_serialNo"],
		EssenceStream:        stream,
	}

	// Return the metadata
	return &avidMeta, nil

}

// findOpAtomStreamWithEssenceData finds the stream in the metadata that contains actual essence data.
// In the Op-Atom format, only one stream contains essence, and the others are empty and contain metadata only
func findOpAtomStreamWithEssenceData(metadata *meta.Meta) *meta.StreamMeta {
	for _, stream := range metadata.Streams {
		if stream.CodecType == "audio" || stream.CodecType == "video" {
			return &stream
		}
	}
	return nil
}
