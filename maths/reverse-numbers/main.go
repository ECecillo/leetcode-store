package main

import (
	"fmt"
	"math"
)

// save at compile time the result so we dont calculate
// each time.
var (
	bound = int(math.Pow(2, 31))
)

func validate(x int) bool {
	return x >= -bound && x <= bound
}

func reverse(x int) int {
	var reverse, lastDigit int

	isNegative := false
	if x < 0 {
		x = -x
		isNegative = true
	}

	for x > 0 {
		lastDigit = x % 10
		reverse = (reverse * 10) + lastDigit
		x = x / 10
	}

	if !validate(reverse) {
		return 0
	}

	if isNegative {
		return -reverse
	}
	return reverse
}

func main() {
	res := reverse(-123)
	fmt.Println(res)
}
