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

	resp := &JsonResponse{
		Ok:   true,
		Data: bks,
	}

	WriteJson(w, resp)
}
