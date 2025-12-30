package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	// POST handlers
	mux.HandleFunc("POST /newUser", s.newUserHandler)
	mux.HandleFunc("POST /postMedia", s.UploadMediaFile)

	// GET handlers
	mux.HandleFunc("GET /getUser", s.getUserHandler)

	// Health
	mux.HandleFunc("GET /health", s.healthHandler)

	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

/////////// Post Methods ///////////

func (s *Server) newUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) UploadMediaFile(w http.ResponseWriter, r *http.Request) {

}

/////////// GET Methods ///////////

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getUserData(w http.ResponseWriter, r *http.Request) {

}
