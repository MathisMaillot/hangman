package github.com/MathisMaillot/hangman

import (
	"fmt"
)

func Displayletter(inputLetter string, currentHangman HangManData, isLetterThere bool, isAscii bool, alphabetAscii map[string]string, height int) []string {
	if isLetterThere {
		for _, v := range containsValueIndex(currentHangman.ToFind, inputLetter) {
			currentHangman.Word[v] = inputLetter
		}
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
	} else {
		fmt.Printf(currentHangman.HangmanPositions[9-currentHangman.Attempts])
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

