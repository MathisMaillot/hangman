package hangman

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func Wordlist(dictionary string) []string {
	var wordList []string
	// var temp string

	file, err := os.Open(dictionary) // on ouvre le fichier
	CheckError(err)                  // on verifie les erreur
	defer file.Close()               // on ferme le fichier après utilisation

	content, err := ioutil.ReadFile(dictionary)
	// scanner := bufio.NewScanner(file)
	CheckError(err)
	str := strings.NewReader(string(content))
	r := bufio.NewReader(str)
	for {
		word, _, err := r.ReadLine()
		if len(word) > 0 {
			wordList = append(wordList, UpperString(string(word)))
		}
		if err != nil {
			break
		}
	}
	return wordList
}
