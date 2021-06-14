package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/spiretechnology/spireav"
)

type Node interface {
	GetOutputTypes() ([]spireav.StreamType, error)
}

type InputNode interface {
	Node
	GetContext() *C.struct_AVFormatContext
	GetStream(streamIndex int) (*InputStream, error)
}

type OutputNode interface {
	Node
	GetContext() *C.struct_AVFormatContext
	GetStream(streamIndex int) (*OutputStream, error)
}

type TransformNode interface {
	Node
	FilterString(inputs []spireav.StreamType) string
}
