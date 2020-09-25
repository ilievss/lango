package rule

import (
	"fmt"
	"testing"
)

func TestVowelPrefixWordRuleApplyValid(t *testing.T) {
	str := "air"
	actual, err := VowelPrefixWordRule.Apply(str)
	expected := "gair"
	if err != nil || actual != expected {
		t.Error(fmt.Sprintf("Failed to correctly apply the vowel prefix rule for '%s'."+
			" Expected: '%s' | Got: '%s' | Error: %s", str, expected, actual, err))
	}
}

func TestVowelPrefixWordRuleApplyStartWithConsonant(t *testing.T) {
	str := "chair"
	actual, err := VowelPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestVowelPrefixWordRuleApplyTwoWords(t *testing.T) {
	str := "iteration two"
	actual, err := VowelPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the vowel prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestVowelPrefixWordRuleApplySentence(t *testing.T) {
	str := "Expectations were high in the room"
	actual, err := VowelPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the vowel prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestVowelPrefixWordRuleApostrophe(t *testing.T) {
	str := "ain't"
	actual, err := VowelPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the vowel prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestVowelPrefixWordRuleNumber(t *testing.T) {
	str := "89712837897123"
	actual, err := VowelPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the vowel prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestVowelPrefixWordRuleSpecialChars(t *testing.T) {
	str := "&^&@%^#"
	actual, err := VowelPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the vowel prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}
