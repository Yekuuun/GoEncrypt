package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// Getting path to root project.
func GetRootPath(indicator string) (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", errors.New("unable to get current dir path")
	}

	for {
		if _, err := os.Stat(filepath.Join(currentPath, indicator)); err == nil {
			return currentPath, nil
		}

		parent := filepath.Dir(currentPath)
		if parent == currentPath {
			return "", errors.New("indicator not found")
		}

		currentPath = parent
	}
}

// check if keys folder is empty
func ContainsKeys() (bool, error) {
	rootPath, err := GetRootPath("go.mod")
	if err != nil {
		return false, errors.New("unable to find root path")
	}

	keyFolder := filepath.Join(rootPath, "data/keys")

	dir, err := os.Open(keyFolder)
	if err != nil {
		return false, errors.New("unable to open keys folder")
	}
	defer dir.Close()

	//read folder content
	files, err := dir.ReadDir(-1)
	if err != nil {
		return false, errors.New("unable to read dir content")
	}

	return !(len(files) == 0), nil
}
