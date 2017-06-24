package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mtso/booker/server/models"
	"github.com/mtso/booker/server/utils"
)

var ErrBookNotFound = errors.New("Missing field: book_id")

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
		WriteErrorResponse(w, ErrBookNotFound)
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
