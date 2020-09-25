package rule

import (
	"fmt"
	"testing"
)

func TestConsonantsQuPrefixWordRuleValidWords(t *testing.T) {
	validWords := map[string]string{
		"square":   "aresquogo",
		"squash":   "ashsquogo",
		"squirrel": "irrelsquogo",
	}

	for word, expected := range validWords {
		actual, err := ConsonantsQuPrefixWordRule.Apply(word)
		if err != nil || actual != expected {
			t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
				" Expected: '%s' | Got: '%s' | Error: %s", word, expected, actual, err))
		}
	}
}

func TestConsonantsQuPrefixWordRuleInvalidWords(t *testing.T) {
	validWords := []string{
		"air",
		"horn",
		"deque",
	}

	for _, word := range validWords {
		actual, err := ConsonantsQuPrefixWordRule.Apply(word)
		if _, ok := err.(RuleNotApplicableError); !ok {
			t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
				" Expected: RuleNotApplicableError | Got: '%s'", word, actual))
		}
	}
}

func TestConsonantsQuPrefixWordRuleEmpty(t *testing.T) {
	str := ""
	actual, err := ConsonantsQuPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsQuPrefixWordRuleTwoWords(t *testing.T) {
	str := "square table"
	actual, err := ConsonantsQuPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsQuPrefixWordRuleSentence(t *testing.T) {
	str := "Squaring the side of a square will give you it's area!"
	actual, err := ConsonantsQuPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsQuPrefixWordRuleApostrophe(t *testing.T) {
	str := "don't"
	actual, err := ConsonantsQuPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsQuPrefixWordRuleNumber(t *testing.T) {
	str := "89712837897123"
	actual, err := ConsonantsQuPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestConsonantsQuPrefixWordRuleSpecialChars(t *testing.T) {
	str := "&^&@%^#"
	actual, err := ConsonantsQuPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the consonants + qu prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}
