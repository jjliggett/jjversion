package jjvercore

import (
	"os"
	"testing"
)

func TestGetWD(t *testing.T) {
	expectedDir, expectedErr := os.Getwd()
	oss := osService{}
	dir, err := oss.getwd()
	if dir != expectedDir || err != expectedErr {
		t.Fatalf(`osService.getwd() = %q, %#v, want match for %q, %#v`, dir, err, expectedDir, expectedErr)
	}
}
