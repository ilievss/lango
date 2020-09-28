package gopher

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type BasicGopherTranslator struct {
	WordTranslator     WordTranslator
	SentenceTranslator SentenceTranslator
}

func NewBasicGopherTranslator() EnTranslator {
	return &BasicGopherTranslator{
		NewRuleBasedWordTranslator(),
		NewSimpleSentenceTranslator(),
	}
}

func (t *BasicGopherTranslator) Translate(text string) (string, error) {
	if strings.Contains(text, stringutils.Space) {
		return t.SentenceTranslator.TranslateSentence(text)
	}
	return t.WordTranslator.TranslateWord(text)
}
