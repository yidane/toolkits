package io

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	path := "/tmp/yidane"

	d, err := CreateDirectory(path, os.FileMode(777))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(d)
}

func TestDirectory_GetDirectories(t *testing.T) {
	path := "/tmp"

	dir, err := NewDirectory(path)
	if err != nil {
		t.Fatal(err)
	}

	dirs, err := dir.GetDirectories()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(dirs))
	for _, v := range dirs {
		fmt.Println(v.FullName)
	}
}

func TestDirectory_GetFiles(t *testing.T) {
	path := "/tmp/yidane"

	dir, err := NewDirectory(path)
	if err != nil {
		t.Fatal(err)
	}

	dirs, err := dir.GetFiles()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(dirs))
	for _, v := range dirs {
		fmt.Println(v.FullName)
	}
}
