package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtso/booker/server/config"
)

func TestApp(t *testing.T) {
	// Set up assertions
	mustEqual := MakeMustEqual(t)
	assertEqual := MakeAssertEqual(t)

	// Start test server
	app := config.InitializeApp()
	ts := httptest.NewServer(app.Handler)
	defer app.Db.Close()
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	mustEqual(err, nil, "make GET request to root: " + ts.URL)
	assertEqual(resp.StatusCode, 200, "reach server route")
}
