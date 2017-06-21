package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
	"github.com/mtso/booker/server/utils"
)

// {"city":"[city]","state":"[state]","password":"[new password]"}
func PostUser(w http.ResponseWriter, r *http.Request) {
	isLocationUpdated := false
	isPasswordUpdated := false

	username, err := GetUsername(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	body := utils.ParseRequestBody(r)
	city, cityOk := body["city"]
	state, stateOk := body["state"]

	if cityOk && stateOk {
		user, err := models.Users.Find(username)
		if err != nil {
			WriteError(w, err)
			return
		}

		err = user.SetLocation(city.(string), state.(string))
		if err != nil {
			WriteError(w, err)
			return
		}

		isLocationUpdated = true
	}

	resp := make(JsonResponse)
	resp["ok"] = isLocationUpdated || isPasswordUpdated
	if isPasswordUpdated {
		resp["isPasswordUpdated"] = isPasswordUpdated
	}
	if isLocationUpdated {
		resp["isLocationUpdated"] = isLocationUpdated
	}

	WriteJson(w, resp)
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

	// if city, err := user.City.Value(); err != nil {
	// 	resp["city"] = city
	// }
	// if state, err := user.State.Value(); err != nil {
	// 	resp["state"] = state
	// }
	city, _ := user.City.Value()
	state, _ := user.State.Value()
	resp["city"] = city
	resp["state"] = state

	WriteJson(w, resp)
}
