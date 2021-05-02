package ftp

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/jlaffaye/ftp"
	"github.com/sebastianappler/fits/common"
)

func Send(fileName string, fileData []byte, toPath common.Path) error {

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

	ftpLocation := filepath.Join(toPath.Url.Path, fileName)
	fmt.Printf("ftp location: %v\n", ftpLocation)

	buf := bytes.NewBuffer(fileData)
	err = c.Stor(ftpLocation, buf)

	if err := c.Quit(); err != nil {
		return fmt.Errorf("error while transfering: %v\n", err)
	}

	fmt.Printf("file uploaded to ftp: %v\n", fileName)

	return nil
}
