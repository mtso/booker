package controllers

import (
	"net/http"
	"log"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
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

	cost, err := bcrypt.Cost(hash)
	if err != nil {
		log.Println(err)
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(pass))
	if err != bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
	} else if err != nil {
		log.Println(err)
	}

	// store hash
	log.Printf("%d\n%s\n", cost, hash)
	log.Println(user, pass)
}
