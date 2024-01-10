package hangman

import (
	"fmt"
	"os"


)

// func main() {
// 	// alphabetAscii, height := hangman.GetAlphabetAscii()
// 	// currentHangman, isAscii := hangman.StartGame(alphabetAscii, height)
// 	wordlist := Wordlist(os.Args[1])
// 	currentHangman := HangManData{
// 		Word: []string{}, 
// 		ToFind: Randomword(wordlist),
// 		Attempts: 10, HangmanPositions: GetHangman("hangman.txt")}	
// 	for i := 0; i < len(currentHangman.ToFind); i++ {
// 		currentHangman.Word = append(currentHangman.Word, "_")
// 	}
// 	currentHangman.Word = DisplayRandomletter(currentHangman)
// 	fmt.Println("Good Luck, you have", currentHangman.Attempts, "attempts.")
// 	game(currentHangman)
// }

func Game(currentHangman HangManData) {
	wordlist := Wordlist(os.Args[1])
	currentHangman = HangManData{
		Word: []string{}, 
		ToFind: Randomword(wordlist),
		Attempts: 10, HangmanPositions: GetHangman("hangman.txt")}	
	for i := 0; i < len(currentHangman.ToFind); i++ {
		currentHangman.Word = append(currentHangman.Word, "_")
	}
	currentHangman.Word = DisplayRandomletter(currentHangman)
	fmt.Println("Good Luck, you have", currentHangman.Attempts, "attempts.")
	var listLetterAlreadyIn string
	var isLetterThere bool
	fmt.Println()
	for !endgame(HangManData(currentHangman)) {
		var input string

		fmt.Println("Choose : ")
		_, err := fmt.Scan(&input) // Read a single line of input
		CheckError(err)
		input = UpperString(input)
		if input == "STOP" {
			Save(currentHangman)
			break
		}
		currentHangman.Attempts, isLetterThere = LetterPresence(input,HangManData(currentHangman), &listLetterAlreadyIn)
		currentHangman.Word = Displayletter(input, HangManData(currentHangman), isLetterThere)
	}
}
