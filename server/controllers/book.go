package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
	"github.com/mtso/booker/server/utils"
)

var ErrInvalidBody = errors.New("Required fields: title, isbn, and image_url")
var ErrNoId = errors.New("Required field: id")

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idparam, ok := vars["id"]
	if !ok {
		WriteErrorResponse(w, ErrNoId)
		return
	}

	id, err := strconv.ParseInt(idparam, 10, 64)
	if WriteErrorResponse(w, err) {
		return
	}

	book, err := models.Books.GetBookResponse(id)

	// If Logged in, get trade status
	username, err := GetUsername(r)
	if err == nil {
		user, err := models.Users.Find(*username)
		if WriteErrorResponse(w, err) {
			return
		}

		trade, err := models.Trades.FindByUser(user.Id, book.Id)
		if err != nil {
			book.Trade.Status = StatusNotTraded
		} else if trade.Status == "" {
			book.Trade.Status = StatusNotTraded
		} else {
			book.Trade.Status = trade.Status
			book.Trade.Id = trade.Id
		}
	}

	resp := &JsonResponse{
		"ok":   true,
		"book": book,
	}
	WriteJson(w, resp)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bks, err := models.Books.GetBooks()
	if WriteErrorResponse(w, err) {
		return
	}

	for _, v := range bks {
		v.ImageUrl = url.QueryEscape(v.ImageUrl)
	}

	resp := &JsonResponse{
		"ok":   true,
		"data": bks,
	}

	WriteJson(w, resp)
}

func GetMyBooks(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(r)
	if WriteErrorResponse(w, err) {
		return
	}

	bks, err := models.Books.GetMyBooks(*username)
	if WriteErrorResponse(w, err) {
		return
	}

	for _, v := range bks {
		v.ImageUrl = url.QueryEscape(v.ImageUrl)
	}

	resp := &JsonResponse{
		"ok":   true,
		"data": bks,
	}

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

	user, err := models.Users.Find(*username)
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

	resp := &JsonResponse{
		"ok":      true,
		"message": "Created book: " + title.(string),
	}

	WriteJson(w, resp)
}
