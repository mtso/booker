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

	body, err := utils.ParseRequestBody(r)
	if WriteErrorResponse(w, err) {
		return
	}
	city, cityOk := body["city"]
	state, stateOk := body["state"]
	displayName, displayNameOk := body["display_name"]
	newPass, newPassOk := body["password"]

	user, err := models.Users.Find(*username)
	if err != nil {
		WriteError(w, err)
		return
	}

	if cityOk && stateOk && displayNameOk {
		err = user.SetLocation(city.(string), state.(string), displayName.(string))
		if err != nil {
			WriteError(w, err)
			return
		}

		isLocationUpdated = true
	}

	if newPassOk {
		err = user.SavePasswordHash([]byte(newPass.(string)))
		if err != nil {
			WriteError(w, err)
			return
		}

		isPasswordUpdated = true
	}

	resp := &JsonResponse{
		"ok": isLocationUpdated || isPasswordUpdated,
	}

	if isPasswordUpdated {
		resp.Set("isPasswordUpdated", isPasswordUpdated)
	}
	if isLocationUpdated {
		resp.Set("isLocationUpdated", isLocationUpdated)
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
	resp["display_name"] = user.DisplayName

	city, _ := user.City.Value()
	state, _ := user.State.Value()
	resp["city"] = city
	resp["state"] = state

	WriteJson(w, resp)
}
