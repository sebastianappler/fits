package ftp

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/jlaffaye/ftp"
)

func Send(fileName string, fileData []byte, url url.URL, username string, password string) error {

	port := url.Port()
	if port == "" {
		port = "21"
	}

	ftpBaseUrl := url.Host + ":" + port
	fmt.Printf("connecting to ftp %v...\n", ftpBaseUrl)
	c, err := ftp.Connect(ftpBaseUrl)
	fmt.Println("connected")

	fmt.Println("logging in...")
	err = c.Login(username, password)
	if err != nil {
		return fmt.Errorf("unable to login: %v\n", err)
	}
	fmt.Println("login success")

	ftpLocation := filepath.Join(url.Path, fileName)
	fmt.Printf("ftp location: %v\n", ftpLocation)

	buf := bytes.NewBuffer(fileData)
	err = c.Stor(ftpLocation, buf)

	if err := c.Quit(); err != nil {
		return fmt.Errorf("error while transfering: %v\n", err)
	}

	fmt.Printf("file uploaded to ftp: %v\n", fileName)

	return nil
}
