// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// DialoguenhanceBuilder Audio Dialogue Enhancement.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#dialoguenhance
type DialoguenhanceBuilder interface {
	filter.Filter
	// Original set original center factor (from 0 to 1) (default 1).
	Original(original float64) DialoguenhanceBuilder
	// OriginalExpr set original center factor (from 0 to 1) (default 1).
	OriginalExpr(original expr.Expr) DialoguenhanceBuilder
	// Enhance set dialogue enhance factor (from 0 to 3) (default 1).
	Enhance(enhance float64) DialoguenhanceBuilder
	// EnhanceExpr set dialogue enhance factor (from 0 to 3) (default 1).
	EnhanceExpr(enhance expr.Expr) DialoguenhanceBuilder
	// Voice set voice detection factor (from 2 to 32) (default 2).
	Voice(voice float64) DialoguenhanceBuilder
	// VoiceExpr set voice detection factor (from 2 to 32) (default 2).
	VoiceExpr(voice expr.Expr) DialoguenhanceBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) DialoguenhanceBuilder
}

// Dialoguenhance creates a new DialoguenhanceBuilder to configure the "dialoguenhance" filter.
func Dialoguenhance(opts ...filter.Option) DialoguenhanceBuilder {
	return &implDialoguenhanceBuilder{
		f: filter.New("dialoguenhance", 1, opts...),
	}
}

type implDialoguenhanceBuilder struct {
	f filter.Filter
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) String() string {
	return dialoguenhanceBuilder.f.String()
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) Outputs() int {
	return dialoguenhanceBuilder.f.Outputs()
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) With(key string, value expr.Expr) filter.Filter {
	return dialoguenhanceBuilder.withOption(key, value)
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) withOption(key string, value expr.Expr) DialoguenhanceBuilder {
	bCopy := *dialoguenhanceBuilder
	bCopy.f = dialoguenhanceBuilder.f.With(key, value)
	return &bCopy
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) Original(original float64) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("original", expr.Double(original))
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) OriginalExpr(original expr.Expr) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("original", original)
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) Enhance(enhance float64) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("enhance", expr.Double(enhance))
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) EnhanceExpr(enhance expr.Expr) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("enhance", enhance)
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) Voice(voice float64) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("voice", expr.Double(voice))
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) VoiceExpr(voice expr.Expr) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("voice", voice)
}

func (dialoguenhanceBuilder *implDialoguenhanceBuilder) Enable(enable expr.Expr) DialoguenhanceBuilder {
	return dialoguenhanceBuilder.withOption("enable", enable)
}
