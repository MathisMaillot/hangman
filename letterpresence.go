package hangman

import "fmt"

type HangManData struct {
	Word             string   // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func LetterPresence(inputLetter string, currentHangman HangManData, letterlist *string) (int, bool) {
	if len(inputLetter) != 1 {
		if string(inputLetter) == string(currentHangman.ToFind) {
			fmt.Println("Good job, the word was :" + currentHangman.ToFind)
			return currentHangman.Attempts, true
		} else {
			currentHangman.Attempts = max(currentHangman.Attempts-2, 0)
			*letterlist = listLetterAlreadyIn(inputLetter, letterlist)
			fmt.Println("Wrong guess, ", currentHangman.Attempts, " attempts remaining")
			return currentHangman.Attempts, false
		}
	} else {
		if containsValue(*letterlist, inputLetter) {
			fmt.Println(inputLetter + " has already been tried")
			return currentHangman.Attempts, false
		} else {
			if containsValue(currentHangman.ToFind, inputLetter) {
				fmt.Println("Good guess :)")
				*letterlist = listLetterAlreadyIn(inputLetter, letterlist)
				return currentHangman.Attempts, true
			} else {
				currentHangman.Attempts = currentHangman.Attempts - 1
				*letterlist = listLetterAlreadyIn(inputLetter, letterlist)
				fmt.Println("Not present in the word :(, ", currentHangman.Attempts, " attempts remaining")
				return currentHangman.Attempts, false
			}
		}
	}
}

func containsValue(s string, value string) bool {
	for j := 0; j < len(s); j++ {
		if value == string(s[j]) {
			return true
		}
	}
	return false
}

func listLetterAlreadyIn(inputletter string, s *string) string {
	for i := 0; i < len(inputletter); i++ {
		if !containsValue(*s, string(inputletter[i])) {
			*s += string(inputletter[i])
		}
	}
	return *s
}
