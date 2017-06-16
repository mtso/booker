package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
)

// type Flash struct {
// 	Code    string `json:"code"`
// 	Message string `json:"message"`
// }

// type ApiResponse struct {
// 	Ok    bool        `json:"ok"`
// 	Data  interface{} `json:"data"`
// 	Flash Flash       `json:"flash"`
// }

func handleAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()

	getSignup := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~ signup here"))
	}

	// s.Path("/signup").Methods("POST").HandlerFunc(getSignup)

	s.HandleFunc("/signup", getSignup).Methods("GET")

	s.HandleFunc("/signup", PostSignup).Methods("POST")
	s.HandleFunc("/login", PostLogin).Methods("POST")
}

// query := r.URL.Query()
// fmt.Printf("%v", query["username"])
func PostSignup(w http.ResponseWriter, r *http.Request) {
	body := DecodeBody(r)
	user := body["username"].(string)
	pass := body["password"].(string)

	newUser := models.User{
		Username: user,
	}
	newUser.SetPasswordHash([]byte(pass))

	var msg string

	err := newUser.Create()
	if err != nil {
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
	body := DecodeBody(r)
	user := body["username"].(string)
	pass := body["password"].(string)

	err := models.Users.Verify(user, []byte(pass))
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

// BodyParser?
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
