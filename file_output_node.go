package spireav

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

type FileOutputNode struct {
	filename string
	ctx      *C.struct_AVFormatContext
	isOpen   bool
}

func NewFileOutputNode(filename string) FileOutputNode {
	return FileOutputNode{
		filename: filename,
		ctx:      nil,
	}
}

func (gfi *FileOutputNode) Open() (*C.struct_AVFormatContext, error) {

	// If it's already open
	if gfi.isOpen {
		return gfi.ctx, nil
	}

	// Create the C string for the filename
	cFilename := C.CString(gfi.filename)
	defer C.free(unsafe.Pointer(cFilename))

	// // Open the file context
	// ret := int(C.avformat_open_input(
	// 	(**C.struct_AVFormatContext)(unsafe.Pointer(&gfi.ctx)),
	// 	cFilename,
	// 	(*C.struct_AVInputFormat)(nil),
	// 	(**C.struct_AVDictionary)(unsafe.Pointer(nil)),
	// ))
	// if ret < 0 {
	// 	return nil, errors.New("Failed to open input file")
	// }

	// It's open now
	gfi.isOpen = true

	// Return the context
	return gfi.ctx, nil

}

func (gfi *FileOutputNode) Close() {

	// If it's not open
	if !gfi.isOpen {
		return
	}

	// Close the input context
	// C.avformat_close_input(&gfi.ctx)

	// It's not open anymore
	gfi.isOpen = false

}

func (gfi FileOutputNode) GetOutputTypes() ([]StreamType, error) {
	return []StreamType{}, nil
}

func (gfi FileOutputNode) FilterString(inputs []StreamType) string {
	return ""
}

func (gfi FileOutputNode) Type() NodeType {
	return NodeTypeOutput
}
