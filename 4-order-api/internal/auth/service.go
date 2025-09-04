package auth

import (
	"math/rand"
	"time"
)

func CreateSession() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	b := make([]rune, 15)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

func CreateCode() string {
	var letters = []rune("0123456789")
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	b := make([]rune, 4)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}
