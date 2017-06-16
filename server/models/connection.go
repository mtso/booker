package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(conn string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", conn)
	if err != nil {
		return
	}

	connects := []func(*sql.DB) error{
		ConnectUsers,
	}

	for _, connect := range connects {
		err = connect(db)
		if err != nil {
			break
		}
	}
	return
}
