package gopher

import (
	"errors"
	"github.com/ilievss/lango/translate/lang/en/gopher"
	"github.com/ilievss/lango/web/helper"
	"net/http"
)

type TranslateWordRequest struct {
	EnglishWord string `json:"english-word"`
}

type TranslateWordResponse struct {
	GopherWord string `json:"gopher-word"`
}

func TranslateEnglishWordHandler(w http.ResponseWriter, r *http.Request) {
	request := TranslateWordRequest{}
	helper.JsonHandler(w, r, &request, func(req *interface{}) (interface{}, int, error) {
		translation, err := gopherTranslator.Translate(request.EnglishWord)
		if err != nil {
			if _, ok := err.(gopher.FailedToTranslateError); ok {
				return nil, 400, errors.New("word is untranslatable")
			}
			return nil, 500, err
		}
		return TranslateWordResponse{translation}, 200, nil
	})
}
