package expr

import (
	"math/big"
	"math/rand"
	"time"
)

type Func func(v ...*big.Int) (*big.Int, error)

func Abs(v ...*big.Int) (*big.Int, error) {
	if len(v) != 1 {
		return nil, ArgCountErr
	}
	return new(big.Int).Abs(v[0]), nil
}

func Binomial(v ...*big.Int) (*big.Int, error) {
	if len(v) != 2 {
		return nil, ArgCountErr
	}
	return new(big.Int).Binomial(v[0].Int64(), v[1].Int64()), nil
}

func Lb(v ...*big.Int) (*big.Int, error) {
	if len(v) != 1 {
		return nil, ArgCountErr
	}
	return big.NewInt(int64(v[0].BitLen())), nil
}

func Exp(v ...*big.Int) (*big.Int, error) {
	switch len(v) {
	case 2:
		return new(big.Int).Exp(v[0], v[1], nil), nil
	case 3:
		return new(big.Int).Exp(v[0], v[1], v[2]), nil
	}
	return nil, ArgCountErr
}

func GCD(v ...*big.Int) (*big.Int, error) {
	if len(v) != 2 {
		return nil, ArgCountErr
	}
	return new(big.Int).GCD(nil, nil, v[0], v[1]), nil
}

func ModInv(v ...*big.Int) (*big.Int, error) {
	if len(v) != 2 {
		return nil, ArgCountErr
	}
	return new(big.Int).ModInverse(v[0], v[1]), nil
}

func Factorial(v ...*big.Int) (*big.Int, error) {
	switch len(v) {
	case 1:
		return new(big.Int).MulRange(2, v[0].Int64()), nil
	case 2:
		return new(big.Int).MulRange(v[0].Int64(), v[1].Int64()), nil
	}
	return nil, ArgCountErr
}

func Rand(v ...*big.Int) (*big.Int, error) {
	switch len(v) {
	case 1:
		return new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), v[0]), nil
	case 2:
		delta := new(big.Int).Sub(v[1], v[0])
		r := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), delta)
		return r.Add(r, v[0]), nil
	}
	return nil, ArgCountErr
}

var FuncMap = map[string]Func{
	"abs":       Abs,
	"binomial":  Binomial,
	"lb":        Lb,
	"pow":       Exp,
	"gcd":       GCD,
	"modinv":    ModInv,
	"factorial": Factorial,
	"rand":      Rand,
}
