package hangman

import (
	"bufio"
	"os"
)

func GetAlphabetAscii() (map[string]string, int) {
	var alphabetAscii = make(map[string]string)
	var temp string
	var i, j, height int
	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	if len(os.Args) > 2 && os.Args[2] == "--letterFile"{
		alphabetFile, err := os.Open(os.Args[3]) // on ouvre le fichier
		CheckError(err)                          // on verifie les erreur
		defer alphabetFile.Close()               // on ferme le fichier après utilisation

		scanner := bufio.NewScanner(alphabetFile)

		for scanner.Scan() {
			if scanner.Text() == "" && len(temp) > 3 {
				alphabetAscii[string(alphabet[i])] = temp[:len(temp)-1]
				temp = ""
				i++
				height = max(height, j)
				j = 0
			} else {
				temp += scanner.Text() + "\n"
				j++
			}
		}
		alphabetAscii["_"] = temp[:len(temp)-1]
	}
	return alphabetAscii, height
}
