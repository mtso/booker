package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mtso/booker/server/models"
	"github.com/mtso/booker/server/utils"
)

var ErrMissingFieldBook = errors.New("Missing field: book_id")
var ErrMissingFieldTrade = errors.New("Missing field: book_id")

func PostTrade(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(r)
	if WriteErrorResponse(w, err) {
		return
	}

	user, err := models.Users.Find(*username)
	if WriteErrorResponse(w, err) {
		return
	}

	body, err := utils.ParseRequestBody(r)
	if WriteErrorResponse(w, err) {
		return
	}

	bookid, ok := body["book_id"]
	if !ok {
		WriteErrorResponse(w, ErrMissingFieldBook)
		return
	}

	err = models.Trades.Create(user.Id, int64(bookid.(float64)))
	if WriteErrorResponse(w, err) {
		return
	}

	resp := &JsonResponse{
		"ok":      true,
		"message": fmt.Sprintf("Made trade request to %s", bookid),
	}

	WriteJson(w, resp)
}

func GetIncomingTrades(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(r)
	if WriteErrorResponse(w, err) {
		return
	}

	user, err := models.Users.Find(*username)
	if WriteErrorResponse(w, err) {
		return
	}

	trs, err := models.Trades.GetIncomingTrades(user.Id)
	if WriteErrorResponse(w, err) {
		return
	}

	resp := &JsonResponse{
		"ok":     true,
		"trades": trs,
	}

	WriteJson(w, resp)
}

func GetOutgoingTrades(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsername(r)
	if WriteErrorResponse(w, err) {
		return
	}

	user, err := models.Users.Find(*username)
	if WriteErrorResponse(w, err) {
		return
	}

	trs, err := models.Trades.GetOutgoingTrades(user.Id)
	if WriteErrorResponse(w, err) {
		return
	}

	resp := &JsonResponse{
		"ok":     true,
		"trades": trs,
	}

	WriteJson(w, resp)
}

var ErrUnauthorized = errors.New("Unauthorized access to resource")

func PutTrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradeid, ok := vars["id"]
	if !ok {
		WriteErrorResponse(w, ErrMissingFieldTrade)
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

	trade, err := models.Trades.FindById(tradeid)
	if WriteErrorResponse(w, err) {
		return
	}

	book, err := models.Books.FindById(trade.BookId)
	if WriteErrorResponse(w, err) {
		return
	}

	if book.UserId != user.Id {
		WriteErrorResponse(w, ErrUnauthorized)
		return
	}

	err = trade.AcceptTrade()
	if WriteErrorResponse(w, err) {
		return
	}

	resp := &JsonResponse{
		"ok":      "true",
		"message": fmt.Sprintf("Accepted trade %d", trade.Id),
	}

	WriteJson(w, resp)
}
