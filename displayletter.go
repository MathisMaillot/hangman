package hangman

import (
	"fmt"
)

func Displayletter(inputLetter string, currentHangman HangManData, isLetterThere bool) []string {
	if isLetterThere {
		for _, v := range containsValueIndex(currentHangman.ToFind, inputLetter) {
			currentHangman.Word[v] = inputLetter
		}

		for _, v := range currentHangman.Word {
			fmt.Printf(v + " ")
		}
		fmt.Println()
	} else {
		fmt.Printf(currentHangman.HangmanPositions[10-currentHangman.Attempts])
	}
	return currentHangman.Word
}

func containsValueIndex(s string, value string) []int {
	var listIndex []int
	for j := 0; j < len(s); j++ {
		if value == string(s[j]) {
			listIndex = append(listIndex, j)
		}
	}
	return listIndex
}
