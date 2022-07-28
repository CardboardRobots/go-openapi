package schemas

import (
	"testing"
)

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
