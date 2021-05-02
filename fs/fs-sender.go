package fs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sebastianappler/fits/common"
)

func Send(fileName string, fileData []byte, toPath common.Path) error {
	to := filepath.Join(toPath.UrlRaw, fileName)

	out, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v\n", err)
	}

	defer out.Close()

	reader := bytes.NewReader(fileData)
	_, err = io.Copy(out, reader)
	if err != nil {
		return fmt.Errorf("writing to output file failed: %v\n", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %v\n", err)
	}

	fmt.Printf("file sent: %v\n", fileName)
	return nil
}

func Remove(fileName string, path common.Path) error {
	err := os.Remove(filepath.Join(path.Url.Path, fileName))
	if err != nil {
		return fmt.Errorf("unable to remove file: %v\n", err)
	}

	return nil

}
