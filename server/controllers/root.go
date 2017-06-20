package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Root = makeRootHandler()

func makeRootHandler() *mux.Router {
	root := mux.NewRouter()
	root.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("hello~"))

		resp := &JsonResponse{
			Ok:      true,
			Message: "test route",
		}
		WriteJson(w, resp)

	})
	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/"))
	})

	handleAuth(root)
	// root.Path("/auth").HandlerFunc()

	handleBooks(root)

	return root
}
