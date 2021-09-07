package main

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func getVersionTags(tagrefs storer.ReferenceIter, tags *object.TagIter) map[string]version {
	versions := make(map[string]version)

	err := tagrefs.ForEach(func(t *plumbing.Reference) error {
		v, ok := GetVersion((t.Name().Short()))
		if ok {
			hash := t.Hash().String()
			if previouslyFoundVersion, ok := versions[hash]; ok {
				if !previouslyFoundVersion.IsGreaterThan(v) {
					versions[hash] = v
				}
			} else {
				versions[hash] = v
			}
		}
		return nil
	})
	CheckIfError(err)

	err = tags.ForEach(func(t *object.Tag) error {
		v, ok := GetVersion((t.Name))
		if ok {
			hash := t.Target.String()
			if previouslyFoundVersion, ok := versions[hash]; ok {
				if !previouslyFoundVersion.IsGreaterThan(v) {
					versions[hash] = v
				}
			} else {
				versions[hash] = v
			}
		}
		return nil
	})
	CheckIfError(err)

	return versions

}
