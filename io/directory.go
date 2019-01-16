package io

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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

//Exists return whether the dir existed
func (dir *Directory) Exists() bool {
	return dir.FileInfo != nil
}

func (dir *Directory) Contains(name string) (bool, error) {
	_, err := os.Stat(path.Join(dir.fullName, name))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}

	return true, nil
}

func (dir *Directory) Parent() (*Directory, error) {
	fmt.Println(path.Join(dir.fullName, ".."))
	return nil, nil
}

func (dir *Directory) Root() (*Directory, error) {
	return nil, nil
}

func NewDirectory(dirPath string) (*Directory, error) {
	dir := new(Directory)
	fullName, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	dir.fullName = fullName

	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return dir, nil
		}
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("argument path must be directory")
	}

	dir.FileInfo = fileInfo
	return dir, nil
}

func (dir *Directory) Create(perm os.FileMode) error {
	if dir.Exists() {
		return nil
	}

	err := os.Mkdir(dir.fullName, perm)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(dir.fullName)
	if err != nil {
		return err
	}

	dir.FileInfo = fileInfo
	return nil
}

func (dir *Directory) CreateFile(name string) (*FileInfo, error) {
	return NewFile(filepath.Join(dir.fullName, name))
}

func newDirectory(path string, fileInfo os.FileInfo) *Directory {
	dir := new(Directory)

	dir.FileInfo = fileInfo
	dir.fullName = path

	return dir
}

func (dir *Directory) CreateSubDirectory(name string, perm os.FileMode) error {
	dirPath := filepath.Join(dir.fullName, name)
	return os.Mkdir(dirPath, perm)
}

func (dir *Directory) Delete() error {
	err := os.RemoveAll(dir.fullName)
	if err != nil {
		return err
	}

	dir.FileInfo = nil
	return nil
}

func (dir *Directory) RemoveSubDirectory(names ...string) error {
	if len(names) == 0 {
		return nil
	}

	for _, name := range names {
		dirPath := filepath.Join(dir.fullName, name)

		fileInfo, err := os.Stat(dirPath)
		if err != nil && os.IsNotExist(err) {
			return err
		}

		if !fileInfo.IsDir() {
			return fmt.Errorf("argument %s must be a directory", name)
		}

		err = os.Remove(dirPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dir *Directory) RemoveAllSubDirectory() error {
	return filepath.Walk(dir.fullName, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return os.Remove(path)
		}

		return nil
	})
}

func (dir *Directory) RemoveAll() error {
	mode := dir.Mode()
	err := os.RemoveAll(dir.fullName)
	if err != nil {
		return err
	}

	return os.Mkdir(dir.fullName, mode)
}

func (dir *Directory) GetDirectory(name string) (*Directory, error) {
	fullName := path.Join(dir.fullName, name)
	return NewDirectory(fullName)
}

func (dir *Directory) GetDirectoriesLike(name string) []*Directory {
	return nil
}

//GetDirectories return all sub directories
func (dir *Directory) GetDirectories() ([]*Directory, error) {
	dirs := make([]*Directory, 0)
	fileInfos, err := ioutil.ReadDir(dir.fullName)

	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			dirs = append(dirs, newDirectory(filepath.Join(dir.fullName, fileInfo.Name()), fileInfo))
		}
	}

	if err != nil {
		return nil, err
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() < dirs[j].Name()
	})

	return dirs, nil
}

func (dir *Directory) GetAllDirectories() []*Directory {
	return nil
}

//GetFiles return all sub files
func (dir *Directory) GetFiles() ([]*FileInfo, error) {
	files := make([]*FileInfo, 0)
	fileInfos, err := ioutil.ReadDir(dir.fullName)

	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() {
			files = append(files, newFile(filepath.Join(dir.fullName, fileInfo.Name()), fileInfo))
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files, err
}

func (dir *Directory) GetFile(name string) (*FileInfo, error) {
	fullPath := filepath.Join(dir.fullName, name)
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}

	if fileInfo.Mode().IsRegular() {
		return newFile(fullPath, fileInfo), nil
	}

	return nil, fmt.Errorf("path %s should be a file", name)
}

//GetAllFiles return all files include sub directory in the directory tree
func (dir *Directory) GetAllFiles() ([]*FileInfo, error) {
	files := make([]*FileInfo, 0)
	err := filepath.Walk(dir.fullName, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			files = append(files, newFile(path, info))
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

//MoveTo 将目录移动到另一个目录
func (dir *Directory) MoveTo(destDirName string) bool {
	return true
}

func (dir *Directory) SetAccessControl(perm os.FileMode) error {
	return os.Chmod(dir.fullName, perm)
}
