package stringutils

import (
	"fmt"
	"testing"
)

func TestGetConsonantsSkipZero(t *testing.T) {
	actual := getConsonants("")
	expected := consonants
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants with no skipping."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipOne(t *testing.T) {
	actual := getConsonants("x")
	expected := "bcdfghjklmnpqrstvwzBCDFGHJKLMNPQRSTVWZ"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping 'x'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipLowerCase(t *testing.T) {
	actual := getConsonants("wtkz")
	expected := "bcdfghjlmnpqrsvxBCDFGHJLMNPQRSVX"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random consonants 'wtkz'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipUpperCase(t *testing.T) {
	actual := getConsonants("QSX")
	expected := "bcdfghjklmnprtvwzBCDFGHJKLMNPRTVWZ"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping consonants 'bnm'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipMixedCase(t *testing.T) {
	actual := getConsonants("dghStWX")
	expected := "bcfjklmnpqrvzBCFJKLMNPQRVZ"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping consonants 'bnm'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipVowels(t *testing.T) {
	actual := getConsonants("auiy")
	expected := "bcdfghjklmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping vowels 'auiy'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipMixedConsonantsAndVowelsLowerCase(t *testing.T) {
	actual := getConsonants("watykuiz")
	expected := "bcdfghjlmnpqrsvxBCDFGHJLMNPQRSVX"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random consonants and vowels 'watykuiz'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipMixedConsonantsAndVowelsUpperCase(t *testing.T) {
	actual := getConsonants("WATYKUIZ")
	expected := "bcdfghjlmnpqrsvxBCDFGHJLMNPQRSVX"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random consonants and vowels 'watykuiz'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipMixedConsonantsAndVowelsMixedCase(t *testing.T) {
	actual := getConsonants("WaTykUiZ")
	expected := "bcdfghjlmnpqrsvxBCDFGHJLMNPQRSVX"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random consonants and vowels 'watykuiz'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipRepeatingConsonantsLowerCase(t *testing.T) {
	actual := getConsonants("wwzzmmjj")
	expected := "bcdfghklnpqrstvxBCDFGHKLNPQRSTVX"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random repeating consonants 'wwzzmmjj'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipRepeatingConsonantsUpperCase(t *testing.T) {
	actual := getConsonants("WWZZMMJJ")
	expected := "bcdfghklnpqrstvxBCDFGHKLNPQRSTVX"
	if actual != expected {
		t.Error(fmt.Sprintf("Failed to get correct consonants when skipping random repeating consonants 'wwzzmmjj'."+
			" Expected: '%s' | Got: '%s'", expected, actual))
	}
}

func TestGetConsonantsSkipRepeatingConsonantsMixedCase(t *testing.T) {
	actual := getConsonants("wWzZmmJJ")
	expected := "bcdfghklnpqrstvxBCDFGHKLNPQRSTVX"
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
