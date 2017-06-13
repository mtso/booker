package models

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

type users struct {
	Verify func(string, []byte) error
	Find   func(string) *User
}

var Users = &users{
	Verify: verify,
	Find:   find,
}

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique"`
	PasswordHash []byte `gorm:"not null"`
	City         string
	State        string
}

func CreateUserSchema(db *gorm.DB) (err error) {
	return db.AutoMigrate(&User{}).Error
}

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
