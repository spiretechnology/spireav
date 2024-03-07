package crop

import (
	"github.com/spiretechnology/spireav/graph/transform"
	"github.com/spiretechnology/spireav/graph/transform/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#toc-crop

func Crop(opts ...transform.Option) transform.Transform {
	return transform.New("crop", 1, opts...)
}

func WithWidth(width expr.Expr) transform.Option {
	return transform.WithOption("w", width)
}

func WithHeight(height expr.Expr) transform.Option {
	return transform.WithOption("h", height)
}

func WithX(x expr.Expr) transform.Option {
	return transform.WithOption("x", x)
}

func WithY(y expr.Expr) transform.Option {
	return transform.WithOption("y", y)
}

func WithKeepAspect(keepAspect bool) transform.Option {
	return transform.WithOption("keep_aspect", expr.Bool(keepAspect))
}

func WithExact(exact bool) transform.Option {
	return transform.WithOption("exact", expr.Bool(exact))
}
