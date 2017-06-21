package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mtso/booker/server/config"
)

func TestLoginLogout(t *testing.T) {
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

	// Login
	buf := BufferUser(User1, Pass1)
	req, err := http.NewRequest("POST", ts.URL+"/auth/login", buf)
	mustEqual(err, nil, "make request: POST /auth/login")
	req.Header["Content-Type"] = append(req.Header["Content-Type"], "application/json")

	res, err := client.Do(req)
	assertEqual(err, nil, "login test user")
	assertEqual(res.StatusCode, 200, "")

	// Save session cookie
	sess_cookie := FilterCookies(res.Cookies(), func(c *http.Cookie) bool {
		return c.Name == "sess_id"
	})
	wu, err := url.Parse(ts.URL)
	mustEqual(err, nil, "Parse testserver's URL")
	client.Jar.SetCookies(wu, sess_cookie)

	// Assert login response body
	resp, err := ParseBody(res)
	mustEqual(err, nil, "login response is JSON encoded")
	assertEqual(resp["ok"], true, "login is ok")
	assertEqual(resp["message"], User1+" logged in.", "Login message matches correct username")

	// Logout with cookie
	req, err = http.NewRequest("POST", ts.URL+"/auth/logout", nil)
	mustEqual(err, nil, "make request: POST /auth/logout")

	res, err = client.Do(req)
	assertEqual(err, nil, "logout test user")

	// Assert logout response body
	assertEqual(res.StatusCode, 200, "")

	resp, err = ParseBody(res)
	mustEqual(err, nil, "logout response is JSON encoded")

	assertEqual(resp["ok"], true, "")
	assertEqual(resp["message"], User1+" logged out.", "logout message matches correct username")
}
