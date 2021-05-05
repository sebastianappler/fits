package gateway

import (
	"fmt"
	"path/filepath"

	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/fs"
	"github.com/sebastianappler/fits/pkg/ftp"
	"github.com/sebastianappler/fits/pkg/ssh"
)

func Send(fileName string, fileData []byte, path common.Path) error {

	scheme := path.Url.Scheme

	if scheme == "" {
		err := fsSend(fileName, fileData, path)
		if err != nil {
			return fmt.Errorf("gateway send fs: %v\n", err)
		}
	}

	if scheme == "ftp" {
		err := ftpSend(fileName, fileData, path)
		if err != nil {
			return fmt.Errorf("gateway send ftp: %v\n", err)
		}
	}

	if scheme == "ssh" {
		err := sshSend(fileName, fileData, path)
		if err != nil {
			return fmt.Errorf("gateway send ssh: %v\n", err)
		}
	}

	return nil
}

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

func Remove(fileName string, path common.Path) error {

	scheme := path.Url.Scheme

	if scheme == "" {
		err := fsRemove(fileName, path)
		if err != nil {
			return fmt.Errorf("Unable to send file: %v", err)
		}
	}

	if scheme == "ftp" {
		// TODO
		fmt.Println("remove files from ftp not implemented")
	}
	if scheme == "ssh" {
		// TODO
		fmt.Println("remove files from ssh not implemented")
	}

	return nil
}

func fsReadFile(fileName string, path common.Path) ([]byte, error) {
	return fs.ReadFile(filepath.Join(path.Url.Path, fileName))
}

func fsGetAllFileNames(path common.Path) ([]string, error) {
	return fs.GetAllFileNames(path.Url.Path)
}
func fsSend(fileName string, fileData []byte, path common.Path) error {
	return fs.Send(fileName, fileData, path.UrlRaw)
}

func ftpSend(fileName string, fileData []byte, path common.Path) error {
	return ftp.Send(fileName, fileData, path.Url, path.Username, path.Password)
}

func sshSend(fileName string, fileData []byte, path common.Path) error {
	return ssh.Send(fileName, fileData, path.Url, path.Username, path.Password)
}

func fsRemove(fileName string, path common.Path) error {
	return fs.Remove(fileName, path.UrlRaw)
}
