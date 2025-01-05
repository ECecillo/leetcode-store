package main

import (
	"math"
	"testing"
)

// On doit trouver le min dans le tableau.

// - Ce min ne peut pas être à la fin sinon on ne fait pas de profit.

// Une fois que l'on a le minimum alors on doit trouver le maximum
// On fait la difference entre les deux et ça nous donne le profit.

// approche lineaire:

// Si le tableau est vide on retourne 0.

// J'init mon min à la valeur de l'élément index 0 du tableau.
// J'init ma variable minIndex à 0.

// J'init un max à la valeur de l'élément index 0 du tabelau.
// J'init ma variable maxIndex à 0.

// Je parcours le tableau:

// - si la valeur à l'index i est inferieur au min alors new min, je change minIndex par l'index i.
// - si la valeur à l'index i est supérieur au max alors new max, je change maxIndex par l'index i.
// Je continue.

func TestMain(t *testing.T) {
	testCases := []struct {
		Name     string
		Given    []int
		Expected int
	}{
		{
			Name:     "return 0 when we only have one element",
			Given:    []int{1},
			Expected: 0,
		},
		{
			Name:     "return 0 when we have more than 10^5 elements",
			Given:    make([]int, int(math.Pow10(5))),
			Expected: 0,
		},
		{
			Name:     "return 0 when the minimum number of the array is at the end",
			Given:    []int{5, 6, 1},
			Expected: 0,
		},
		{
			Name:     "return 0 when all values are the same",
			Given:    []int{1, 1, 1, 1},
			Expected: 0,
		},
		{
			Name:     "return 0 as the maximum profit made",
			Given:    []int{7, 6, 4, 3, 1},
			Expected: 0,
		},
		{
			Name:     "return 5 as the maximum profit made",
			Given:    []int{7, 1, 5, 3, 6, 4},
			Expected: 5,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			got := maxProfit(test.Given)

			if got != test.Expected {
				t.Errorf("expected %d but got %d", test.Expected, got)
			}

		})
	}
}
