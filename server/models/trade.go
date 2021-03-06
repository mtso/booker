package models

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	StatusRequested = "StatusRequested"
	StatusAccepted  = "StatusAccepted"
	StatusCanceled  = "StatusCanceled"
)

const (
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
			users.display_name,
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

	GetOutgoing = `SELECT DISTINCT ON(trades.id)
			trades.id,
			trades.book_id,
			trades.user_id,
			trades.status,
			users.username,
			users.display_name,
			users.city,
			users.state,
			books.title,
			books.image_url
		FROM Trades, Users, Books
		WHERE trades.user_id = $1
		AND trades.status = 'StatusRequested'
		AND users.id = books.user_id
		AND books.id = trades.book_id
		ORDER BY trades.id DESC`

	GetTrade = `SELECT DISTINCT ON(trades.id) trades.id, books.title FROM Trades, Users, Books WHERE trades.user_id = $1 OR books.user_id = $1`

	InsertTrade = `INSERT INTO Trades (user_id, book_id) SELECT $1, $2
		WHERE NOT EXISTS (
			SELECT * FROM Trades
			WHERE user_id = $1
			AND book_id = $2
			AND status = status 'StatusRequested'
		)`

	// find trade by tradeid
	// validate that trades.book_id's book.user_id is userid
	// update trades
	// update books
	UpdateTrade = `UPDATE Trades
		SET status = CASE
			WHEN id = $1 THEN status 'StatusAccepted'
			ELSE status 'StatusCanceled'
		END
		WHERE book_id = $2`

	CancelTradeExec = `UPDATE Trades
		SET status = 'StatusCanceled'
		WHERE id = $1 AND user_id = $2`

	SelectById = `SELECT id, book_id, user_id, status
		FROM Trades
		WHERE id = $1
		LIMIT 1`

	SelectByUser = `SELECT id, book_id, user_id, status
		FROM Trades
		WHERE user_id = $1
			AND book_id = $2
		ORDER BY id DESC
		LIMIT 1`
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
	Id     int64  `json:"id"`
	BookId int64  `json:"book_id"`
	UserId int64  `json:"user_id"`
	Status string `json:"status"`
}

// Initializer that stores a reference to the db connection.
func ConnectTrades(conn *sql.DB) (err error) {
	Trades.db = conn
	_, err = conn.Exec(CreateTableTrades)
	return
}

type TradeResponse struct {
	Id     int64  `json:"id"`
	BookId int64  `json:"book_id"`
	UserId int64  `json:"user_id"`
	Status string `json:"status"`

	User struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		City        string `json:"city"`
		State       string `json:"state"`
	} `json:"user"`

	Book struct {
		Title    string `json:"title"`
		ImageUrl string `json:"image_url"`
	} `json:"book"`
}

func (s TradeSchema) GetIncomingTrades(userid int64) ([]TradeResponse, error) {
	rows, err := s.db.Query(GetIncoming, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trades := make([]TradeResponse, 0)

	var tr TradeResponse
	for scanTradeResponse(rows, &tr) == nil {
		trades = append(trades, tr)
	}

	return trades, nil
}

func (s TradeSchema) GetOutgoingTrades(userid int64) ([]TradeResponse, error) {
	rows, err := s.db.Query(GetOutgoing, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trades := make([]TradeResponse, 0)

	var tr TradeResponse
	for scanTradeResponse(rows, &tr) == nil {
		trades = append(trades, tr)
	}

	return trades, nil
}

func scanTradeResponse(r *sql.Rows, t *TradeResponse) error {
	if r.Next() {
		var city sql.NullString
		var state sql.NullString

		err := r.Scan(&t.Id, &t.BookId, &t.UserId, &t.Status,
			&t.User.Username, &t.User.DisplayName, &city, &state,
			&t.Book.Title, &t.Book.ImageUrl)

		if c, _ := city.Value(); c != nil {
			t.User.City = c.(string)
		}

		if s, _ := state.Value(); s != nil {
			t.User.State = s.(string)
		}

		return err
	}

	return ErrNotFoundTrade
}

func (s TradeSchema) Create(userid, bookid int64) (err error) {
	_, err = s.db.Exec(InsertTrade, userid, bookid)
	return
}

// Trade-Accepting funcs

func (s TradeSchema) FindById(id string) (t Trade, err error) {
	rows, err := s.db.Query(SelectById, id)
	if err != nil {
		return
	}
	defer rows.Close()

	err = scanTrade(rows, &t)
	return
}

func (s TradeSchema) FindByUser(id int64, bookid int64) (t Trade, err error) {
	rows, err := s.db.Query(SelectByUser, id, bookid)
	if err != nil {
		return
	}
	defer rows.Close()

	err = scanTrade(rows, &t)
	return
}

func scanTrade(r *sql.Rows, t *Trade) (err error) {
	if r.Next() {
		err = r.Scan(&t.Id, &t.BookId, &t.UserId, &t.Status)
	} else {
		err = ErrNotFoundTrade
	}
	return
}

func (t *Trade) AcceptTrade() error {
	if _, err := Trades.db.Exec(UpdateTrade, t.Id, t.BookId); err != nil {
		return err
	}
	if _, err := Books.db.Exec(UpdateBookUser, t.BookId, t.UserId); err != nil {
		return err
	}
	return nil
}

func (s TradeSchema) CancelTrade(id string, userid int64) (int64, error) {
	res, err := s.db.Exec(CancelTradeExec, id, userid)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
