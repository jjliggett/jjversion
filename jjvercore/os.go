package jjvercore

import "os"

type osInterface interface {
	getwd() (dir string, err error)
}

type osService struct{}

func (oss osService) getwd() (dir string, err error) {
	return os.Getwd()
}
