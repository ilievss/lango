package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

const xrPrefix = "xr"

type xrPrefixWordRuleInner struct{}

var XrPrefixWordRule = &xrPrefixWordRuleInner{}

func (r *xrPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return stringutils.Empty, RuleNotApplicableError{"Supplied string is not a word: " + word}
	}
	if strings.Contains(word, stringutils.Apostrophe) {
		return stringutils.Empty, RuleNotApplicableError{"Word has apostrophes: " + word}
	}
	if !strings.HasPrefix(word, xrPrefix) {
		return stringutils.Empty, RuleNotApplicableError{"Word doesn't have an 'xr' prefix: " + word}
	}
	return "ge" + word, nil
}
