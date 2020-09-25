package rule

import (
	"fmt"
	"testing"
)

func TestConsonantsPrefixWordRuleSingleConsonant(t *testing.T) {
	str := "hair"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	expected := "airhogo"
	if err != nil || actual != expected {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonant prefix rule for '%s'."+
			" Expected: '%s' | Got: '%s' | Error: %s", str, expected, actual, err))
	}
}

func TestConsonantsPrefixWordRuleDoubleConsonant(t *testing.T) {
	str := "chair"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	expected := "airchogo"
	if err != nil || actual != expected {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonant prefix rule for '%s'."+
			" Expected: '%s' | Got: '%s' | Error: %s", str, expected, actual, err))
	}
}

func TestConsonantsPrefixWordRuleTripleConsonant(t *testing.T) {
	str := "split"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	expected := "itsplogo"
	if err != nil || actual != expected {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonant prefix rule for '%s'."+
			" Expected: '%s' | Got: '%s' | Error: %s", str, expected, actual, err))
	}
}

func TestConsonantsPrefixWordRuleVowel(t *testing.T) {
	str := "air"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsPrefixWordRuleEmpty(t *testing.T) {
	str := ""
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsPrefixWordRuleTwoWord(t *testing.T) {
	str := "air blast"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsPrefixWordRuleSentence(t *testing.T) {
	str := "That horn has got some air blast!"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsPrefixWordRuleApostrophe(t *testing.T) {
	str := "don't"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonant prefix rule for '%s'."+
			" Expected RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsPrefixWordRuleNumber(t *testing.T) {
	str := "89712837897123"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsPrefixWordRuleSpecialChars(t *testing.T) {
	str := "&^&@%^#"
	actual, err := ConsonantsPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}
