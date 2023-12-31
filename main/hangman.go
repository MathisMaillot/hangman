package main

import (
	"fmt"
	"os"

	"github.com/MathisMaillot/hangman"
)

func main() {
	// alphabetAscii, height := hangman.GetAlphabetAscii()
	// currentHangman, isAscii := hangman.StartGame(alphabetAscii, height)
	wordlist := hangman.Wordlist(os.Args[1])
	currentHangman := hangman.HangManData{
		Word: "", 
		ToFind: hangman.Randomword(wordlist),
		Attempts: 10, HangmanPositions: hangman.GetHangman("hangman.txt")}
	for i := 0; i < len(currentHangman.ToFind); i++ {
		currentHangman.Word += "_"
	}
	currentHangman.Word = hangman.DisplayRandomletter(currentHangman)
	fmt.Println("Good Luck, you have", currentHangman.Attempts, "attempts.")
	game(currentHangman)
}

func game(currentHangman hangman.HangManData) {
	var listLetterAlreadyIn string
	var isLetterThere bool

	for !endgame(hangman.HangManData(currentHangman)) {
		var input string

		fmt.Println("Choose : ")
		_, err := fmt.Scan(&input) // Read a single line of input
		hangman.CheckError(err)
		input = hangman.UpperString(input)
		if input == "STOP" {
			hangman.Save(currentHangman)
			break
		}
		currentHangman.Attempts, isLetterThere = hangman.LetterPresence(input, hangman.HangManData(currentHangman), &listLetterAlreadyIn)
		currentHangman.Word = hangman.Displayletter(input, hangman.HangManData(currentHangman), isLetterThere)
	}
}

func endgame(currentHangman hangman.HangManData) bool {
	var word string
	var gameIsFinished bool

	for _, v := range currentHangman.Word {
		word += string(string(v)[0])
	}

	if currentHangman.Attempts == 0 {
		gameIsFinished = true
		fmt.Println("You lose, the word was : " + currentHangman.ToFind)
	} else if word == currentHangman.ToFind {
		gameIsFinished = true
		fmt.Println("Good job, the word was : " + currentHangman.ToFind)
	}
	return gameIsFinished
}
