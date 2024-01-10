package hangman

import "fmt"

func Endgame(currentHangman HangManData) (bool, bool) {
	var word string
	var IsWin bool
	var gameIsFinished bool

	for _, v := range currentHangman.Word {
		word += string(string(v)[0])
	}

	if currentHangman.Attempts == 0 {
		gameIsFinished = true
		fmt.Println("You lose, the word was : " + currentHangman.ToFind)
	} else if word == currentHangman.ToFind {
		IsWin = true
		gameIsFinished = true
		fmt.Println("Good job, the word was : " + currentHangman.ToFind)
	}
	return IsWin, gameIsFinished
}
