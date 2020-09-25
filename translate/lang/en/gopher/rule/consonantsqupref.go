package rule

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type consonantsQuPrefixWordRuleInner struct{}

var ConsonantsQuPrefixWordRule = &consonantsQuPrefixWordRuleInner{}

func (r *consonantsQuPrefixWordRuleInner) Apply(word string) (string, error) {
	if !stringutils.IsWord(word) {
		return empty, RuleNotApplicableError{"rule does not apply to this word: " + word}
	}

	cluster := stringutils.GetConsonantClusterPrefix(word, "q")
	if cluster == empty {
		return empty, RuleNotApplicableError{"Word does not start with a consonant"}
	}

	removedCluster := strings.Replace(word, cluster, "", 1)
	if !strings.HasPrefix(removedCluster, "qu") {
		return empty, RuleNotApplicableError{"Word does not contain 'qu' after consonants"}
	}
	return strings.ReplaceAll(removedCluster, "qu", "") + cluster + "quogo", nil
}
