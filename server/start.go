package main

import (
	"log"
	"net/http"

	"github.com/mtso/booker/server/controllers"
	"github.com/mtso/booker/server/models"
)

func main() {
	db, err := models.InitializeDb("host=localhost user=wiggs dbname=booker sslmode=disable password=cupcakes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/", controllers.Root)
	http.ListenAndServe(":3750", nil)
}
