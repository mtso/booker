package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
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

// var db *gorm.DB

// func InitializeDb(conn string) (*gorm.DB, error) {
// 	newdb, err := gorm.Open("postgres", conn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	db = newdb

// 	if err := CreateUserSchema(db); err != nil {
// 		return nil, err
// 	}
// 	return newdb, nil
// }
