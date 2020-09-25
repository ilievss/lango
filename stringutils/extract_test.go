package stringutils

import (
	"fmt"
	"testing"
)

func TestGetConsonantsSkipZero(t *testing.T) {
	actual := getConsonants("")
	expected := "qwrtpsfghjklzxcvbnm"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants with no skipping."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipOne(t *testing.T) {
	actual := getConsonants("x")
	expected := "qwrtpsfghjklzcvbnm"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping 'x'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipConsecutiveBeginning(t *testing.T) {
	actual := getConsonants("qwr")
	expected := "tpsfghjklzxcvbnm"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping consecutive consonants'qwr'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipRandom(t *testing.T) {
	actual := getConsonants("wtkz")
	expected := "qrpsfghjlxcvbnm"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random consonants 'wtkz'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipConsecutiveEnding(t *testing.T) {
	actual := getConsonants("bnm")
	expected := "qwrtpsfghjklzxcv"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping consonants 'bnm'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipVowels(t *testing.T) {
	actual := getConsonants("auiy")
	expected := "qwrtpsfghjklzxcvbnm"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping vowels 'auiy'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipMixedConsonantsAndVowels(t *testing.T) {
	actual := getConsonants("watykuiz")
	expected := "qrpsfghjlxcvbnm"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random consonants and vowels 'watykuiz'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipRepeatingConsonants(t *testing.T) {
	actual := getConsonants("wwzzmmjj")
	expected := "qrtpsfghklxcvbn"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random repeating consonants 'wwzzmmjj'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantClusterSingleChar(t *testing.T) {
	str := "bird"
	actual := GetConsonantClusterPrefix(str, "")
	expected := "b"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonant cluster from '%s'."+
			" Expected: '%s' | Got: '%s'", str, expected, actual))
	}
}

func TestGetConsonantClusterTwoChars(t *testing.T) {
	str := "chair"
	actual := GetConsonantClusterPrefix(str, "")
	expected := "ch"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonant cluster from '%s'."+
			" Expected: '%s' | Got: '%s'", str, expected, actual))
	}
}

func TestGetConsonantClusterThreeChars(t *testing.T) {
	str := "split"
	actual := GetConsonantClusterPrefix(str, "")
	expected := "spl"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonant cluster from '%s'."+
			" Expected: '%s' | Got: '%s'", str, expected, actual))
	}
}

func TestGetConsonantClusterVowel(t *testing.T) {
	str := "acute"
	actual := GetConsonantClusterPrefix(str, "")
	expected := ""
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonant cluster from '%s'."+
			" Expected: '%s' | Got: '%s'", str, expected, actual))
	}
}
