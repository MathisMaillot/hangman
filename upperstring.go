package github.com/MathisMaillot/hangman

import "strings"

func UpperString(s string) string {
	s = strings.Replace(s, "à", "a", -1)
	s = strings.Replace(s, "â", "a", -1)
	s = strings.Replace(s, "ä", "a", -1)

	s = strings.Replace(s, "é", "e", -1)
	s = strings.Replace(s, "è", "e", -1)
	s = strings.Replace(s, "ê", "e", -1)
	s = strings.Replace(s, "ë", "e", -1)

	s = strings.Replace(s, "î", "i", -1)
	s = strings.Replace(s, "ï", "i", -1)
	s = strings.Replace(s, "ì", "i", -1)

	s = strings.Replace(s, "î", "i", -1)
	s = strings.Replace(s, "ï", "i", -1)
	s = strings.Replace(s, "ì", "i", -1)

	s = strings.Replace(s, "ò", "o", -1)
	s = strings.Replace(s, "ô", "o", -1)
	s = strings.Replace(s, "ö", "o", -1)

	s = strings.Replace(s, "ù", "u", -1)
	s = strings.Replace(s, "û", "u", -1)
	s = strings.Replace(s, "ü", "u", -1)

	s = strings.Replace(s, "ç", "c", -1)

	s = strings.ToUpper(s)

	s = strings.Replace(s, "À", "A", -1)
	s = strings.Replace(s, "Â", "A", -1)
	s = strings.Replace(s, "Ä", "A", -1)

	s = strings.Replace(s, "È", "E", -1)
	s = strings.Replace(s, "É", "E", -1)
	s = strings.Replace(s, "Ê", "E", -1)
	s = strings.Replace(s, "Ë", "E", -1)

	s = strings.Replace(s, "Ì", "I", -1)
	s = strings.Replace(s, "Í", "I", -1)
	s = strings.Replace(s, "Î", "I", -1)
	s = strings.Replace(s, "Ï", "I", -1)

	s = strings.Replace(s, "Ò", "O", -1)
	s = strings.Replace(s, "Ô", "O", -1)
	s = strings.Replace(s, "Ö", "O", -1)

	s = strings.Replace(s, "Ù", "U", -1)
	s = strings.Replace(s, "Û", "U", -1)
	s = strings.Replace(s, "Ü", "U", -1)

	s = strings.Replace(s, "Ç", "C", -1)

	return s
}
