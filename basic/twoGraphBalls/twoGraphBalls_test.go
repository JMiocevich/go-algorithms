package twoGraphBalls

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {

	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	if got := twoCrystalBalls(data); got != idx {
		t.Errorf("two_crystal_balls(data) = %v; want %v", got, idx)
	}

	secondData := make([]bool, 821)
	if got := twoCrystalBalls(secondData); got != -1 {
		t.Errorf("two_crystal_balls(secondData) = %v; want -1", got)
	}
}
