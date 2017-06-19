package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"net/http"
	// "net/url"
	"net/http/httptest"
	// "path"
	"bytes"

	"github.com/mtso/booker/server/controllers"
)

func parseBody(r *http.Response) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}

	return v.(map[string]interface{})
}

func TestTest(t *testing.T) {
	ts := httptest.NewServer(controllers.Root)
	defer ts.Close()

	buf := bytes.NewBuffer([]byte("hello~"))

	res, err := http.Post(ts.URL + "/test", "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	resp := parseBody(res)

	if !resp["ok"].(bool) {
		t.Error("expected response ok to be true")
	}
}

// func TestApp(t *testing.T) {
// 	app := httptest.NewServer(controllers.Root)
// 	defer app.Close()

// 	cred := struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}{
// 		"wiggs",
// 		"cupcakes",
// 	}

// 	raw, err := json.Marshal(cred)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// u, err := url.Parse(app.URL)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// base := app.URL // http://127.0.0.1:xxxxx hostname

// 	// t.Error(app.URL)

// 	// u.Path = "/auth/login"

// 	buf := bytes.NewBuffer(raw)
// 	// t.Error(u.String())
// 	// res, err := http.Post(u.String(), "application/json", buf)
// 	res, err := http.Post(app.URL + "/auth/login", "application/json", buf)
// 	if err != nil {
// 		// t.Fatalf("%v", err)
// 		t.Fatal(err)
// 	}
// 	// t.Error(res)

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Error(body)

// 	// res, err := http.Get(app.URL)
// 	// t.Error(res)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	// body, err := ioutil.ReadAll(res.Body)
// 	// res.Body.Close()

// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }

// 	// t.Error(string(body))
// }
