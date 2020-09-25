package stringutils

import "strings"

const space = " "

func IsWord(word string) bool {
	return len(word) > 0 && len(strings.Split(word, space)) == 1
}
