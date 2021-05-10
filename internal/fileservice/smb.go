package fileservice

import (
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/smb"
)

type SmbFileService struct{}

func (SmbFileService) List(path common.Path) ([]string, error) {
	return nil, nil //smb.List(path.UrlRaw)
}

func (SmbFileService) Read(filename string, path common.Path) ([]byte, error) {
	return nil, nil //smb.Read(filepath.Join(path.Url.Path, fileName))
}

func (SmbFileService) Send(filename string, data []byte, path common.Path) error {
	return smb.Send(filename, data, path.Url, path.Username, path.Password)
}

func (SmbFileService) Remove(filename string, path common.Path) error {
	return nil //smb.Remove(fileName, path.UrlRaw)
}
