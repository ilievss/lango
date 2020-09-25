package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type consonantsQuPrefixWordRuleInner struct{}

var ConsonantsQuPrefixWordRule = &consonantsQuPrefixWordRuleInner{}

func (r *consonantsQuPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return stringutils.Empty, RuleNotApplicableError{"Supplied string is not a word: " + word}
	}
	if strings.Contains(word, stringutils.Apostrophe) {
		return stringutils.Empty, RuleNotApplicableError{"Word has apostrophes: " + word}
	}
	cluster := stringutils.GetConsonantClusterPrefix(word, "q")
	if cluster == stringutils.Empty {
		return stringutils.Empty, RuleNotApplicableError{"Word does not start with a consonant"}
	}

	removedCluster := strings.Replace(word, cluster, "", 1)
	if !strings.HasPrefix(removedCluster, "qu") {
		return stringutils.Empty, RuleNotApplicableError{"Word does not contain 'qu' after consonants"}
	}
	return strings.ReplaceAll(removedCluster, "qu", "") + cluster + "quogo", nil
}
