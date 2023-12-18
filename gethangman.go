package hangman

import (
	"io/ioutil"
	"os"
)

func GetHangman(file string) [10]string {
	var hangmanPosition [10]string
	var j int
	var temp string

	hangman, err := os.Open(file) // on ouvre le fichier
	CheckError(err)               // on verifie les erreur
	defer hangman.Close()         // on ferme le fichier après utilisation

	content, err := ioutil.ReadFile(file)
	CheckError(err)

	for i, a := range content {
		if i%71 == 0 && i != 0 {
			hangmanPosition[j] += string(temp)
			temp = " "
			j++
		} else {
			temp += string(a)
		}
	}
	hangmanPosition[j] += string(temp)
	return hangmanPosition
}
