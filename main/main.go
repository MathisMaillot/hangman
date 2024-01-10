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
		Word: []string{}, 
		ToFind: hangman.Randomword(wordlist),
		Attempts: 10, HangmanPositions: hangman.GetHangman("hangman.txt")}	
	for i := 0; i < len(currentHangman.ToFind); i++ {
		currentHangman.Word = append(currentHangman.Word, "_")
	}
	currentHangman.Word = hangman.DisplayRandomletter(currentHangman)
	fmt.Println("Good Luck, you have", currentHangman.Attempts, "attempts.")
	game(currentHangman)
}

func game(currentHangman hangman.HangManData) {
	var listLetterAlreadyIn string
	var isLetterThere bool
	fmt.Println()
	for !hangman.Endgame(currentHangman) {
		var input string

		fmt.Println("Choose : ")
		_, err := fmt.Scan(&input) // Read a single line of input
		hangman.CheckError(err)
		input = hangman.UpperString(input)
		if input == "STOP" {
			hangman.Save(currentHangman)
			break
		}
		currentHangman.Attempts, isLetterThere = hangman.LetterPresence(input, currentHangman, &listLetterAlreadyIn)
		currentHangman.Word = hangman.Displayletter(input, currentHangman, isLetterThere)
	}
}
