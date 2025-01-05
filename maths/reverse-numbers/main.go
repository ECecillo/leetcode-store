package main

import (
	"fmt"
	"math"
)

func validate(x int) bool {
	bound := int(math.Pow(2, 31))

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
