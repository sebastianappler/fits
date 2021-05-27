package ftp

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	fmt.Printf("connecting to ftp %v... ", ftpBaseUrl)
	c, err := ftp.Connect(ftpBaseUrl)
	err = c.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("unable to login: %v\n", err)
	}
	fmt.Printf("success\n")

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

	return nil
}

func List(path string, url url.URL, username string, password string) ([]string, error) {
	c, err := connect(url, username, password)
	files, err := c.List(path)
	if err != nil {
		return nil, fmt.Errorf("unable to list files", err)
	}

	filenames := []string{}
	for _, file := range files {

		if file.Type == ftp.EntryTypeFile {
			filenames = append(filenames, file.Name)
		}
	}

	return filenames, nil
}

func Read(filename string, url url.URL, username string, password string) ([]byte, error) {
	c, err := connect(url, username, password)
	ftpLocation := filepath.Join(url.Path, filename)
	r, err := c.Retr(ftpLocation)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v\n", err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	return buf, nil
}

func Remove(filename string, url url.URL, username string, password string) error {
	c, err := connect(url, username, password)
	ftpLocation := filepath.Join(url.Path, filename)
	err = c.Delete(ftpLocation)
	if err != nil {
		return fmt.Errorf("unable to delete file: %v\n", err)
	}
	return nil
}
