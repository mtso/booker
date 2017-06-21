package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"golang.org/x/net/publicsuffix"
)

// Login helper
func AuthenticateSession(ts *httptest.Server, client *http.Client, user ...string) error {
	username := User1
	password := Pass1
	if len(user) > 1 {
		username = user[0]
		password = user[1]
	}
	buf := BufferUser(username, password)

	// Login
	req, err := http.NewRequest("POST", ts.URL+"/auth/login", buf)
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

func BufferUser(user, pass string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(`{"username":"` + user + `","password":"` + pass + `"}`))
}

func MapCookies(cookies []*http.Cookie) map[string]*http.Cookie {
	m := make(map[string]*http.Cookie)
	for _, c := range cookies {
		m[c.Name] = c
	}
	return m
}

func FilterCookies(cookies []*http.Cookie, cb func(*http.Cookie) bool) []*http.Cookie {
	cc := make([]*http.Cookie, 0)
	for _, c := range cookies {
		if cb(c) {
			cc = append(cc, c)
		}
	}
	return cc
}

func MakeAssertEqual(t *testing.T) func(interface{}, interface{}, string) {
	return func(got, want interface{}, m string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got=%v want=%v", m, got, want)
		}
	}
}

func MakeMustEqual(t *testing.T) func(interface{}, interface{}, string) {
	return func(got, want interface{}, m string) {
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%s: got=%v want=%v", m, got, want)
		}
	}
}

func ParseBody(r *http.Response) (js map[string]interface{}, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var buf interface{}
	err = json.Unmarshal(body, &buf)
	if err != nil {
		return
	}

	js = buf.(map[string]interface{})
	return
}

func MakeCookieMonster() *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}
	return &http.Client{Jar: jar}
}
