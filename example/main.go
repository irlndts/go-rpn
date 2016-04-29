package main

import (
	"fmt"
	"github.com/irlndts/go-rpn"
)

func main() {
	fmt.Println(rpn.Parse("8 * 3 - 5 + ( 1 - 8 * 2 )"))
	fmt.Println(rpn.Calc(rpn.Parse("8 * 3 - 5 + ( 1 - 8 * 2 )")))

}
