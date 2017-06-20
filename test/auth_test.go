package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"reflect"

	"github.com/mtso/booker/server"
	"github.com/mtso/booker/server/controllers"
	"github.com/mtso/booker/server/utils"
)

func MakeAssertEquals(t *testing.T) func(interface{}, interface{}, string) {
	return func(got, want interface{}, m string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got=%v want=%v", m, got, want)
		}
	}
}

func TestTest(t *testing.T) {
	ts := httptest.NewServer(controllers.Root)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte("hello~"))

	res, err := http.Post(ts.URL + "/test", "application/json", buf)
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
	assertEquals := MakeAssertEquals(t)
	app := main.InitializeApp()
	defer app.Db.Close()

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte(`{"username":"wiggs","password":"cupcakes"}`))
	res, err := http.Post(ts.URL + "/auth/login", "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := utils.ParseBody(res)
	if err != nil {
		t.Fatal(err)
	}

	assertEquals(err, nil, "No error in parsing JSON")
	assertEquals(res.StatusCode, 200, "Status code 200 for login")
	assertEquals(resp["ok"], true, "Response is ok")
	assertEquals(resp["message"], "wiggs logged in.", "Login message matches correct username")
}
