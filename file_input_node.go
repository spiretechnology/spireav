package spireav

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"errors"
	"unsafe"
)

type FileInputNode struct {
	filename string
	ctx      *C.struct_AVFormatContext
	isOpen   bool
}

func NewFileInputNode(filename string) FileInputNode {
	return FileInputNode{
		filename: filename,
		ctx:      nil,
		isOpen:   false,
	}
}

func (gfi *FileInputNode) Open() (*C.struct_AVFormatContext, error) {

	// If it's already open
	if gfi.isOpen {
		return gfi.ctx, nil
	}

	// Create the C string for the filename
	cFilename := C.CString(gfi.filename)
	defer C.free(unsafe.Pointer(cFilename))

	// Open the file context
	ret := int(C.avformat_open_input(
		(**C.struct_AVFormatContext)(unsafe.Pointer(&gfi.ctx)),
		cFilename,
		(*C.struct_AVInputFormat)(nil),
		(**C.struct_AVDictionary)(unsafe.Pointer(nil)),
	))
	if ret < 0 {
		return nil, errors.New("Failed to open input file")
	}

	// It's open now
	gfi.isOpen = true

	// Return the context
	return gfi.ctx, nil

}

func (gfi *FileInputNode) Close() {

	// If it's not open
	if !gfi.isOpen {
		return
	}

	// Close the input context
	C.avformat_close_input(&gfi.ctx)

	// It's not open anymore
	gfi.isOpen = false

}

func (gfi FileInputNode) GetOutputTypes() ([]StreamType, error) {

	// Make sure the file is open
	_, err := gfi.Open()
	if err != nil {
		return nil, err
	}

	// Load the stream information
	ret := int(C.avformat_find_stream_info(gfi.ctx, nil))
	if ret < 0 {
		return nil, errors.New("Cannot find stream information")
	}

	// Create the array of stream types
	var types []StreamType

	// Loop through the streams
	for i := uint(0); i < uint(gfi.ctx.nb_streams); i++ {

		// Dereference the stream at the offset
		stream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(gfi.ctx.streams)) + uintptr(i)*unsafe.Sizeof(*gfi.ctx.streams)))

		// Get the type for the stream
		var streamType StreamType
		if stream.codecpar.codec_type == C.AVMEDIA_TYPE_VIDEO {
			streamType = StreamTypeVideo
		} else if stream.codecpar.codec_type == C.AVMEDIA_TYPE_AUDIO {
			streamType = StreamTypeAudio
		} else {
			streamType = StreamTypeUnknown
		}

		// Add the stream type to the slice
		types = append(types, streamType)

	}

	// Return the slice of types
	return types, nil

}

func (gfi FileInputNode) FilterString(inputs []StreamType) string {
	return ""
}

func (gfi FileInputNode) Type() NodeType {
	return NodeTypeInput
}
