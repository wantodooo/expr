package main

import (
	"github.com/zephyrtronium/expr"
	"fmt"
	"os"
)

func main() {
	z, err := expr.EvalString(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(z)
	}
}
