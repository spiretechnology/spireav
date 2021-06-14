package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/spiretechnology/spireav"
)

type FileOutputNode struct {
	OutputNode
	filename string
	ctx      *C.struct_AVFormatContext
	isOpen   bool
	streams  []*OutputStream
}

func NewFileOutputNode(filename string) *FileOutputNode {
	return &FileOutputNode{
		filename: filename,
		ctx:      nil,
	}
}

func (gfo *FileOutputNode) Open() (*C.struct_AVFormatContext, error) {

	// If it's already open
	if gfo.isOpen {
		return gfo.ctx, nil
	}

	// Create the C string for the filename
	cFilename := C.CString(gfo.filename)
	defer C.free(unsafe.Pointer(cFilename))

	// Open the file output context
	C.avformat_alloc_output_context2(
		(**C.struct_AVFormatContext)(unsafe.Pointer(&gfo.ctx)),
		nil,
		nil,
		cFilename,
	)
	if gfo.ctx == nil {
		return nil, errors.New("failed to create output file context")
	}

	// It's open now
	gfo.isOpen = true

	// Return the context
	return gfo.ctx, nil

}

func (gfo *FileOutputNode) Close() {

	// If it's not open
	if !gfo.isOpen {
		return
	}

	// Close the input context
	// C.avformat_close_output(&gfo.ctx)

	// It's not open anymore
	gfo.isOpen = false

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

func (gfo *FileOutputNode) AddStream(encoder *Encoder) (*OutputStream, error) {

	// Open the file output
	ofmtCtx, err := gfo.Open()
	if err != nil {
		return nil, err
	}

	// Create the output stream
	avStream := C.avformat_new_stream(ofmtCtx, nil)
	if avStream == nil {
		return nil, errors.New("failed allocating output stream")
	}

	// Create the wrapped stream instance
	stream := OutputStream{
		avStream: avStream,
		encoder:  encoder,
	}

	// Create the encoder context for the stream
	encCtx, err := stream.OpenEncoderContext()
	if err != nil {
		return nil, err
	}

	// If the global header flag exists, add it to the output
	if ofmtCtx.oformat.flags&C.AVFMT_GLOBALHEADER > 0 {
		encCtx.flags |= C.AV_CODEC_FLAG_GLOBAL_HEADER
	}

	// Add the stream to the slice
	gfo.streams = append(gfo.streams, &stream)

	// Return without error
	return &stream, nil

}

func (gfo *FileOutputNode) GetStream(streamIndex int) (*OutputStream, error) {
	if streamIndex < 0 || streamIndex >= len(gfo.streams) {
		return nil, errors.New("stream index out of bounds")
	}
	return gfo.streams[streamIndex], nil

}

func (gfo *FileOutputNode) WriteHeader() error {

	// Open the file output
	ofmtCtx, err := gfo.Open()
	if err != nil {
		return err
	}

	// Create the C string for the filename
	cFilename := C.CString(gfo.filename)
	defer C.free(unsafe.Pointer(cFilename))

	// Dump the output format to the console
	C.av_dump_format(ofmtCtx, 0, cFilename, 1)

	// As long as there's actually a file
	if ofmtCtx.oformat.flags&C.AVFMT_NOFILE == 0 {
		ret := int(C.avio_open(
			(**C.struct_AVIOContext)(unsafe.Pointer(&ofmtCtx.pb)),
			cFilename,
			C.AVIO_FLAG_WRITE,
		))
		if ret < 0 {
			return errors.New("failed to open file output")
		}
	}

	// Write the file header
	ret := int(C.avformat_write_header(ofmtCtx, nil))
	if ret < 0 {
		return errors.New("error occurred while opening output file")
	}

	// Return without error
	return nil

}
