package web

import (
	"github.com/gorilla/mux"
	"github.com/ilievss/lango/web/translate/lang/en/gopher"
	"log"
	"net/http"
)

type LangoServer struct {
	server http.Server
}

func New(address string) LangoServer {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/word", gopher.TranslateEnglishWordHandler).Methods("POST")
	r.HandleFunc("/sentence", gopher.TranslateEnglishSentenceHandler).Methods("POST")
	r.HandleFunc("/history", gopher.RetrieveTranslationHistoryHandler).Methods("GET")

	return LangoServer{server: http.Server{Addr: address, Handler: r}}
}

func (s *LangoServer) Start() {
	log.Fatal(s.server.ListenAndServe())
}
