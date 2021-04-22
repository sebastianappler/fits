package senders

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jlaffaye/ftp"
	"github.com/sebastianappler/fits/common"
)

func FtpSend(fileLocalPath string, toPath common.Path) error {

	ftpBaseUrl := toPath.Url.Hostname() + ":21"
	fmt.Printf("Connecting to ftp %v...\n", ftpBaseUrl)
	c, err := ftp.Connect(ftpBaseUrl)
	fmt.Println("Connected!")

	fmt.Println("Logging in...")
	err = c.Login(toPath.Username, toPath.Password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Login success!")

	fmt.Printf("File full path: %v\n", fileLocalPath)
	file, err := os.Open(fileLocalPath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	filename := filepath.Base(fileLocalPath)
	ftpLocation := filepath.Join(toPath.Url.Path, filename)
	fmt.Printf("FtpLocation: %v\n", ftpLocation)

	buf := bytes.NewBuffer(data)
	err = c.Stor(ftpLocation, buf)

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %v uploaded to ftp!\n", filename)

	err = os.Remove(fileLocalPath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil
}
