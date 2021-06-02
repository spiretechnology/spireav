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

// GetCodec gets the codec used by this input stream
func (is *InputStream) GetCodec() *Codec {
	return CodecWithID(is.avStream.codecpar.codec_id)
}

// OpenDecoderContext opens a new decoder context that can be used to decode data from this stream
func (is *InputStream) OpenDecoderContext() (*C.struct_AVCodecContext, error) {
	if is.decCtx == nil {

		// Get the decoder for this input stream
		dec, err := is.GetCodec().GetDecoder()
		if err != nil {
			return nil, err
		}

		// Allocate a context for decoding
		is.decCtx = C.avcodec_alloc_context3(dec.avCodec)
		if is.decCtx == nil {
			return nil, errors.New("failed to allocate decoder context for stream")
		}

		// Setup the codec context parameters
		ret := int(C.avcodec_parameters_to_context(is.decCtx, is.avStream.codecpar))
		if ret < 0 {
			return nil, errors.New("failed to copy decoder paramaters to decoder context")
		}

		// If it's video, predict the frame rate
		if is.decCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
			is.decCtx.framerate = C.av_guess_frame_rate(is.ifmtCtx, is.avStream, nil)
			is.decCtx.time_base = C.struct_AVRational{
				num: 1,
				den: 90000,
			}
		}

	}
	return is.decCtx, nil
}
