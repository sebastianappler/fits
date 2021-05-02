package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sebastianappler/fits/common"
)

func GetAllFileNames(path common.Path) ([]string, error) {

	files, err := ioutil.ReadDir(path.Url.Path)
	fileNames := []string{}
	if err != nil {
		return nil, fmt.Errorf("unable to read files: %v\n", err)
	}

	for _, file := range files {
		if file.IsDir() != true {
			fmt.Printf("found file %v\n", file.Name())
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func ReadFile(fileName string, path common.Path) ([]byte, error) {
	return readFile(filepath.Join(path.Url.Path, fileName))
}

func readFile(filePath string) ([]byte, error) {
	fmt.Printf("file full path: %v\n", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v\n", err)
	}

	return data, nil
}
