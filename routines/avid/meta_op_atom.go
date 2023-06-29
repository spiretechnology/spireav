package avid

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spiretechnology/spireav"
	"github.com/spiretechnology/spireav/mxf2raw"
)

// AvidMxfMeta is the metadata specific to AVID that can be gleaned from looking at the
// metadata pulled directly from the file.
//
// Reel is the equivalent of Spire's "tape"
// Material Package is the equivalent of Spire's "clip"
type AvidMxfMeta struct {
	FilePackageUmid      string
	ReelName             string
	ReelUmid             string
	MaterialPackageName  string
	MaterialPackageUmid  string
	MediaComposerVersion string
	CameraManufacturer   string
	CameraModel          string
	CameraSerialNum      string
	LensModelName        string
	Timecode             string
	CreationDate         *time.Time
	EssenceStream        *spireav.StreamMeta
}

// GetMetadata gets the metadata for the given MXF file. It uses two methods for getting the metadata (ffprobe and mxf2raw)
// and combines the results to get the most complete metadata possible.
func GetMetadata(ctx context.Context, filename string) (*AvidMxfMeta, *spireav.Meta, error) {
	// Get the metadata for the file using the usual route, and the Avid-specific metadata
	var genericMetadata *spireav.Meta
	var mxfMetadata *mxf2raw.Mxf2RawResult
	var err1, err2 error
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		genericMetadata, err1 = spireav.GetMetadata(ctx, filename)
	}()
	go func() {
		defer wg.Done()
		mxfMetadata, err2 = mxf2raw.GetMetadata(ctx, filename, nil)
	}()
	wg.Wait()

	// If there was an error getting the metadata, return it
	if err1 != nil {
		return nil, nil, fmt.Errorf("getting metadata: %s", err1)
	}

	// If there was an error getting the Avid-specific metadata, log it
	if err2 != nil {
		log.Printf("Error getting MXF metadata: %s", err2)
	}

	// Return the combined metadata
	avidMeta, err := ParseAvidMxfOpAtomMeta(genericMetadata, mxfMetadata)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing Avid metadata: %s", err)
	}
	return avidMeta, genericMetadata, nil
}

// ParseAvidMxfOpAtomMeta parses a single op-atom file's metadata and returns the Avid-specific identifiers
// and names for the tape and clip (reel and material package)
func ParseAvidMxfOpAtomMeta(metadata *spireav.Meta, mxfMetadata *mxf2raw.Mxf2RawResult) (*AvidMxfMeta, error) {
	// Get the stream that actually contains essence data, and thusly is represented by this Op-Atom file
	stream := findOpAtomStreamWithEssenceData(metadata)
	if stream == nil {
		return nil, errors.New("op-atom MXF file contains no stream with essence")
	}

	// Get the tape ID (reel name)
	reelName, ok := metadata.Format.Tags["comment_TapeID"]
	if !ok {
		reelName = stream.Tags["reel_name"]
	}

	// Create the starting point for the metadata
	avidMeta := AvidMxfMeta{
		FilePackageUmid:      stream.Tags["file_package_umid"],
		ReelName:             reelName,
		ReelUmid:             stream.Tags["reel_umid"],
		MaterialPackageName:  metadata.Format.Tags["material_package_name"],
		MaterialPackageUmid:  metadata.Format.Tags["material_package_umid"],
		MediaComposerVersion: metadata.Format.Tags["product_name"],
		CameraManufacturer:   metadata.Format.Tags["comment_manufacturer"],
		CameraModel:          metadata.Format.Tags["comment_modelName"],
		CameraSerialNum:      metadata.Format.Tags["comment_serialNo"],
		LensModelName:        metadata.Format.Tags["comment_LensModelName"],
		Timecode:             stream.Tags["timecode"],
		EssenceStream:        stream,
	}

	// Parse the creation date of the MXF file
	if creationDate, ok := metadata.Format.Tags["comment_CreationDate"]; ok {
		creation, err := time.Parse("2006-01-02T15:04:05-07:00", creationDate)
		if err != nil {
			fmt.Println("warning: could not parse MXF creation date: ", creationDate)
		} else {
			avidMeta.CreationDate = &creation
		}
	}

	// If there is mxf-specific metadata
	if mxfMetadata != nil {
		// Use the timecode from the mxf metadata if it exists
		if tc := mxfMetadata.Clip.StartTimecodes.PhysicalSource.Timecode; tc != "" {
			avidMeta.Timecode = tc
		}
	}

	// Return the metadata
	return &avidMeta, nil
}

// findOpAtomStreamWithEssenceData finds the stream in the metadata that contains actual essence data.
// In the Op-Atom format, only one stream contains essence, and the others are empty and contain metadata only
func findOpAtomStreamWithEssenceData(metadata *spireav.Meta) *spireav.StreamMeta {
	for _, stream := range metadata.Streams {
		if stream.CodecType == "audio" || stream.CodecType == "video" {
			return &stream
		}
	}
	return nil
}
