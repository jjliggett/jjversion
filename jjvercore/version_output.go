package jjvercore

import (
	"fmt"
)

type VersionOutput struct {
	Major uint64
	Minor uint64
	Patch uint64
	Sha   string
}

func (v *VersionOutput) Json() string {
	return fmt.Sprintf("{\n  \"Major\": %d,\n  \"Minor\": %d,\n  \"Patch\": %d,\n  \"MajorMinorPatch\": \"%d.%d.%d\",\n  \"Sha\": \"%s\",\n  \"ShortSha\": \"%s\"\n}", v.Major, v.Minor, v.Patch, v.Major, v.Minor, v.Patch, v.Sha, v.Sha[:7])
}
