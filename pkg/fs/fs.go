package fs

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func List(path string) ([]string, error) {

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

func Read(fullPath string) ([]byte, error) {
	fmt.Printf("file full path: %v\n", fullPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v\n", err)
	}

	return data, nil
}

func Send(filename string, data []byte, path string) error {
	to := filepath.Join(path, filename)

	out, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v\n", err)
	}

	defer out.Close()

	reader := bytes.NewReader(data)
	_, err = io.Copy(out, reader)
	if err != nil {
		return fmt.Errorf("writing to output file failed: %v\n", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %v\n", err)
	}

	fmt.Printf("file sent: %v\n", filename)
	return nil
}

func Remove(filename string, path string) error {
	err := os.Remove(filepath.Join(path, filename))
	if err != nil {
		return fmt.Errorf("unable to remove file: %v\n", err)
	}

	return nil
}
