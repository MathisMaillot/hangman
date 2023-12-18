package github.com/MathisMaillot/hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func StartGame(alphabetAscii map[string]string, height int) (HangManData, bool) {
	var currentHangman HangManData
	var content []byte
	var isAscii bool

	if len(os.Args) == 3 && os.Args[1] == "--startWith" {
		saveFile, err := os.Open(os.Args[2]) // on ouvre le fichier
		CheckError(err)                      // on verifie les erreur
		defer saveFile.Close()               // on ferme le fichier après utilisation

		scanner := bufio.NewScanner(saveFile)
		CheckError(err)
		for scanner.Scan() {
			content = scanner.Bytes()
		}

		err = json.Unmarshal(content, &currentHangman)
		CheckError(err)
		fmt.Println("Welcome Back, you have", currentHangman.Attempts, "attempts remaining.")

		for _, letter := range currentHangman.Word {
			fmt.Printf(letter)
			fmt.Printf(" ")
		}
		fmt.Println()
	} else if len(os.Args) == 2 {
		wordlist := Wordlist(os.Args[1])
		currentHangman = HangManData{Word: []string{}, ToFind: Randomword(wordlist), Attempts: 10, HangmanPositions: GetHangman("hangman.txt")}

		for i := 0; i < len(currentHangman.ToFind); i++ {
			currentHangman.Word = append(currentHangman.Word, "_")
		}
		currentHangman.Word = DisplayRandomletter(currentHangman, isAscii, alphabetAscii, height)
		fmt.Println("Good Luck, you have", currentHangman.Attempts, "attempts.")
	} else if len(os.Args) == 4 && os.Args[2] == "--letterFile" {
		isAscii = true
		wordlist := Wordlist(os.Args[1])
		currentHangman = HangManData{Word: []string{}, ToFind: Randomword(wordlist), Attempts: 10, HangmanPositions: GetHangman("hangman.txt")}

		for i := 0; i < len(currentHangman.ToFind); i++ {
			currentHangman.Word = append(currentHangman.Word, "_")
		}
		currentHangman.Word = DisplayRandomletter(currentHangman, isAscii, alphabetAscii, height)
		fmt.Println()
		fmt.Println("Good Luck, you have", currentHangman.Attempts, "attempts.")

	} else {
		fmt.Println("Error syntax, please retry")
		os.Exit(1)
	}
	return currentHangman, isAscii
}
