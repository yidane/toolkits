package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
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

	f, err := dir.Contains("1")
	if err != nil {
		t.Fatal(err)
	}

	if f {
		err = dir.RemoveSubDirectory("1")
		if err != nil {
			t.Fatal(err)
		}
	}

	err = dir.CreateSubDirectory("1", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

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

	err = dir.Create(os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

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
	dir, err := NewDirectory(filepath.Join(basePath, "FullName"))
	if err != nil {
		t.Fatal(err)
	}

	fullName := dir.FullName()
	if len(fullName) == 0 {
		t.Fatal("fullName is nothing")
	}

	fmt.Println(fullName)
}

func TestDirectory_Name(t *testing.T) {
	name := "Name"
	dir, err := NewDirectory(filepath.Join(basePath, name))
	if err != nil {
		t.Fatal(err)
	}

	if name != dir.Name() {
		t.Fatal("Name should be equal ", name)
	}
}

func TestDirectory_Exists(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "Exists"))
	if err != nil {
		t.Fatal(err)
	}

	if dir.Exists() {
		t.Fatal("dir ", dir.FullName(), " should be not exists")
	}

	err = dir.Create(os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		t.Fatal("dir ", dir.FullName(), " should be exists")
	}

	err = dir.Delete()
	if err != nil {
		t.Fatal(err)
	}
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
	dir, err := NewDirectory(filepath.Join(basePath, "CreateSubDirectory"))
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		err = dir.Create(os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	u := time.Now().UnixNano()
	var i int64 = 0
	for ; i < 10; i++ {
		name := strconv.FormatInt(u+i, 10)
		err = dir.CreateSubDirectory(name, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDirectory_Delete(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "Delete"))
	if err != nil {
		t.Fatal(err)
	}

	exists := dir.Exists()
	fmt.Println("directory ", dir.FullName(), "exists :", exists)

	err = dir.Delete()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("delete 1st")

	err = dir.Create(os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("directory ", dir.FullName(), "created")

	exists = dir.Exists()
	fmt.Println("directory ", dir.FullName(), "exists :", exists)

	err = dir.Delete()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("delete 2nd")

	exists = dir.Exists()
	fmt.Println("directory ", dir.FullName(), "exists :", exists)
}

func TestDirectory_RemoveSubDirectory(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "RemoveSubDirectory"))
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		err = dir.Create(os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	const name = "SubDirectory"

	reportContain := func() {
		f, err := dir.Contains(name)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(name, "contain :", f)
	}

	reportContain()

	err = dir.CreateSubDirectory(name, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("CreateSubDirectory")

	reportContain()

	err = dir.RemoveSubDirectory(name)
	if err != nil {
		t.Fatal(err)
	}

	reportContain()
}

func TestDirectory_RemoveAllSubDirectory(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "RemoveSubDirectory"))
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		err = dir.Create(os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	reportDir := func() {
		subDirs, err := dir.GetDirectories()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("contain subdirectory : ", len(subDirs))
	}

	reportDir()

	u := time.Now().UnixNano()
	var i int64 = 0
	for ; i < 10; i++ {
		name := strconv.FormatInt(u+i, 10)
		err = dir.CreateSubDirectory(name, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	reportDir()

	err = dir.RemoveAllSubDirectory()
	if err != nil {
		t.Fatal(err)
	}

	reportDir()
}

func TestDirectory_RemoveAll(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "RemoveAll"))
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		err = dir.Create(os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	var reportDir = func() {
		subDirs, err := dir.GetDirectories()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("contain subdirectory : ", len(subDirs))

		files, err := dir.GetFiles()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("contain files : ", len(files))
	}

	reportDir()

	u := time.Now().UnixNano()
	var i int64 = 0
	for ; i < 20; i++ {
		name := strconv.FormatInt(u+i, 10)
		if i%2 == 0 {
			err = dir.CreateSubDirectory(name, os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}
		} else {
			_, err = dir.CreateFile(name)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	reportDir()

	err = dir.RemoveAll()
	if err != nil {
		t.Fatal(err)
	}

	reportDir()
}

func TestDirectory_GetDirectory(t *testing.T) {
	dir, err := NewDirectory(filepath.Join(basePath, "GetDirectory"))
	if err != nil {
		t.Fatal(err)
	}

	if !dir.Exists() {
		err = dir.Create(os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	const name = "yidane"

	reportExists := func() {
		directory, err := dir.GetDirectory(name)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("directory ", name, " exists :", directory.Exists())
	}

	reportExists()

	err = dir.CreateSubDirectory(name, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("CreateSubDirectory")
	reportExists()

	err = dir.RemoveSubDirectory(name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("RemoveSubDirectory")

	reportExists()
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

func TestDirectory_GetFile(t *testing.T) {

}
