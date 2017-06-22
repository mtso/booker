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

func TEMPGetApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// need this to be nullable
	u, _ := GetUsername(r)

	state := struct {
		Username *string `json:"username"`
	}{
		Username: u,
	}

	js, err := json.Marshal(state)
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

var ServeStatic = http.StripPrefix("/", http.FileServer(http.Dir("./dist/")))
