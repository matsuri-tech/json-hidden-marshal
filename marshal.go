package json_hidden_marshal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func scanToMap(val interface{}) (map[string]interface{}, error) {
	strct := reflect.TypeOf(val)
	values := reflect.ValueOf(val)

	// pointerであればderefする
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}

	m := map[string]interface{}{}
	for i := 0; i < strct.NumField(); i++ {
		field := strct.Field(i)
		hidden := field.Tag.Get("hidden")

		// Marshal時にmap[string]interface{}に置き換えるため、失われるjsonタグの対応をする必要がある
		// 現状は名前の変更のみ対応
		jsonTags := strings.Split(field.Tag.Get("json"), ",")
		fieldName := field.Name
		if len(jsonTags) >= 1 {
			if jsonTags[0] != "" {
				fieldName = jsonTags[0]
			}
		}

		if hidden == "" {
			ival := values.Field(i)
			if ival.Kind() == reflect.Struct {
				// PERF: 再帰をやめる
				v, err := scanToMap(ival.Interface())
				if err != nil {
					return nil, err
				}

				m[fieldName] = v
			} else {
				// 構造体でないときはbase typeなのでそのまま
				m[fieldName] = ival.Interface()
			}
		} else if hidden == "-" || hidden == "true" {
			continue
		} else if hidden == "mask" {
			m[fieldName] = strings.Repeat("*", len(values.Field(i).String()))
		} else {
			return nil, errors.New(fmt.Sprintf("unsupported hidden tag: %v", hidden))
		}
	}

	return m, nil
}

func Marshal(val interface{}) ([]byte, error) {
	m, err := scanToMap(val)
	if err != nil {
		return nil, err
	}

	return json.Marshal(m)
}
