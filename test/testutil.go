package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func MapCookies(cookies []*http.Cookie) map[string]*http.Cookie {
	m := make(map[string]*http.Cookie)
	for _, c := range cookies {
		m[c.Name] = c
	}
	return m
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
