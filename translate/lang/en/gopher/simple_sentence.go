package gopher

import (
	"github.com/ilievss/lango/stringutils"
	"strings"
)

type SimpleSentenceTranslator struct {
	WordTranslator
}

func NewSimpleSentenceTranslator() SentenceTranslator {
	return &SimpleSentenceTranslator{
		NewRuleBasedWordTranslator(),
	}
}

func (t *SimpleSentenceTranslator) TranslateSentence(text string) (string, error) {
	if text == stringutils.Empty {
		return stringutils.Empty, FailedToTranslateError{}
	}
	words := strings.Split(text, stringutils.Space)
	// the last word will contain a punctuation mark,
	// so we have to remove that character from the word
	// before we translate it
	numWords := len(words)
	lastWordIdx := numWords - 1
	lastWordLen := len(words[lastWordIdx])
	punctuationMark := words[lastWordIdx][lastWordLen-1 : lastWordLen]
	// remove the punctuation mark
	words[lastWordIdx] = words[lastWordIdx][0 : lastWordLen-1]

	translatedWords := make([]string, 0, numWords)
	for _, word := range words {
		result, err := t.TranslateWord(word)
		if err != nil {
			return stringutils.Empty, err
		}
		translatedWords = append(translatedWords, result)
	}

	// add the punctuation mark
	translatedWords[numWords-1] = translatedWords[numWords-1] + punctuationMark
	return strings.Join(translatedWords, stringutils.Space), nil
}
