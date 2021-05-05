package gateway

import (
	"fmt"

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

func fsSend(fileName string, fileData []byte, path common.Path) error {
	return fs.Send(fileName, fileData, path.UrlRaw)
}

func ftpSend(fileName string, fileData []byte, path common.Path) error {
	return ftp.Send(fileName, fileData, path.Url, path.Username, path.Password)
}

func sshSend(fileName string, fileData []byte, path common.Path) error {
	return ssh.Send(fileName, fileData, path.Url, path.Username, path.Password)
}
