package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

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

