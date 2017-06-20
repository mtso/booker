package models

// import (
// 	"database/sql"
// 	"errors"

// 	_ "github.com/lib/pq"
// )

// const (
// 	// STATUS SHOULD BE SOME KIND OF ENUM
// 	CreateTableTrades = `CREATE TABLE IF NOT EXISTS Trades (
// 		id      bigserial NOT NULL UNIQUE,
// 		user_id bigint    NOT NULL,
// 		book_id bigint    NOT NULL,
// 		status  text      NOT NULL
// 	)`
// 	SelectTradesByUserId = `SELECT id, user_id, book_id, status FROM Trades
// 		WHERE user_id = $1, status = xxx ORDER DESC LIMIT 30`
// 	InsertTrade = `INSERT INTO Trades (user_id, book_id, status) VALUES ($1, $2, $3)`
// 	UpdateTrade = `UPDATE Trades
// 		SET status = $2
// 		WHERE id = $1`
// )

// // Singleton handle to UserSchema.
// var Trades TradeSchema

// var ErrNotFoundTrade = errors.New("Trade not found.")

// // Contains the sql.DB connection.
// type TradeSchema struct {
// 	db *sql.DB
// }

// // User model.
// type Trade struct {
// 	Id     int64  `sql:"id"`
// 	UserId int64 `sql:"user_id"`
// 	BookId int64 `sql:"book_id"`
// 	Status int64 `sql:"status"`
// }

// // Initializer that stores a reference to the db connection.
// func ConnectTrades(conn *sql.DB) (err error) {
// 	Trades.db = conn
// 	_, err = conn.Exec(CreateTableTrades)
// 	return
// }

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
