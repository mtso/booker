package config

import (
	"database/sql"
	"os"
	"net/http"
	"log"

	"github.com/mtso/booker/server/controllers"
	"github.com/mtso/booker/server/models"
)

type App struct {
	Db      *sql.DB
	Handler http.Handler
}

func InitializeApp() *App {
	db, err := models.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return &App{
		db,
		controllers.Root,
	}
}
