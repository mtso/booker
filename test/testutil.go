package test

import (
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
