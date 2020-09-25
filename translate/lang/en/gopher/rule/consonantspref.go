package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type consonantsPrefixWordRuleInner struct{}

var ConsonantsPrefixWordRule = &consonantsPrefixWordRuleInner{}

func (r *consonantsPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return stringutils.Empty, RuleNotApplicableError{"Given string is not a word: " + word}
	}
	if strings.Contains(word, stringutils.Apostrophe) {
		return stringutils.Empty, RuleNotApplicableError{"Word has apostrophes: " + word}
	}
	cluster := stringutils.GetConsonantClusterPrefix(word, "x")
	if cluster == stringutils.Empty {
		return stringutils.Empty, RuleNotApplicableError{"Word does not start with a consonant: " + word}
	}
	translated := strings.ReplaceAll(word, cluster, "") + cluster + "ogo"
	return translated, nil
}
