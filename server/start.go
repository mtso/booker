package main

import (
	"net/http"
	"os"

	"github.com/mtso/booker/server/config"
)

func main() {
	app := config.InitializeApp()
	defer app.Db.Close()

	http.ListenAndServe(":"+os.Getenv("PORT"), app.Handler)
}
