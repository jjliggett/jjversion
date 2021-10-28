package jjvercore

import (
	"regexp"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func incrementVersionByCommitMessages(cIter object.CommitIter, versions map[string]version, vs versionSettings, v version) version {
	type actionType int

	const (
		SetVersion actionType = iota
		IncrementMajor
		IncrementMinor
		IncrementPatch
	)

	type action struct {
		atype actionType
		sv    version
	}

	if vs.Commit_Message_Incrementing_Enabled {
		actions := make([]action, 0)
		err := cIter.ForEach(func(c *object.Commit) error {
			commit_version, ok := versions[c.Hash.String()]
			success1, _ := regexp.MatchString(vs.Major_Version_Bump_Message, c.Message)
			success2, _ := regexp.MatchString(vs.Minor_Version_Bump_Message, c.Message)
			success3, _ := regexp.MatchString(vs.Patch_Version_Bump_Message, c.Message)
			if ok {
				var a action = action{
					atype: SetVersion,
					sv:    commit_version,
				}
				actions = append(actions, a)
			} else if success1 {
				actions = append(actions, action{atype: IncrementMajor})
			} else if success2 {
				actions = append(actions, action{atype: IncrementMinor})
			} else if success3 {
				actions = append(actions, action{atype: IncrementPatch})
			}
			return nil
		})
		CheckIfError(err)
		for i := len(actions) - 1; i >= 0; i-- {
			if actions[i].atype == SetVersion {
				v = actions[i].sv
			} else if actions[i].atype == IncrementMajor {
				v.IncrementMajor()
			} else if actions[i].atype == IncrementMinor {
				v.IncrementMinor()
			} else if actions[i].atype == IncrementPatch {
				v.IncrementPatch()
			}
		}
	}

	return v
}
