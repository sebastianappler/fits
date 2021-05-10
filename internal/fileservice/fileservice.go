package fileservice

import "github.com/sebastianappler/fits/internal/common"

type FileService interface {
	List(path common.Path) ([]string, error)
	Read(filename string, path common.Path) ([]byte, error)
	Send(filename string, data []byte, path common.Path) error
	Remove(filename string, path common.Path) error
}
