// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SMTPEBarsBuilder corresponds to the "smtpebars" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#smtpebars
type SMTPEBarsBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) SMTPEBarsBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) SMTPEBarsBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) SMTPEBarsBuilder
}

// SMTPEBars creates a new SMTPEBarsBuilder to configure the "smtpebars" filter.
func SMTPEBars(opts ...filter.Option) SMTPEBarsBuilder {
	return &implSMTPEBarsBuilder{
		f: filter.New("smtpebars", 1, opts...),
	}
}

type implSMTPEBarsBuilder struct {
	f filter.Filter
}

func (b *implSMTPEBarsBuilder) String() string {
	return b.f.String()
}

func (b *implSMTPEBarsBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implSMTPEBarsBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implSMTPEBarsBuilder) withOption(key string, value expr.Expr) SMTPEBarsBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implSMTPEBarsBuilder) Size(width, height int) SMTPEBarsBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implSMTPEBarsBuilder) FrameRate(frameRate expr.Expr) SMTPEBarsBuilder {
	return b.withOption("r", frameRate)
}

func (b *implSMTPEBarsBuilder) Duration(duration string) SMTPEBarsBuilder {
	return b.withOption("d", expr.String(duration))
}
