package fileservice

import (
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/ssh"
)

type SshFileService struct{}

func (SshFileService) List(path common.Path) ([]string, error) {
	return nil, nil //ssh.List(path.UrlRaw)
}

func (SshFileService) Read(fileName string, path common.Path) ([]byte, error) {
	return nil, nil //ssh.Read(filepath.Join(path.Url.Path, fileName))
}

func (SshFileService) Send(fileName string, fileData []byte, path common.Path) error {
	return ssh.Send(fileName, fileData, path.Url, path.Password, path.Username)
}

func (SshFileService) Remove(fileName string, path common.Path) error {
	return nil //ssh.Remove(fileName, path.UrlRaw)
}
