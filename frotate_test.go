package frotate_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tharindu96/frotate-go"
)

func TestNewRotateFile(t *testing.T) {
	ti := time.Now()
	b := frotate.NewRotateFile(ti, "xx", ".xx", "./")
	assert.NotNil(t, b)
}

func TestGetRotateFileFromFile(t *testing.T) {
	// create temp file
	testFileName := "./backup-2020-01-01-02-02-02.zip"

	assert.Nil(t, createFile(testFileName))

	backup, err := frotate.GetRotateFileFromFile(testFileName, "backup")
	assert.Nil(t, err)
	assert.NotNil(t, backup)

	assert.Nil(t, deleteFile(testFileName))
}

func TestRotateFileSave(t *testing.T) {
	fileName := "test.xx"

	// create temp file
	assert.Nil(t, createFile(fileName))

	backup, err := frotate.SaveRotateFile(fileName, "backup", "./")
	assert.Nil(t, err)
	assert.NotNil(t, backup)

	assert.Nil(t, backup.Delete())
}

func TestRotateFileSaveError(t *testing.T) {
	backup, err := frotate.SaveRotateFile("imageinaryfasf.xx", "backup", "./")
	assert.NotNil(t, err)
	assert.Nil(t, backup)
}

func TestLoadRotateFilesFromDirectory(t *testing.T) {
	createFile("./backup-2020-01-01-02-02-02.zip")

	rfiles, err := frotate.LoadRotateFilesFromDirectory("./", "backup")
	assert.Nil(t, err)
	assert.NotNil(t, rfiles)

	assert.Len(t, rfiles, 1)

	for i, j := 0, 1; j < len(rfiles); i++ {
		a := rfiles[i].Date()
		b := rfiles[j].Date()
		assert.True(t, a.Before(*b))
		j++
	}

	assert.Nil(t, rfiles[0].Delete())

	rfiles, err = frotate.LoadRotateFilesFromDirectory("./", "backup")
	assert.Nil(t, err)
	assert.NotNil(t, rfiles)

	assert.Len(t, rfiles, 0)
}

func TestAddFile(t *testing.T) {

	createFile("a.txt")
	createFile("b.txt")
	createFile("c.txt")

	assert.Nil(t, frotate.AddFile("a.txt", "backup", "./", 2))

	time.Sleep(1 * time.Second)

	assert.Nil(t, frotate.AddFile("b.txt", "backup", "./", 2))

	time.Sleep(1 * time.Second)

	assert.Nil(t, frotate.AddFile("c.txt", "backup", "./", 2))

	rfiles, err := frotate.LoadRotateFilesFromDirectory("./", "backup")
	assert.Nil(t, err)
	assert.NotNil(t, rfiles)

	assert.Len(t, rfiles, 2)

	for _, rfile := range rfiles {
		assert.Nil(t, rfile.Delete())
	}

}

func createFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(""), 0666)
}

func deleteFile(filename string) error {
	return os.Remove(filename)
}
