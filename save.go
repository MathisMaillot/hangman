package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Save(hangman HangManData) HangManData {
	var currentHangman HangManData
	dataHangman, err := json.Marshal(hangman)
	CheckError(err)
	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("save.txt", dataHangman, 0777)
	CheckError(err)
	// err = json.Unmarshal(b, &currentHangman)
	// CheckError(err)
	fmt.Println("Game Saved in save.txt.")
	return currentHangman
}
