package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Root = makeRootHandler()

func makeRootHandler() *mux.Router {
	root := mux.NewRouter()
	root.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello~"))
	})
	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/"))
	})

	handleAuth(root)
	// root.Path("/auth").HandlerFunc()

	return root
}
