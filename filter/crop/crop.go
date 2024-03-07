package crop

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#toc-crop

func Crop(opts ...filter.Option) filter.Filter {
	return filter.New("crop", 1, opts...)
}

func WithWidth(width expr.Expr) filter.Option {
	return filter.WithOption("w", width)
}

func WithHeight(height expr.Expr) filter.Option {
	return filter.WithOption("h", height)
}

func WithX(x expr.Expr) filter.Option {
	return filter.WithOption("x", x)
}

func WithY(y expr.Expr) filter.Option {
	return filter.WithOption("y", y)
}

func WithKeepAspect(keepAspect bool) filter.Option {
	return filter.WithOption("keep_aspect", expr.Bool(keepAspect))
}

func WithExact(exact bool) filter.Option {
	return filter.WithOption("exact", expr.Bool(exact))
}
