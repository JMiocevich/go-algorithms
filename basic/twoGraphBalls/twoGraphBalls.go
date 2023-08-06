package twoGraphBalls

import (
	"fmt"
	"math"
)

func twoCrystalBalls(breaks []bool) int {

	jmpAmount := int(math.Sqrt(float64(len(breaks))))
	fmt.Println("here", len(breaks))

	i := jmpAmount
	for ; i < (len(breaks)); i += jmpAmount {

		if breaks[i] {
			break
		}

	}
	i -= jmpAmount

	for j := 0; j < len(breaks) && i < len(breaks); i, j = i+1, j+1 {

		if breaks[i] {
			return i
		}

	}
	return -1
}
