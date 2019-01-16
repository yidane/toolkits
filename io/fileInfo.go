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

func (file *FileInfo) Exists() bool {
	return file.FileInfo != nil
}

func NewFile(path string) (*FileInfo, error) {
	fi := new(FileInfo)
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			//TODO:应该需要判断位置是否合法
			fi.fullName = path
			return fi, nil
		}

		return nil, err
	}

	if !fileInfo.Mode().IsRegular() {
		return nil, errors.New("argument path should be a file path")
	}

	fi.fullName = path
	fi.FileInfo = fileInfo
	return fi, nil
}

func newFile(path string, fileInfo os.FileInfo) *FileInfo {
	fi := new(FileInfo)
	fi.fullName = path
	fi.FileInfo = fileInfo

	return fi
}

func (file *FileInfo) Create() error {
	_, err := os.Create(file.fullName)

	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(file.fullName)
	if err != nil {
		return err
	}
	file.FileInfo = fileInfo

	return nil
}

func (file *FileInfo) Delete() error {
	err := os.Remove(file.fullName)
	if os.IsNotExist(err) {
		return nil
	}

	return err
}

func (file *FileInfo) Rename(newName string) bool {
	return true
}

func (file *FileInfo) Move() bool {
	return true
}
