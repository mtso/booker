package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtso/booker/server/config"
)

func TestTrade(t *testing.T) {
	// Set up assertions
	assertEqual := MakeAssertEqual(t)
	mustEqual := MakeMustEqual(t)

	// Init client with cookie jar
	client2 := MakeCookieMonster()

	// Start test server
	app := config.InitializeApp()
	ts := httptest.NewServer(app.Handler)
	defer app.Db.Close()
	defer ts.Close()

	err := AuthenticateSession(ts, client2, User2, Pass2)

	data := bytes.NewBuffer([]byte(`{"isRequest":true}`))
	req, err := http.NewRequest("POST", ts.URL+"/api/trade/1", data)
	mustEqual(err, nil, "prep POST /api/trade/1")

	res, err := client2.Do(req)
	mustEqual(err, nil, "reach server")

	resp, err := ParseBody(res)
	mustEqual(err, nil, "respond with JSON")

	assertEqual(resp["ok"], true, "validate trade request")
}
