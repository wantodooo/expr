package expr

import "errors"

var (
	UnknownTokenErr = errors.New("don't know what to do with token")
	UnknownOpErr    = errors.New("don't know what to do with operator")
	UnknownLitErr   = errors.New("don't know what to do with literal")

	DivideByZero  = errors.New("division by zero")
	NegativeShift = errors.New("negative shift")

	ArgCountErr = errors.New("incorrect number of arguments")
)
