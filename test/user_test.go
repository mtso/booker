package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtso/booker/server/config"
)

func TestUpdateLocation(t *testing.T) {
	// Set up assertions
	assertEqual := MakeAssertEqual(t)
	mustEqual := MakeMustEqual(t)

	// Init client with cookie jar
	client := MakeCookieMonster()

	// Start test server
	app := config.InitializeApp()
	defer app.Db.Close()

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	err := AuthenticateSession(ts, client)
	mustEqual(err, nil, "authenticate session")

	data := []byte(`{"city":"Test City","state":"Test State"}`)
	req, err := http.NewRequest("POST", ts.URL+"/api/user", bytes.NewBuffer(data))
	mustEqual(err, nil, "prep post request to /api/user")

	res, err := client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err := ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["ok"], true, "good request")

	req, err = http.NewRequest("GET", ts.URL+"/api/user", nil)
	mustEqual(err, nil, "prep get request to /api/user")

	res, err = client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err = ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["username"], "wiggs", "recognize username")
	assertEqual(body["city"], "Test City", "save city")
	assertEqual(body["state"], "Test State", "save state")
}

func TestUpdatePassword(t *testing.T) {
	// Set up assertions
	assertEqual := MakeAssertEqual(t)
	mustEqual := MakeMustEqual(t)

	// Init client with cookie jar
	client := MakeCookieMonster()

	// Start test server
	app := config.InitializeApp()
	defer app.Db.Close()

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	err := AuthenticateSession(ts, client)
	mustEqual(err, nil, "authenticate session")

	// Update password
	data := []byte(`{"password":"muffins"}`)
	req, err := http.NewRequest("POST", ts.URL+"/api/user", bytes.NewBuffer(data))
	mustEqual(err, nil, "prep post request to /api/user")

	res, err := client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err := ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["ok"], true, "get a JSON API response")
	assertEqual(body["message"], "password changed", "successfully change password")

	// Logout
	req, err = http.NewRequest("POST", ts.URL+"/auth/logout", nil)
	mustEqual(err, nil, "prep post request to /auth/logout")

	res, err = client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err = ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")
	assertEqual(body["ok"], true, "get a JSON API response")

	// Log in with new password
	err = AuthenticateSession(ts, client, "wiggs", "muffins")
	mustEqual(err, nil, "authenticate with new password")

	// Change password back
	data = []byte(`{"password":"cupcakes"}`)
	req, err = http.NewRequest("POST", ts.URL+"/api/user", bytes.NewBuffer(data))
	mustEqual(err, nil, "prep post request to /api/user")

	res, err = client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err = ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["ok"], true, "get a JSON API response")
	assertEqual(body["message"], "password changed", "successfully change password")
}