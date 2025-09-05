package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	chSl := make(chan int)
	chSq := make(chan int)

	go makeSlice(chSl)
	go squareSlice(chSl, chSq)

	for i := range chSq {
		fmt.Printf("%d ", i)
	}
}

func makeSlice(chSl chan int) {
	sl := make([]int, 10)
	for i := range sl {
		sl[i] = rand.IntN(101)
		chSl <- sl[i]
	}
	close(chSl)
}

func squareSlice(chSl chan int, chSq chan int) {
	for i := range chSl {
		chSq <- i * i
	}
	close(chSq)
}
