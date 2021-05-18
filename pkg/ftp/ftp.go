package ftp

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/jlaffaye/ftp"
)

func connect(url url.URL, username string, password string) (*ftp.ServerConn, error) {
	port := url.Port()
	if port == "" {
		port = "21"
	}

	ftpBaseUrl := url.Hostname() + ":" + port
	fmt.Printf("connecting to ftp %v...\n", ftpBaseUrl)
	c, err := ftp.Connect(ftpBaseUrl)
	fmt.Println("connected")

	fmt.Println("logging in...")
	err = c.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("unable to login: %v\n", err)
	}
	fmt.Println("login success")

	return c, nil
}

func Send(filename string, data []byte, url url.URL, username string, password string) error {
	c, err := connect(url, username, password)
	if err != nil {
		return err
	}

	ftpLocation := filepath.Join(url.Path, filename)
	fmt.Printf("ftp location: %v\n", ftpLocation)

	buf := bytes.NewBuffer(data)
	err = c.Stor(ftpLocation, buf)

	if err := c.Quit(); err != nil {
		return fmt.Errorf("error while transfering: %v\n", err)
	}

	fmt.Printf("file uploaded to ftp: %v\n", filename)
	return nil
}

func List(path string, url url.URL, username string, password string) ([]string, error) {
	c, err := connect(url, username, password)
	if err != nil {
		return nil, err
	}

	files, err := c.List(path)
	filenames := []string{}
	for _, file := range files {

		if file.Type == ftp.EntryTypeFile {
			filenames = append(filenames, file.Name)
		}
	}

	return filenames, nil
}
