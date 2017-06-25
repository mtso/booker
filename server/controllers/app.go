package controllers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mtso/booker/server/models"
)

var indexTemplBytes, _ = ioutil.ReadFile("./server/views/index.template.html")
var indexTempl = template.Must(template.New("").Parse(string(indexTemplBytes)))

var ServeStatic = http.StripPrefix("/static/", http.FileServer(http.Dir("./dist/")))

type UserState struct {
	Username *string `json:"username"`
	DisplayName string `json:"display_name"`
	City string `json:"city"`
	State string `json:"state"`
}

func preloadState(w http.ResponseWriter, r *http.Request) *interface{} {
	var userState UserState

	username, err := GetUsername(r)
	if err == nil {
		user, err := models.Users.Find(*username)
		if WriteErrorResponse(w, err) {
			return nil
		}

		city, _ := user.City.Value()
		state, _ := user.State.Value()

		userState.Username = username
		userState.DisplayName = user.DisplayName
		userState.City = city.(string)
		userState.State = state.(string)
	}

	state := struct {
		User UserState `json:"user"`
	}{
		User: userState,
	}

	var i interface{}
	i = state
	return &i
}

func ServeApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	st := preloadState(w, r)

	js, err := json.Marshal(st)
	if WriteErrorResponse(w, err) {
		return
	}

	data := struct {
		State template.JS
	}{
		template.JS(string(js)),
	}

	err = indexTempl.Execute(w, data)
	if err != nil {
		log.Println(err)
		WriteErrorResponse(w, err)
	}
}
