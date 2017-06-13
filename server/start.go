package main

import (
	"log"
	"net/http"
	// "os"

	"github.com/mtso/booker/server/controllers"
	"github.com/mtso/booker/server/models"
)

func main() {
	db, err := models.InitializeDb("host=localhost user=wiggs dbname=booker sslmode=disable password=cupcakes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// if err := models.InitializeDb(os.Getenv("DATABASE_URL")); err != nil {
	// 	log.Fatal(err)
	// }

	http.Handle("/", controllers.Root)
	http.ListenAndServe(":3750", nil)
}

