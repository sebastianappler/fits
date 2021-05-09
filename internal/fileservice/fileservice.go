package fileservice

import "github.com/sebastianappler/fits/internal/common"

type FileService interface {
	List(path common.Path) ([]string, error)
	Read(fileName string, path common.Path) ([]byte, error)
	Send(fileName string, fileData []byte, path common.Path) error
	Remove(fileName string, path common.Path) error
}
