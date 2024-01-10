package hangman

import "fmt"

func Endgame(currentHangman HangManData) (bool) {
	var word string
	var IsWin bool

	for _, v := range currentHangman.Word {
		word += string(string(v)[0])
	}

	if currentHangman.Attempts == 0 {
		fmt.Println("You lose, the word was : " + currentHangman.ToFind)
	} else if word == currentHangman.ToFind {
		IsWin = true
		fmt.Println("Good job, the word was : " + currentHangman.ToFind)
	}
	return IsWin
}
