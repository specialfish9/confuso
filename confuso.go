package confuso

import (
	"errors"
	"reflect"
)

// LoadConf read the configuration from a file called fileName, parse the
// content and populate the fields of the struct out. It will return any
// error it encounter.
func LoadConf(fileName string, out interface{}) error {
	config, err := readConfig(fileName)

	if err != nil {
		return err
	}

	value := reflect.ValueOf(out).Elem()

	if value.Type().Kind() != reflect.Struct {
		return errors.New("unable to load env because input was not a go structure")
	}

	return populateStruct(value, "", config)
}

func populateStruct(object reflect.Value, namespace string, config map[string]string) error {
	if object.Kind() != reflect.Struct {
		return errors.New("cannot populate a non struct")
	}
	for i := 0; i < object.NumField(); i++ {
		fieldType := object.Type().Field(i)
		fieldValue := object.Field(i)
		fieldName := mkNamespace(namespace, fieldType)
		if fieldType.Type.Kind() == reflect.Struct {
			if err := populateStruct(fieldValue, fieldName, config); err != nil {
				return err
			}
		} else {
			value, ok := config[fieldName]
			if !ok {
				return errors.New("field " + fieldName + " not found!")
			}

			if envVar := matchEnvVar(value); envVar != "" {
				if err := setEnvField(fieldValue, envVar); err != nil {
					return err
				}
			} else {
				if err := setField(object.Field(i), value); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
