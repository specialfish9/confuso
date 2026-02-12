package confuso

import "testing"

const testConfig string = "./test-config/test-config.yaml"

func Test_readYAML(t *testing.T) {
	input := NewYAMLInput(testConfig)
	out, err := input.read()
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	this, ok := out["this"]
	assert(t, ok)

	is, ok := this.(map[string]any)["is"]
	assert(t, ok)

	a, ok := is.(map[string]any)["a"]
	assert(t, ok)

	aMap := a.(map[string]any)
	assert_eq(t, aMap["string"], "present")
	assert_eq(t, aMap["bool"], true)
	assert_eq(t, aMap["number"], 1)

	other, ok := out["other"]
	assert(t, ok)

	object, ok := other.(map[string]any)["object"]
	assert(t, ok)
	assert_eq(t, object, "here")
}

func Test_Do(t *testing.T) {
	out := Config{}

	err := Do(testConfig, &out)
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	assert_eq(t, out.This.Is.A.String, "present")
	assert_eq(t, out.This.Is.A.Bool, true)
	assert_eq(t, out.This.Is.A.Number, 1)
}

func Test_DoWithOptionals(t *testing.T) {
	out := ConfigWithOptional{}

	err := Do(testConfig, &out)
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	assert(t, out.This.Is.A.String.Ok, "expected string to be present")
	assert_eq(t, out.This.Is.A.String.MustVal(), "present")
	assert(t, !out.This.Is.A.OptString.Ok, "expected optString to be absent")

	assert(t, out.This.Is.A.Bool.Ok, "expected bool to be present")
	assert_eq(t, out.This.Is.A.Bool.MustVal(), true)
	assert(t, !out.This.Is.A.OptBool.Ok, "expected optBool to be absent")

	assert(t, out.This.Is.A.Number.Ok, "expected number to be present")
	assert_eq(t, out.This.Is.A.Number.MustVal(), 1)
	assert(t, !out.This.Is.A.OptNumber.Ok, "expected optNumber to be absent")
}

func assert(t *testing.T, pred bool, message ...string) {
	if !pred {
		t.Fatal("assertion failed", message[0])
	}
}

func assert_eq(t *testing.T, got, expected any) {
	if got != expected {
		t.Fatalf("expected: `%v`, but got `%v`", expected, got)
	}
}
