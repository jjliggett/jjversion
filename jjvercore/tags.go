package jjvercore

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func getVersionTags(tagrefs storer.ReferenceIter, tags *object.TagIter) map[string]version {
	versions := make(map[string]version)

	err := tagrefs.ForEach(func(t *plumbing.Reference) error {
		v, ok := GetVersion(t.Name().Short())
		if ok {
			hash := t.Hash().String()
			if shouldSaveVersionTag(versions, v, hash) {
				versions[hash] = v
			}
		}
		return nil
	})
	CheckIfError(err)

	err = tags.ForEach(func(t *object.Tag) error {
		v, ok := GetVersion(t.Name)
		if ok {
			hash := t.Target.String()
			if shouldSaveVersionTag(versions, v, hash) {
				versions[hash] = v
			}
		}
		return nil
	})
	CheckIfError(err)

	return versions
}

func shouldSaveVersionTag(versions map[string]version, v version, hash string) bool {
	if previouslyFoundVersion, ok := versions[hash]; ok {
		if !previouslyFoundVersion.IsGreaterThan(v) {
			return true
		}
	} else {
		return true
	}
	return false
}
