package main

import (
	"net/http"

	"github.com/mtso/booker/server/controllers"
)

func main() {
	http.Handle("/", controllers.Root)
	http.ListenAndServe(":3750", nil)
}

