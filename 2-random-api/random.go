package main

import (
	"math/rand/v2"
	"strconv"
)

func Random() string {
	i := rand.IntN(7)
	if i == 0 {
		i += 1
	}
	return strconv.Itoa(i)
}
