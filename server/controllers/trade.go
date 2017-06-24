package controllers

import (
	"net/http"

	"github.com/mtso/booker/server/models"
)

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
		"ok": true,
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
		"ok": true,
		"trades": trs,
	}

	WriteJson(w, resp)
}

