// This is free software. It comes without any warranty, to the extent
// permitted by applicable law. You can redistribute it and/or modify it under
// the termos of the Do What the Fuck You Want To Public License, Version 2,
// as published by Sam Hocevar. See COPYING for more details.

package main

import (
	"github.com/zephyrtronium/expr"
	"fmt"
	"os"
	"strings"
)

func main() {
	z, err := expr.EvalString(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(z)
	}
}
