package confuso

import (
	"reflect"
)

// Optional is a wrapper type that indicates a Value may or may not be present.
type Optional[T any] struct {
	// Value holds the actual Value of type T. It is only valid if Ok is true.
	// If Ok is false, Value is set to the zero Value and should not be accessed.
	Value T
	// Ok indicates whether the Value is present (true) or not (false).
	Ok bool
}

func (o Optional[T]) isOptional() bool {
	return true
}

// Val returns the Value and a boolean indicating whether the Value is present
// (Ok is true) or not (Ok is false).
func (o Optional[T]) Val() (T, bool) {
	return o.Value, o.Ok
}

// MustVal is like Val but panics if the Value is not present
// (Ok is false). It should only be used when the caller is certain that the
// Value is present.
func (o Optional[T]) MustVal() T {
	if o.Ok {
		return o.Value
	}
	panic("Optional Value is not present")
}

// Or returns the Value if it is present (Ok is true), or the provided defaultValue
// if it is not present (Ok is false).
func (o Optional[T]) Or(defaultValue T) T {
	if o.Ok {
		return o.Value
	}
	return defaultValue
}

func (o Optional[T]) String() string {
	if o.Ok {
		return "Optional(" + reflect.ValueOf(o.Value).String() + ")"
	}
	return "Optional(empty)"
}

// isOptional checks if the given reflect.Value implements the optional interface,
// which indicates that it is an Optional type.
type optional interface {
	isOptional() bool
}

func isOptional(field reflect.Value) bool {
	optionalInterface := reflect.TypeOf((*optional)(nil)).Elem()
	return field.Type().Implements(optionalInterface)
}
