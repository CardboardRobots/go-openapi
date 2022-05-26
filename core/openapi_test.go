package core

import (
	"testing"
)

func TestPackage(t *testing.T) {
	testPackage := Package{
		Name: "test",
	}

	value := testPackage.String()

	want := "package test\n\n"
	if value != want {
		t.Errorf("received: %v, expected %v", value, want)
	}
}

func TestStruct(t *testing.T) {
	testStruct := Struct{
		Name: "test",
	}

	value := testStruct.String()

	want := "type test struct {\n}\n\n"
	if value != want {
		t.Errorf("received: %v, expected %v", value, want)
	}
}

func TestType(t *testing.T) {
	testType := Type{
		Name:    "string",
		Pointer: true,
	}

	value := testType.String()

	want := "*string"
	if value != want {
		t.Errorf("received: %v, expected %v", value, want)
	}
}

func TestField(t *testing.T) {
	field := Field{
		Name: "test",
		Type: Type{
			Name:    "string",
			Pointer: true,
		},
		Tag: "",
	}

	value := field.String()

	want := "\ttest *string\n"
	if value != want {
		t.Errorf("received: %v, expected %v", value, want)
	}
}

func TestFullStruct(t *testing.T) {
	testStruct := Struct{
		Name: "Test",
		Fields: []Field{{
			Name: "FirstName",
			Type: Type{
				Name:    "string",
				Pointer: true,
			},
			Tag: Json("firstName"),
		}},
	}

	value := testStruct.String()

	want := "type Test struct {\n\tFirstName *string `json:\"firstName\"`\n}\n\n"
	if value != want {
		t.Errorf("received: %v, expected %v", value, want)
	}
}
