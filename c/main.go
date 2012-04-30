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
