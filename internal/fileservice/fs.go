package fileservice

import (
	"path/filepath"

	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/pkg/fs"
)

type FsFileService struct{}

func (FsFileService) List(path common.Path) ([]string, error) {
	return fs.List(path.UrlRaw)
}

func (FsFileService) Read(fileName string, path common.Path) ([]byte, error) {
	return fs.Read(filepath.Join(path.Url.Path, fileName))
}

func (FsFileService) Send(fileName string, fileData []byte, path common.Path) error {
	return fs.Send(fileName, fileData, path.UrlRaw)
}

func (FsFileService) Remove(fileName string, path common.Path) error {
	return fs.Remove(fileName, path.UrlRaw)
}
