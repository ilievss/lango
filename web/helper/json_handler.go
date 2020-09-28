package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func JsonHandler(w http.ResponseWriter,
	r *http.Request,
	expectedBody interface{},
	responseBodySupplier func(*interface{}) (interface{}, int, error)) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Could not read request body"))
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
