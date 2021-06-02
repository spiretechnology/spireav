package graph

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <libavformat/avformat.h>
import "C"
import "github.com/spiretechnology/spireav"

type NodeType uint

const (
	NodeTypeInput     NodeType = 0
	NodeTypeOutput    NodeType = 1
	NodeTypeTransform NodeType = 2
)

type Node interface {
	FilterString(inputs []spireav.StreamType) string
	GetOutputTypes() ([]spireav.StreamType, error)
	Type() NodeType
}

type InputNode interface {
	Open() (*C.struct_AVFormatContext, error)
	GetStream(streamIndex int) (*InputStream, error)
}

type OutputNode interface {
	Open() (*C.struct_AVFormatContext, error)
	AddStream(encoder *Encoder) (*OutputStream, error)
	GetStream(streamIndex int) (*OutputStream, error)
}
