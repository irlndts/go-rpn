package rpn

import (
	"fmt"
	"testing"
)

var parse_tests = map[string]string{
	"3 + 4 * 2 / ( 1 - 5 )":     "3 4 2 * 1 5 - / +",
	"8 * 3 - 5 + ( 1 - 8 * 2 )": "8 3 * 5 - 1 8 2 * - +",
}

var calc_tests = map[string]float64{
	"3 4 2 * 1 5 - / +":     1,
	"8 3 * 5 - 1 8 2 * - +": 4,
}

func TestParse(t *testing.T) {
	for k, v := range parse_tests {
		assert(t, Parse(k) == v, fmt.Sprintf("Result looked for %s, and received %s ", v, Parse(k)))
	}
	/*
		check(t, err)
	*/
}

func TestCalc(t *testing.T) {
	for k, v := range calc_tests {
		res, err := Calc(k)
		check(t, err)
		assert(t, res == v, fmt.Sprintf("Result looked for %s, and received %s ", v, res))
	}
	/*
	 */
}

func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

func assert(t *testing.T, condition bool, assertion string) {
	if !condition {
		t.Errorf("Assertion failed: %v", assertion)
	}
}
