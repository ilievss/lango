package gopher

import (
	"fmt"
	"testing"
)

func TestGopherToEnTranslator_TranslateValidWords(t *testing.T) {
	textToTranslation := map[string]string{
		"squander":    "andersquogo",
		"squarespace": "arespacesquogo",
		"squeeze":     "eezesquogo",
		"arrow":       "garrow",
		"chest":       "estchogo",
		"xray":        "gexray",
	}

	translator := New()
	for text, translation := range textToTranslation {
		actual, err := translator.Translate(text)
		if err != nil || actual != translation {
			t.Error(fmt.Sprintf("Failed to translate '%s'."+
				" Expected: '%s' | Got: '%s' | Error: %s", text, translation, actual, err))
		}
	}
}

func TestGopherToEnTranslator_TranslateInvalidWords(t *testing.T) {
	untranslatable := []string{
		"don't",
		"some phrase",
		"A sentence used only for testing.",
		"878728347234",
		"*@#&*&*&@*#",
		"235235*@#&*&*&@*#",
		"*@#&*&23355*&@*#35235",
		"872348jasdhjasdhj",
		"word767",
		"",
	}

	translator := New()
	for _, text := range untranslatable {
		actual, err := translator.Translate(text)
		if _, ok := err.(FailedToTranslateError); !ok {
			t.Error(fmt.Sprintf("Expected error when translating '%s'."+
				" Expected: FailedToTranslateError | Got: '%s'", text, actual))
		}
	}
}
