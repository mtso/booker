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
	SelectBookById = `SELECT books.id, title, isbn, image_url, user_id, username FROM Books, Users
		WHERE books.id = $1 AND users.id = books.user_id
		LIMIT 1`
	SelectBookByIsbn = `SELECT books.id, title, isbn, image_url, user_id FROM Books
		WHERE isbn = $1
		LIMIT 1`
	SelectBooks = `SELECT
		DISTINCT ON (books.id) 
			books.id
			, title
			, isbn
			, image_url
			, username
		FROM Books, Users
		WHERE users.id = books.user_id
		ORDER BY books.id DESC
		OFFSET $1 LIMIT $2`
	SelectMyBooks = `SELECT
		DISTINCT ON (books.id) 
			books.id
			, title
			, isbn
			, image_url
			, username
		FROM Books, Users
		WHERE users.id = books.user_id AND users.username = $1
		ORDER BY books.id DESC`
	SelectBooksByUserId = `SELECT id, title, isbn, image_url, user_id FROM Books
		WHERE user_id = $1 ORDER DESC LIMIT 10`
	InsertBook     = `INSERT INTO Books (title, isbn, image_url, user_id) VALUES ($1, $2, $3, $4)`
	UpdateBookUser = `UPDATE Books SET user_id = $2 WHERE id = $1`
)

// Singleton handle to UserSchema.
var Books BookSchema

var ErrNotFoundBook = errors.New("Book not found.")

// Contains the sql.DB connection.
type BookSchema struct {
	db *sql.DB
}

// Book model.
type Book struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	ImageUrl string `json:"image_url"`
	UserId   int64  `json:"user_id,omitempty"`
	Username string `json:"username"`
}

type BookTest struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	ImageUrl string `json:"image_url"`
	UserId   int64  `json:"user_id,omitempty"`
	User struct {
		Name string `json:"username"`
	} `json:"user"`
	// Username string `json:"username"`
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

	rows, err := s.db.Query(SelectBooks, offset, count)
	if err != nil {
		return nil, err
	}

	bks := make([]Book, 0)
	for rows.Next() {
		var bk Book
		err := rows.Scan(&bk.Id, &bk.Title, &bk.Isbn, &bk.ImageUrl, &bk.Username)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	return bks, nil
}

// Sketchy GetMyBooks where user.username == username
func (s BookSchema) GetMyBooks(username string) ([]Book, error) {
	rows, err := s.db.Query(SelectMyBooks, username)
	if err != nil {
		return nil, err
	}

	bks := make([]Book, 0)
	for rows.Next() {
		var bk Book
		err := rows.Scan(&bk.Id, &bk.Title, &bk.Isbn, &bk.ImageUrl, &bk.Username)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	return bks, nil
}

func (s BookSchema) FindById(id int64) (book Book, err error) {
	rows, err := s.db.Query(SelectBookById, id)
	if err != nil {
		return
	}

	err = scanFullBook(rows, &book)
	return
}

func (s BookSchema) Find(isbn string) (book Book, err error) {
	rows, err := s.db.Query(SelectBookByIsbn, isbn)
	if err != nil {
		return
	}

	err = scanBook(rows, &book)
	return
}

func (b *Book) UpdateUser(userId int64) (err error) {
	_, err = Books.db.Exec(UpdateBookUser, b.Id, userId)
	return
}

func (b *Book) Create() (err error) {
	_, err = Users.db.Exec(InsertBook, b.Title, b.Isbn, b.ImageUrl, b.UserId)
	return
}

func scanFullBook(r *sql.Rows, b *Book) (err error) {
	if r.Next() {
		err = r.Scan(&b.Id, &b.Title, &b.Isbn, &b.ImageUrl, &b.UserId, &b.Username)
	} else {
		err = ErrNotFoundUser
	}
	return
}

// SQL scanner helper
func scanBook(r *sql.Rows, b *Book) (err error) {
	if r.Next() {
		err = r.Scan(&b.Id, &b.Title, &b.Isbn, &b.ImageUrl, &b.UserId)
	} else {
		err = ErrNotFoundUser
	}
	return
}
