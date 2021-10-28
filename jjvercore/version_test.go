package jjvercore

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
	s := "v0.42.2"
	v, success := GetVersion(s)
	if v.Major != 0 || v.Minor != 42 || v.Patch != 2 || !success {
		t.Fatalf(`GetVersion("v0.42.2") = %#v %v, want match for version {0, 42, 2}, true`, v, success)
	}
}

func TestGetVersionFailure(t *testing.T) {
	t.Run("not a version", testGetVersionFailureFunc("not a version"))
	t.Run("v99999999999999999999.0.0", testGetVersionFailureFunc("v99999999999999999999.0.0"))
	t.Run("v0.99999999999999999999.0", testGetVersionFailureFunc("v0.99999999999999999999.0"))
	t.Run("v0.0.99999999999999999999", testGetVersionFailureFunc("v0.0.99999999999999999999"))
}

func testGetVersionFailureFunc(s string) func(*testing.T) {
	return func(t *testing.T) {
		v, success := GetVersion(s)
		if v.Major != 0 || v.Minor != 0 || v.Patch != 0 || success {
			t.Fatalf(`GetVersion("%v") = %#v %v, want match for version {0, 0, 0}, false`, s, v, success)
		}
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

func TestGetVersionFromReleaseBranchSuccess(t *testing.T) {
	s := "release/42.10.20"
	v, success := GetVersionFromReleaseBranch(s)
	if v.Major != 42 || v.Minor != 10 || v.Patch != 20 || !success {
		t.Fatalf(`GetVersionFromReleaseBranch("release/42.10.20") = %#v %v, want match for version {42, 10, 20}, true`, v, success)
	}
}

func TestGetVersionFromReleaseBranchFailures(t *testing.T) {
	t.Run("not a release branch", testGetVersionFromReleaseBranchFailureFunc("not a release branch"))
	t.Run("r99999999999999999999.0.0", testGetVersionFromReleaseBranchFailureFunc("release/99999999999999999999.0.0"))
	t.Run("r0.99999999999999999999.0", testGetVersionFromReleaseBranchFailureFunc("release/0.99999999999999999999.0"))
	t.Run("r0.0.99999999999999999999", testGetVersionFromReleaseBranchFailureFunc("release/0.0.99999999999999999999"))
}

func testGetVersionFromReleaseBranchFailureFunc(s string) func(*testing.T) {
	return func(t *testing.T) {
		v, success := GetVersionFromReleaseBranch(s)
		if v.Major != 0 || v.Minor != 0 || v.Patch != 0 || success {
			t.Fatalf(`GetVersionFromReleaseBranch("%v") = %#v %v, want match for version {0, 0, 0}, false`, s, v, success)
		}
	}
}
