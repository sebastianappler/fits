package senders

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jlaffaye/ftp"
	"github.com/sebastianappler/fits/common"
)

func FtpSend(fileLocalPath string, toPath common.Path) error {

	port := toPath.Url.Port()
	if port == "" {
		port = "21"
	}

	ftpBaseUrl := toPath.Url.Host + ":" + port
	fmt.Printf("connecting to ftp %v...\n", ftpBaseUrl)
	c, err := ftp.Connect(ftpBaseUrl)
	fmt.Println("connected")

	fmt.Println("logging in...")
	err = c.Login(toPath.Username, toPath.Password)
	if err != nil {
		return fmt.Errorf("unable to login: %v\n", err)
	}
	fmt.Println("login success")

	fmt.Printf("file full path: %v\n", fileLocalPath)
	data, err := os.ReadFile(fileLocalPath)
	if err != nil {
		return fmt.Errorf("unable to read file: %v\n", err)
	}

	filename := filepath.Base(fileLocalPath)
	ftpLocation := filepath.Join(toPath.Url.Path, filename)
	fmt.Printf("ftp location: %v\n", ftpLocation)

	buf := bytes.NewBuffer(data)
	err = c.Stor(ftpLocation, buf)

	if err := c.Quit(); err != nil {
		return fmt.Errorf("error while transfering: %v\n", err)
	}

	fmt.Printf("file %v uploaded to ftp\n", filename)

	return nil
}
