package main

import (
	"regexp"
	"testing"
)

func TestStringSuccess(t *testing.T) {
	v := version{0, 10, 1}
	value := v.String()
	want := regexp.MustCompile(`\b` + "0.10.1" + `\b`)
	if !want.MatchString(value) {
		t.Fatalf(`String() = %q, want match for %#q`, value, want)
	}
}

func TestGetVersionSuccess(t *testing.T) {
	text := "v0.42.2"
	result, success := GetVersion(text)
	if success == false || result.Major != 0 || result.Minor != 42 || result.Patch != 2 {
		t.Fatalf(`GetVersion(text) = %q %v, want match for Version { 0, 42, 2 }, nil`, result, true)
	}
}
