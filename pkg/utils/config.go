package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Printing main banner.
func PrintBanner() {
	fmt.Printf("\n")

	blue := "\033[34m"
	reset := "\033[0m"

	banner := `          ____       _____                             _   
         / ___| ___ | ____|_ __   ___ _ __ _   _ _ __ | |_ 
        | |  _ / _ \|  _| | '_ \ / __| '__| | | | '_ \| __|
        | |_| | (_) | |___| | | | (__| |  | |_| | |_) | |_ 
         \____|\___/|_____|_| |_|\___|_|   \__, | .__/ \__|
                                           |___/|_|          
        ---simple cli file encryption tool written in GO---`

	fmt.Println(blue + banner + reset)
}

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
