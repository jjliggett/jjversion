package main

import (
	"fmt"
	"os"
)

var oss osInterface
var gs gitServiceInterface

func init() {
	oss = osService{}
	gs = gitService{}
}

func calculate_version() {
	path, err := oss.getwd()
	CheckIfError(err)

	vs := getVersionSettings()

	r, err := gs.openRepository(path)
	CheckIfError(err)

	ref, err := gs.getRepositoryHead(r)
	CheckIfError(err)

	if ref.Name().IsBranch() {
		branchVersion, success := GetVersionFromReleaseBranch(ref.Name().Short())
		if success {
			fmt.Println(branchVersion.Json(ref.Hash().String()))
			os.Exit(0)
		}
	}

	tagrefs, err := gs.getRepositoryTags(r)
	CheckIfError(err)

	tags, err := gs.getRepositoryTagObjects(r)
	CheckIfError(err)

	versions := getVersionTags(tagrefs, tags)

	commit_version, ok := versions[ref.Hash().String()]
	if ok {
		fmt.Println(commit_version.Json(ref.Hash().String()))
		os.Exit(0)
	}

	cIter, err := gs.getRepositoryCommits(r)
	CheckIfError(err)

	v := version{0, 0, 0}

	v = incrementVersionByCommitMessages(cIter, versions, vs, v)

	fmt.Println(v.Json(ref.Hash().String()))

}
