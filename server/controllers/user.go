package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
	// "github.com/mtso/booker/server/utils"
)

// {"city":"[city]","state":"[state]","password":"[new password]"}
func PostUser(w http.ResponseWriter, r *http.Request) {
	// body := utils.ParseRequestBody(r)
	// city, cityOk := body["city"]
	// state, stateOk := body["state"]

	// if cityOk && stateOk {
	// 	models.Users.Find()
	// }
}

// {"username":"[username]"}
// respond with username, city, state
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		WriteError(w, ErrNoUsername)
		return
	}

	user, err := models.Users.Find(username)
	if err != nil {
		WriteError(w, err)
		return
	}

	resp := make(JsonResponse)
	resp["ok"] = true
	resp["username"] = user.Username
	resp["city"] = user.City.String
	resp["state"] = user.State.String

	WriteJson(w, resp)
}
