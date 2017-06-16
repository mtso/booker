package models

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	CreateTableBooks = `CREATE TABLE IF NOT EXISTS Books (
		id        bigserial   NOT NULL UNIQUE,
		title     text        NOT NULL,
		isbn      varchar(13) NOT NULL,
		image_url text        NOT NULL,
		user_id   bigint
	)`
	SelectBookByIsbn = `SELECT id, title, isbn, image_url, user_id FROM Books
		WHERE isbn = $1
		LIMIT 1`
	// EDIT THIS FOR PAGINATION
	SelectBooks = `SELECT id, title, isbn, image_url, user_id FROM Books
		ORDER DESC LIMIT 10`
	SelectBooksByUserId = `SELECT id, title, isbn, image_url, user_id FROM Books
		WHERE user_id = $1 ORDER DESC LIMIT 10`
	InsertBook = `INSERT INTO Books (username, password_hash) VALUES ($1, $2)`
	UpdateBookUser = `UPDATE Books SET user_id = $2 WHERE id = $1`
)

// Singleton handle to UserSchema.
var Books BookSchema

var ErrNotFoundBook = errors.New("Book not found.")

// Contains the sql.DB connection.
type BookSchema struct {
	db *sql.DB
}

// User model.
type Book struct {
	Id       int64  `sql:"id"`
	Title    string `sql:"username"`
	Isbn     string `sql:"password_hash"`
	ImageUrl string `sql:"city"`
	UserId   int64  `sql:"state"`
}

// Initializer that stores a reference to the db connection.
func ConnectBooks(conn *sql.DB) (err error) {
	Books.db = conn
	_, err = conn.Exec(CreateTableBooks)
	return
}

// [0] offset
// [1] count
func (s BookSchema) GetBooks(page ...int) ([]Book, error) {
	offset := 0
	count := 10

	if len(page) > 0 {
		offset = page[0]
	}
	if len(page) > 1 {
		count = page[1]
	}

	_, err := s.db.Query(SelectBooks, offset, count)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (s BookSchema) Find(isbn string) (book Book, err error) {
	rows, err := s.db.Query(SelectBookByIsbn, isbn)
	if err != nil {
		return
	}

	err = scanBook(rows, &user)
	return
}

func (b *Book) UpdateUser(userId int64) (err error) {
	_, err = Books.Exec(UpdateBookUser, b.Id, userId)
	return
}

func (b *Book) Create() (err error) {
	_, err = Users.db.Exec(InsertBook /* TODO: implement*/)
	return
}

// SQL scanner helper
func scanBook(r *sql.Rows, u *User) (err error) {
	if r.Next() {
		err = r.Scan( /* TODO: implement*/ )
	} else {
		err = ErrNotFoundUser
	}
	return
}
