// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// EqBuilder Adjust brightness, contrast, gamma, and saturation.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#eq
type EqBuilder interface {
	filter.Filter
	// Contrast set the contrast adjustment, negative values give a negative image (default "1.0").
	Contrast(contrast string) EqBuilder
	// ContrastExpr set the contrast adjustment, negative values give a negative image (default "1.0").
	ContrastExpr(contrast expr.Expr) EqBuilder
	// Brightness set the brightness adjustment (default "0.0").
	Brightness(brightness string) EqBuilder
	// BrightnessExpr set the brightness adjustment (default "0.0").
	BrightnessExpr(brightness expr.Expr) EqBuilder
	// Saturation set the saturation adjustment (default "1.0").
	Saturation(saturation string) EqBuilder
	// SaturationExpr set the saturation adjustment (default "1.0").
	SaturationExpr(saturation expr.Expr) EqBuilder
	// Gamma set the initial gamma value (default "1.0").
	Gamma(gamma string) EqBuilder
	// GammaExpr set the initial gamma value (default "1.0").
	GammaExpr(gamma expr.Expr) EqBuilder
	// GammaR gamma value for red (default "1.0").
	GammaR(gammaR string) EqBuilder
	// GammaRExpr gamma value for red (default "1.0").
	GammaRExpr(gammaR expr.Expr) EqBuilder
	// GammaG gamma value for green (default "1.0").
	GammaG(gammaG string) EqBuilder
	// GammaGExpr gamma value for green (default "1.0").
	GammaGExpr(gammaG expr.Expr) EqBuilder
	// GammaB gamma value for blue (default "1.0").
	GammaB(gammaB string) EqBuilder
	// GammaBExpr gamma value for blue (default "1.0").
	GammaBExpr(gammaB expr.Expr) EqBuilder
	// GammaWeight set the gamma weight which reduces the effect of gamma on bright areas (default "1.0").
	GammaWeight(gammaWeight string) EqBuilder
	// GammaWeightExpr set the gamma weight which reduces the effect of gamma on bright areas (default "1.0").
	GammaWeightExpr(gammaWeight expr.Expr) EqBuilder
	// Eval specify when to evaluate expressions (from 0 to 1) (default init).
	Eval(eval int) EqBuilder
}

// Eq creates a new EqBuilder to configure the "eq" filter.
func Eq(opts ...filter.Option) EqBuilder {
	return &implEqBuilder{
		f: filter.New("eq", 1, opts...),
	}
}

type implEqBuilder struct {
	f filter.Filter
}

func (eqBuilder *implEqBuilder) String() string {
	return eqBuilder.f.String()
}

func (eqBuilder *implEqBuilder) Outputs() int {
	return eqBuilder.f.Outputs()
}

func (eqBuilder *implEqBuilder) With(key string, value expr.Expr) filter.Filter {
	return eqBuilder.withOption(key, value)
}

func (eqBuilder *implEqBuilder) withOption(key string, value expr.Expr) EqBuilder {
	bCopy := *eqBuilder
	bCopy.f = eqBuilder.f.With(key, value)
	return &bCopy
}

func (eqBuilder *implEqBuilder) Contrast(contrast string) EqBuilder {
	return eqBuilder.withOption("contrast", expr.String(contrast))
}

func (eqBuilder *implEqBuilder) ContrastExpr(contrast expr.Expr) EqBuilder {
	return eqBuilder.withOption("contrast", contrast)
}

func (eqBuilder *implEqBuilder) Brightness(brightness string) EqBuilder {
	return eqBuilder.withOption("brightness", expr.String(brightness))
}

func (eqBuilder *implEqBuilder) BrightnessExpr(brightness expr.Expr) EqBuilder {
	return eqBuilder.withOption("brightness", brightness)
}

func (eqBuilder *implEqBuilder) Saturation(saturation string) EqBuilder {
	return eqBuilder.withOption("saturation", expr.String(saturation))
}

func (eqBuilder *implEqBuilder) SaturationExpr(saturation expr.Expr) EqBuilder {
	return eqBuilder.withOption("saturation", saturation)
}

func (eqBuilder *implEqBuilder) Gamma(gamma string) EqBuilder {
	return eqBuilder.withOption("gamma", expr.String(gamma))
}

func (eqBuilder *implEqBuilder) GammaExpr(gamma expr.Expr) EqBuilder {
	return eqBuilder.withOption("gamma", gamma)
}

func (eqBuilder *implEqBuilder) GammaR(gammaR string) EqBuilder {
	return eqBuilder.withOption("gamma_r", expr.String(gammaR))
}

func (eqBuilder *implEqBuilder) GammaRExpr(gammaR expr.Expr) EqBuilder {
	return eqBuilder.withOption("gamma_r", gammaR)
}

func (eqBuilder *implEqBuilder) GammaG(gammaG string) EqBuilder {
	return eqBuilder.withOption("gamma_g", expr.String(gammaG))
}

func (eqBuilder *implEqBuilder) GammaGExpr(gammaG expr.Expr) EqBuilder {
	return eqBuilder.withOption("gamma_g", gammaG)
}

func (eqBuilder *implEqBuilder) GammaB(gammaB string) EqBuilder {
	return eqBuilder.withOption("gamma_b", expr.String(gammaB))
}

func (eqBuilder *implEqBuilder) GammaBExpr(gammaB expr.Expr) EqBuilder {
	return eqBuilder.withOption("gamma_b", gammaB)
}

func (eqBuilder *implEqBuilder) GammaWeight(gammaWeight string) EqBuilder {
	return eqBuilder.withOption("gamma_weight", expr.String(gammaWeight))
}

func (eqBuilder *implEqBuilder) GammaWeightExpr(gammaWeight expr.Expr) EqBuilder {
	return eqBuilder.withOption("gamma_weight", gammaWeight)
}

func (eqBuilder *implEqBuilder) Eval(eval int) EqBuilder {
	return eqBuilder.withOption("eval", expr.Int(eval))
}
