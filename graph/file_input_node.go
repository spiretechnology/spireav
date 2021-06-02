package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/spiretechnology/spireav"
)

// FileInputNode is a node that represents an input to a graph
type FileInputNode struct {
	filename string
	ctx      *C.struct_AVFormatContext
	isOpen   bool
}

// NewFileInputNode creates a FileInputNode from the provided filename
func NewFileInputNode(filename string) *FileInputNode {
	return &FileInputNode{
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
		return nil, errors.New("failed to open input file")
	}

	// Load the stream information
	ret = int(C.avformat_find_stream_info(gfi.ctx, nil))
	if ret < 0 {
		return nil, errors.New("error finding stream information")
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

func (gfi *FileInputNode) GetOutputTypes() ([]spireav.StreamType, error) {

	// Make sure the file is open
	_, err := gfi.Open()
	if err != nil {
		return nil, err
	}

	// Create the array of stream types
	var types []spireav.StreamType

	// Loop through the streams
	for i := uint(0); i < uint(gfi.ctx.nb_streams); i++ {

		// Dereference the stream at the offset
		stream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(gfi.ctx.streams)) + uintptr(i)*unsafe.Sizeof(*gfi.ctx.streams)))

		// Get the type for the stream
		var streamType spireav.StreamType
		if stream.codecpar.codec_type == C.AVMEDIA_TYPE_VIDEO {
			streamType = spireav.StreamTypeVideo
		} else if stream.codecpar.codec_type == C.AVMEDIA_TYPE_AUDIO {
			streamType = spireav.StreamTypeAudio
		} else {
			streamType = spireav.StreamTypeUnknown
		}

		// Add the stream type to the slice
		types = append(types, streamType)

	}

	// Return the slice of types
	return types, nil

}

func (gfi *FileInputNode) FilterString(inputs []spireav.StreamType) string {
	return ""
}

func (gfi *FileInputNode) Type() NodeType {
	return NodeTypeInput
}

func (gfi *FileInputNode) GetStream(streamIndex int) (*InputStream, error) {

	// Open the file context
	ifmtCtx, err := gfi.Open()
	if err != nil {
		return nil, err
	}

	// If the index is out of bounds
	if streamIndex < 0 || streamIndex >= int(ifmtCtx.nb_streams) {
		return nil, fmt.Errorf("stream index #%d out of bounds", streamIndex)
	}

	// Dereference the stream at the index
	stream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ifmtCtx.streams)) + uintptr(streamIndex)*unsafe.Sizeof(*ifmtCtx.streams)))

	// Return the stream
	return &InputStream{
		ifmtCtx:  ifmtCtx,
		avStream: stream,
	}, nil

}
