package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
import "C"
import (
	"errors"
)

type InputStream struct {
	ifmtCtx  *C.struct_AVFormatContext
	avStream *C.struct_AVStream
	decCtx   *C.struct_AVCodecContext
}

func NewInputStream(
	ifmtCtx *C.struct_AVFormatContext,
	avStream *C.struct_AVStream,
) (*InputStream, error) {

	// Get the decoder for this input stream
	dec, err := CodecWithID(avStream.codecpar.codec_id).GetDecoder()
	if err != nil {
		return nil, err
	}

	// Allocate a context for decoding
	decCtx := C.avcodec_alloc_context3(dec.avCodec)
	if decCtx == nil {
		return nil, errors.New("failed to allocate decoder context for stream")
	}

	// Setup the codec context parameters
	ret := int(C.avcodec_parameters_to_context(decCtx, avStream.codecpar))
	if ret < 0 {
		return nil, errors.New("failed to copy decoder paramaters to decoder context")
	}

	// If it's video, predict the frame rate
	if decCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
		decCtx.framerate = C.av_guess_frame_rate(ifmtCtx, avStream, nil)
		decCtx.time_base = avStream.time_base
		// decCtx.time_base = C.struct_AVRational{
		// 	num: 1,
		// 	den: 90000,
		// }
	}

	// Open the encoder
	if int(C.avcodec_open2(decCtx, decCtx.codec, nil)) < 0 {
		return nil, errors.New("failed to open video decoder for stream")
	}

	// Return the input stream
	return &InputStream{
		ifmtCtx:  ifmtCtx,
		avStream: avStream,
		decCtx:   decCtx,
	}, nil

}

func (is *InputStream) Close() error {
	C.avcodec_free_context(&is.decCtx)
	return nil
}

func (is *InputStream) GetWidth() int {
	return int(is.decCtx.width)
}

func (is *InputStream) GetHeight() int {
	return int(is.decCtx.width)
}

func (is *InputStream) GetAspectRatio() (int, int) {
	ratio := is.decCtx.sample_aspect_ratio
	return int(ratio.num), int(ratio.den)
}

func (is *InputStream) GetPixFmt() int32 {
	return is.decCtx.pix_fmt
}

func (is *InputStream) GetFrameRate() (int, int) {
	rate := is.decCtx.framerate
	return int(rate.num), int(rate.den)
}

func (is *InputStream) GetTimeBase() (int, int) {
	timebase := is.decCtx.time_base
	return int(timebase.num), int(timebase.den)
}
