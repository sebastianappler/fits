package fileservice

import (
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/ssh"
)

type SshFileService struct{}

func (SshFileService) List(path common.Path) ([]string, error) {
	return nil, nil //ssh.List(path.UrlRaw)
}

func (SshFileService) Read(filename string, path common.Path) ([]byte, error) {
	return nil, nil //ssh.Read(filepath.Join(path.Url.Path, fileName))
}

func (SshFileService) Send(filename string, data []byte, path common.Path) error {
	return ssh.Send(filename, data, path.Url, path.Password, path.Username)
}

func (SshFileService) Remove(filename string, path common.Path) error {
	return nil //ssh.Remove(fileName, path.UrlRaw)
}
