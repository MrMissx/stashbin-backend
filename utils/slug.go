package utils

import "math/rand"

const length = 10

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func CreateSlug() string {
	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
