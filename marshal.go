package json_hidden_marshal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Marshal(val interface{}) ([]byte, error) {
	strct := reflect.TypeOf(val)
	values := reflect.ValueOf(val)

	m := map[string]interface{}{}
	for i := 0; i < strct.NumField(); i++ {
		field := strct.Field(i)
		hidden := field.Tag.Get("hidden")

		// Marshal時にmap[string]interface{}に置き換えるため、失われるjsonタグの対応をする必要がある
		// 現状は名前の変更のみ対応
		jsonTags := strings.Split(field.Tag.Get("json"), ",")
		fieldName := field.Name
		if len(jsonTags) >= 1 {
			fieldName = jsonTags[0]
		}

		if hidden == "" {
			m[fieldName] = values.Field(i).Interface()
		} else if hidden == "-" || hidden == "true" {
			continue
		} else if hidden == "mask" {
			m[fieldName] = strings.Repeat("*", len(values.Field(i).String()))
		} else {
			return nil, errors.New(fmt.Sprintf("unsupported hidden tag: %v", hidden))
		}
	}

	return json.Marshal(m)
}
