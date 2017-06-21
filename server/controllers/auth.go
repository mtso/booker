package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"github.com/mtso/booker/server/models"
)

const SessionId = "sess_id"

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

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
	WriteJson(w, &JsonResponse{
		Ok:      true,
		Message: u + " is logged into redirecting endpoint",
	})
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
	msg := user + " logged in."

	if err != nil {
		log.Println(err)
		msg = err.Error()
	}

	response := &JsonResponse{
		Ok:      success,
		Message: msg,
	}

	// Save session.
	session, err := store.Get(r, SessionId)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	session.Values["username"] = user
	if err := session.Save(r, w); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteJson(w, response)
}

func WriteError(w http.ResponseWriter, err error, code ...int) {
	log.Println(code, err)
	WriteErrorResponse(w, err, code...)
}

func WriteErrorResponse(w http.ResponseWriter, err error, args ...int) bool {
	if err == nil {
		return false
	}
	code := http.StatusInternalServerError
	if len(args) > 0 {
		code = args[0]
	}
	errorResponse := &JsonResponse{
		Ok:      false,
		Message: err.Error(),
	}
	WriteJson(w, errorResponse, code)
	return true
}

// func PostPassword(w http.ResponseWriter, r *http.Request) {
// 	body, err := ParseBody(r)
// 	if err != nil {
// 		WriteError(w, err)
// 		return
// 	}
// 	_, _ := body["username"]
// 	_, _ := body["previous"]
// 	_, _ := body["new"]
// 	// user, err := models.Users.FindAndVerify(username)
// 	// user.SetPasswordHash([]byte(previous))
// }

func PostLogout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SessionId)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	username, err := GetUsername(r)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	delete(session.Values, "username")

	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &JsonResponse{
		Ok:      true,
		Message: username + " logged out.",
	}
	WriteJson(w, resp)
}

func WriteJson(w http.ResponseWriter, response interface{}, code ...int) {
	js, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var ErrNoUsername = errors.New("No username found for session")

func GetUsername(r *http.Request) (string, error) {
	session, err := store.Get(r, SessionId)
	if err != nil {
		return "", err
	}
	if found, ok := session.Values["username"]; ok && found != nil {
		return found.(string), nil
	} else {
		return "", ErrNoUsername
	}
}

func IsLoggedIn(r *http.Request) (bool, error) {
	username, err := GetUsername(r)
	return username != "", err
}

func IsLoggedInMiddleware(next http.HandlerFunc, args ...string) http.HandlerFunc {
	redirectUrl := "/"
	if len(args) > 0 {
		redirectUrl = args[0]
	}

	return func(w http.ResponseWriter, r *http.Request) {
		isLoggedIn, _ := IsLoggedIn(r)

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
