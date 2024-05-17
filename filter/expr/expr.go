package expr

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// PI is the area of the unit disc, approximately 3.14.
	PI = Var("PI")
	// E is exp(1) (Euler’s number), approximately 2.718.
	E = Var("E")
	// PHI is the golden ratio (1+sqrt(5))/2, approximately 1.618.
	PHI = Var("PHI")
)

// Expr is an expression in a complex filter node.
type Expr interface {
	String() string
}

// Raw is a raw expression that is not escaped.
type Raw string

func (r Raw) String() string {
	return string(r)
}

// Var is a variable in an expression.
type Var string

func (v Var) String() string {
	return string(v)
}

// Size is an image size with width and height.
type Size struct {
	Width, Height int
}

func (s Size) String() string {
	return fmt.Sprintf("%dx%d", s.Width, s.Height)
}

// Rational is a video rate or rational value with a numerator and denominator.
type Rational struct {
	Num int
	Den int
}

func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}

// String is a string literal in an expression.
type String string

func (s String) String() string {
	str := string(s)
	if strings.Contains(str, ":") {
		str = strings.ReplaceAll(str, ":", "\\:")
	}
	if strings.Contains(str, " ") {
		str = fmt.Sprintf("'%s'", str)
	}
	return str
}

// Bool is a boolean literal in an expression.
type Bool bool

func (b Bool) String() string {
	if b {
		return "1"
	}
	return "0"
}

// Int is an integer literal in an expression.
type Int int

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Int64 is an int64 literal in an expression.
type Int64 int64

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Float is a float literal in an expression.
type Float float32

func (f Float) String() string {
	return fmt.Sprintf("%0.2f", f)
}

// Double is a double literal in an expression.
type Double float64

func (d Double) String() string {
	return fmt.Sprintf("%0.2f", d)
}

// Color is a color literal in an expression.
type Color string

func (c Color) WithOpacity(opacity float64) Color {
	return Color(fmt.Sprintf("%s@%0.2f", c, opacity))
}

func (c Color) String() string {
	return string(c)
}

// Duration is a duration literal in an expression.
type Duration time.Duration

func (d Duration) String() string {
	return fmt.Sprintf("%0.3f", time.Duration(d).Seconds())
}

// PixFmt is a pixel format literal in an expression.
type PixFmt string

func (p PixFmt) String() string {
	return string(p)
}

// SampleFmt is a sample format literal in an expression.
type SampleFmt string

func (s SampleFmt) String() string {
	return string(s)
}

// ChannelLayout is a channel layout literal in an expression.
type ChannelLayout string

func (c ChannelLayout) String() string {
	return string(c)
}

// Dictionary is a dictionary literal in an expression.
type Dictionary map[string]Expr

func (d Dictionary) String() string {
	var parts []string
	for key, value := range d {
		parts = append(parts, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(parts, "\\:")
}

// Flags is a list of flags in an expression.
type Flags []string

func (f Flags) String() string {
	return strings.Join(f, "+")
}
