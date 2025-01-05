package main

import (
	"math"
	"reflect"
	"testing"
)

func TestCountingBits(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "with n=2",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "with n=5",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name:     "n can not be greater than 10 power of 5",
			n:        2 * int(math.Pow10(5)),
			expected: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := countingBits(tC.n)

			if !reflect.DeepEqual(got, tC.expected) {
				t.Errorf("expected %v but got %v", tC.expected, got)
			}
		})
	}
}

func BenchmarkCountingBitsRegular(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countingBits(i)
	}
}
