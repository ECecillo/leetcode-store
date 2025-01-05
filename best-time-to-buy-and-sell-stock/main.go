package main

import (
	"fmt"
	"math"
)

var (
	pricesLengthLimit = int(math.Pow10(5))
)

// maxProfit return the maximum amount of profit that can be made
// given a pricing array.
//
// # Constraints
//
// - 1 <= prices.length <= 10^5
//
// - 0 <= prices[i] <= 10^4
func maxProfit(prices []int) int {
	if len(prices) <= 1 || len(prices) >= pricesLengthLimit {
		return 0
	}

	var (
		min      int
		minIndex int

		max      int
		maxIndex int
	)

	min = prices[0]
	minIndex = 0

	for i, price := range prices {
		if price < min {
			min = price
			minIndex = i

			// Start to save who is the max number
			max = prices[i]
			maxIndex = i
		}

		if price > max {
			max = price
			maxIndex = i
		}
	}

	// if the minimum is at the last position of the array then we return 0.
	if minIndex == len(prices)-1 {
		return 0
	}

	fmt.Printf("Best day to buy is %d and best day to sell is %d \n", minIndex, maxIndex)
	return max - min
}

func main() {
	fmt.Println("Best time to buy and and sell stock")
}
