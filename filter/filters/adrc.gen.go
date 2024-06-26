// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// AdrcBuilder Audio Spectral Dynamic Range Controller.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#adrc
type AdrcBuilder interface {
	filter.Filter
	// Transfer set the transfer expression (default "p").
	Transfer(transfer expr.Expr) AdrcBuilder
	// Attack set the attack (from 1 to 1000) (default 50).
	Attack(attack float64) AdrcBuilder
	// AttackExpr set the attack (from 1 to 1000) (default 50).
	AttackExpr(attack expr.Expr) AdrcBuilder
	// Release set the release (from 5 to 2000) (default 100).
	Release(release float64) AdrcBuilder
	// ReleaseExpr set the release (from 5 to 2000) (default 100).
	ReleaseExpr(release expr.Expr) AdrcBuilder
	// Channels set channels to filter (default "all").
	Channels(channels string) AdrcBuilder
	// ChannelsExpr set channels to filter (default "all").
	ChannelsExpr(channels expr.Expr) AdrcBuilder
	// Enable expression to enable or disable the filter.
	Enable(enable expr.Expr) AdrcBuilder
}

// Adrc creates a new AdrcBuilder to configure the "adrc" filter.
func Adrc(opts ...filter.Option) AdrcBuilder {
	return &implAdrcBuilder{
		f: filter.New("adrc", 1, opts...),
	}
}

type implAdrcBuilder struct {
	f filter.Filter
}

func (adrcBuilder *implAdrcBuilder) String() string {
	return adrcBuilder.f.String()
}

func (adrcBuilder *implAdrcBuilder) Outputs() int {
	return adrcBuilder.f.Outputs()
}

func (adrcBuilder *implAdrcBuilder) With(key string, value expr.Expr) filter.Filter {
	return adrcBuilder.withOption(key, value)
}

func (adrcBuilder *implAdrcBuilder) withOption(key string, value expr.Expr) AdrcBuilder {
	bCopy := *adrcBuilder
	bCopy.f = adrcBuilder.f.With(key, value)
	return &bCopy
}

func (adrcBuilder *implAdrcBuilder) Transfer(transfer expr.Expr) AdrcBuilder {
	return adrcBuilder.withOption("transfer", transfer)
}

func (adrcBuilder *implAdrcBuilder) Attack(attack float64) AdrcBuilder {
	return adrcBuilder.withOption("attack", expr.Double(attack))
}

func (adrcBuilder *implAdrcBuilder) AttackExpr(attack expr.Expr) AdrcBuilder {
	return adrcBuilder.withOption("attack", attack)
}

func (adrcBuilder *implAdrcBuilder) Release(release float64) AdrcBuilder {
	return adrcBuilder.withOption("release", expr.Double(release))
}

func (adrcBuilder *implAdrcBuilder) ReleaseExpr(release expr.Expr) AdrcBuilder {
	return adrcBuilder.withOption("release", release)
}

func (adrcBuilder *implAdrcBuilder) Channels(channels string) AdrcBuilder {
	return adrcBuilder.withOption("channels", expr.String(channels))
}

func (adrcBuilder *implAdrcBuilder) ChannelsExpr(channels expr.Expr) AdrcBuilder {
	return adrcBuilder.withOption("channels", channels)
}

func (adrcBuilder *implAdrcBuilder) Enable(enable expr.Expr) AdrcBuilder {
	return adrcBuilder.withOption("enable", enable)
}
