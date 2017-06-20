package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleUser(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/user", IsLoggedInMiddleware(PostUser)).Methods("POST")
	api.HandleFunc("/user", GetUser).Methods("POST")
}

// {"city":"[city]","state":"[state]","password":"[new password]"}
func PostUser(w http.ResponseWriter, r *http.Request) {

}

// {"username":"[username]"}
// respond with username, city, state
func GetUser(w http.ResponseWriter, r *http.Request) {

}
