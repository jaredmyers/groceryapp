package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaredmyers/groceryapp/user_service/go_version/models"
	"github.com/jaredmyers/groceryapp/user_service/go_version/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/selections", makeHandleFunc(s.handleSelections))

	log.Println("API server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleSelections(w http.ResponseWriter, r *http.Request) error {

	newSelections := &models.Selections{}
	if err := json.NewDecoder(r.Body).Decode(newSelections); err != nil {
		log.Println(err)
		return err
	}

	if err := s.store.AddNewSelections(newSelections); err != nil {
		return err
	}

	return WriteJSON(w, r, http.StatusOK, newSelections)

}

// helper functions

// write to json in one place
func WriteJSON(w http.ResponseWriter, r *http.Request, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Println(r.URL, r.Method, status)
	return json.NewEncoder(w).Encode(v)
}

// decorator/adapter for handling errors in one spot
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err)
			WriteJSON(w, r, http.StatusBadRequest, ApiError{err.Error()})
		}

	}
}
