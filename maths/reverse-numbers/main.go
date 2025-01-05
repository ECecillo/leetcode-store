package main

import "fmt"

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

	if isNegative {
		return -reverse
	}
	return reverse
}

func main() {
	res := reverse(-123)
	fmt.Println(res)
}
