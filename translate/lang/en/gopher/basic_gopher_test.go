package gopher

import (
	"fmt"
	"testing"
)

func TestBasicEnTranslator_TranslateSentence_ValidWords(t *testing.T) {
	textToTranslation := map[string]string{
		"squats are good for your health!": "atssquogo gare oodgogo orfogo gyour ealthhogo!",
		"We will be having lunch!":         "eWogo illwogo ebogo avinghogo unchlogo!",
	}

	translator := NewBasicGopherTranslator()
	for text, translation := range textToTranslation {
		actual, err := translator.Translate(text)
		if err != nil || actual != translation {
			t.Error(fmt.Sprintf("Failed to translate '%s'."+
				" Expected: '%s' | Got: '%s' | Error: %s", text, translation, actual, err))
		}
	}
}

func TestBasicEnTranslator_TranslateSentence_InvalidWords(t *testing.T) {
	untranslatable := []string{
		"don't look at me!",
		"878728347234",
		"*@#&*&*&@*#",
		"235235*@#&*&*&@*#",
		"*@#&*&23355*&@*#35235",
		"872348jasdhjasdhj",
		"word767",
		"",
	}

	translator := NewBasicGopherTranslator()
	for _, text := range untranslatable {
		actual, err := translator.Translate(text)
		if _, ok := err.(FailedToTranslateError); !ok {
			t.Error(fmt.Sprintf("Expected error when translating '%s'."+
				" Expected: FailedToTranslateError | Got: '%s'", text, actual))
		}
	}
}
