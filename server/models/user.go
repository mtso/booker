package models

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	CreateTableUsers = `CREATE TABLE IF NOT EXISTS Users (
		id            bigserial    NOT NULL UNIQUE,
		username      varchar(64)  NOT NULL UNIQUE,
		display_name  varchar(64)  NOT NULL,
		password_hash varchar(64)  NOT NULL,
		city          varchar(128),
		state         varchar(64)
	)`
	SelectUserByName = `SELECT id, username, display_name, password_hash, city, state FROM Users
		WHERE username = $1
		LIMIT 1`
	InsertUser         = `INSERT INTO Users (username, display_name, password_hash) VALUES ($1, $1, $2)`
	UpdateUserLocation = `UPDATE Users
		SET city = $2,
		    state = $3
		WHERE username = $1`
	UpdateUserPassword = `UPDATE Users
		SET password_hash = $2
		WHERE username = $1`
)

// Singleton handle to UserSchema.
var Users UserSchema

var ErrNotFoundUser = errors.New("User not found.")

// Contains the sql.DB connection.
type UserSchema struct {
	db *sql.DB
}

// User model.
type User struct {
	Id           int64          `sql:"id"`
	Username     string         `sql:"username"`
	DisplayName  string         `sql:"display_name"`
	PasswordHash []byte         `sql:"password_hash"`
	City         sql.NullString `sql:"city"`
	State        sql.NullString `sql:"state"`
}

// Initializer that stores a reference to the db connection.
func ConnectUsers(conn *sql.DB) (err error) {
	Users.db = conn
	_, err = conn.Exec(CreateTableUsers)
	return
}

func (u UserSchema) Verify(username string, password []byte) error {
	user, err := u.Find(username)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(user.PasswordHash, password)
}

func (u UserSchema) FindAndVerify(username string, password []byte) (user User, err error) {
	user, err = u.Find(username)
	if err != nil {
		return
	}
	return user, bcrypt.CompareHashAndPassword(user.PasswordHash, password)
}

func (u UserSchema) Find(username string) (user User, err error) {
	rows, err := u.db.Query(SelectUserByName, username)
	if err != nil {
		return
	}

	err = scanUser(rows, &user)
	return
}

func (u *User) Create() (err error) {
	_, err = Users.db.Exec(InsertUser, u.Username, u.PasswordHash)
	return
}

func (u *User) SetLocation(city string, state string) (err error) {
	_, err = Users.db.Exec(UpdateUserLocation, u.Username, city, state)
	return
}

func (u *User) SavePasswordHash(password []byte) error {
	hash, err := bcrypt.GenerateFromPassword(password, -1)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	_, err = Users.db.Exec(UpdateUserPassword, u.Username, hash)
	return err
}

func (u *User) SetPasswordHash(password []byte) error {
	hash, err := bcrypt.GenerateFromPassword(password, -1)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	return nil
}

// SQL scanner helper
func scanUser(r *sql.Rows, u *User) (err error) {
	if r.Next() {
		err = r.Scan(&u.Id, &u.Username, &u.DisplayName, &u.PasswordHash, &u.City, &u.State)
	} else {
		err = ErrNotFoundUser
	}
	return
}
