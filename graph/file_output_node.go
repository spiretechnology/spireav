package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
/*
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>

typedef int (*OutCallbackRead)(void* ptr, uint8_t* buf, int buf_size);
typedef int (*OutCallbackWrite)(void* ptr, uint8_t* buf, int buf_size);
typedef int64_t (*OutCallbackSeek)(void* ptr, int64_t pos, int whence);

extern int goOutputCallbackWrite(void* ptr, uint8_t* buf, int buf_size);
extern int64_t goOutputCallbackSeek(void* ptr, int64_t pos, int whence);
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

var fileOutputNodes = map[uint64]*FileOutputNode{}
var fileOutputNodesMutex sync.Mutex

func unusedFileOutputNodeKey() uint64 {
	for {
		r := rand.Uint64()
		if _, exists := fileOutputNodes[r]; !exists {
			return r
		}
	}
}

//export goOutputCallbackSeek
func goOutputCallbackSeek(ptr *C.void, pos C.int64_t, whence C.int) C.int64_t {
	if whence&0x10000 > 0 {
		return -1
	}
	fileOutputIndex := *(*uint64)(unsafe.Pointer(ptr))
	node, exists := fileOutputNodes[fileOutputIndex]
	if !exists {
		return -1
	}
	n, err := node.writer.Seek(int64(pos), int(whence))
	if err != nil {
		fmt.Println("SEEK error: ", err.Error())
		return -1
	}
	return C.int64_t(n)
}

//export goOutputCallbackWrite
func goOutputCallbackWrite(ptr *C.void, buf *C.uint8_t, buf_size C.int) C.int {
	fileOutputIndex := *(*uint64)(unsafe.Pointer(ptr))
	// fmt.Println("WRITE called: ", fileOutputIndex, buf, buf_size)
	node, exists := fileOutputNodes[fileOutputIndex]
	if !exists {
		return -1
	}
	buffer := make([]byte, buf_size)
	for i := 0; i < int(buf_size); i++ {
		buffer[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(i)))
	}
	n, err := node.writer.Write(buffer)
	if err != nil {
		fmt.Println("WRITE error: ", err.Error())
		return -1
	}
	return C.int(n)
}

type FileOutputNode struct {
	OutputNode
	writer    io.WriteSeeker
	key       uint64
	ioContext *C.struct_AVIOContext
	ofmtCtx   *C.struct_AVFormatContext
	streams   []*OutputStream
}

func NewOutputNode(
	writer io.WriteSeeker,
	format *Format,
	streams []*StreamDescription,
) (*FileOutputNode, error) {

	// Define your buffer size
	bufferSize := 8192

	// Alloc a buffer for the stream
	buffer := C.av_malloc(C.size_t(bufferSize))

	fileOutputNodesMutex.Lock()

	// Create the file output node
	node := FileOutputNode{
		key:    unusedFileOutputNodeKey(),
		writer: writer,
	}

	// Add the writer to the slice
	fileOutputNodes[node.key] = &node

	fileOutputNodesMutex.Unlock()

	// Get a AVContext stream
	node.ioContext = C.avio_alloc_context(
		(*C.uchar)(buffer),        // Buffer
		C.int(bufferSize),         // Buffer size
		1,                         // Buffer is only readable - set to 1 for read/write
		unsafe.Pointer(&node.key), // User (your) specified data
		C.OutCallbackRead(nil),    // Function - Reading Packets (see example)
		C.OutCallbackWrite(C.goOutputCallbackWrite), // Function - Write Packets
		C.OutCallbackSeek(C.goOutputCallbackSeek),   // Function - Seek to position in stream (see example)
	)
	if node.ioContext == nil {
		return nil, errors.New("failed to allocate IO context")
	}

	// Open the file output context
	C.avformat_alloc_output_context2(
		(**C.struct_AVFormatContext)(unsafe.Pointer(&node.ofmtCtx)),
		format.GetOutputFormat(),
		nil,
		nil,
	)
	if node.ofmtCtx == nil {
		return nil, errors.New("failed to create output file context")
	}

	// Apply the custom data destination to the format context
	node.ofmtCtx.pb = node.ioContext
	node.ofmtCtx.flags |= C.AVFMT_FLAG_CUSTOM_IO

	// Add output streams here according to the supplied arguments
	for _, stream := range streams {
		node.addStream(stream)
	}

	// Write the header to the output
	if err := node.writeHeader(); err != nil {
		return nil, err
	}

	// Return the open node
	return &node, nil

}

func (gfo *FileOutputNode) GetContext() *C.struct_AVFormatContext {
	return gfo.ofmtCtx
}

func (gfo *FileOutputNode) Close() {

	// Close the output context
	if gfo.ofmtCtx != nil {
		// C.avformat_close_output(&gfo.ofmtCtx)
		C.avformat_free_context(gfo.ofmtCtx)
		gfo.ofmtCtx = nil
	}

}

func (gfo *FileOutputNode) GetOutputTypes() ([]spireav.StreamType, error) {
	types := make([]spireav.StreamType, len(gfo.streams))
	for i := range gfo.streams {
		if gfo.streams[i].encCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
			types[i] = spireav.StreamTypeVideo
		} else if gfo.streams[i].encCtx.codec_type == C.AVMEDIA_TYPE_AUDIO {
			types[i] = spireav.StreamTypeAudio
		} else {
			types[i] = spireav.StreamTypeUnknown
		}
	}
	return types, nil
}

func (gfo *FileOutputNode) addStream(desc *StreamDescription) (*OutputStream, error) {

	// Create the stream from the description
	stream, err := NewOutputStream(
		gfo.ofmtCtx,
		desc,
	)
	if err != nil {
		return nil, err
	}

	// Add the stream to the slice
	gfo.streams = append(gfo.streams, stream)

	// Return without error
	return stream, nil

}

func (gfo *FileOutputNode) GetStream(streamIndex int) (*OutputStream, error) {
	if streamIndex < 0 || streamIndex >= len(gfo.streams) {
		return nil, errors.New("stream index out of bounds")
	}
	return gfo.streams[streamIndex], nil
}

func (gfo *FileOutputNode) writeHeader() error {

	// // Create the C string for the filename
	// cFilename := C.CString("")
	// defer C.free(unsafe.Pointer(cFilename))

	// // Dump the output format to the console
	// C.av_dump_format(gfo.ofmtCtx, 0, cFilename, 1)

	// // As long as there's actually a file
	// if gfo.ofmtCtx.oformat.flags&C.AVFMT_NOFILE == 0 {
	// 	ret := int(C.avio_open(
	// 		(**C.struct_AVIOContext)(unsafe.Pointer(&gfo.ofmtCtx.pb)),
	// 		cFilename,
	// 		C.AVIO_FLAG_WRITE,
	// 	))
	// 	if ret < 0 {
	// 		return errors.New("failed to open file output")
	// 	}
	// }

	// Write the file header
	ret := int(C.avformat_write_header(gfo.ofmtCtx, nil))
	if ret < 0 {
		return errors.New("error occurred while opening output file")
	}

	// Return without error
	return nil

}
