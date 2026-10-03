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

		// Unexported fields can't be set via reflection
		if !fieldType.IsExported() {
			continue
		}

		fieldName := getFieldName(fieldType)
		isOpt := isOptional(fieldValue)

		configValue, ok := config[fieldName]
		if !ok {
			if isOpt {
				continue
			}
			return fmt.Errorf("field %q is missing from provided config", fieldName)
		}

		// Optionals are special structs, so they are handled before the generic path
		if isOpt {
			if err := setOptionalField(fieldName, fieldValue, configValue); err != nil {
				return err
			}
			continue
		}

		if err := populateValue(fieldValue, configValue); err != nil {
			return fmt.Errorf("%q->%w", fieldName, err)
		}
	}

	return nil
}

func populateSlice(object reflect.Value, config []any) error {
	if object.Kind() != reflect.Slice {
		return errors.New("cannot populate a non slice")
	}
	if !object.CanSet() {
		return errors.New("cannot populate a non settable slice (pass an addressable value)")
	}

	// Build a fresh slice of the right length. Elements of a slice are
	// always addressable, so we can populate them in place.
	result := reflect.MakeSlice(object.Type(), len(config), len(config))

	for i, configItem := range config {
		if err := populateValue(result.Index(i), configItem); err != nil {
			return fmt.Errorf("[%d]->%w", i, err)
		}
	}

	object.Set(result)
	return nil
}

// populateValue fills dst (which must be settable) from item, branching on
// the destination's type rather than the config item's type.
func populateValue(dst reflect.Value, item any) error {
	if item == nil {
		return nil // leave the zero value
	}

	// Allocate through pointers so []*T and nested pointers work.
	for dst.Kind() == reflect.Pointer {
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		dst = dst.Elem()
	}

	switch dst.Kind() {
	case reflect.Struct:
		subConfig, ok := item.(map[string]any)
		if !ok {
			return fmt.Errorf("expected object for %s but got %T (%v)", dst.Type(), item, item)
		}
		return populateStruct(dst, subConfig)

	case reflect.Slice:
		sliceConfig, ok := item.([]any)
		if !ok {
			return fmt.Errorf("expected list for %s but got %T (%v)", dst.Type(), item, item)
		}
		return populateSlice(dst, sliceConfig)

	default:
		if err := setField(dst, item); err != nil {
			return fmt.Errorf("setting %s: %w", dst.Type(), err)
		}
		return nil
	}
}

func getFieldName(field reflect.StructField) string {
	if tag := field.Tag.Get("confuso"); tag != "" {
		return tag
	}
	return field.Name
}
