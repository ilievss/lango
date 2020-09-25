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
		return empty, RuleNotApplicableError{"Supplied word is not a word: " + word}
	}
	if !strings.HasPrefix(word, xrPrefix) {
		return empty, RuleNotApplicableError{"Word doesn't have an 'xr' prefix: " + word}
	}
	return "ge" + word, nil
}
