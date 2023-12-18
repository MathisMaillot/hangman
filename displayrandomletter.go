package hangman

import (
	"fmt"
	"math/rand"
)

func DisplayRandomletter(currentHangman HangManData, isAscii bool, alphabetAscii map[string]string, height int ) []string {
	var randomIndex int = rand.Intn(len(currentHangman.ToFind) - 1)
	currentHangman.Word[randomIndex] = string(currentHangman.ToFind[randomIndex])
	if isAscii {
		var word string
			for _, v := range currentHangman.Word {
				word += v
			}
			DisplyAscii(word,alphabetAscii, height )
	} else {
		for _, v := range currentHangman.Word {
			fmt.Printf(v + " ")
		}
		fmt.Println()
	}
	return currentHangman.Word
}

