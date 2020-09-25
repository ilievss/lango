package gopher

import (
	"github.com/ilievss/lango/stringutils"
	"github.com/ilievss/lango/translate/lang/en/gopher/rule"
)

type GopherToEnTranslator struct {
	rules []rule.WordRule
}

func New() GopherToEnTranslator {
	return GopherToEnTranslator{
		[]rule.WordRule{
			rule.VowelPrefixWordRule,
			rule.XrPrefixWordRule,
			rule.ConsonantsQuPrefixWordRule,
			rule.ConsonantsPrefixWordRule,
		},
	}
}

func (t *GopherToEnTranslator) Translate(text string) (string, error) {
	for _, currentRule := range t.rules {
		result, err := currentRule.Apply(text)
		if err == nil {
			return result, nil
		}
	}
	return stringutils.Empty, FailedToTranslateError{}
}
