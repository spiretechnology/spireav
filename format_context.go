package spireav

//#cgo LDFLAGS: -lavformat -lavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"errors"
	"unsafe"
)

// FormatContext is a context representing a multimedia file
type FormatContext struct {
	handle           *C.struct_AVFormatContext
	streamInfoLoaded bool
}

// OpenFileContext opens a format context from a file
func OpenFileContext(filename string) (*FormatContext, error) {

	// The format context handle
	var fmtCtx *C.struct_AVFormatContext

	// Get a C handle to the filename
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	// Open the input
	ret := int(C.avformat_open_input(
		(**C.struct_AVFormatContext)(unsafe.Pointer(&fmtCtx)),
		cFilename,
		(*C.struct_AVInputFormat)(nil),
		(**C.struct_AVDictionary)(unsafe.Pointer(nil)),
	))

	// If there was an error
	if ret != 0 {
		return nil, errors.New("Failed to open input file: " + filename)
	}

	// Create the new instance
	return &FormatContext{
		handle: fmtCtx,
	}, nil

}

// Close closes the format context
func (ctx *FormatContext) Close() {

	// Close the input
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx.handle)))

}

// loadStreamInfo analyzes the format context and pulls in information about the streams in the format
func (ctx *FormatContext) loadStreamInfo() error {
	if ctx.streamInfoLoaded {
		return nil
	}
	ctx.streamInfoLoaded = true
	ret := int(C.avformat_find_stream_info(ctx.handle, nil))
	if ret < 0 {
		return errors.New("cannot find stream information")
	}
	return nil
}

// Metadata gets the metadata from the format context
func (ctx *FormatContext) Metadata() (map[string]string, error) {

	// Load all of the stream information
	ret := int(C.avformat_find_stream_info(ctx.handle, nil))
	if ret < 0 {
		return nil, errors.New("Cannot find stream information")
	}

	// Create the key matching string (empty prefix)
	cKeyStr := C.CString("")
	defer C.free(unsafe.Pointer(cKeyStr))

	// Create a map to write into
	meta := make(map[string]string)

	// A pointer to the current entry
	var entry *C.struct_AVDictionaryEntry

	// Loop through the entries
	for {

		// Load the next entry
		entry = C.av_dict_get(ctx.handle.metadata, cKeyStr, (*C.struct_AVDictionaryEntry)(unsafe.Pointer(entry)), C.AV_DICT_IGNORE_SUFFIX)
		if entry == nil {
			break
		}

		// Add the key and value to the map
		meta[C.GoString(entry.key)] = C.GoString(entry.value)

	}

	// Return the metadata
	return meta, nil

}

// StreamsCount gets the number of streams in the format context
func (ctx *FormatContext) StreamsCount() uint {
	return uint(ctx.handle.nb_streams)
}

// Streams gets the slice of streams in the format context
func (ctx *FormatContext) Streams() ([]*Stream, error) {

	// Load the stream information
	if err := ctx.loadStreamInfo(); err != nil {
		return nil, err
	}

	// Count the number of streams
	count := ctx.StreamsCount()

	// Create the slice of streams to be returned
	streams := make([]*Stream, count)

	// Get an unsafe pointer to the streams array on the format context
	streamsPtr := uintptr(unsafe.Pointer(ctx.handle.streams))

	// Loop through the streams count
	for i := uint(0); i < count; i++ {

		// Dereference the stream at the offset
		streamHandle := *(**C.struct_AVStream)(unsafe.Pointer(streamsPtr + uintptr(i)*unsafe.Sizeof(*ctx.handle.streams)))

		// Create the stream wrapper instance
		streams[i] = &Stream{
			handle: streamHandle,
		}

	}

	// Return the streams
	return streams, nil

}
