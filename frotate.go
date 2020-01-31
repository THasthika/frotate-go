package frotate

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"time"
)

// RotateFile item
type RotateFile struct {
	date      time.Time
	prefix    string
	ext       string
	directory string
}

// save item to disk
// delete item from disk
// given a list of items sort them to date
// read a directory and get all backup files
// add a file to directory with date format and remove if the size is greater than the limit

// NewRotateFile create a new backup file in memory
func NewRotateFile(date time.Time, prefix string, ext string, directory string) *RotateFile {
	return &RotateFile{
		date,
		prefix,
		ext,
		directory,
	}
}

// SaveRotateFile save backup file
func SaveRotateFile(oldFilePath string, prefix string, directory string) (*RotateFile, error) {
	t := time.Now()
	ext := path.Ext(oldFilePath)
	filename := getFormattedFileName(t, prefix, ext)
	newFilePath := path.Join(directory, filename)
	if err := os.Rename(oldFilePath, newFilePath); err != nil {
		return nil, err
	}
	RotateFile := NewRotateFile(t, prefix, ext, directory)
	return RotateFile, nil
}

// GetRotateFileFromFile reads file name and extract RotateFile
func GetRotateFileFromFile(filename string, prefix string) (*RotateFile, error) {
	directory, filename := path.Split(filename)
	re := regexp.MustCompile(`^` + prefix + `-(\d{4}-\d{2}-\d{2}-\d{2}-\d{2}-\d{2})\..*`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) != 2 {
		return nil, errors.New("Not matching the format")
	}
	dateString := matches[1]
	ext := path.Ext(filename)
	date, err := time.Parse("2006-01-02-15-04-05", dateString)
	if err != nil {
		return nil, err
	}

	return NewRotateFile(date, prefix, ext, directory), nil
}

// LoadRotateFilesFromDirectory load files to array
func LoadRotateFilesFromDirectory(directory string, prefix string) ([]*RotateFile, error) {
	// list files
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	rotateFiles := make([]*RotateFile, 0)

	// load the list of files to Build
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fname := f.Name()
		fpath := path.Join(directory, fname)
		rfile, err := GetRotateFileFromFile(fpath, prefix)
		if err != nil {
			continue
		}
		rotateFiles = append(rotateFiles, rfile)
	}

	// sort the list by date - latest file is at the back
	sort.Slice(rotateFiles[:], func(i, j int) bool {
		return rotateFiles[i].date.Before(rotateFiles[i].date)
	})

	return rotateFiles, nil
}

// AddFile add the file in filepath to the directory and remove old files if limit is exceeded
func AddFile(filepath string, prefix string, directory string, limit uint) error {
	// list files
	rfiles, err := LoadRotateFilesFromDirectory(directory, prefix)
	if err != nil {
		return err
	}

	rfile, err := SaveRotateFile(filepath, prefix, directory)
	if err != nil {
		return err
	}

	rfiles = append(rfiles, rfile)

	for uint(len(rfiles)) > limit {
		var rfile *RotateFile
		rfile, rfiles = rfiles[0], rfiles[1:]
		if err := rfile.Delete(); err != nil {
			return err
		}
	}

	return nil

}

// Date returns the date
func (rotateFile *RotateFile) Date() *time.Time {
	return &rotateFile.date
}

// Delete delete backup file
func (rotateFile *RotateFile) Delete() error {
	fileName := getFormattedFileName(rotateFile.date, rotateFile.prefix, rotateFile.ext)
	filePath := path.Join(rotateFile.directory, fileName)
	return os.Remove(filePath)
}

func getFormattedFileName(date time.Time, prefix string, ext string) string {
	dateString := date.Format("2006-01-02-15-04-05")
	return prefix + "-" + dateString + ext
}
