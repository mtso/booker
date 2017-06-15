package models

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	CreateTableUser = `CREATE TABLE IF NOT EXISTS Users (
		id bigserial NOT NULL UNIQUE,
		username varchar(64) NOT NULL UNIQUE,
		password_hash varchar(64) NOT NULL,
		city varchar(128),
		state varchar(64),
		created_at timestamp NOT NULL DEFAULT NOW()
	)`
	SelectUserByName = `SELECT * FROM Users
		WHERE username = $1
		LIMIT 1`
	InsertUser = `INSERT Users (username, passwordhash) VALUES ($1, $2)`
	UpdateUser = `UPDATE Users
		SET city = $2,
		    state = $3
		WHERE username = $1`
)

var Users UsersSchema

var ErrNotFoundUser = errors.New("User not found.")

type UsersSchema struct {
	db *sql.DB
}

type User struct {
	Id           int64     `sql:"id"`
	Username     string    `sql:"username"`
	PasswordHash []byte    `sql:"password_hash"`
	City         string    `sql:"city"`
	State        string    `sql:"state"`
	CreatedAt    time.Time `sql:"created_at"`
}

func ConnectUsers(conn *sql.DB) (err error) {
	db = conn
	err = db.Exec(CreateTableUser)
	return
}

func (u Users) Verify(username string, password []byte) error {

}

func (u Users) Find(username string) (u *User, err error) {
	rows, err := u.db.Query(SelectUserByName, username)
	if err != nil {
		return nil, err
	}

	err = ScanUser(rows, u)
	return
	// if rows.Next() {
	// 	var u User
	// 	err := rows.Scan(&u.Id, &u.Username, &u.PasswordHash, &u.City, &u.State, &u.CreatedAt)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &u, nil
	// } else {
	// 	return nil, ErrNotFoundUser
	// }
}

func ScanUser(r *sql.Rows, u *User) (err error) {
	if r.Next() {
		err = rows.Scan(u.Id, u.Username, u.PasswordHash, u.City, u.State, u.CreatedAt)
	} else {
		err = ErrNotFoundUser
	}
	return
}

// type users struct {
// 	Verify func(string, []byte) error
// 	Find   func(string) *User
// }

// var Users = &users{
// 	Verify: verify,
// 	Find:   find,
// }

// type User struct {
// 	gorm.Model
// 	Username     string `gorm:"not null;unique"`
// 	PasswordHash []byte `gorm:"not null"`
// 	City         string
// 	State        string
// }

// func CreateUserSchema(db *gorm.DB) (err error) {
// 	return db.AutoMigrate(&User{}).Error
// }

func (u *User) StoreHash(password []byte) error {
	hash, err := bcrypt.GenerateFromPassword(password, -1)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	return nil
}

func (u *User) Create() (err error) {
	err = db.Create(u).Error
	return
}

// cost, err := bcrypt.Cost(hash)
// err = bcrypt.CompareHashAndPassword(hash, []byte(pass))
// bcrypt.ErrMismatchedHashAndPassword
func verify(username string, password []byte) error {
	u := User{}
	db.Where("Username = ?", username).First(&u)
	// TODO: what if user does not exist?
	// if db.Error != nil {
	// 	// handle not found!
	// }
	return bcrypt.CompareHashAndPassword(u.PasswordHash, password)
}

func find(username string) (u *User) {
	db.Where("Username = ?", username).First(u)
	return
}

func (u *User) Save() error {
	return db.Save(u).Error
}
