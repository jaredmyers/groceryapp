package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaredmyers/groceryapp/backend/models"
	"github.com/jaredmyers/groceryapp/backend/storage"
	"github.com/rs/cors"
)

type Server struct {
	listenAddr string
	store      storage.Storage
	services   map[string]string
}

func NewServer(listenAddr string, store storage.Storage, services map[string]string) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		services:   services,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/products", makeHandleFunc(s.handleGroceryPage))
	router.HandleFunc("/selections", makeHandleFunc(s.handleGroceryList))
	router.HandleFunc("/login", makeHandleFunc(s.handleLogin))

	corsRouter := cors.Default().Handler(router)
	log.Println("API Server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, corsRouter)
}

// handlers

func (s *Server) handleGroceryPage(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGroceryPageProducts(w, r)
	default:
		/*
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed"))
		*/
	}
	return WriteJSON(w, r, http.StatusMethodNotAllowed, ApiError{Error: fmt.Errorf("not allowed").Error()})
}

func (s *Server) handleGroceryList(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleGroceryListSelections(w, r)
	default:
		/*
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed"))
		*/
	}
	return WriteJSON(w, r, http.StatusMethodNotAllowed, ApiError{Error: fmt.Errorf("not allowed").Error()})
}

func (s *Server) handleGroceryPageProducts(w http.ResponseWriter, r *http.Request) error {

	products, _ := s.store.GetGroceryPageProducts()
	return WriteJSON(w, r, http.StatusOK, products)
}

func (s *Server) handleGroceryListSelections(w http.ResponseWriter, r *http.Request) error {

	selections := s.store.GroceryListSelections()
	if err := json.NewDecoder(r.Body).Decode(selections); err != nil {
		return err
	}

	// just checking
	log.Println(selections)

	// conform selections to JSON (bytes),
	// for sending to user_service
	var jsonBuf bytes.Buffer
	json.NewEncoder(&jsonBuf).Encode(selections)

	// send to kafka here,
	// doing REST for now
	if err := sendToService(s.services["userService"]+"/selections", &jsonBuf); err != nil {
		log.Println(err)
		return err
	}

	return WriteJSON(w, r, http.StatusOK, selections)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "POST" {
		return WriteJSON(w, r, http.StatusMethodNotAllowed, ApiError{Error: fmt.Errorf("not allowed").Error()})
	}

	login := &models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(login); err != nil {
		return err
	}

	// conform creds to JSON (bytes)
	// for sending to user_service
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(login)

	// send to kafka here,
	// doing REST for now
	if err := sendToService(s.services["userService"]+"/login", &buf); err != nil {
		return err
	}

	return nil
}

// helpers

func sendToService(service string, jsonBuf *bytes.Buffer) error {
	_, err := http.Post(service, "application/json", jsonBuf)
	if err != nil {
		return err
	}
	return nil
}

// one spot for writing out json
func WriteJSON(w http.ResponseWriter, r *http.Request, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Println(r.URL, r.Method, status)
	return json.NewEncoder(w).Encode(v)
}

// decorator/adapter so errors can be handled in one spot
type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, r, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}
