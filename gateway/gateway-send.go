package gateway

import (
	"fmt"

	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/fs"
	"github.com/sebastianappler/fits/ftp"
	"github.com/sebastianappler/fits/ssh"
)

func Send(fileName string, fileData []byte, path common.Path) error {

	scheme := path.Url.Scheme

	if scheme == "" {
		err := fs.Send(fileName, fileData, path)
		if err != nil {
			return fmt.Errorf("gateway.send fs: %v\n", err)
		}
	}

	if scheme == "ftp" {
		err := ftp.Send(fileName, fileData, path)
		if err != nil {
			return fmt.Errorf("gateway.send ftp: %v\n", err)
		}
	}

	if scheme == "ssh" {
		err := ssh.Send(fileName, fileData, path)
		if err != nil {
			return fmt.Errorf("gateway.send ssh: %v\n", err)
		}
	}

	return nil
}
