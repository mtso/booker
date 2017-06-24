package models

import (
	"log"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	// StatusRequested = iota
	// StatusAccepted
	// StatusCanceled
	StatusRequested = "StatusRequested"
	StatusAccepted = "StatusAccepted"
	StatusCanceled = "StatusCanceled"
)

const (
	// STATUS SHOULD BE SOME KIND OF ENUM
	CreateTableTrades = `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
			CREATE TYPE status AS ENUM ('StatusRequested', 'StatusAccepted', 'StatusCanceled');
		END IF;
	END$$;

	CREATE TABLE IF NOT EXISTS Trades (
		id      bigserial NOT NULL UNIQUE,
		user_id bigint    NOT NULL,
		book_id bigint    NOT NULL,
		status  status    NOT NULL DEFAULT 'StatusRequested'
	)`

	GetIncoming = `SELECT DISTINCT ON(trades.id)
			trades.id,
			trades.book_id,
			trades.user_id,
			trades.status,
			users.username,
			users.city,
			users.state,
			books.title,
			books.image_url
		FROM Trades, Users, Books
		WHERE books.user_id = $1
		AND trades.book_id = books.id
		AND trades.status = 'StatusRequested'
		AND trades.user_id = users.id
		ORDER BY trades.id DESC`

	// GetIncomingTest = `SELECT DISTINCT ON(trades.id)
	// 		trades.id,
	// 		trades.user_id
	// 	FROM Trades, Users, Books
	// 	WHERE books.user_id = $1
	// 	AND trades.book_id = books.id
	// 	AND trades.status = 'StatusRequested'
	// 	AND trades.user_id = users.id
	// 	ORDER BY trades.id DESC`

	GetOutgoing = `SELECT DISTINCT ON(trades.id)
			trades.id,
			trades.book_id,
			trades.user_id,
			trades.status,
			users.username,
			users.city,
			users.state,
			books.title,
			books.image_url
		FROM Trades, Users, Books
		WHERE trades.user_id = $1
		AND trades.status = 'StatusRequested'
		AND users.id = books.user_id
		ORDER BY trades.id DESC`

	GetTrade = `SELECT DISTINCT ON(trades.id) trades.id, books.title FROM Trades, Users, Books WHERE trades.user_id = $1 OR books.user_id = $1`

	InsertTrade = `INSERT INTO Trades (user_id, book_id) VALUES ($1, $2)`
	UpdateTrade = `UPDATE Trades
		SET status = $2
		WHERE id = $1`
)

// Singleton handle to UserSchema.
var Trades TradeSchema

var ErrNotFoundTrade = errors.New("Trade not found.")

// Contains the sql.DB connection.
type TradeSchema struct {
	db *sql.DB
}

// Trade model.
type Trade struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"user_id"`
	BookId int64 `json:"book_id"`
	Status string `json:"status"`
}

// Initializer that stores a reference to the db connection.
func ConnectTrades(conn *sql.DB) (err error) {
	Trades.db = conn
	_, err = conn.Exec(CreateTableTrades)
	return
}

// type 

type TradeResponse struct{
	Id int64 `json:"id"`
	BookId int64 `json:"book_id"`
	UserId int64 `json:"user_id"`
	Status string `json:"status"`

	User struct {
		Username string `json:"username"`
		City string `json:"city"`
		State string `json:"state"`
	} `json:"user"`

	Book struct {
		Title string `json:"title"`
		ImageUrl string `json:"image_url"`
	} `json:"book"`
}

func (s TradeSchema) GetIncomingTrades(userid int64) ([]TradeResponse, error) {
	rows, err := s.db.Query(GetIncoming, userid)
	if err != nil {
		return nil, err
	}

	trades := make([]TradeResponse, 0)

	var tr TradeResponse
	for scanIncomingTrade(rows, &tr) == nil {
		trades = append(trades, tr)
	}
	log.Printf("")

	return trades, nil
}

func (s TradeSchema) GetOutgoingTrades(userid int64) ([]TradeResponse, error) {
	rows, err := s.db.Query(GetOutgoing, userid)
	if err != nil {
		return nil, err
	}

	trades := make([]TradeResponse, 0)

	var tr TradeResponse
	for scanIncomingTrade(rows, &tr) == nil {
		trades = append(trades, tr)
	}
	// log.Println(trades)

	return trades, nil
}

func scanIncomingTrade(r *sql.Rows, t *TradeResponse) error {
	if r.Next() {
		var city sql.NullString
		var state sql.NullString

		err := r.Scan(&t.Id, &t.BookId, &t.UserId, &t.Status,
			&t.User.Username, &city, &state,
			&t.Book.Title, &t.Book.ImageUrl)

		if c, err := city.Value(); err != nil {
			t.User.City = c.(string)
		}
		if s, err := state.Value(); err != nil {
			t.User.State = s.(string)
		}
		return err
	}

	return ErrNotFoundTrade
}

// // [0] offset
// // [1] count
// func (s TradeSchema) GetTrades(page ...int) ([]Trade, error) {
// 	offset := 0
// 	count := 10

// 	if len(page) > 0 {
// 		offset = page[0]
// 	}
// 	if len(page) > 1 {
// 		count = page[1]
// 	}

// 	_, err := s.db.Query(SelectTrades, offset, count)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, err
// }

// func (s TradeSchema) Find(isbn string) (book Trade, err error) {
// 	rows, err := s.db.Query(SelectTradeByIsbn, isbn)
// 	if err != nil {
// 		return
// 	}

// 	err = scanTrade(rows, &user)
// 	return
// }

// func (t *Trade) UpdateTrade(userId int64) (err error) {
// 	_, err = Trades.Exec(UpdateTradeUser, b.Id, userId)
// 	return
// }

// func (t *Trade) Create() (err error) {
// 	_, err = Users.db.Exec(InsertTrade /* TODO: implement*/)
// 	return
// }

// // SQL scanner helper
// func scanTrade(r *sql.Rows, u *User) (err error) {
// 	if r.Next() {
// 		err = r.Scan( /* TODO: implement*/ )
// 	} else {
// 		err = ErrNotFoundUser
// 	}
// 	return
// }
