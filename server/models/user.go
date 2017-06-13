// package models

// import (
// 	// "crypto/tls"

// 	"github.com/go-pg/pg"
// 	"github.com/go-pg/pg/orm"
// )

// var db *pg.DB

// func InitializeDb(conn string) (err error) {
// 	opts, err := pg.ParseURL(conn)
// 	if err != nil {
// 		return err
// 	}
// 	opts.TLSConfig = nil
// 	// opts.TLSConfig = &tls.Config{
// 	// 	InsecureSkipVerify: true,
// 	// }
// 	db = pg.Connect(opts)
// 	err = CreateUserSchema(db)
// 	return
// }

// type User struct {
// 	TableName struct{} `sql:"User"`
// 	Id int64// `sql:",pk"`
// 	Username string //`sql:",unique,notnull"`
// 	PasswordHash string //`sql:",notnull"`
// 	City string
// 	State string
// }

// func CreateUserSchema(db *pg.DB) (err error) {
// 	err = db.CreateTable(&User{}, &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	})
// 	return
// }

// func (u *User) Create() (err error) {
// 	err = db.Insert(u)
// 	return
// }

package models

import (
	// "crypto/tls"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/go-pg/pg/orm"
)

var db *gorm.DB

func InitializeDb(conn string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	// err = CreateUserSchema(db).Error
	if err := CreateUserSchema(db); err != nil {
		return nil, err
	}
	return db, nil

	// opts, err := pg.ParseURL(conn)
	// if err != nil {
	// 	return err
	// }
	// opts.TLSConfig = nil
	// // opts.TLSConfig = &tls.Config{
	// // 	InsecureSkipVerify: true,
	// // }
	// db = pg.Connect(opts)
	// return nil
}

type User struct {
	gorm.Model
	ID uint64 `gorm:"primary_key"`
	Username string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`
	City string
	State string
}

func CreateUserSchema(db *gorm.DB) (err error) {
	// if err = db.DropTable("users").Error; err != nil {
	// 	return
	// }
	return db.Error
	// return db.Set("gorm:table_options", "IF NOT EXISTS").CreateTable(&User{}).Error
	// return db.CreateTable(&User{}).Error
	// return db.AutoMigrate(&User{}).Error
	// err = db.CreateTable(&Users{}, &orm.CreateTableOptions{})
	// return
}

func (u *User) Create() (err error) {
	err = db.Create(u).Error
	return
}
