package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtso/booker/server"
	"github.com/mtso/booker/server/controllers"
	"github.com/mtso/booker/server/utils"
)

func TestTest(t *testing.T) {
	ts := httptest.NewServer(controllers.Root)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte("hello~"))

	res, err := http.Post(ts.URL+"/test", "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := utils.ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	if !resp["ok"].(bool) {
		t.Error("expected response ok to be true")
	}
}

func TestApp(t *testing.T) {
	assertEqual := MakeAssertEqual(t)
	app := main.InitializeApp()
	defer app.Db.Close()

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte(`{"username":"wiggs","password":"cupcakes"}`))
	res, err := http.Post(ts.URL+"/auth/login", "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := utils.ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(err, nil, "No error in parsing JSON")
	assertEqual(res.StatusCode, 200, "Status code 200 for login")
	assertEqual(resp["ok"], true, "Response is ok")
	assertEqual(resp["message"], "wiggs logged in.", "Login message matches correct username")

	cookies := MapCookies(res.Cookies())
	sess_cookie := cookies["sess_id"]

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

	assertEqual(res.StatusCode, 200, "Status 200 for /testroute")

	resp, err = utils.ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(resp["ok"], true, "Response is ok")
	assertEqual(resp["message"], "wiggs is logged into redirecting endpoint", "Saves cookie session")
}
