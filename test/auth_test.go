package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mtso/booker/server"
	"github.com/mtso/booker/server/controllers"
)

func TestTest(t *testing.T) {
	ts := httptest.NewServer(controllers.Root)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte("hello~"))

	res, err := http.Post(ts.URL+"/test", "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	if !resp["ok"].(bool) {
		t.Error("expected response ok to be true")
	}
}

func TestApp(t *testing.T) {
	// Set up assertions
	assertEqual := MakeAssertEqual(t)

	// Start test server
	app := main.InitializeApp()
	defer app.Db.Close()

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	// Login
	buf := bytes.NewBuffer([]byte(`{"username":"wiggs","password":"cupcakes"}`))
	res, err := http.Post(ts.URL+"/auth/login", "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	// Assert response body
	assertEqual(err, nil, "No error in parsing JSON")
	assertEqual(res.StatusCode, 200, "Status code 200 for login")
	assertEqual(resp["ok"], true, "Response is ok")
	assertEqual(resp["message"], "wiggs logged in.", "Login message matches correct username")

	cookies := MapCookies(res.Cookies())
	sess_cookie := cookies["sess_id"]

	// Try testroute
	req, err := http.NewRequest("GET", ts.URL+"/auth/testroute", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(sess_cookie)

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Assert response body
	assertEqual(res.StatusCode, 200, "Status 200 for /testroute")

	resp, err = ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(resp["ok"], true, "Response is ok")
	assertEqual(resp["message"], "wiggs is logged into redirecting endpoint", "Saves cookie session")
}

func TestLoginLogout(t *testing.T) {
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

	// Login
	req, err := http.NewRequest("POST", ts.URL+"/auth/login", bytes.NewBuffer([]byte(`{"username":"wiggs","password":"cupcakes"}`)))
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
	assertEqual(resp["message"], "wiggs logged in.", "Login message matches correct username")

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
	assertEqual(resp["message"], "wiggs logged out.", "logout message matches correct username")
}

func TestLoginTest(t *testing.T) {
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

	err := AuthenticateSession(ts, client)
	mustEqual(err, nil, "authenticate session")

	req, err := http.NewRequest("GET", ts.URL+"/auth/test", nil)
	mustEqual(err, nil, "prep request to /auth/test")
	resp, err := client.Do(req)
	assertEqual(err, nil, "test fine")

	body, err := ioutil.ReadAll(resp.Body)
	mustEqual(err, nil, "validate body")
	assertEqual(string(body), "wiggs is logged in", "verify correct username")
}

func AuthenticateSession(ts *httptest.Server, client *http.Client) error {
	// Login
	req, err := http.NewRequest("POST", ts.URL+"/auth/login", bytes.NewBuffer([]byte(`{"username":"wiggs","password":"cupcakes"}`)))
	if err != nil {
		return err
	}
	req.Header["Content-Type"] = append(req.Header["Content-Type"], "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Save session cookie
	sess_cookie := FilterCookies(res.Cookies(), func(c *http.Cookie) bool {
		return c.Name == "sess_id"
	})
	cookieurl, err := url.Parse(ts.URL)
	if err != nil {
		return err
	}

	client.Jar.SetCookies(cookieurl, sess_cookie)
	return nil
}