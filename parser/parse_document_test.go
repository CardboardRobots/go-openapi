package parser

import "testing"

func TestGetPropertyName(t *testing.T) {
	value := GetPropertyName("")
	if value != "" {
		t.Errorf("%v", value)
	}

	value = GetPropertyName("ab")
	if value != "Ab" {
		t.Errorf("%v", value)
	}
}
