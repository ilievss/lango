package rule

import (
	"fmt"
	"testing"
)

func TestXrPrefixWordRuleApplyValidWords(t *testing.T) {
	validWords := map[string]string{
		"xremote": "gexremote",
		"xray":    "gexray",
		"xrender": "gexrender",
	}

	for word, expected := range validWords {
		actual, err := XrPrefixWordRule.Apply(word)
		if err != nil || actual != expected {
			t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
				" Expected: '%s' | Got: '%s' | Error: %s", word, expected, actual, err))
		}
	}
}

func TestXrPrefixWordRuleApplyInvalidWords(t *testing.T) {
	validWords := []string{
		"bunny",
		"rabbit",
		"hair",
	}

	for _, word := range validWords {
		actual, err := XrPrefixWordRule.Apply(word)
		if _, ok := err.(RuleNotApplicableError); !ok {
			t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
				" Expected: RuleNotApplicableError | Got: '%s'", word, actual))
		}
	}
}

func TestXrPrefixWordRuleApplyTwoWords(t *testing.T) {
	str := "iteration two"
	actual, err := XrPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestXrPrefixWordRuleApplySentence(t *testing.T) {
	str := "Expectations were high in the room"
	actual, err := XrPrefixWordRule.Apply(str)
	if err == nil {
		t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
			" Expected: error | Got: '%s'", str, actual))
	}
}

func TestXrPrefixWordRuleApostrophe(t *testing.T) {
	str := "don't"
	actual, err := XrPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestXrPrefixWordRuleNumber(t *testing.T) {
	str := "89712837897123"
	actual, err := XrPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}

func TestXrPrefixWordRuleSpecialChars(t *testing.T) {
	str := "&^&@%^#"
	actual, err := XrPrefixWordRule.Apply(str)
	if _, ok := err.(RuleNotApplicableError); !ok {
		t.Error(fmt.Sprintf("Failed to correctly apply the 'xr' prefix rule for '%s'."+
			" Expected: RuleNotApplicableError | Got: '%s'", str, actual))
	}
}
