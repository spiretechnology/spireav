package vidsrc

import (
	"fmt"

	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// https://ffmpeg.org/ffmpeg-filters.html#nullsrc

func AllRGB(opts ...filter.Option) filter.Filter {
	return filter.New("allrgb", 1, opts...)
}

func AllYUV(opts ...filter.Option) filter.Filter {
	return filter.New("allyuv", 1, opts...)
}

func Color(opts ...filter.Option) filter.Filter {
	return filter.New("color", 1, opts...)
}

func ColorChart(opts ...filter.Option) filter.Filter {
	return filter.New("colorchart", 1, opts...)
}

func ColorSpectrum(opts ...filter.Option) filter.Filter {
	return filter.New("colorspectrum", 1, opts...)
}

func HaldCLUT(opts ...filter.Option) filter.Filter {
	return filter.New("haldclutsrc", 1, opts...)
}

func Null(opts ...filter.Option) filter.Filter {
	return filter.New("nullsrc", 1, opts...)
}

func PAL75Bars(opts ...filter.Option) filter.Filter {
	return filter.New("pal75bars", 1, opts...)
}

func PAL100Bars(opts ...filter.Option) filter.Filter {
	return filter.New("pal100bars", 1, opts...)
}

func RGBTest(opts ...filter.Option) filter.Filter {
	return filter.New("rgbtestsrc", 1, opts...)
}

func SMPTEBars(opts ...filter.Option) filter.Filter {
	return filter.New("smptebars", 1, opts...)
}

func SMPTEHDBars(opts ...filter.Option) filter.Filter {
	return filter.New("smptehdbars", 1, opts...)
}

func TestSource(opts ...filter.Option) filter.Filter {
	return filter.New("testsrc", 1, opts...)
}

func TestSource2(opts ...filter.Option) filter.Filter {
	return filter.New("testsrc2", 1, opts...)
}

func YUVTestSource(opts ...filter.Option) filter.Filter {
	return filter.New("testsrc2", 1, opts...)
}

func WithSize(width, height int) filter.Option {
	return filter.WithOption("s", expr.String(fmt.Sprintf("%dx%d", width, height)))
}

func WithFrameRate(frameRate expr.Expr) filter.Option {
	return filter.WithOption("r", frameRate)
}

func WithDuration(duration expr.Expr) filter.Option {
	return filter.WithOption("d", duration)
}

func WithColor(color expr.Expr) filter.Option {
	return filter.WithOption("c", color)
}
