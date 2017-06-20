package main

import (
	// "database/sql"
	"log"
	"net/http"
	"os"

	"github.com/mtso/booker/server/controllers"
	"github.com/mtso/booker/server/models"
)

func main() {
	db, err := models.Connect(os.Getenv("DATABASE_URL"))
	// db, err := models.InitializeDb("host=localhost user=wiggs dbname=booker sslmode=disable password=cupcakes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/", controllers.Root)
	http.ListenAndServe(":3750", nil)
}

// type App struct {
// 	Db      *sql.DB
// 	Handler http.Handler
// }

// func InitializeApp() *App {
// 	db, err := models.Connect(os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &App{
// 		db,
// 		controllers.Root,
// 	}
// }
