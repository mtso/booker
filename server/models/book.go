package models

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	CreateTableBooks = `CREATE TABLE IF NOT EXISTS Users (
		id        bigserial   NOT NULL UNIQUE,
		title     text        NOT NULL,
		isbn      varchar(13) NOT NULL,
		image_url text        NOT NULL,
		user_id   bigint
	)`
	SelectBookByUserId = ``
	SelectBooksDesc = `SELECT id, title, isbn`

	SelectUserByName = `SELECT id, username, password_hash, city, state FROM Users
		WHERE username = $1
		LIMIT 1`
	InsertUser = `INSERT INTO Users (username, password_hash) VALUES ($1, $2)`
	UpdateUser = `UPDATE Users
		SET city = $2,
		    state = $3
		WHERE username = $1`
)

// Singleton handle to UserSchema.
var Books BookSchema

var ErrNotFoundBook = errors.New("User not found.")

// Contains the sql.DB connection.
type BookSchema struct {
	db *sql.DB
}

// User model.
type Book struct {
	Id        int64  `sql:"id"`
	Title     string `sql:"username"`
	Isbn      string `sql:"password_hash"`
	ImageUrl  string `sql:"city"`
	UserId    int64  `sql:"state"`
}

// Initializer that stores a reference to the db connection.
func ConnectBooks(conn *sql.DB) (err error) {
	Books.db = conn
	_, err = conn.Exec(CreateTableBooks)
	return
}

// [0] offset
// [1] count
func (s BookSchema) GetBooks(page ...int) []Book, error {
	offset := 0
	count := 10

	if len(page) > 0 {
		offset = page[0]
	}
	if len(page) > 1 {
		count = page[1]
	}

	rows, err := s.db.Query(SelectBooks, offset, count)
	if err != nil {
		return nil, err
	}

}

func (s BookSchema) Find(username string) (user User, err error) {
	rows, err := s.db.Query(SelectUserByName, username)
	if err != nil {
		return
	}

	err = scanUser(rows, &user)
	return
}

func (b *Book) Create() (err error) {
	_, err = Users.db.Exec(InsertBook/* TODO: implement*/)
	return
}

// SQL scanner helper
func scanBook(r *sql.Rows, u *User) (err error) {
	if r.Next() {
		err = r.Scan(/* TODO: implement*/)
	} else {
		err = ErrNotFoundUser
	}
	return
}
