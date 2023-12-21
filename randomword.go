package hangman

import (
	"math/rand"
)

func Randomword(l []string) string {
	word := l[rand.Intn(len(l)-1)]
	return UpperString(word)
}
