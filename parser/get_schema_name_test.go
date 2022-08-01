package parser

import "testing"

func TestGetSchemaName(t *testing.T) {
	cases := []struct {
		Component string
		Result    string
	}{{
		Component: "#/components/schemas/Test",
		Result:    "Test",
	}, {
		Component: "#/Components/Schemas/Test",
		Result:    "Test",
	}}
	for _, c := range cases {
		result := GetSchemaName(c.Component)
		if c.Result != result {
			t.Errorf("Case: %v, Received: %v, Expected %v", c.Component, result, c.Result)
		}
	}
}
