package gateway

import (
	"fmt"
	"path/filepath"

	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/fs"
)

func GetAllFileNames(path common.Path) ([]string, error) {
	scheme := path.Url.Scheme

	if scheme == "" {
		fileNames, err := fsGetAllFileNames(path)
		if err != nil {
			return nil, fmt.Errorf("unable get file names: %v\n", err)
		}
		return fileNames, nil
	}

	if scheme == "ftp" {
		// TODO
		fmt.Println("get files from ftp not implemented")
	}
	if scheme == "ssh" {
		// TODO
		fmt.Println("get files from ssh not implemented")
	}
	return nil, nil
}

func ReadFile(fileName string, path common.Path) ([]byte, error) {
	scheme := path.Url.Scheme

	if scheme == "" {
		data, err := fsReadFile(fileName, path)
		if err != nil {
			return nil, fmt.Errorf("unable to send file: %v\n", err)
		}
		return data, nil
	}

	if scheme == "ftp" {
		// TODO
		fmt.Println("read files from ftp not implemented")
	}
	if scheme == "ssh" {
		// TODO
		fmt.Println("read files from ssh not implemented")
	}

	return nil, nil
}

func fsReadFile(fileName string, path common.Path) ([]byte, error) {
	return fs.ReadFile(filepath.Join(path.Url.Path, fileName))
}

func fsGetAllFileNames(path common.Path) ([]string, error) {
	return fs.GetAllFileNames(path.Url.Path)
}
