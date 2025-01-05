package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	maxNumberAllowed float64 = math.Pow10(5)
)

func detectOne(binaryString string) int {
	numberOfOne := 0
	for _, char := range binaryString {
		if char == '1' {
			numberOfOne++
		}
	}
	return numberOfOne
}

func countingBits(n int) (ans []int) {

	if n > int(maxNumberAllowed) {
		return nil
	}

	// reduce memory re-allocation
	ans = make([]int, n+1)

	for i := range n + 1 {
		iBinNumber := strconv.FormatInt(int64(i), 2)
		ans[i] = detectOne(iBinNumber)
	}
	return ans
}

func main() {
	fmt.Println("counting bits")
}
