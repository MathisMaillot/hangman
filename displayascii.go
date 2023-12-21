package github.com/MathisMaillot/hangman

import "fmt"

func DisplyAscii(word string, alphabetAscii map[string]string, height int) {
	var matrix []string
	for i := 0; i < height - 2; i++ {
		line := ""
		for _, letter := range word {
			art, exists := alphabetAscii[string(letter)]
			if exists {
				lines := splitLines(art)
				line += lines[i] + "  "
			}
		}
		matrix = append(matrix, line)
	}
	for _, line := range matrix {
		fmt.Println(line)
	}
}

func splitLines(s string) []string {
	var lines []string
	currentLine := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLine += string(r)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}
