package gopher

import (
	"encoding/json"
	"errors"
	"github.com/ilievss/lango/translate/lang/en/gopher"
	"io/ioutil"
	"net/http"
)

var gopherTranslator = gopher.New()

type TranslateWordRequest struct {
	EnglishWord string `json:"english-word"`
}

type TranslateWordResponse struct {
	GopherWord string `json:"gopher-word"`
}

func JsonHandler(w http.ResponseWriter,
	r *http.Request,
	expectedBody interface{},
	responseBodySupplier func(*interface{}) (interface{}, int, error)) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Could not request body"))
		return
	}
	err = json.Unmarshal(body, expectedBody)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Could not unmarshal JSON: " + string(body)))
		return
	}

	responseBody, statusCode, err := responseBodySupplier(&expectedBody)
	if err != nil {
		w.WriteHeader(statusCode)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

func TranslateEnglishWordHandler(w http.ResponseWriter, r *http.Request) {
	request := TranslateWordRequest{}
	JsonHandler(w, r, &request, func(req *interface{}) (interface{}, int, error) {
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
