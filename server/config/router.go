package config

import (
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/mtso/booker/server/controllers"
)

var catchall = func(*http.Request, *mux.RouteMatch) bool {
	return true
}

var exactroot = func(r *http.Request, _ *mux.RouteMatch) bool {
	// Match exact path
	return r.URL.Path == "/"
}

func makeRootHandler() *mux.Router {
	router := mux.NewRouter()

	// Authentication endpoint
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", PostSignup).Methods("POST")
	auth.HandleFunc("/login", PostLogin).Methods("POST")
	auth.HandleFunc("/logout", PostLogout).Methods("POST")

	// API endpoint
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/user", IsLoggedInMiddleware(PostUser)).Methods("POST")
	api.HandleFunc("/user", GetUser).Methods("POST")

	// API book* endpoints
	api.HandleFunc("/book/{id:[0-9]+}", GetBook).Methods("GET")
	api.HandleFunc("/book", IsLoggedInMiddleware(PostBook)).Methods("POST")
	// api.HandleFunc("/book", IsLoggedInMiddleware(DeleteBook)).Methods("DELETE")
	books := api.PathPrefix("/books").Subrouter()
	books.HandleFunc("/mybooks", IsLoggedInMiddleware(GetMyBooks)).Methods("GET")

	// /books endpoint root catchall
	books.Methods("GET").MatcherFunc(catchall).HandlerFunc(GetBooks)

	user := api.PathPrefix("/user").Subrouter()
	user.Path("/{username:[a-z0-9A-Z]*}").HandlerFunc(GetUser)

	api.HandleFunc("/trades/incoming", IsLoggedInMiddleware(GetIncomingTrades)).Methods("GET")
	api.HandleFunc("/trades/outgoing", IsLoggedInMiddleware(GetOutgoingTrades)).Methods("GET")
	api.HandleFunc("/trade", IsLoggedInMiddleware(PostTrade)).Methods("POST")
	api.HandleFunc("/trade/{id:[0-9]+}", IsLoggedInMiddleware(PutTrade)).Methods("PUT")
	// api.HandleFunc("/trade/{id:[0-9]+}", IsLoggedInMiddleware(DeleteTrade)).Methods("DELETE")

	// Serve app
	router.PathPrefix("/").MatcherFunc(exactroot).HandlerFunc(TEMPGetApp)
	router.PathPrefix("/").Handler(ServeStatic)

	return router
}
