package jsonhelpers

import (
	"encoding/json"
	"net/http/httptest"
)

// EmptyObject is an obj that when compared with DeepEqual matches an
// Unmarshaled empty json object
var EmptyObject = make(map[string]interface{})

// EmptyArray is an obj that when compared with DeepEqual matches an
// Unmarshaled empty json array
var EmptyArray = make([]interface{}, 0)

// Unmarshal returns interface from json response object
func Unmarshal(resp *httptest.ResponseRecorder) (obj interface{}) {
	json.Unmarshal(resp.Body.Bytes(), &obj)
	return
}

// Marshal returns a *bytes.Buffer from a marshalled []byte
func Marshal(obj interface{}) (s string) {
	jsonByte, err := json.Marshal(obj)
	if err != nil {
		// should raise error?
		return
	}
	// why add newline? see: https://github.com/golang/go/issues/7767
	return string(jsonByte[:]) + "\n"
}
