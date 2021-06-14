package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
import "C"
import (
	"errors"
)

type OutputStream struct {
	avStream *C.struct_AVStream
	encCtx   *C.struct_AVCodecContext
	encoder  *Encoder
}

func NewOutputStream(
	ofmtCtx *C.struct_AVFormatContext,
	desc *StreamDescription,
) (*OutputStream, error) {

	// Get the encoder for the stream
	encoder, err := desc.Codec.GetEncoder()
	if err != nil {
		return nil, err
	}

	// Create the output stream
	avStream := C.avformat_new_stream(ofmtCtx, encoder.avCodec)
	if avStream == nil {
		return nil, errors.New("failed allocating output stream")
	}

	// Allocate a context for encoding
	encCtx := C.avcodec_alloc_context3(encoder.avCodec)
	if encCtx == nil {
		return nil, errors.New("failed to allocate encoder context for stream")
	}

	// If the global header flag exists, add it to the output
	if ofmtCtx.oformat.flags&C.AVFMT_GLOBALHEADER > 0 {
		encCtx.flags |= C.AV_CODEC_FLAG_GLOBAL_HEADER
	}

	// fmt.Println("HEY: ", avStream.time_base)
	// encCtx.time_base = avStream.time_base

	// Create the output stream
	os := OutputStream{
		avStream: avStream,
		encCtx:   encCtx,
		encoder:  encoder,
	}

	// Set the time base for the stream
	// TODO: unsure if this is needed
	// avStream.time_base = encCtx.time_base

	// Run the handler on the stream
	desc.Handler(&os)

	// Open the encoder
	ret := int(C.avcodec_open2(encCtx, encoder.avCodec, (**C.struct_AVDictionary)(nil)))
	if ret < 0 {
		return nil, errors.New("failed to open video encoder for stream")
	}

	// Copy encoder parameters to the output stream
	ret = int(C.avcodec_parameters_from_context(avStream.codecpar, encCtx))
	if ret < 0 {
		return nil, errors.New("failed to copy encoder parameters to output stream")
	}

	// Return the output stream
	return &os, nil

}

func (os *OutputStream) Close() error {
	C.avcodec_free_context(&os.encCtx)
	return nil
}

func (os *OutputStream) GetEncoder() *Encoder {
	return os.encoder
}

func (os *OutputStream) SetWidth(width int) {
	os.encCtx.width = C.int(width)
}

func (os *OutputStream) SetHeight(height int) {
	os.encCtx.height = C.int(height)
}

func (os *OutputStream) SetAspectRatio(num, den int) {
	os.encCtx.sample_aspect_ratio = C.struct_AVRational{
		num: C.int(num),
		den: C.int(den),
	}
	os.avStream.sample_aspect_ratio = os.encCtx.sample_aspect_ratio
}

func (os *OutputStream) SetPixFmt(pixFmt int32) {
	os.encCtx.pix_fmt = pixFmt
}

func (os *OutputStream) SetFrameRate(num, den int) {
	os.encCtx.framerate = C.struct_AVRational{
		num: C.int(num),
		den: C.int(den),
	}
}

func (os *OutputStream) SetTimeBase(num, den int) {
	os.encCtx.time_base = C.struct_AVRational{
		num: C.int(num),
		den: C.int(den),
	}
	os.avStream.time_base = os.encCtx.time_base
}
