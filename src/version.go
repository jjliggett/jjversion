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

func (v *version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *version) Json(hash string) string {
	return fmt.Sprintf("{\n  \"Major\": %d,\n  \"Minor\": %d,\n  \"Patch\": %d,\n  \"MajorMinorPatch\": \"%d.%d.%d\",\n  \"Sha\": \"%s\",\n  \"ShortSha\": \"%s\"\n}", v.Major, v.Minor, v.Patch, v.Major, v.Minor, v.Patch, hash, hash[:7])
}

func GetVersion(value string) (version, bool) {
	var expression = regexp.MustCompile(`^(v|V)(?P<major>0|([1-9][0-9]*)).(?P<minor>0|([1-9][0-9]*)).(?P<patch>0|([1-9][0-9]*))$`)
	match := expression.FindStringSubmatch(value)
	if match == nil {
		return version{}, false
	}
	version_parts := make(map[string]string)
	for i, name := range expression.SubexpNames() {
		if i != 0 && name != "" {
			version_parts[name] = match[i]
		}
	}
	major, err := strconv.ParseUint(version_parts["major"], 10, 64)
	if err != nil {
		return version{}, false
	}
	minor, err := strconv.ParseUint(version_parts["minor"], 10, 64)
	if err != nil {
		return version{}, false
	}
	patch, err := strconv.ParseUint(version_parts["patch"], 10, 64)
	if err != nil {
		return version{}, false
	}
	v := version{major, minor, patch}
	return v, true
}

func GetVersionFromReleaseBranch(shortBranchName string) (version, bool) {
	var expression = regexp.MustCompile(`^release\/(?P<major>0|([1-9][0-9]*)).(?P<minor>0|([1-9][0-9]*)).(?P<patch>0|([1-9][0-9]*))$`)
	match := expression.FindStringSubmatch(shortBranchName)
	if match == nil {
		return version{}, false
	}
	version_parts := make(map[string]string)
	for i, name := range expression.SubexpNames() {
		if i != 0 && name != "" {
			version_parts[name] = match[i]
		}
	}
	major, err := strconv.ParseUint(version_parts["major"], 10, 64)
	if err != nil {
		return version{}, false
	}
	minor, err := strconv.ParseUint(version_parts["minor"], 10, 64)
	if err != nil {
		return version{}, false
	}
	patch, err := strconv.ParseUint(version_parts["patch"], 10, 64)
	if err != nil {
		return version{}, false
	}
	v := version{major, minor, patch}
	return v, true
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
