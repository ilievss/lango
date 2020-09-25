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
	cluster := stringutils.GetConsonantClusterPrefix(word, "x")
	if cluster == empty {
		return empty, RuleNotApplicableError{"rule does not apply to this word: " + word}
	}
	translated := strings.ReplaceAll(word, cluster, "") + cluster + "ogo"
	return translated, nil
}
