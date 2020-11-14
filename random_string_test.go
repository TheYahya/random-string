package randomstring_test

import (
	"github.com/theyahya/random-string"
	"testing"
)

func TestRandomStringLength(t *testing.T) {
	length := 10
	randomString := randomstring.Generate(10)
	expected := len(randomString)
	if length != expected {
		t.Errorf("expected %v", expected)
	}
}
