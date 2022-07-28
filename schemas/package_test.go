package schemas

import "testing"

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
