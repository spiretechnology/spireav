package graph

type StreamDescription struct {
	Codec   *Codec
	Handler func(stream *OutputStream)
}

func NewStreamDescription(
	codecName string,
	handler func(stream *OutputStream),
) *StreamDescription {
	return &StreamDescription{
		Codec:   CodecWithName(codecName),
		Handler: handler,
	}
}
