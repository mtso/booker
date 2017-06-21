package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseResponseBody(r *http.Response) (js map[string]interface{}, err error) {
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

// BodyParser?
func ParseRequestBody(r *http.Request) (m map[string]interface{}) {
	decoder := json.NewDecoder(r.Body)
	var raw interface{}

	err := decoder.Decode(&raw)
	if err != nil {
		log.Println(err)
		return
	}

	m = raw.(map[string]interface{})
	return
}
