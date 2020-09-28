package gopher

import (
	"errors"
	"github.com/ilievss/lango/translate/lang/en/gopher"
	"github.com/ilievss/lango/web/helper"
	"github.com/ilievss/lango/web/translate"
	"net/http"
)

type TranslateSentenceRequest struct {
	EnglishSentence string `json:"english-sentence"`
}

type TranslateSentenceResponse struct {
	GopherSentence string `json:"gopher-sentence"`
}

func TranslateEnglishSentenceHandler(w http.ResponseWriter, r *http.Request) {
	request := TranslateSentenceRequest{}
	helper.JsonHandler(w, r, &request, func(req *interface{}) (interface{}, int, error) {
		translation, err := gopherTranslator.Translate(request.EnglishSentence)
		if err != nil {
			if _, ok := err.(gopher.FailedToTranslateError); ok {
				return nil, 400, errors.New("sentence is untranslatable")
			}
			return nil, 500, err
		}
		translationHistory.Add(translate.HistoryEntry{Input: request.EnglishSentence, Translation: translation})
		return TranslateSentenceResponse{translation}, 200, nil
	})
}
