package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavcodec/avcodec.h>
import "C"

type Decoder struct {
	codec   *Codec
	avCodec *C.struct_AVCodec
}

// GetCodec gets the codes to which this decoder belongs
func (dec *Decoder) GetCodec() *Codec {
	return dec.codec
}
