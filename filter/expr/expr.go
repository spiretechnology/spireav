package expr

import (
	"fmt"
	"strconv"
	"strings"
)

// Expr is an expression in a complex filter node.
type Expr string

// String gets a string representation of the expression.
func (e Expr) String() string {
	return string(e)
}

func Size(width, height int) Expr {
	return Expr(fmt.Sprintf("%dx%d", width, height))
}

// String creates a string expression.
func String(str string) Expr {
	if strings.Contains(str, ":") {
		str = strings.ReplaceAll(str, ":", "\\:")
		str = fmt.Sprintf("'%s'", str)
	}
	return Expr(str)
}

// Bool creates an expression for a boolean value.
func Bool(val bool) Expr {
	if val {
		return Expr("1")
	} else {
		return Expr("0")
	}
}

// Var creates an expression for a variable.
func Var(name string) Expr {
	return Expr(name)
}

// Int creates an expression for an integer.
func Int(val int) Expr {
	return Expr(strconv.FormatInt(int64(val), 10))
}

// Float creates an expression for a float.
func Float(val float64) Expr {
	return Expr(fmt.Sprintf("%0.2f", val))
}

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

func chainMath(op string, exprs ...Expr) Expr {
	strs := make([]string, len(exprs))
	for i, e := range exprs {
		strs[i] = e.String()
	}
	return Expr(fmt.Sprintf("(%s)", strings.Join(strs, op)))
}

func functionCall(name string, args ...Expr) Expr {
	strs := make([]string, len(args))
	for i, e := range args {
		strs[i] = e.String()
	}
	return Expr(fmt.Sprintf("%s(%s)", name, strings.Join(strs, ",")))
}

func Sin(x Expr) Expr {
	return functionCall("sin", x)
}

func Cos(x Expr) Expr {
	return functionCall("cos", x)
}

func Sin0to1(x Expr) Expr {
	num := Add(Sin(Mul(x, Var("PI"))), Int(1))
	return Div(num, Int(2))
}
