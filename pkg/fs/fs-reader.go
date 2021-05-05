package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func GetAllFileNames(path string) ([]string, error) {

	files, err := ioutil.ReadDir(path)
	fileNames := []string{}
	if err != nil {
		return nil, fmt.Errorf("unable to read files: %v\n", err)
	}

	for _, file := range files {
		if file.IsDir() != true {
			nowUtc := time.Now().UTC()
			fileModTimeUtc := file.ModTime().UTC()

			// Hack to prevent transfering of incomplete files.
			// Can happen when files are created empty and writes
			// are appended multiple times before the file is
			// complete.
			const backOffSeconds = 5
			recentlyEdited := fileModTimeUtc.Add(time.Second * time.Duration(backOffSeconds)).After(nowUtc)
			fmt.Printf("Recently edited, backing off: %v\n", recentlyEdited)

			if recentlyEdited == false {
				fileNames = append(fileNames, file.Name())
			}
		}
	}

	return fileNames, nil
}

func ReadFile(fullPath string) ([]byte, error) {
	fmt.Printf("file full path: %v\n", fullPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v\n", err)
	}

	return data, nil
}
