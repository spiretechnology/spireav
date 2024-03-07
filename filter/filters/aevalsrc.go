package filters

import (
	"strings"

	"github.com/spiretechnology/spireav/filter/expr"
)

func (b *implAudioEvalSourceBuilder) ChannelExpressions(exps ...expr.Expr) AudioEvalSourceBuilder {
	var expsStr []string
	for _, exp := range exps {
		expsStr = append(expsStr, exp.String())
	}
	return b.withOption("exprs", expr.Expr(strings.Join(expsStr, "|")))
}

func (b *implAudioEvalSourceBuilder) ChannelLayouts(layouts ...string) AudioEvalSourceBuilder {
	return b.withOption("c", expr.String(strings.Join(layouts, "|")))
}
