package rpn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// A list of operations with priorities which are currently available
var opa = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

// Parsing a string via shunting-yard algorithm to the rpn expression
// input: expression string
// returns RPN style string
// for more information about Shunting-Yard look at https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func Parse(e string) (rpn string) {
	var stack []string // holds operators and left parenthesis
	for _, tok := range strings.Fields(e) {
		switch tok {
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					rpn += " " + op              // add it to result
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // add operand to result
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

// Revers Polish Notation calculator
// input is a string of rpn-style expression
// returns a result
// for more details about RPN look at https://en.wikipedia.org/wiki/Reverse_Polish_notation
func Calc(expr string) (float64, error) {
	tokens := strings.Split(expr, " ")
	stack := make([]int, 0)

	for _, token := range tokens {
		i, err := strconv.Atoi(token)

		if err == nil {
			// token is a digit, so push it to stack
			stack = append(stack, i)
		} else {
			// token is not an int digit
			switch token {
			case "+", "-", "*", "/", "^":
				stack, err = operation(stack, token)
				if err != nil {
					return 0, err
				}
				break
			case " ":
				break
			default:
				return 0, errors.New(fmt.Sprintf("Error: symbol '%s' is not available for rpn calculator", token))
			}
		}
	}

	result, _ := pop(stack)
	return float64(result), nil
}

func operation(stack []int, op string) ([]int, error) {
	a, stack := pop(stack)
	b, stack := pop(stack)
	switch op {
	case "+":
		return append(stack, a+b), nil
	case "-":
		return append(stack, b-a), nil
	case "*":
		return append(stack, a*b), nil
	case "^":
		return append(stack, a^b), nil
	case "/":
		if a == 0 {
			return nil, errors.New(fmt.Sprintf("Error: Division by Zero"))
		}
		return append(stack, b/a), nil
	default:
		return nil, errors.New(fmt.Sprintf("Error: symbol '%s' is not available for rpn calculator", op))
	}
}

func pop(a []int) (int, []int) {
	return a[len(a)-1], a[:len(a)-1]
}
