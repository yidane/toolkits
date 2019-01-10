package io

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Directory struct {
	os.FileInfo
	fullName string
}

func (dir *Directory) FullName() string {
	return dir.fullName
}

func NewDirectory(path string) (*Directory, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("argument path must be directory")
	}

	return newDirectory(path, fileInfo)
}

func newDirectory(path string, fileInfo os.FileInfo) (*Directory, error) {
	dir := new(Directory)
	dir.FileInfo = fileInfo
	dir.fullName = path

	return dir, nil
}

func CreateDirectory(path string, perm os.FileMode) (*Directory, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, perm)
		if err != nil {
			return nil, err
		}
	}

	return NewDirectory(path)
}

func (dir *Directory) CreateSubDirectory(name string, perm os.FileMode) error {
	path := filepath.Join(dir.fullName, name)
	return os.Mkdir(path, perm)
}

func (dir *Directory) Delete() error {
	return os.RemoveAll(dir.fullName)
}

func (dir *Directory) RemoveSubDirectory(names ...string) error {
	if len(names) == 0 {
		return nil
	}

	for _, name := range names {
		path := filepath.Join(dir.fullName, name)

		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			return fmt.Errorf("argument %s must be a directory", name)
		}

		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dir *Directory) RemoveAllSubDirectory() error {
	err := filepath.Walk(dir.fullName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return os.Remove(path)
		}

		return nil
	})

	return err
}

func (dir *Directory) RemoveAll() error {
	mode := dir.Mode()
	err := os.RemoveAll(dir.fullName)
	if err != nil {
		return err
	}

	return os.Mkdir(dir.fullName, mode)
}

func (dir *Directory) GetDirectories() ([]*Directory, error) {
	dirs := make([]*Directory, 0)
	err := filepath.Walk(dir.fullName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if os.SameFile(info, dir.FileInfo) {
				return nil
			}

			directory, err := newDirectory(path, info)
			if err != nil {
				return err
			}

			dirs = append(dirs, directory)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() < dirs[j].Name()
	})

	return dirs, nil
}

func (dir *Directory) GetFiles() ([]*FileInfo, error) {
	files := make([]*FileInfo, 0)
	err := filepath.Walk(dir.fullName, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			fileInfo, err := newFile(path, info)
			if err != nil {
				return err
			}

			files = append(files, fileInfo)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files, nil
}
