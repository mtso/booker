package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
)

var Users = models.Users

func handleAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()

	getSignup := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~ signup here"))
	}
	s.HandleFunc("/signup", getSignup).Methods("GET")

	s.HandleFunc("/signup", PostSignup).Methods("POST")
	s.HandleFunc("/login", PostLogin).Methods("POST")
}

// query := r.URL.Query()
// fmt.Printf("%v", query["username"])
func PostSignup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var raw interface{}

	err := decoder.Decode(&raw)
	if err != nil {
		log.Println("PostSignup Error:", err)
		return
	}

	body := raw.(map[string]interface{})
	user := body["username"].(string)
	pass := body["password"].(string)

	newUser := models.User{
		Username: user,
	}
	newUser.StoreHash([]byte(pass))

	var msg string

	if err := newUser.Create(); err != nil {
		msg = err.Error()
		log.Println("CreateUser error", err)
	} else {
		msg = "created user: " + user
	}

	success := err == nil

	response := struct {
		Success bool   `json:"success"`
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

func PostLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var raw interface{}

	err := decoder.Decode(&raw)
	if err != nil {
		log.Println("PostSignup Error:", err)
		return
	}

	body := raw.(map[string]interface{})
	user := body["username"].(string)
	pass := body["password"].(string)

	err = Users.Verify(user, []byte(pass))
	success := err == nil
	var msg string

	if err != nil {
		log.Println(err)
		msg = err.Error()
	}

	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message,omitempty"`
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

func DecodeBody(r *http.Request) (m map[string]interface{}) {
	decoder := json.NewDecoder(r.Body)
	var raw interface{}

	err := decoder.Decode(&raw)
	if err != nil {
		log.Println(err)
		return
	}

	m = raw.(map[string]interface{})
	return
}
