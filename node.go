package spireav

type NodeType uint

const (
	NodeTypeInput     NodeType = 0
	NodeTypeOutput    NodeType = 1
	NodeTypeTransform NodeType = 2
)

type Node interface {
	FilterString(inputs []StreamType) string
	GetOutputTypes() ([]StreamType, error)
	Type() NodeType
}
