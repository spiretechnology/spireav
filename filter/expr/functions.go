package expr

import (
	"fmt"
	"strings"
)

// Div is a divide expression.
func Div(exprs ...Expr) Expr {
	return chainMath("/", exprs...)
}

// Mul is a multiply expression.
func Mul(exprs ...Expr) Expr {
	return chainMath("*", exprs...)
}

// Add is an addition expression.
func Add(exprs ...Expr) Expr {
	return chainMath("+", exprs...)
}

// Sub is a subtract expression.
func Sub(exprs ...Expr) Expr {
	return chainMath("-", exprs...)
}

// ToBool converts x to a boolean value.
func ToBool(x Expr) Expr {
	return If(Gte(x, Int(1)), Int(1), Int(0))
}

// And returns the logical AND of all the expressions.
func And(x ...Expr) Expr {
	bools := make([]Expr, len(x))
	for i, e := range x {
		bools[i] = ToBool(e)
	}
	return Mul(bools...)
}

// Or returns the logical OR of all the expressions.
func Or(x ...Expr) Expr {
	bools := make([]Expr, len(x))
	for i, e := range x {
		bools[i] = ToBool(e)
	}
	return ToBool(Add(bools...))
}

func chainMath(op string, exprs ...Expr) Expr {
	strs := make([]string, len(exprs))
	for i, e := range exprs {
		strs[i] = e.String()
	}
	return Raw(fmt.Sprintf("(%s)", strings.Join(strs, op)))
}

func functionCall(name string, args ...Expr) Expr {
	strs := make([]string, len(args))
	for i, e := range args {
		strs[i] = e.String()
	}
	return Raw(fmt.Sprintf("%s(%s)", name, strings.Join(strs, ",")))
}

// Abs computes absolute value of x.
func Abs(x Expr) Expr {
	return functionCall("abs", x)
}

// Acos computes arccosine of x.
func Acos(x Expr) Expr {
	return functionCall("acos", x)
}

// Asin computes arcsine of x.
func Asin(x Expr) Expr {
	return functionCall("asin", x)
}

// Atan computes arctangent of x.
func Atan(x Expr) Expr {
	return functionCall("atan", x)
}

// Atan2 computes arctangent of y/x.
func Atan2(y, x Expr) Expr {
	return functionCall("atan2", y, x)
}

// Between returns 1 if x is in the interval [min, max], 0 otherwise.
func Between(x, min, max Expr) Expr {
	return functionCall("between", x, min, max)
}

// BitAnd computes bitwise AND of x and y.
func BitAnd(x, y Expr) Expr {
	return functionCall("bitand", x, y)
}

// BitOr computes bitwise OR of x and y.
func BitOr(x, y Expr) Expr {
	return functionCall("bitor", x, y)
}

// Ceil rounds the value of expression expr upwards to the nearest integer. For example, "ceil(1.5)" is "2.0".
func Ceil(x Expr) Expr {
	return functionCall("ceil", x)
}

// Clip returns the value of x clipped between min and max.
func Clip(x, min, max Expr) Expr {
	return functionCall("clip", x, min, max)
}

// Cos computes cosine of x.
func Cos(x Expr) Expr {
	return functionCall("cos", x)
}

// Cosh computes hyperbolic cosine of x.
func Cosh(x Expr) Expr {
	return functionCall("cosh", x)
}

// Eq returns 1 if x is equal to y, 0 otherwise.
func Eq(x, y Expr) Expr {
	return functionCall("eq", x, y)
}

// Exp computes the exponential function of x, which is e raised to the power of x.
func Exp(x Expr) Expr {
	return functionCall("exp", x)
}

// Floor rounds the value of expression expr downwards to the nearest integer. For example, "floor(1.5)" is "1.0".
func Floor(x Expr) Expr {
	return functionCall("floor", x)
}

// Gauss computes Gauss function of x, corresponding to exp(-x*x/2) / sqrt(2*PI).
func Gauss(x Expr) Expr {
	return functionCall("gauss", x)
}

// Gcd return the greatest common divisor of x and y. If both x and y are 0 or either or both are less than zero then behavior is undefined.
func Gcd(x, y Expr) Expr {
	return functionCall("gcd", x, y)
}

// Gt returns 1 if x is greater than y, 0 otherwise.
func Gt(x, y Expr) Expr {
	return functionCall("gt", x, y)
}

// Gte returns 1 if x is greater than or equal to y, 0 otherwise.
func Gte(x, y Expr) Expr {
	return functionCall("gte", x, y)
}

// Hypot returns the length of the hypotenuse of a right-angled triangle with sides of length x and y.
func Hypot(x, y Expr) Expr {
	return functionCall("hypot", x, y)
}

// If evaluates x, and if the result is non-zero return the evaluation result of y, otherwise the evaluation result of z.
func If(x, y, z Expr) Expr {
	return functionCall("if", x, y, z)
}

// IfNot evaluates x, and if the result is zero return the evaluation result of y, otherwise the evaluation result of z.
func IfNot(x, y, z Expr) Expr {
	return functionCall("ifnot", x, y, z)
}

// IsInf returns 1.0 if x is +/-INFINITY, 0.0 otherwise.
func IsInf(x Expr) Expr {
	return functionCall("isinf", x)
}

