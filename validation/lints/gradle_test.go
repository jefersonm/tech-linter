package lints

import (
	"testing"
)

func TestHasSpecificDependency(t *testing.T) {
	if HasSpecificDependency("Junit") != true {
		t.Error("Junit dependency not found")
	}
}