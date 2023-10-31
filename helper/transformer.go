package helper

import (
	"fmt"
	"reflect"
)

func StructToMap(data interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		switch field.Kind() {
		case reflect.String:
			result[fieldName] = field.String()
		case reflect.Float64:
			result[fieldName] = fmt.Sprintf("%.2f", field.Float())
		}
	}

	return result
}
