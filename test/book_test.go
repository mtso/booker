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

	assertEqual(res.StatusCode, 200, "is valid route /api/books")

	resp, err := ParseBody(res)
	mustEqual(err, nil, "body is encoded in JSON")

	books, ok := resp["data"].([]interface{})
	assertEqual(ok, true, "data property is an array")
	assertEqual(len(books) <= 10, true, "return length of max 10 books per page")

	mustEqual(len(books) > 0, true, "need a book to test")

	raw := books[0]
	book := raw.(map[string]interface{})

	_, ok = book["isbn"]
	assertEqual(ok, true, "book object has isbn")
	_, ok = book["image_url"]
	assertEqual(ok, true, "book object has image_url")
	_, ok = book["username"]
	assertEqual(ok, true, "book object has username")
	_, ok = book["id"]
	assertEqual(ok, true, "book object has id")
	title, ok := book["title"]
	mustEqual(ok, true, "book object has a title")

	switch title.(type) {
	case string:
		break
	default:
		assertEqual(false, true, "title is a string")
	}
}
