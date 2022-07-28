package schemas

import "testing"

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
