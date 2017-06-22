package controllers

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var indexTemplBytes, _ = ioutil.ReadFile("./server/views/index.template.html")
var indexTempl = template.Must(template.New("").Parse(string(indexTemplBytes)))

func TEMPGetApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct {
		State template.JS
	}{
		template.JS("{}"),
	}

	err := indexTempl.Execute(w, data)
	if err != nil {
		log.Println(err)
		WriteErrorResponse(w, err)
	}
}

var ServeStatic = http.StripPrefix("/", http.FileServer(http.Dir("./dist/")))
