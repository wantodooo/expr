package expr

import (
	"go/ast"
	"go/parser"
	"go/token"
	"math/big"
)

// Post-order traversal, equivalent to postfix notation.
func Eval(node interface{}) (*big.Int, error) {
	switch nn := node.(type) {
	case *ast.BinaryExpr:
		z := new(big.Int)
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
	case *ast.UnaryExpr:
		z := new(big.Int)
		switch nn.Op {
		case token.SUB: // -x
			return z.Neg(nn.X), nil
		case token.XOR: // ^x
			return z.Not(nn.X), nil
		case token.ADD: // +x (useless)
			return nn.X, nil
		}
	case *ast.BasicLit:
		z := new(big.Int)
		switch nn.Kind {
		case token.INT:
			z.SetString(nn.Value, 0)
			return z, nil
		default:
			return nil, UnknownLitErr
		}
	case *ast.ParenExpr:
		z, err := Eval(nn.X)
		if err != nil {
			return nil, err
		}
		return z, nil
	case *ast.CallExpr:
		ident, ok := nn.Fun.(*ast.Ident)
		if !ok {
			return nil, UnknownTokenErr // quarter to four am; dunno correct error
		}
		var f Func
		f, ok = FuncMap[ident.Name]
		if !ok {
			return nil, UnknownFuncErr
		}
		var aerr error
		args := make([]*big.Int, len(nn.Args))
		for i, a := range nn.Args {
			if args[i], aerr = Eval(a); aerr != nil {
				return nil, aerr
			}
		}
		x, xerr := f.Call(args...)
		if xerr != nil {
			return nil, xerr
		}
		return x, nil
	}
	return nil, UnknownTokenErr
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
