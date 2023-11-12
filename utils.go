package confuso

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

func setField(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.Int:
		return setInt(field, value)
	case reflect.Bool:
		return setBool(field, value)
	case reflect.String:
		return setStr(field, value)
	}
	return errors.New("Unsupported struct field type: " + field.Kind().String())
}

func setEnvField(field reflect.Value, envName string) error {
	envVar := os.Getenv(envName)
	return setField(field, envVar)
}

func setInt(field reflect.Value, value string) error {
	conv, err := strconv.Atoi(value)

	if err != nil {
		return err
	}

	field.SetInt(int64(conv))
	return nil
}

func setStr(field reflect.Value, value string) error {
	field.SetString(value)
	return nil
}

func setBool(field reflect.Value, value string) error {
	if value != "true" && value != "false" {
		return errors.New(value + " is not a valid boolean!")
	}

	field.SetBool(value == "true")

	return nil
}

func mkNamespace(base string, field reflect.StructField) string {
	name := ""
	tag := field.Tag.Get("confuso")
	if tag != "" {
		name = tag
	} else {
		name = field.Name
	}
	if base == "" {
		return name
	}
	return fmt.Sprintf("%s.%s", base, name)
}

func matchEnvVar(s string) string {
	r := regexp.MustCompile("\\${([a-zA-Z_0-9]+)}")
	submatches := r.FindStringSubmatch(s)
	if len(submatches) >= 2 {
		return submatches[1]
	}
	return ""
}
