package io

import (
	"errors"
	"os"
)

type FileInfo struct {
	os.FileInfo
	fullName string
}

func (file *FileInfo) FullName() string {
	return file.fullName
}

func NewFile(path string) (*FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !fileInfo.Mode().IsRegular() {
		return nil, errors.New("argument path should be a file path")
	}

	return newFile(path, fileInfo), nil
}

func newFile(path string, fileInfo os.FileInfo) *FileInfo {
	f := new(FileInfo)
	f.FileInfo = fileInfo
	f.fullName = path

	return f
}
