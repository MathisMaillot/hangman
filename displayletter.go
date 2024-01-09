package hangman

import (
	"fmt"
)

func Displayletter(inputLetter string, currentHangman HangManData, isLetterThere bool) string {
	if isLetterThere {
		for _, v := range containsValueIndex(currentHangman.ToFind, inputLetter) {
			for i := 0; i < len(currentHangman.Word); i++ {
				if  i == v {
					currentHangman.Word += string(currentHangman.ToFind[v]) + " "
					
				} else {
					currentHangman.Word += "_ "
				}
			}
		}
		// if isAscii {
		// 	var word string
		// 	for _, v := range currentHangman.Word {
		// 		word += v
		// 	}
		// 	DisplyAscii(word,alphabetAscii, height )
		// } else {
		fmt.Println(currentHangman.Word)
		// }
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

