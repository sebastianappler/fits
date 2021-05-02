package gateway

import (
	"fmt"

	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/fs"
)

func Remove(fileName string, path common.Path) error {

	scheme := path.Url.Scheme

	if scheme == "" {
		err := fs.Remove(fileName, path)
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
