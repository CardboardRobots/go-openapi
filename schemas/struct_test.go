package schemas

import "testing"

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
