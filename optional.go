package confuso

import (
	"reflect"
)

type Optional[T any] struct {
	Value T
	Ok    bool
}

func (o Optional[T]) isOptional() bool {
	return true
}

type optional interface {
	isOptional() bool
}

func isOptional(field reflect.Value) bool {
	optionalInterface := reflect.TypeOf((*optional)(nil)).Elem()
	return field.Type().Implements(optionalInterface)
}
