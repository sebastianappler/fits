package fileservice

import (
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/ftp"
)

type FtpFileService struct{}

func (FtpFileService) List(path common.Path) ([]string, error) {
	return nil, nil //ftp.List(path.UrlRaw)
}

func (FtpFileService) Read(filename string, path common.Path) ([]byte, error) {
	return nil, nil //ftp.Read(filepath.Join(path.Url.Path, fileName))
}

func (FtpFileService) Send(filename string, data []byte, path common.Path) error {
	return ftp.Send(filename, data, path.Url, path.Password, path.Username)
}

func (FtpFileService) Remove(filename string, path common.Path) error {
	return nil //ftp.Remove(fileName, path.UrlRaw)
}
