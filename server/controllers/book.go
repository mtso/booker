package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
)

func handleBooks(r *mux.Router) {
	sub := r.PathPrefix("/api/books").Subrouter()
	sub.HandleFunc("", GetBooks).Methods("GET")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bks, err := models.Books.GetBooks()
	if WriteErrorResponse(w, err) {
		return
	}

	resp := &JsonResponse{
		Ok:   true,
		Data: bks,
	}

	WriteJson(w, resp)
}
