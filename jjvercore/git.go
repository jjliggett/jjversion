package jjvercore

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type gitServiceInterface interface {
	openRepository(path string) (*git.Repository, error)
	getRepositoryHead(r *git.Repository) (*plumbing.Reference, error)
	getRepositoryTags(r *git.Repository) (storer.ReferenceIter, error)
	getRepositoryTagObjects(r *git.Repository) (*object.TagIter, error)
	getRepositoryCommits(r *git.Repository) (object.CommitIter, error)
}

type gitService struct{}

func (gs gitService) openRepository(path string) (*git.Repository, error) {
	return git.PlainOpen(path)
}

func (gs gitService) getRepositoryHead(r *git.Repository) (*plumbing.Reference, error) {
	return r.Head()
}

func (gs gitService) getRepositoryTags(r *git.Repository) (storer.ReferenceIter, error) {
	return r.Tags()
}

func (gs gitService) getRepositoryTagObjects(r *git.Repository) (*object.TagIter, error) {
	return r.TagObjects()
}

func (gs gitService) getRepositoryCommits(r *git.Repository) (object.CommitIter, error) {
	return r.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
}
