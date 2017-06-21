package config

import (
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/mtso/booker/server/controllers"
)

var catchall = func(*http.Request, *mux.RouteMatch) bool {
	return true
}

func makeRootHandler() *mux.Router {
	router := mux.NewRouter()

	// Authentication endpoint
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", PostSignup).Methods("POST")
	auth.HandleFunc("/login", PostLogin).Methods("POST")
	auth.HandleFunc("/logout", PostLogout).Methods("POST")
	auth.HandleFunc("/test", TestLogin).Methods("GET")
	auth.HandleFunc("/testroute", IsLoggedInMiddleware(TestEndpoint)).Methods("GET")

	// API endpoint
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user", IsLoggedInMiddleware(PostUser)).Methods("POST")
	api.HandleFunc("/user", GetUser).Methods("POST")

	books := api.PathPrefix("/books").Subrouter()

	books.HandleFunc("/what", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello~"))
	}).Methods("GET")

	// /books endpoint root catchall
	books.Methods("GET").MatcherFunc(catchall).HandlerFunc(GetBooks)

	user := api.PathPrefix("/user").Subrouter()
	user.Path("/{username:[a-z0-9A-Z]*}").HandlerFunc(GetUser)
	// user.PathPrefix("").HandlerFunc(GetUser)
	// {username:asdf}
	// .Methods("GET")
	return router
}
