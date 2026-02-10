package confuso

import (
	"errors"
	"fmt"
	"reflect"
)

type Input interface {
	read() (map[string]any, error)
}

func Do(fileName string, out any) error {
	input := NewYAMLInput(fileName)
	config, err := input.read()
	if err != nil {
		return fmt.Errorf("confuso: %w", err)
	}

	value := reflect.ValueOf(out).Elem()

	if value.Type().Kind() != reflect.Struct {
		return errors.New("confuso: input value is not a Go struct")
	}

	if err := populateStruct(value, config); err != nil {
		return fmt.Errorf("confuso: %w", err)
	}

	return nil
}

func populateStruct(object reflect.Value, config map[string]any) error {
	if object.Kind() != reflect.Struct {
		return errors.New("cannot populate a non struct")
	}

	for i := range object.NumField() {
		fieldType := object.Type().Field(i)
		fieldValue := object.Field(i)

		fieldName := getFieldName(fieldType)

		// Check if the value is optional
		isOpt := isOptional(fieldValue)

		configValue, ok := config[fieldName]
		if !ok {
			if isOpt {
				continue
			} else {
				return fmt.Errorf("field '%s' is missing from provided config!", fieldName)
			}
		}

		if isOpt {
			// Optionals are special structs, so they must be handled in advance
			if err := setOptionalField(fieldName, fieldValue, configValue); err != nil {
				return err
			}
		} else if fieldType.Type.Kind() == reflect.Struct {
			// If the field is a struct, we expect `configValue` to be a map
			subConfig, ok := configValue.(map[string]any)
			if !ok {
				return fmt.Errorf("unexpected value '%v' provided for field '%s'", configValue, fieldName)
			}
			// Propagate the visit to the rest of the struct
			if err := populateStruct(fieldValue, subConfig); err != nil {
				return err
			}
		} else {
			// Otherwise, we can just try to set the value as is
			if err := setField(fieldName, fieldValue, configValue); err != nil {
				return err
			}
		}
	}
	return nil
}

func getFieldName(field reflect.StructField) string {
	name := ""
	tag := field.Tag.Get("confuso")
	if tag != "" {
		name = tag
	} else {
		name = field.Name
	}
	return name
}
