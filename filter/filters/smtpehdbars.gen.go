// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// SMTPEHDBarsBuilder corresponds to the "smtpehdbars" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#smtpehdbars
type SMTPEHDBarsBuilder interface {
	filter.Filter
	// Size sets the "s" option on the filter.
	Size(width, height int) SMTPEHDBarsBuilder
	// FrameRate sets the "r" option on the filter.
	FrameRate(frameRate expr.Expr) SMTPEHDBarsBuilder
	// Duration sets the "d" option on the filter.
	Duration(duration string) SMTPEHDBarsBuilder
}

// SMTPEHDBars creates a new SMTPEHDBarsBuilder to configure the "smtpehdbars" filter.
func SMTPEHDBars(opts ...filter.Option) SMTPEHDBarsBuilder {
	return &implSMTPEHDBarsBuilder{
		f: filter.New("smtpehdbars", 1, opts...),
	}
}

type implSMTPEHDBarsBuilder struct {
	f filter.Filter
}

func (b *implSMTPEHDBarsBuilder) String() string {
	return b.f.String()
}

func (b *implSMTPEHDBarsBuilder) Outputs() int {
	return b.f.Outputs()
}

func (b *implSMTPEHDBarsBuilder) With(key string, value expr.Expr) filter.Filter {
	return b.withOption(key, value)
}

func (b *implSMTPEHDBarsBuilder) withOption(key string, value expr.Expr) SMTPEHDBarsBuilder {
	bCopy := *b
	bCopy.f = b.f.With(key, value)
	return &bCopy
}

func (b *implSMTPEHDBarsBuilder) Size(width, height int) SMTPEHDBarsBuilder {
	return b.withOption("s", expr.Size(width, height))
}

func (b *implSMTPEHDBarsBuilder) FrameRate(frameRate expr.Expr) SMTPEHDBarsBuilder {
	return b.withOption("r", frameRate)
}

func (b *implSMTPEHDBarsBuilder) Duration(duration string) SMTPEHDBarsBuilder {
	return b.withOption("d", expr.String(duration))
}
