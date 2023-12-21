package hangman

func CheckError(e error) { // fonction pour arrêter le programme si erreur
	if e != nil {
		panic(e)
	}
}
