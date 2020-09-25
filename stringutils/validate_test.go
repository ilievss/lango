package stringutils

import (
	"fmt"
	"testing"
)

func TestIsWordValid(t *testing.T) {
	str := "hair"
	actual := IsWord(str)
	expected := true
	if actual != expected {
		t.Error(fmt.Sprintf("Function IsWord returned wrong result for string '%s'."+
			" Expected: %t | Got: %t", str, expected, actual))
	}
}

func TestIsWordTwoWord(t *testing.T) {
	str := "long hair"
	actual := IsWord(str)
	expected := false
	if actual != expected {
		t.Error(fmt.Sprintf("Function IsWord returned wrong result for string '%s'."+
			" Expected: %t | Got: %t", str, expected, actual))
	}
}

func TestIsWordSentence(t *testing.T) {
	str := "Sam has long hair"
	actual := IsWord(str)
	expected := false
	if actual != expected {
		t.Error(fmt.Sprintf("Function IsWord returned wrong result for string '%s'."+
			" Expected: %t | Got: %t", str, expected, actual))
	}
}

func TestIsWordEmpty(t *testing.T) {
	str := ""
	actual := IsWord(str)
	expected := false
	if actual != expected {
		t.Error(fmt.Sprintf("Function IsWord returned wrong result for string '%s'."+
			" Expected: %t | Got: %t", str, expected, actual))
	}
}
