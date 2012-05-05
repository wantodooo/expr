// This is free software. It comes without any warranty, to the extent
// permitted by applicable law. You can redistribute it and/or modify it under
// the termos of the Do What the Fuck You Want To Public License, Version 2,
// as published by Sam Hocevar. See COPYING for more details.

package expr

import "errors"

var (
	UnknownTokenErr = errors.New("don't know what to do with token")
	UnknownOpErr    = errors.New("don't know what to do with operator")
	UnknownLitErr   = errors.New("don't know what to do with literal")
	UnknownFuncErr  = errors.New("unknown function name")

	DivideByZero  = errors.New("division by zero")
	NegativeShift = errors.New("negative shift")

	ArgCountErr = errors.New("incorrect number of arguments")
)
