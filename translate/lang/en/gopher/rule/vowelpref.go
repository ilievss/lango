package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

const empty = ""
const vowels string = "aeiouy"

type vowelPrefixWordRuleInner struct{}

var VowelPrefixWordRule = &vowelPrefixWordRuleInner{}

func (r *vowelPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return empty, RuleNotApplicableError{"Supplied word is not a word: " + word}
	}
	if strings.IndexAny(word, vowels) != 0 {
		return empty, RuleNotApplicableError{"Word does not start with a vowel: " + word}
	}
	return "g" + word, nil
}
