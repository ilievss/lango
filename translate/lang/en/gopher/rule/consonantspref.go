package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type consonantsPrefixWordRuleInner struct{}

var ConsonantsPrefixWordRule = &consonantsPrefixWordRuleInner{}

func (r *consonantsPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return empty, RuleNotApplicableError{"Given string is not a word: " + word}
	}
	if strings.Contains(word, apostrophe) {
		return empty, RuleNotApplicableError{"Word has apostrophes: " + word}
	}
	cluster := stringutils.GetConsonantClusterPrefix(word, "x")
	if cluster == empty {
		return empty, RuleNotApplicableError{"Word does not start with a consonant: " + word}
	}
	translated := strings.ReplaceAll(word, cluster, "") + cluster + "ogo"
	return translated, nil
}
