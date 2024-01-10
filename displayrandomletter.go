package hangman

import (
	"fmt"
	"math/rand"
)

func DisplayRandomletter(currentHangman HangManData) []string {
	var randomIndex int = rand.Intn(len(currentHangman.ToFind) - 1)
	currentHangman.Word[randomIndex] = string(currentHangman.ToFind[randomIndex])
	for _, v := range currentHangman.Word {
		fmt.Printf(v + " ")
	}
	fmt.Println()
	return currentHangman.Word
}
