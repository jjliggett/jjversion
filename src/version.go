package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type version struct {
	Major uint64
	Minor uint64
	Patch uint64
}

func (v *version) IncrementMajor() {
	if v.Major >= 1 {
		v.Major = v.Major + 1
		v.Minor = 0
	} else {
		v.Minor = v.Minor + 1
	}
	v.Patch = 0
}

func (v *version) IncrementMinor() {
	v.Minor = v.Minor + 1
	v.Patch = 0
}

func (v *version) IncrementPatch() {
	v.Patch = v.Patch + 1
}

func (value1 *version) IsGreaterThan(value2 version) bool {
	if value1.Major > value2.Major {
		return true
	}
	if value1.Major == value2.Major && value1.Minor > value2.Minor {
		return true
	}
	if value1.Major == value2.Major && value1.Minor == value2.Minor && value1.Patch > value2.Patch {
		return true
	}
	return false
}

func (v *version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *version) Json(hash string) string {
	return fmt.Sprintf("{\n  \"Major\": %d,\n  \"Minor\": %d,\n  \"Patch\": %d,\n  \"MajorMinorPatch\": \"%d.%d.%d\",\n  \"Sha\": \"%s\",\n  \"ShortSha\": \"%s\"\n}", v.Major, v.Minor, v.Patch, v.Major, v.Minor, v.Patch, hash, hash[:7])
}

func GetVersion(s string) (version, bool) {
	expression := `^(v|V)(?P<major>0|([1-9][0-9]*)).(?P<minor>0|([1-9][0-9]*)).(?P<patch>0|([1-9][0-9]*))$`
	return getVersion(s, expression)
}

func GetVersionFromReleaseBranch(shortBranchName string) (version, bool) {
	expression := `^release\/(?P<major>0|([1-9][0-9]*)).(?P<minor>0|([1-9][0-9]*)).(?P<patch>0|([1-9][0-9]*))$`
	return getVersion(shortBranchName, expression)
}

func getVersion(s string, expression string) (version, bool) {
	e := regexp.MustCompile(expression)
	match := e.FindStringSubmatch(s)
	if match == nil {
		return version{}, false
	}
	versionParts := make(map[string]string)
	for i, name := range e.SubexpNames() {
		if i != 0 && name != "" {
			versionParts[name] = match[i]
		}
	}
	major, err := strconv.ParseUint(versionParts["major"], 10, 64)
	if err != nil {
		return version{}, false
	}
	minor, err := strconv.ParseUint(versionParts["minor"], 10, 64)
	if err != nil {
		return version{}, false
	}
	patch, err := strconv.ParseUint(versionParts["patch"], 10, 64)
	if err != nil {
		return version{}, false
	}
	v := version{major, minor, patch}
	return v, true
}
