package senders

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sebastianappler/fits/common"
)

func FsSend(fileLocalPath string, toPath common.Path) error {
	to := filepath.Join(toPath.UrlRaw, filepath.Base(fileLocalPath))
	in, err := os.Open(fileLocalPath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	out, err := os.Create(to)
	if err != nil {
		in.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	in.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %s", err)
	}

	si, err := os.Stat(fileLocalPath)
	if err != nil {
		return fmt.Errorf("Stat error: %s", err)
	}

	err = os.Chmod(to, si.Mode())
	if err != nil {
		return fmt.Errorf("Chmod error: %s", err)
	}

	return nil
}
