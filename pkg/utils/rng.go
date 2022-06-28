package utils

import (
	"math/rand"
)

func Pick(m int) int {
	return rand.Intn(m)
}
