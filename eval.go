package expr

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math/big"
)

// Post-order traversal, equivalent to postfix notation.
func Eval(node interface{}) (*big.Int, error) {
	z := &big.Int{}
	switch nn := node.(type) {
	case *ast.BinaryExpr:
		x, xerr := Eval(nn.X)
		if xerr != nil {
			return nil, xerr
		}
		if ystar, ok := nn.Y.(*ast.StarExpr); ok && nn.Op == token.MUL {
			// exponentiation
			y, yerr := Eval(ystar.X)
			if yerr != nil {
				return nil, yerr
			}
			return z.Exp(x, y, nil), nil
		}
		y, yerr := Eval(nn.Y)
		if yerr != nil {
			return nil, yerr
		}
		switch nn.Op {
		case token.ADD:
			return z.Add(x, y), nil
		case token.SUB:
			return z.Sub(x, y), nil
		case token.MUL:
			return z.Mul(x, y), nil
		case token.QUO:
			if y.Cmp(z) == 0 { // 0 denominator
				return nil, DivideByZero
			}
			return z.Quo(x, y), nil
		case token.REM:
			if y.Cmp(z) == 0 {
				return nil, DivideByZero
			}
			return z.Rem(x, y), nil
		case token.AND:
			return z.And(x, y), nil
		case token.OR:
			return z.Or(x, y), nil
		case token.XOR:
			return z.Xor(x, y), nil
		case token.SHL:
			if y.Cmp(z) < 0 { // negative shift
				return nil, NegativeShift
			}
			return z.Lsh(x, uint(y.Int64())), nil
		case token.SHR:
			if y.Cmp(z) < 0 {
				return nil, NegativeShift
			}
			return z.Rsh(x, uint(y.Int64())), nil
		case token.AND_NOT:
			return z.AndNot(x, y), nil
		default:
			return nil, UnknownOpErr
		}
	case *ast.BasicLit:
		switch nn.Kind {
		case token.INT:
			z.SetString(nn.Value, 0)
			return z, nil
		default:
			return nil, UnknownLitErr
		}
	case *ast.ParenExpr:
		x, xerr := Eval(nn.X)
		if xerr != nil {
			return nil, xerr
		}
		return x, nil
	default:
		return nil, UnknownTokenErr
	}
	panic("unreachable")
}

// Evaluate an expression in a string.
func EvalString(expr string) (*big.Int, error) {
	tree, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, err
	}
	var z *big.Int
	z, err = Eval(tree)
	return z, err
}
