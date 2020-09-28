package gopher

import (
	"github.com/ilievss/lango/stringutils"
	"github.com/ilievss/lango/translate/lang/en/gopher/rule"
)

type RuleBasedWordTranslator struct {
	rules []rule.WordRule
}

func NewRuleBasedWordTranslator() WordTranslator {
	return &RuleBasedWordTranslator{
		[]rule.WordRule{
			rule.VowelPrefixWordRule,
			rule.XrPrefixWordRule,
			rule.ConsonantsQuPrefixWordRule,
			rule.ConsonantsPrefixWordRule,
		},
	}
}

func (t *RuleBasedWordTranslator) TranslateWord(word string) (string, error) {
	if word == stringutils.Empty {
		return stringutils.Empty, FailedToTranslateError{}
	}
	for _, currentRule := range t.rules {
		result, err := currentRule.Apply(word)
		if err == nil {
			return result, nil
		}
	}
	return stringutils.Empty, FailedToTranslateError{}
}
