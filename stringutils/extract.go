package stringutils

import "strings"

const consonants = "bcdfghjklmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ"

func getConsonants(skip string) string {
	newConsonants := ""
	skipBothCases := strings.ToLower(skip) + strings.ToUpper(skip)
	for i := 0; i < len(consonants); i++ {
		char := consonants[i : i+1]
		if strings.IndexAny(skipBothCases, char) == -1 {
			newConsonants += char
		}
	}
	return newConsonants
}

func GetConsonantClusterPrefix(str, skip string) string {
	desiredConsonants := getConsonants(skip)
	cluster := ""
	for i := 0; i < len(str); i++ {
		subPrefix := str[i : i+1]
		if strings.IndexAny(subPrefix, desiredConsonants) == 0 {
			cluster += subPrefix
		} else {
			break
		}
	}
	return cluster
}
