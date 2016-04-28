package main

import (
	"fmt"
	"github.com/irlndts/go-rpn"
)

func main() {
	fmt.Println(rpn.Parse("3 + 4 * 2 / ( 1 - 5 ) "))
	fmt.Println(rpn.Calc(rpn.Parse("3 + 4 * 2 / ( 1 - 5 )")))

}
