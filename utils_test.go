package confuso

import (
	"reflect"
	"testing"
)

type setFieldTest struct {
	field reflect.Value
}

func assertNotPanic[T any](t *testing.T, f func() T) T {
	defer func(t *testing.T) {
		if err := recover(); err != nil {
			t.Errorf("it is not ok to panic")
		}
	}(t)

	return f()
}

func assertPanic[T any](t *testing.T, f func() T) T {
	defer func(t *testing.T) {
		if err := recover(); err == nil {
			t.Errorf("it is not ok to not panic")
		}
	}(t)

	return f()
}

func assertNotError(t *testing.T, f func() error) {
	err := f()
	if err != nil {
		t.Errorf("got unexpected error: %v", err)
	}
}

func assertError(t *testing.T, f func() error) {
	err := f()
	if err == nil {
		t.Errorf("expected error but got nothing")
	}
}

func assertEquals[T comparable](t *testing.T, expected, got T) {
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestSetBool_true(t *testing.T) {
	variable := false
	field := reflect.ValueOf(&variable).Elem()

	assertNotError(t, func() error { return setBool(field, "true") })
	assertNotPanic(t, func() bool { return field.Bool() == true })
	assertEquals(t, true, variable)
}

func TestSetBool_false(t *testing.T) {
	variable := true
	field := reflect.ValueOf(&variable).Elem()

	assertNotError(t, func() error { return setBool(field, "false") })
	assertNotPanic(t, func() bool { return field.Bool() == false })
	assertEquals(t, false, variable)
}

func TestSetBool_invalid(t *testing.T) {
	variable := true
	field := reflect.ValueOf(&variable).Elem()

	assertError(t, func() error { return setBool(field, "invalid") })
}

func TestSetStr_ok(t *testing.T) {
	variable := "test"
	field := reflect.ValueOf(&variable).Elem()
	n := "some content"

	assertNotError(t, func() error { return setStr(field, n) })
	assertNotPanic(t, func() bool { return field.String() == n })
	assertEquals(t, n, variable)
}

func TestSetInt_ok(t *testing.T) {
	variable := int64(42)
	field := reflect.ValueOf(&variable).Elem()
	n := int64(43)

	assertNotError(t, func() error { return setInt(field, "43") })
	assertNotPanic(t, func() bool { return field.Int() == n })
	assertEquals(t, n, variable)
}

func TestSetInt_invalid(t *testing.T) {
	variable := 42
	field := reflect.ValueOf(&variable).Elem()

	assertError(t, func() error { return setInt(field, "aaaa") })
}

func TestSetInt_negative(t *testing.T) {
	variable := int64(42)
	field := reflect.ValueOf(&variable).Elem()
	n := int64(-42)

	assertNotError(t, func() error { return setInt(field, "-42") })
	assertNotPanic(t, func() bool { return field.Int() == n })
	assertEquals(t, n, variable)
}

func TestMatchEnvVar_ok(t *testing.T) {
	type Test struct {
		in  string
		out string
	}
	tests := []Test{
		{"${CIAO}", "CIAO"},
		{"${NUM83R2}", "NUM83R2"},
		{"${UNDER_SCORE}", "UNDER_SCORE"},
		{"${lowercase}", "lowercase"},
		{"${_}", "_"},
		{"$invalid", ""},
		{"invalid", ""},
		{"${NOT VALID}", ""},
		{"${VALID?}", ""},
		{"${INVALID!}", ""},
		{"", ""},
		{"${}", ""},
	}

	for _, test := range tests {
		assertEquals(t, matchEnvVar(test.in), test.out)
	}
}
