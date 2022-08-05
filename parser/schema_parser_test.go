package parser

import (
	"testing"

	"github.com/cardboardrobots/go-openapi/entity"
)

func TestKeyToName(t *testing.T) {
	name := KeyToName(entity.VERB_GET, "/test/{testId}")
	want := "GetTestByTestId"
	if name != want {
		t.Errorf("Received: %v, Expected: %v", name, want)
	}
}
