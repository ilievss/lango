package gopher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var translatedWordUrl = "http://localhost:8080/word"

func TestTranslateEnglishWordBadJson(t *testing.T) {
	invalidJsonBody := []byte("{\"english\": \"bad json\"")
	req := httptest.NewRequest("POST", translatedWordUrl, bytes.NewBuffer(invalidJsonBody))
	w := httptest.NewRecorder()
	TranslateEnglishWordHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != 400 {
		t.Fatal(fmt.Sprintf("Expected status code 400, but got %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to read response body: %s", err))
	}

	errorMessage := string(body)
	if errorMessage != "Could not unmarshal JSON: {\"english\": \"bad json\"" {
		t.Fatal(fmt.Sprintf("Wrong error message in response: %s", errorMessage))
	}
}

func TestTranslateEnglishWordMissingJsonField(t *testing.T) {
	invalidJsonBody := []byte("{\"english\": \"invalid\"}")
	req := httptest.NewRequest("POST", translatedWordUrl, bytes.NewBuffer(invalidJsonBody))
	w := httptest.NewRecorder()
	TranslateEnglishWordHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != 400 {
		t.Fatal(fmt.Sprintf("Expected status code 400, but got %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to read response body: %s", err))
	}

	errorMessage := string(body)
	if errorMessage != "word is untranslatable" {
		t.Fatal(fmt.Sprintf("Wrong error message in response: %s", errorMessage))
	}
}

func TestTranslateEnglishWordValidWord(t *testing.T) {
	requestWord := "Spain"
	expected := "ainSpogo"
	request := TranslateWordRequest{requestWord}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		t.Fatal("Failed to marshal request")
	}

	req := httptest.NewRequest("POST", translatedWordUrl, bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	TranslateEnglishWordHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatal(fmt.Sprintf("Expected status code 200, but got %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to read response body: %s", err))
	}
	var responseBody TranslateWordResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to unmarshal response body: %s", err))
	}

	if responseBody.GopherWord != expected {
		t.Fatal(fmt.Sprintf("Expected : '%s', but got '%s'", expected, responseBody.GopherWord))
	}
}

func TestTranslateEnglishWordUntranslatableWord(t *testing.T) {
	requestWord := "don't"
	request := TranslateWordRequest{requestWord}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		t.Fatal("Failed to marshal request")
	}

	req := httptest.NewRequest("POST", translatedWordUrl, bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	TranslateEnglishWordHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != 400 {
		t.Fatal(fmt.Sprintf("Expected status code 500, but got %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(fmt.Sprintf("Failed to read response body: %s", err))
	}

	errorMessage := string(body)
	if errorMessage != "word is untranslatable" {
		t.Fatal(fmt.Sprintf("Wrong error message: %s", errorMessage))
	}
}
