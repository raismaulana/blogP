package util

import (
	"encoding/json"

	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator

// MustJSON is converter from interface{} to string
// Warning! this function will always assume the convertion is success
// if you are not sure the convertion is always succeed then use ToJSON
func MustJSON(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}