// IsNan returns 1.0 if x is NAN, 0.0 otherwise.
func IsNan(x Expr) Expr {
	return functionCall("isnan", x)
}

// Load loads the value of the internal variable with index idx, which was previously stored with st(idx, expr). The function returns the loaded value.
func Load(idx Expr) Expr {
	return functionCall("ld", idx)
}

// Lerp returns linear interpolation between x and y by amount of z.
func Lerp(x, y, a Expr) Expr {
	return functionCall("lerp", x, y, a)
}

// Log computes natural logarithm of x.
func Log(x Expr) Expr {
	return functionCall("log", x)
}

// Lt returns 1 if x is lesser than y, 0 otherwise.
func Lt(x, y Expr) Expr {
	return functionCall("lt", x, y)
}

// Lte returns 1 if x is lesser than or equal to y, 0 otherwise.
func Lte(x, y Expr) Expr {
	return functionCall("lte", x, y)
}

// Max returns the maximum between x and y.
func Max(x Expr, y ...Expr) Expr {
	if len(y) == 0 {
		return x
	}
	if len(y) == 1 {
		return functionCall("max", x, y[0])
	}
	return functionCall("max", x, Max(y[0], y[1:]...))
}

// Min returns the minimum between x and y.
func Min(x Expr, y ...Expr) Expr {
	if len(y) == 0 {
		return x
	}
	if len(y) == 1 {
		return functionCall("min", x, y[0])
	}
	return functionCall("min", x, Min(y[0], y[1:]...))
}

// Mod computes the remainder of division of x by y.
func Mod(x, y Expr) Expr {
	return functionCall("mod", x, y)
}

// Not returns 1.0 if expr is zero, 0.0 otherwise.
func Not(x Expr) Expr {
	return functionCall("not", x)
}

// Pow computes the power of x elevated y, it is equivalent to "(x)^(y)".
func Pow(x, y Expr) Expr {
	return functionCall("pow", x, y)
}

// Print prints the value of expression t with loglevel l. If l is not specified then a default log level is used. Return the value of the expression printed.
func Print(x Expr) Expr {
	return functionCall("print", x)
}

// Random returns a pseudo random value between 0.0 and 1.0. idx is the index of the internal variable used to save the seed/state, which can be previously stored with st(idx).
func Random(idx Expr) Expr {
	return functionCall("random", idx)
}

// RandomI returns a pseudo random value in the interval between min and max. idx is the index of the internal variable which will be used to save the seed/state, which can be previously stored with st(idx).
func RandomI(idx, min, max Expr) Expr {
	return functionCall("randomi", idx, min, max)
}

// Root finds an input value for which the function represented by expr with argument ld(0) is 0 in the interval 0..max.
// The expression in expr must denote a continuous function or the result is undefined.
// ld(0) is used to represent the function input value, which means that the given expression will be evaluated multiple times with various input values that the expression can access through ld(0). When the expression evaluates to 0 then the corresponding input value will be returned.
func Root(expr, max Expr) Expr {
	return functionCall("root", expr, max)
}

// Round rounds the value of expression expr to the nearest integer. For example, "round(1.5)" is "2.0".
func Round(x Expr) Expr {
	return functionCall("round", x)
}

// Sign computes the sign of x.
func Sign(x Expr) Expr {
	return functionCall("sgn", x)
}

// Sin computes sin of x.
func Sin(x Expr) Expr {
	return functionCall("sin", x)
}

// Sinh computes hyperbolic sine of x.
func Sinh(x Expr) Expr {
	return functionCall("sinh", x)
}

// Sqrt computes square root of x.
func Sqrt(x Expr) Expr {
	return functionCall("sqrt", x)
}

// Squish computes the expression 1/(1 + exp(4*x)).
func Squish(x Expr) Expr {
	return functionCall("squish", x)
}

// Store stores the value of the expression expr in an internal variable. idx specifies the index of the variable where to store the value, and it is a value ranging from 0 to 9. The function returns the value stored in the internal variable.
// The stored value can be retrieved with ld(var).
func Store(idx, expr Expr) Expr {
	return functionCall("st", idx, expr)
}

// Tan computes tangent of x.
func Tan(x Expr) Expr {
	return functionCall("tan", x)
}

// Tanh computes hyperbolic tangent of x.
func Tanh(x Expr) Expr {
	return functionCall("tanh", x)
}

// Taylor evaluates a Taylor series at x, given an expression representing the ld(idx)-th derivative of a function at 0.
// When the series does not converge the result is undefined.
// ld(idx) is used to represent the derivative order in expr, which means that the given expression will be evaluated multiple times with various input values that the expression can access through ld(idx). If idx is not specified then 0 is assumed.
// Note, when you have the derivatives at y instead of 0, taylor(expr, x-y) can be used.
func Taylor(expr, x, idx Expr) Expr {
	return functionCall("taylor", expr, x, idx)
}

// Time returns the current (wallclock) time in seconds.
func Time() Expr {
	return functionCall("time", Int(0))
}

// Trunc rounds the value of expression expr towards zero to the nearest integer. For example, "trunc(-1.5)" is "-1.0".
func Trunc(x Expr) Expr {
	return functionCall("trunc", x)
}

// While evaluates expression expr while the expression cond is non-zero, and returns the value of the last expr evaluation, or NAN if cond was always false.
func While(cond, expr Expr) Expr {
	return functionCall("while", cond, expr)
}
