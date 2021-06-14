package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
/*
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>

typedef int (*InCallbackRead)(void* ptr, uint8_t* buf, int buf_size);
typedef int (*InCallbackWrite)(void* ptr, uint8_t* buf, int buf_size);
typedef int64_t (*InCallbackSeek)(void* ptr, int64_t pos, int whence);

extern int goInputCallbackRead(void* ptr, uint8_t* buf, int buf_size);
extern int64_t goInputCallbackSeek(void* ptr, int64_t pos, int whence);
*/
import "C"
import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"unsafe"

	"github.com/spiretechnology/spireav"
)

var fileInputNodes = map[uint64]*FileInputNode{}
var fileInputNodesMutex sync.Mutex

func unusedFileInputNodeKey() uint64 {
	for {
		r := rand.Uint64()
		if _, exists := fileInputNodes[r]; !exists {
			return r
		}
	}
}

//export goInputCallbackRead
func goInputCallbackRead(ptr *C.void, buf *C.uint8_t, buf_size C.int) C.int {
	fileInputIndex := *(*uint64)(unsafe.Pointer(ptr))
	node, exists := fileInputNodes[fileInputIndex]
	if !exists {
		return -1
	}
	buffer := make([]byte, buf_size)
	n, err := node.reader.Read(buffer)
	if err != nil {
		fmt.Println("READ error: ", err.Error())
		return -1
	}
	for i := 0; i < n; i++ {
		*(*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(i))) = C.uint8_t(buffer[i])
	}
	return C.int(n)
}

//export goInputCallbackSeek
func goInputCallbackSeek(ptr *C.void, pos C.int64_t, whence C.int) C.int64_t {
	if whence&0x10000 > 0 {
		return -1
	}
	fileInputIndex := *(*uint64)(unsafe.Pointer(ptr))
	node, exists := fileInputNodes[fileInputIndex]
	if !exists {
		return 0
	}
	n, err := node.reader.Seek(int64(pos), int(whence))
	if err != nil {
		fmt.Println("SEEK error: ", err.Error())
		return 0
	}
	return C.int64_t(n)
}

// FileInputNode is a node that represents an input to a graph
type FileInputNode struct {
	InputNode
	reader     io.ReadSeeker
	readBuffer unsafe.Pointer
	key        uint64
	ioContext  *C.struct_AVIOContext
	ifmtCtx    *C.struct_AVFormatContext
}

func NewInputNode(reader io.ReadSeeker, format *Format) (*FileInputNode, error) {

	fileInputNodesMutex.Lock()

	// Create the input node
	node := FileInputNode{
		key:    unusedFileInputNodeKey(),
		reader: reader,
	}

	// Add the reader to the slice
	if fileInputNodes == nil {
		fileInputNodes = map[uint64]*FileInputNode{}
	}
	fileInputNodes[node.key] = &node

	fileInputNodesMutex.Unlock()

	// Alloc a buffer for the stream
	bufferSize := C.size_t(8192)
	node.readBuffer = C.av_malloc(bufferSize)

	// Get a AVContext stream
	node.ioContext = C.avio_alloc_context(
		(*C.uchar)(node.readBuffer),             // Buffer
		C.int(bufferSize),                       // Buffer size
		0,                                       // Buffer is only readable - set to 1 for read/write
		unsafe.Pointer(&node.key),               // User (your) specified data
		C.InCallbackRead(C.goInputCallbackRead), // Function - Reading Packets (see example)
		C.InCallbackWrite(nil),                  // Function - Write Packets
		C.InCallbackSeek(C.goInputCallbackSeek), // Function - Seek to position in stream (see example)
	)
	if node.ioContext == nil {
		return nil, errors.New("failed to allocate IO context")
	}

	// Allocate an AVFormatContext with a custom pb data source
	node.ifmtCtx = C.avformat_alloc_context()
	node.ifmtCtx.pb = node.ioContext
	node.ifmtCtx.flags |= C.AVFMT_FLAG_CUSTOM_IO // we set up our own IO

	// Open "file" (open our custom IO)
	// Empty string is where filename would go. Doesn't matter since we aren't reading a file
	// NULL params are format and demuxer settings, respectively
	if int(C.avformat_open_input(&node.ifmtCtx, nil, format.GetInputFormat(), nil)) < 0 {
		return nil, errors.New("failed to open input stream")
	}

	// Load the stream information
	ret := int(C.avformat_find_stream_info(node.ifmtCtx, nil))
	if ret < 0 {
		return nil, errors.New("error finding stream information")
	}

	// Create and return the input node
	return &node, nil

}

func (gfi *FileInputNode) GetContext() *C.struct_AVFormatContext {
	return gfi.ifmtCtx
}

func (gfi *FileInputNode) Close() {

	// Close the input context
	if gfi.ifmtCtx != nil {
		C.avformat_close_input(&gfi.ifmtCtx)
		gfi.ifmtCtx = nil
	}

	// Free the read buffer
	C.av_free(gfi.readBuffer)

	// Remove the file input node from the map
	fileInputNodesMutex.Lock()
	delete(fileInputNodes, gfi.key)
	fileInputNodesMutex.Unlock()

}

func (gfi *FileInputNode) GetOutputTypes() ([]spireav.StreamType, error) {

	// Create the array of stream types
	var types []spireav.StreamType

	// Loop through the streams
	for i := uint(0); i < uint(gfi.ifmtCtx.nb_streams); i++ {

		// Dereference the stream at the offset
		stream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(gfi.ifmtCtx.streams)) + uintptr(i)*unsafe.Sizeof(*gfi.ifmtCtx.streams)))

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

func (gfi *FileInputNode) GetStream(streamIndex int) (*InputStream, error) {

	// If the index is out of bounds
	if streamIndex < 0 || streamIndex >= int(gfi.ifmtCtx.nb_streams) {
		return nil, fmt.Errorf("stream index #%d out of bounds", streamIndex)
	}

	// Dereference the stream at the index
	avStream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(gfi.ifmtCtx.streams)) + uintptr(streamIndex)*unsafe.Sizeof(*gfi.ifmtCtx.streams)))

	// Create the new input stream
	return NewInputStream(gfi.ifmtCtx, avStream)

}
