package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
import "C"

type Encoder struct {
	codec   *Codec
	avCodec *C.struct_AVCodec
}

// GetCodec gets the codes to which this encoder belongs
func (enc *Encoder) GetCodec() *Codec {
	return enc.codec
}
