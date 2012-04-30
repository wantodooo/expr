package expr

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math/big"
)

// Post-order traversal, equivalent to postfix notation.
func Eval(node interface{}) (*big.Int, error) {
	z := big.NewInt(0)
	switch nn := node.(type) {
	case *ast.BinaryExpr:
		x, xerr := Eval(nn.X)
		if xerr != nil {
			return nil, xerr
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
			return z.Quo(x, y), nil
		case token.REM:
			return z.Rem(x, y), nil
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
