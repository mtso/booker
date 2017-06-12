package controllers

import (
	"net/http"
	"log"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
)

func handleAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()

	getSignup := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~ signup here"))
	}
	s.HandleFunc("/signup", getSignup).Methods("GET")

	s.HandleFunc("/signup", PostSignup).Methods("POST")
}

// query := r.URL.Query()
// fmt.Printf("%v", query["username"])

// cost, err := bcrypt.Cost(hash)
// err = bcrypt.CompareHashAndPassword(hash, []byte(pass))
// bcrypt.ErrMismatchedHashAndPassword
func PostSignup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u interface{}

	err := decoder.Decode(&u)
	if err != nil {
		log.Println("PostSignup Error:", err)
		return
	}

	ut := u.(map[string]interface{})
	user := ut["username"].(string)
	pass := ut["password"].(string)

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), -1)
	if err != nil {
		log.Println("bcrypt error:", err)
		return
	}

	nu := models.User{
		Username: user,
		PasswordHash: string(hash),
	}
	msg := "created user: " + user

	err = nu.Create()
	if err != nil {
		msg = err.Error()
		log.Println("CreateUser error", err)
	}

	success := err == nil
	response := struct{
		Success bool `json:"success"`
		Message string `json:"message"`
	}{
		success,
		msg,
	}

	js, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
