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

	err := AuthenticateSession(ts, client, User1, Pass1)
	mustEqual(err, nil, "authenticate session")

	data := []byte(`{"city":"Test City","state":"Test State"}`)
	req, err := http.NewRequest("POST", ts.URL+"/api/user", bytes.NewBuffer(data))
	mustEqual(err, nil, "prep POST request to /api/user")

	res, err := client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err := ParseBody(res)
	mustEqual(err, nil, "encode POST response in JSON")

	assertEqual(body["ok"], true, "good request")
	assertEqual(body["isLocationUpdated"], true, "updated location")

	req, err = http.NewRequest("GET", ts.URL+"/api/user/"+User1, nil)
	mustEqual(err, nil, "prep GET request to /api/user")

	res, err = client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err = ParseBody(res)
	mustEqual(err, nil, "encode GET response in JSON")

	assertEqual(body["username"], User1, "recognize username")
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
	ts := httptest.NewServer(app.Handler)
	defer app.Db.Close()
	defer ts.Close()

	err := AuthenticateSession(ts, client, User3, Pass3)
	mustEqual(err, nil, "authenticate session")

	newPass := "testpass3new"

	// Update password
	data := []byte(`{"password":"` + newPass + `"}`)
	req, err := http.NewRequest("POST", ts.URL+"/api/user", bytes.NewBuffer(data))
	mustEqual(err, nil, "prep post request to /api/user")

	res, err := client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err := ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["ok"], true, "get a JSON API response")
	assertEqual(body["isPasswordUpdated"], true, "successfully change password")

	// Logout
	req, err = http.NewRequest("POST", ts.URL+"/auth/logout", nil)
	mustEqual(err, nil, "prep post request to /auth/logout")

	res, err = client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err = ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")
	assertEqual(body["ok"], true, "get a JSON API response")

	// Log in with new password
	err = AuthenticateSession(ts, client, User3, newPass)
	mustEqual(err, nil, "authenticate with new password")

	// Change password back
	data = []byte(`{"password":"` + Pass3 + `"}`)
	req, err = http.NewRequest("POST", ts.URL+"/api/user", bytes.NewBuffer(data))
	mustEqual(err, nil, "prep post request to /api/user")

	res, err = client.Do(req)
	mustEqual(err, nil, "reach server")

	body, err = ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["ok"], true, "get a JSON API response")
	assertEqual(body["isPasswordUpdated"], true, "successfully change password")
}

func TestGetUser(t *testing.T) {
	assertEqual := MakeAssertEqual(t)
	mustEqual := MakeMustEqual(t)

	// Start test server
	app := config.InitializeApp()
	ts := httptest.NewServer(app.Handler)
	defer app.Db.Close()
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/user/" + User1)
	mustEqual(err, nil, "execute GET request to /api/user")
	assertEqual(res.StatusCode, 200, "finds path")

	body, err := ParseBody(res)
	mustEqual(err, nil, "encode response in JSON")

	assertEqual(body["ok"], true, "succeed request")
	username, ok := body["username"]
	assertEqual(ok, true, "username exists")
	assertEqual(username, User1, "correct username")
	_, ok = body["city"]
	assertEqual(ok, true, "city exists")
	_, ok = body["state"]
	assertEqual(ok, true, "state exists")
}
