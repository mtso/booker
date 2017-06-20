package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtso/booker/server"
)

func TestGetBooks(t *testing.T) {
	// Set up assertions
	assertEqual := MakeAssertEqual(t)
	mustEqual := MakeMustEqual(t)

	// Init client with cookie jar
	client := MakeCookieMonster()

	// Start test server
	app := main.InitializeApp()
	defer app.Db.Close()

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/api/books", nil)
	mustEqual(err, nil, "prep request to /api/books")

	res, err := client.Do(req)
	mustEqual(err, nil, "execute GET /api/books")

	resp, err := ParseBody(res)
	mustEqual(err, nil, "body is encoded in JSON")

	books, ok := resp["data"].([]interface{})
	assertEqual(ok, true, "data property is an array")
	assertEqual(len(books), 10, "return length of 10 books per page")
}
