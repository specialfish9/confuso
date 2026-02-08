package confuso

import (
	"fmt"
	"reflect"
	"regexp"
)

func setField(fieldName string, field reflect.Value, value any) error {
	val := reflect.ValueOf(value)

	if !val.Type().AssignableTo(field.Type()) {
		return fmt.Errorf("cannot assign value '%v' to field %s", value, fieldName)
	}

	field.Set(val)
	return nil
}

func setOptionalField(fieldName string, optField reflect.Value, value any) error {
	valueField := optField.FieldByName("Value")
	okField := optField.FieldByName("Ok")

	val := reflect.ValueOf(value)

	if !val.Type().AssignableTo(valueField.Type()) {
		return fmt.Errorf("cannot assign value '%v' to optional field %s", value, fieldName)
	}

	valueField.Set(val)
	okField.SetBool(true)
	return nil
}

func matchEnvVar(s string) string {
	r := regexp.MustCompile("\\${([a-zA-Z_0-9]+)}")
	submatches := r.FindStringSubmatch(s)
	if len(submatches) >= 2 {
		return submatches[1]
	}
	return ""
}
