package controllers

import (
	"net/http"
	"net/url"

	"github.com/mtso/booker/server/models"
)

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
