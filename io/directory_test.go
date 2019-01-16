package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const dirName = "toolkit"

var basePath = filepath.Join(os.TempDir(), dirName)

func TestDirectory_GetDirectories(t *testing.T) {
	dir, err := NewDirectory(basePath)
	if err != nil {
		t.Fatal(err)
	}

	err = dir.Create(os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal()

	f, err := dir.Contains("1")
	if err != nil {
		t.Fatal(err)
	}

	if f {
		dir.RemoveSubDirectory("1")
	}

	dir.CreateSubDirectory("1", os.ModePerm)

	fmt.Println(dir.Name())

	dirs, err := dir.GetDirectories()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(dirs))
	for _, v := range dirs {
		fmt.Println(v.FullName())
	}
}

func TestDirectory_GetFiles(t *testing.T) {
	dir, err := NewDirectory(basePath)
	if err != nil {
		t.Fatal(err)
	}

	if dir.Exists() {
		err = dir.Delete()
		if err != nil {
			t.Fatal(err)
		}
	}

	dir.Create(os.ModePerm)
	files, err := dir.GetFiles()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range files {
		fmt.Println(v.FullName())
	}

	fs, _ := ioutil.ReadDir(basePath)
	fmt.Println(len(fs))
	for _, v := range fs {
		fmt.Println(v.Name())
	}
}

func TestDirectory_FullName(t *testing.T) {

}

func TestDirectory_Name(t *testing.T) {

}

func TestDirectory_Exists(t *testing.T) {

}

func TestDirectory_Parent(t *testing.T) {

}

func TestDirectory_Root(t *testing.T) {

}

func TestNewDirectory(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "NewDirectory"))
	if err != nil {
		t.Fatal(err)
	}

	if dir.Exists() {
		fmt.Println(fmt.Sprintf("directory %s exists", dir.FullName()))
	} else {
		fmt.Println(fmt.Sprintf("directory %s do not exists", dir.FullName()))
	}
}

func TestDirectory_Create(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "Create"))
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		fmt.Println("directory not exists :", dir.FullName())
		err = dir.Create(os.ModePerm)
		if err != nil {
			t.Fatal(err)
		} else {
			fmt.Println("directory created :", dir.FullName())
		}
	} else {
		fmt.Println("directory exists :", dir.FullName())
	}

	err = dir.Delete()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("directory deleted :", dir.FullName())
}

func TestDirectory_CreateSubDirectory(t *testing.T) {

}

func TestDirectory_Delete(t *testing.T) {

}

func TestDirectory_RemoveSubDirectory(t *testing.T) {

}

func TestDirectory_RemoveAllSubDirectory(t *testing.T) {

}

func TestDirectory_RemoveAll(t *testing.T) {

}

func TestDirectory_GetDirectory(t *testing.T) {

}

func TestDirectory_GetDirectoriesLike(t *testing.T) {

}

func TestDirectory_GetAllDirectories(t *testing.T) {

}

func TestDirectory_GetAllFiles(t *testing.T) {

}

func TestDirectory_MoveTo(t *testing.T) {

}

func TestDirectory_SetAccessControl(t *testing.T) {

}

func TestDirectory_Contains(t *testing.T) {

}

func TestDirectory_CreateFile(t *testing.T) {

}

func Test_newDirectory(t *testing.T) {

}

func TestDirectory_GetFile(t *testing.T) {

}
