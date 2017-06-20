package controllers

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
)

func handleBooks(r *mux.Router) {
	sub := r.PathPrefix("/api/books").Subrouter()

	sub.HandleFunc("/what", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello~"))
	}).Methods("GET")

	sub.Methods("GET").MatcherFunc(func(*http.Request, *mux.RouteMatch) bool {
		return true
	}).HandlerFunc(GetBooks)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bks, err := models.Books.GetBooks()
	if WriteErrorResponse(w, err) {
		return
	}

	for _, v := range bks {
		v.ImageUrl = url.QueryEscape(v.ImageUrl)
	}

	resp := &JsonResponse{
		Ok:   true,
		Data: bks,
	}

	WriteJson(w, resp)
}
