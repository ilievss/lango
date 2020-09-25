package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

const vowels string = "aeiouy"

type vowelPrefixWordRuleInner struct{}

var VowelPrefixWordRule = &vowelPrefixWordRuleInner{}

func (r *vowelPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return stringutils.Empty, RuleNotApplicableError{"Supplied word is not a word: " + word}
	}
	if strings.Contains(word, stringutils.Apostrophe) {
		return stringutils.Empty, RuleNotApplicableError{"Word has apostrophes: " + word}
	}
	if strings.IndexAny(word, vowels) != 0 {
		return stringutils.Empty, RuleNotApplicableError{"Word does not start with a vowel: " + word}
	}
	return "g" + word, nil
}
