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

// GetCodec gets the codec used by this input stream
func (os *OutputStream) GetCodec() *Codec {
	return os.encoder.codec
}

func (os *OutputStream) OpenEncoderContext() (*C.struct_AVCodecContext, error) {
	if os.encCtx == nil {

		// Allocate a context for encoding
		os.encCtx = C.avcodec_alloc_context3(os.encoder.avCodec)
		if os.encCtx == nil {
			return nil, errors.New("failed to allocate encoder context for stream")
		}

	}
	return os.encCtx, nil
}
