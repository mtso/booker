package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/mtso/booker/server/models"
)

const SessionId = "sess_id"

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

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
	s.HandleFunc("/test", TestLogin).Methods("GET")
	s.HandleFunc("/testroute", IsLoggedInMiddleware(TestEndpoint)).Methods("GET")
}

func TestLogin(w http.ResponseWriter, r *http.Request) {
	// test that we save session ID properly
	ok, err := IsLoggedIn(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if ok {
		u, _ := GetUsername(r)
		w.Write([]byte(u + " is logged in"))
	} else {
		w.Write([]byte("not logged in"))
	}
}

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	u, _ := GetUsername(r)
	w.Write([]byte(u + " is logged into redirecting endpoint"))
}

// query := r.URL.Query()
// fmt.Printf("%v", query["username"])
func PostSignup(w http.ResponseWriter, r *http.Request) {
	body := ParseBody(r)
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
	body := ParseBody(r)
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

	// Save session.
	session, err := store.Get(r, SessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["username"] = user
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetUsername(r *http.Request) (string, error) {
	session, err := store.Get(r, SessionId)
	if err != nil {
		return "", err
	}
	return session.Values["username"].(string), nil
}

func IsLoggedIn(r *http.Request) (bool, error) {
	session, err := store.Get(r, SessionId)
	if err != nil {
		return false, err
	}
	return session.Values["username"] != nil, nil
}

func IsLoggedInMiddleware(next http.HandlerFunc, args ...string) http.HandlerFunc {
	redirectUrl := "/"
	if len(args) > 0 {
		redirectUrl = args[0]
	}

	return func(w http.ResponseWriter, r *http.Request) {
		isLoggedIn, err := IsLoggedIn(r)
		if err != nil {
			http.Error(w, redirectUrl, http.StatusFound)
			return
		}

		if isLoggedIn {
			next(w, r)
		} else {
			http.Redirect(w, r, redirectUrl, http.StatusFound)
		}
	}
}

// BodyParser?
func ParseBody(r *http.Request) (m map[string]interface{}) {
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
