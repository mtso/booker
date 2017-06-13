package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitializeDb(conn string) (*gorm.DB, error) {
	newdb, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db = newdb

	if err := CreateUserSchema(db); err != nil {
		return nil, err
	}
	return newdb, nil
}
