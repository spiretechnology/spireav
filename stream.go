package spireav

//#cgo LDFLAGS: -lavformat -lavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"unsafe"
)

type Stream struct {
	handle *C.struct_AVStream
}

// Metadata gets the metadata specific to this stream
func (s *Stream) Metadata() (map[string]string, error) {

	// Create the key matching string (empty prefix)
	cKeyStr := C.CString("")
	defer C.free(unsafe.Pointer(cKeyStr))

	// Create a map to write into
	meta := make(map[string]string)

	// A pointer to the current entry
	var entry *C.struct_AVDictionaryEntry

	// Loop through the entries
	for {

		// Load the next entry
		entry = C.av_dict_get(s.handle.metadata, cKeyStr, (*C.struct_AVDictionaryEntry)(unsafe.Pointer(entry)), C.AV_DICT_IGNORE_SUFFIX)
		if entry == nil {
			break
		}

		// Add the key and value to the map
		meta[C.GoString(entry.key)] = C.GoString(entry.value)

	}

	// Return the metadata
	return meta, nil

}

// CodecID gets the identifier of the codec
func (s *Stream) CodecID() uint {
	return 0
}

// Type gets the codec type of this stream (audio, video, etc.)
func (s *Stream) Type() StreamType {
	rawtype := s.handle.codecpar.codec_type
	if rawtype == C.AVMEDIA_TYPE_AUDIO {
		return StreamTypeAudio
	} else if rawtype == C.AVMEDIA_TYPE_VIDEO {
		return StreamTypeVideo
	}
	return StreamTypeUnknown
}
