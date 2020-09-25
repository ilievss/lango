package stringutils

import "strings"

const space = " "
const nonWordCharacters = "1234567890`_=+[];\"\\|,./<>?!@#$%^&*()"

func IsWord(word string) bool {
	return len(word) > 0 &&
		len(strings.Split(word, space)) == 1 &&
		strings.IndexAny(word, nonWordCharacters) == -1
}
