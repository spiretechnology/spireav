package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

// Format is a container format for output files
type Format struct {
	name string
}

// FormatWithName returns an output format instance with the given name
func FormatWithName(name string) *Format {
	return &Format{name}
}

func (f *Format) GetInputFormat() *C.struct_AVInputFormat {

	// Create the C string for the format name
	cName := C.CString(f.name)
	defer C.free(unsafe.Pointer(cName))

	// Find the format with the name
	return C.av_find_input_format(cName)

}

func (f *Format) GetOutputFormat() *C.struct_AVOutputFormat {

	// Create the C string for the format name
	cName := C.CString(f.name)
	defer C.free(unsafe.Pointer(cName))

	// Find the format with the name
	return C.av_guess_format(cName, nil, nil)

}
