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
		t.Fatalf(`GetVersion(text) = %#v %v, want match for version {0, 42, 2}, true`, result, success)
	}
}

func TestIncrementMajorInitialVersion(t *testing.T) {
	v := version{0, 2, 1}
	v.IncrementMajor()
	if v.Major != 0 || v.Minor != 3 || v.Patch != 0 {
		t.Fatalf(`version {0, 2, 1} .IncrementMajor() = %#v, want match for version {0, 3, 0}`, v)
	}
}

func TestIncrementMajor(t *testing.T) {
	v := version{1, 2, 3}
	v.IncrementMajor()
	if v.Major != 2 || v.Minor != 0 || v.Patch != 0 {
		t.Fatalf(`version {1, 2, 3} .IncrementMajor() = %#v, want match for version {2, 0, 0}`, v)
	}
}

func TestIncrementMinor(t *testing.T) {
	v := version{1, 2, 3}
	v.IncrementMinor()
	if v.Major != 1 || v.Minor != 3 || v.Patch != 0 {
		t.Fatalf(`version {1, 2, 3} .IncrementMinor() = %#v, want match for version {1, 3, 0}`, v)
	}
}

func TestIncrementPatch(t *testing.T) {
	v := version{1, 2, 3}
	v.IncrementPatch()
	if v.Major != 1 || v.Minor != 2 || v.Patch != 4 {
		t.Fatalf(`version {1, 2, 3} .IncrementPatch() = %#v, want match for version {1, 2, 4}`, v)
	}
}
