package rng

import (
	"math/rand"
)

type MapRng struct {
}

func Pick(m int) int {
	return rand.Intn(m)
}
