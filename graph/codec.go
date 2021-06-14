package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavcodec/avcodec.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// Codec is a way of encoding or decoding audio or video stream data. A codec doesn't do anything on its own,
// all the action takes place in either the Decoder or Encoder instance associated with the codec
type Codec struct {
	id   uint32
	name *string
}

// CodecWithID returns a codec instance with the given codec id
func CodecWithID(id uint32) *Codec {
	return &Codec{id, nil}
}

// CodecWithName returns a codec instance with the given codec name
func CodecWithName(name string) *Codec {
	return &Codec{0, &name}
}

// GetDecoder gets a decoder instance for this codec
func (codec *Codec) GetDecoder() (*Decoder, error) {
	var decoder *C.struct_AVCodec
	if codec.name != nil {
		name := C.CString(*codec.name)
		defer C.free(unsafe.Pointer(name))
		decoder = C.avcodec_find_decoder_by_name(name)
	} else {
		decoder = C.avcodec_find_decoder(codec.id)
	}
	if decoder == nil {
		return nil, fmt.Errorf("failed to find decoder with id #%d", codec.id)
	}
	return &Decoder{
		codec:   codec,
		avCodec: decoder,
	}, nil
}

// GetEncoder gets a decoder instance for this codec
func (codec *Codec) GetEncoder() (*Encoder, error) {
	var encoder *C.struct_AVCodec
	if codec.name != nil {
		name := C.CString(*codec.name)
		defer C.free(unsafe.Pointer(name))
		encoder = C.avcodec_find_encoder_by_name(name)
	} else {
		encoder = C.avcodec_find_encoder(codec.id)
	}
	if encoder == nil {
		return nil, fmt.Errorf("failed to find encoder with id #%d", codec.id)
	}
	return &Encoder{
		codec:   codec,
		avCodec: encoder,
	}, nil
}
