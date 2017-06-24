package controllers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var indexTemplBytes, _ = ioutil.ReadFile("./server/views/index.template.html")
var indexTempl = template.Must(template.New("").Parse(string(indexTemplBytes)))

var ServeStatic = http.StripPrefix("/static/", http.FileServer(http.Dir("./dist/")))

func preloadState(r *http.Request) *interface{} {
	u, _ := GetUsername(r)

	state := struct {
		Username *string `json:"username"`
	}{
		Username: u,
	}

	var i interface{}
	i = state
	return &i
}

func TEMPGetApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	st := preloadState(r)

	js, err := json.Marshal(st)
	if WriteErrorResponse(w, err) {
		return
	}

	data := struct {
		State template.JS
	}{
		template.JS(string(js)),
		// template.JS("{}"),
	}

	err = indexTempl.Execute(w, data)
	if err != nil {
		log.Println(err)
		WriteErrorResponse(w, err)
	}
}
