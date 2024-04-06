package service

import (
	"reflect"
	"strings"
)

func MapTo(dto interface{}, body map[string]interface{}) {
	v := reflect.ValueOf(dto).Elem()
	typeOfDto := v.Type()
	for i := 0; i < v.NumField(); i++ {
		valueOfDto, ok := body[strings.ToLower(typeOfDto.Field(i).Name)]
		if ok {
			v.Field(i).SetString(valueOfDto.(string))
		}
	}
}
