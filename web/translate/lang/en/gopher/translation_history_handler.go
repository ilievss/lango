package gopher

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

func RetrieveTranslationHistoryHandler(w http.ResponseWriter, r *http.Request) {
	entries := translationHistory.Entries()
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Input) < strings.ToLower(entries[j].Input)
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
