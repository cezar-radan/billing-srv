package cfile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	filePerm   os.FileMode = 0640
	folderPerm os.FileMode = 0755
)

// function to move a file (source -> target)
func MoveFile(source, target string) error {
	err := os.Rename(source, target)
	if err != nil {
		return err
	}
	return nil
}

// function to create a file
func WriteFile(fileName string, data []byte) error {
	err := ioutil.WriteFile(fileName, data, filePerm)
	if err != nil {
		return err
	}
	return nil
}

// function to remove a file
func RemoveFile(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("%s is a directory not a file. Cannot remove a directory", filename)
	}

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}

//function to copy a file (source -> target)
func CopyFile(source, target string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(target, input, filePerm)
	if err != nil {
		return err
	}
	return nil
}

// function to check if file exists
func CheckFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// function to check if folder exists
func CheckFolderExists(base string) bool {
	info, err := os.Stat(base)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// function to create folder if not exists
func CreateFolder(base string) error {
	info, err := os.Stat(base)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(base, folderPerm)
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}

	if !info.IsDir() {
		return errors.New("object already exists, but is not folder")
	}

	return nil
}

// function to remove/delete folder if exists
func RemoveFolder(base string) error {
	_, err := os.Stat(base)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}

	err = os.RemoveAll(base)
	if err != nil {
		return err
	}
	return nil
}
