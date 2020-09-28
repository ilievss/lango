package gopher

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type BasicEnTranslator struct {
	WordTranslator     WordTranslator
	SentenceTranslator SentenceTranslator
}

func NewBasicGopherTranslator() EnTranslator {
	return &BasicEnTranslator{
		NewRuleBasedWordTranslator(),
		NewSimpleSentenceTranslator(),
	}
}

func (t *BasicEnTranslator) Translate(text string) (string, error) {
	if strings.Contains(text, stringutils.Space) {
		return t.SentenceTranslator.TranslateSentence(text)
	}
	return t.WordTranslator.TranslateWord(text)
}
