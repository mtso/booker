package controllers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/mtso/booker/server/models"
	"github.com/mtso/booker/server/utils"
)

var ErrInvalidBody = errors.New("Required fields: title, isbn, and image_url")

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bks, err := models.Books.GetBooks()
	if WriteErrorResponse(w, err) {
		return
	}

	for _, v := range bks {
		v.ImageUrl = url.QueryEscape(v.ImageUrl)
	}

	resp := make(JsonResponse)
	resp["ok"] = true
	resp["data"] = bks

	WriteJson(w, resp)
}

func GetMyBooks(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(r)
	if WriteErrorResponse(w, err) {
		return
	}

	bks, err := models.Books.GetMyBooks(username)
	if WriteErrorResponse(w, err) {
		return
	}

	for _, v := range bks {
		v.ImageUrl = url.QueryEscape(v.ImageUrl)
	}

	resp := make(JsonResponse)
	resp["ok"] = true
	resp["data"] = bks

	WriteJson(w, resp)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ParseRequestBody(r)
	if WriteErrorResponse(w, err) {
		return
	}

	title, titleOk := body["title"]
	isbn, isbnOk := body["isbn"]
	imageUrl, imageUrlOk := body["image_url"]

	if !titleOk || !isbnOk || !imageUrlOk {
		WriteErrorResponse(w, ErrInvalidBody)
		return
	}

	username, err := GetUsername(r)
	if WriteErrorResponse(w, err) {
		return
	}

	user, err := models.Users.Find(username)
	if WriteErrorResponse(w, err) {
		return
	}

	book := &models.Book{
		Title:    title.(string),
		Isbn:     isbn.(string),
		ImageUrl: imageUrl.(string),
		UserId:   user.Id,
	}

	err = book.Create()
	if WriteErrorResponse(w, err) {
		return
	}

	resp := make(JsonResponse)
	resp["ok"] = true
	resp["message"] = "Created book: " + title.(string)

	WriteJson(w, resp)
}
