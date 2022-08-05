package parser

import "testing"

func TestKeyToName(t *testing.T) {
	name := KeyToName("/test")
	want := "Test"
	if name != want {
		t.Errorf("Received: %v, Expected: %v", name, want)
	}
}
