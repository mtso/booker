package main

import (
	"net/http"

	"github.com/mtso/booker/server/config"
)

func main() {
	app := config.InitializeApp()
	defer app.Db.Close()

	http.ListenAndServe(":3750", app.Handler)
}
